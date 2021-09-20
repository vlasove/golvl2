package wget

import (
	"errors"
	"net/http"

	"github.com/vlasove/golvl2/develop/wget/internal/app/managers"
)

var (
	errManagerBuild = errors.New("wget: errors when creating directory")
	errHTTPGetBad   = errors.New("wget: can not make get request")
)

// WGet ...
type WGet struct {
	manager managers.Manager
	BaseURL string
}

// New ...
func New(url string, manager managers.Manager) *WGet {
	return &WGet{
		manager: manager,
		BaseURL: url,
	}
}

// Parse ...
// паттерн транспортный уровень ->
//		клиент делает вызов
//		получает респонс и ошибку
//		wget инкапсулирует Do клиента и отдает пользователю ответ
func (w *WGet) Parse() (*http.Response, error) {
	_, err := w.manager.Build()
	if err != nil {
		return nil, errManagerBuild
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	res, err := client.Get(w.BaseURL)
	if err != nil {
		return nil, errHTTPGetBad
	}

	return res, nil
}
