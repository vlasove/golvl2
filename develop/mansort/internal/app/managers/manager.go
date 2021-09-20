package managers

// Manager ...
type Manager interface {
	Read() ([]string, error)
	Write([]string) error
}
