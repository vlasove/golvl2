package mangrep

import (
	"log"

	"github.com/vlasove/materials/tasks_2/utils/mangrep/internal/app/managers"
)

// ManGrep ...
type ManGrep struct {
	numLinePrintedHandlerDone bool
	fixedTemplateHandlerDone  bool
	inversionHandlerDone      bool
	regularHandlerDone        bool
	counterHandlerDone        bool
	amountAfterHandlerDone    bool
	amountBeforeHandlerDone   bool
	amountSidesHandlerDone    bool

	inputManager  managers.Manager
	outputManager managers.Manager
	options       Options
	data          []string
	result        []string
	template      string
}

// New ...
func New(input, output managers.Manager, template string) *ManGrep {
	return &ManGrep{
		inputManager:  input,
		outputManager: output,
		template:      template,
	}
}

// ApplyOptions ...
func (m *ManGrep) ApplyOptions(options ...Option) *ManGrep {
	opts := GetDefaultOptions()
	for _, opt := range options {
		if opt != nil {
			if err := opt(&opts); err != nil {
				log.Fatal(err)
			}
		}
	}
	m.options = opts
	return m
}

// Search ...
func (m *ManGrep) Search() error {
	data, err := m.inputManager.Read()
	if err != nil {
		return err
	}
	m.data = data
	beforeHandler := &amountBeforeHandler{}

	afterHander := &amountAfterHandler{}
	afterHander.setNext(beforeHandler)
	sidesHandler := &amountSidesHandler{}
	sidesHandler.setNext(afterHander)

	counterHandler := &counterHandler{}
	counterHandler.setNext(sidesHandler)

	linePrintedHandler := &numLinePrintedHandler{}
	linePrintedHandler.setNext(counterHandler)

	inverseHandler := &inversionHandler{}
	inverseHandler.setNext(linePrintedHandler)

	fixedHandler := &fixedTemplateHandler{}
	fixedHandler.setNext(inverseHandler)

	regularHandler := &regularHandler{}
	regularHandler.setNext(fixedHandler)

	regularHandler.execute(m)
	return nil
}

// OutputResult ...
func (m *ManGrep) OutputResult() error {
	if err := m.outputManager.Write(m.result); err != nil {
		return err
	}
	return nil
}
