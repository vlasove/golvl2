package managers

import "net/http"

// Manager ...
type Manager interface {
	Build() (string, error)
	WriteResponse(*http.Response) (int64, error)
}
