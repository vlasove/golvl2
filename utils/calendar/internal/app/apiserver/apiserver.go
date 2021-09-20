package apiserver

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/vlasove/materials/tasks_2/utils/calendar/internal/app/models"
	"github.com/vlasove/materials/tasks_2/utils/calendar/internal/app/store"
)

var (
	errBadRequestByMethod     = errors.New("method not allowed for this url")
	errQueryParamNotProvided  = errors.New("should provided 'date' as YYYY-MM-DD")
	errInvalidQueryDate       = errors.New("date should be YYYY-MM-DD")
	errNotProvidedIDInForm    = errors.New("request body should contains id:int")
	errNotPovidedUserIDInForm = errors.New("request body should contains user_id:int")
	errNotProvidedDateInForm  = errors.New("request body should conatains date:YYYY-MM-DD")
	errNotProvidedInfoInForm  = errors.New("request body should contains info:string")
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *log.Logger
	router *http.ServeMux
	store  *store.Store
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: log.Default(),
		router: http.NewServeMux(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Println("logger configurated successfully")

	if err := s.configureStore(); err != nil {
		return err
	}
	defer s.store.Close()
	s.logger.Println("database configurated successfully")

	r := s.configureRouter()
	s.logger.Println("router configurated successfully")

	s.logger.Println("starting api server at port:", s.config.BindAddr)
	return http.ListenAndServe(s.config.BindAddr, r)
}

func (s *APIServer) configureLogger() error {
	file, err := os.OpenFile(s.config.LogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	mw := io.MultiWriter(os.Stdout, file)
	s.logger.SetOutput(mw)
	return nil
}

func (s *APIServer) configureStore() error {
	st := store.New()
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

func (s *APIServer) configureRouter() http.Handler {
	s.router.HandleFunc("/create_event", s.handleCreate())
	s.router.HandleFunc("/update_event", s.handleUpdate())
	s.router.HandleFunc("/delete_event", s.handleDelete())
	s.router.HandleFunc("/events_for_day", s.handleGetForDay())
	s.router.HandleFunc("/events_for_week", s.handleGetForWeek())
	s.router.HandleFunc("/events_for_month", s.handleGetForMonth())
	return s.loggingMiddleware(s.router)
}

func (s *APIServer) handleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if r.Header.Get("content-type") != "application/x-www-form-urlencoded" {
				s.error(w, r, http.StatusUnsupportedMediaType, nil)
				return
			}
			if err := r.ParseForm(); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}

			eventR, err := s.decodeFormCreate(r.Form)
			if err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}

			if err := eventR.Validate(); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}

			event := models.NewEventFromRequest(eventR)
			if err := s.store.EventRepository().CreateEvent(event); err != nil {
				s.error(w, r, 503, err)
				return
			}

			s.respond(w, r, http.StatusCreated, nil)
			return
		}
		s.error(w, r, http.StatusBadRequest, errBadRequestByMethod)
	}
}

func (s *APIServer) handleUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if r.Header.Get("content-type") != "application/x-www-form-urlencoded" {
				s.error(w, r, http.StatusUnsupportedMediaType, nil)
				return
			}
			if err := r.ParseForm(); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}

			eventR, err := s.decodeFormUpdate(r.Form)
			if err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}

			if err := eventR.Validate(); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}

			event := models.NewEventFromRequest(eventR)
			if err := s.store.EventRepository().UpdateEvent(event); err != nil {
				s.error(w, r, http.StatusServiceUnavailable, err)
				return
			}

			s.respond(w, r, http.StatusAccepted, nil)
			return
		}
		s.error(w, r, http.StatusBadRequest, errBadRequestByMethod)
	}
}

func (s *APIServer) handleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if r.Header.Get("content-type") != "application/x-www-form-urlencoded" {
				s.error(w, r, http.StatusUnsupportedMediaType, nil)
				return
			}
			if err := r.ParseForm(); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}
			val, ok := r.Form["id"]
			if !ok {
				s.error(w, r, http.StatusBadRequest, errNotProvidedIDInForm)
				return
			}
			id, err := strconv.Atoi(val[0])
			if err != nil {
				s.error(w, r, http.StatusBadRequest, errNotProvidedIDInForm)
				return
			}
			if err := s.store.EventRepository().DeleteEvent(id); err != nil {
				s.error(w, r, 503, err)
				return
			}

			s.respond(w, r, http.StatusAccepted, nil)
			return
		}
		s.error(w, r, http.StatusBadRequest, errBadRequestByMethod)
	}
}

func (s *APIServer) handleGetForDay() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			vals := r.URL.Query()
			date, ok := vals["date"]
			log.Println(date)
			if !ok {
				s.error(w, r, http.StatusBadRequest, errQueryParamNotProvided)
				return
			}
			dateStringCurrent, err := time.Parse(store.BaseTimeSample, date[0])
			if err != nil {
				s.error(w, r, http.StatusBadRequest, errInvalidQueryDate)
				return
			}
			dateStringFuture := dateStringCurrent.AddDate(0, 0, 1).Format(store.BaseTimeSample)

			events, err := s.store.EventRepository().GetEventsForDates(date[0], dateStringFuture)
			if err != nil {
				s.error(w, r, http.StatusServiceUnavailable, err)
				return
			}
			s.respond(w, r, http.StatusOK, map[string]interface{}{"events": events})
			return
		}
		s.error(w, r, http.StatusBadRequest, errBadRequestByMethod)
	}
}

func (s *APIServer) handleGetForWeek() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			vals := r.URL.Query()
			date, ok := vals["date"]
			log.Println(date)
			if !ok {
				s.error(w, r, http.StatusBadRequest, errQueryParamNotProvided)
				return
			}
			dateStringCurrent, err := time.Parse(store.BaseTimeSample, date[0])
			if err != nil {
				s.error(w, r, http.StatusBadRequest, errInvalidQueryDate)
				return
			}
			dateStringFuture := dateStringCurrent.AddDate(0, 0, 7).Format(store.BaseTimeSample)

			events, err := s.store.EventRepository().GetEventsForDates(date[0], dateStringFuture)
			if err != nil {
				s.error(w, r, http.StatusServiceUnavailable, err)
				return
			}
			s.respond(w, r, http.StatusOK, map[string]interface{}{"events": events})
			return
		}
		s.error(w, r, http.StatusBadRequest, errBadRequestByMethod)
	}
}

func (s *APIServer) handleGetForMonth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			vals := r.URL.Query()
			date, ok := vals["date"]
			if !ok {
				s.error(w, r, http.StatusBadRequest, errQueryParamNotProvided)
				return
			}
			dateStringCurrent, err := time.Parse(store.BaseTimeSample, date[0])
			if err != nil {
				s.error(w, r, http.StatusBadRequest, errInvalidQueryDate)
				return
			}
			dateStringFuture := dateStringCurrent.AddDate(0, 1, 0).Format(store.BaseTimeSample)

			events, err := s.store.EventRepository().GetEventsForDates(date[0], dateStringFuture)
			if err != nil {
				s.error(w, r, http.StatusServiceUnavailable, err)
				return
			}
			s.respond(w, r, http.StatusOK, map[string]interface{}{"events": events})
			return
		}
		s.error(w, r, http.StatusBadRequest, errBadRequestByMethod)
	}
}

func (s *APIServer) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *APIServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}

func (s *APIServer) decodeFormUpdate(form url.Values) (*models.EventRequest, error) {
	eventR := new(models.EventRequest)
	valID, ok := form["id"]
	if !ok {
		return nil, errNotProvidedIDInForm
	}
	id, err := strconv.Atoi(valID[0])
	if err != nil {
		return nil, errNotProvidedIDInForm
	}
	eventR.ID = id

	valUserID, ok := form["userID"]
	if !ok {
		return nil, errNotPovidedUserIDInForm
	}
	userID, err := strconv.Atoi(valUserID[0])
	if err != nil {
		return nil, errNotPovidedUserIDInForm
	}
	eventR.UserID = userID

	valDate, ok := form["date"]
	if !ok {
		return nil, errNotProvidedDateInForm
	}
	eventR.Date = valDate[0]

	valInfo, ok := form["info"]
	if !ok {
		return nil, errNotProvidedInfoInForm
	}
	eventR.Info = valInfo[0]
	return eventR, nil
}

func (s *APIServer) decodeFormCreate(form url.Values) (*models.EventRequest, error) {
	eventR := new(models.EventRequest)

	valUserID, ok := form["userID"]
	if !ok {
		return nil, errNotPovidedUserIDInForm
	}
	userID, err := strconv.Atoi(valUserID[0])
	if err != nil {
		return nil, errNotPovidedUserIDInForm
	}
	eventR.UserID = userID

	valDate, ok := form["date"]
	if !ok {
		return nil, errNotProvidedDateInForm
	}
	eventR.Date = valDate[0]

	valInfo, ok := form["info"]
	if !ok {
		return nil, errNotProvidedInfoInForm
	}
	eventR.Info = valInfo[0]
	return eventR, nil
}
