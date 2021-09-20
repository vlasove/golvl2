package shell

// Shell ...
type Shell struct {
	Args     []string
	Executor Executor
}

// New ...
func New() *Shell {
	return &Shell{}
}

// SetArgs ...
func (s *Shell) SetArgs(args []string) {
	s.Args = args
}

// SetExecutor ...
func (s *Shell) SetExecutor(e Executor) {
	s.Executor = e
}

// Start ...
func (s *Shell) Start() (string, error) {
	return s.Executor.Execute(s)
}
