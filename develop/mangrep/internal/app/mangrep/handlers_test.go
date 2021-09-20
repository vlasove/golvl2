package mangrep

import (
	"reflect"
	"testing"
)

type finisher struct {
}

func (f *finisher) execute(m *ManGrep) {
}

func (f *finisher) setNext(next handler) {
}

func TestMangrep_AmountBeforeHandler(t *testing.T) {
	mg := &ManGrep{
		data:   []string{"bob", "alex", "joshex"},
		result: []string{"alex"},
		options: Options{
			amountBefore: 1,
		},
	}
	mockHandler := &finisher{}
	amountHandler := &amountBeforeHandler{}
	amountHandler.setNext(mockHandler)

	wanted := []string{"bob", "alex"}

	amountHandler.execute(mg)

	if !reflect.DeepEqual(mg.result, wanted) {
		t.Errorf("got %v want %v", mg.result, wanted)
	}
}

func TestMangrep_AmountAfterHandler(t *testing.T) {
	mg := &ManGrep{
		data:   []string{"bob", "alex", "joshex"},
		result: []string{"alex"},
		options: Options{
			amountAfter: 1,
		},
	}
	mockHandler := &finisher{}
	amountHandler := &amountAfterHandler{}
	amountHandler.setNext(mockHandler)

	wanted := []string{"alex", "joshex"}

	amountHandler.execute(mg)

	if !reflect.DeepEqual(mg.result, wanted) {
		t.Errorf("got %v want %v", mg.result, wanted)
	}
}

func TestMangrep_NumLinePrintedHandler(t *testing.T) {
	mg := &ManGrep{
		data:   []string{"bob", "alex", "joshex"},
		result: []string{"bob", "alex", "joshex"},
		options: Options{
			isNumLinePrinted: true,
		},
	}
	mockHandler := &finisher{}
	numericHandler := &numLinePrintedHandler{}
	numericHandler.setNext(mockHandler)

	wanted := []string{"1 bob", "2 alex", "3 joshex"}

	numericHandler.execute(mg)

	if !reflect.DeepEqual(mg.result, wanted) {
		t.Errorf("got %v want %v", mg.result, wanted)
	}
}

func TestMangrep_FixedTemplateHandler(t *testing.T) {
	mg := &ManGrep{
		data: []string{"bob", "alex", "joshex"},
		options: Options{
			isFixedTemplate: true,
		},
	}
	mockHandler := &finisher{}
	fixedHandler := &fixedTemplateHandler{}
	fixedHandler.setNext(mockHandler)

	testCases := []struct {
		name     string
		template string
		want     []string
	}{
		{
			name:     "valid only one",
			template: `bo`,
			want:     []string{"bob"},
		},
		{
			name:     "valid double",
			template: `ex`,
			want:     []string{"alex", "joshex"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			mg.fixedTemplateHandlerDone = false
			mg.result = []string{}
			mg.template = test.template
			fixedHandler.execute(mg)
			if !reflect.DeepEqual(mg.result, test.want) {
				t.Errorf("got %v want %v", mg.result, test.want)
			}
		})
	}
}

func TestMangrep_InversionHandler(t *testing.T) {
	mg := &ManGrep{
		data: []string{"bob", "alex", "josh"},
		options: Options{
			isInversion: true,
		},
	}
	mockHandler := &finisher{}
	inversionHandler := &inversionHandler{}
	inversionHandler.setNext(mockHandler)

	testCases := []struct {
		name         string
		result       []string
		wantedResult []string
	}{
		{
			name:         "valid only one",
			result:       []string{"bob", "josh"},
			wantedResult: []string{"alex"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			mg.inversionHandlerDone = false
			mg.result = test.result
			inversionHandler.execute(mg)
			if !reflect.DeepEqual(mg.result, test.wantedResult) {
				t.Errorf("got %v want %v", mg.result, test.wantedResult)
			}
		})
	}
}

func TestMangrep_RegularHandler(t *testing.T) {
	mg := &ManGrep{
		data: []string{"bob", "alex", "josh"},
	}
	mockHandler := &finisher{}
	regularHandler := &regularHandler{}
	regularHandler.setNext(mockHandler)

	testCases := []struct {
		name     string
		template string
		want     []string
	}{
		{
			name:     "valid only one",
			template: `b.b`,
			want:     []string{"bob"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			mg.regularHandlerDone = false
			mg.template = test.template
			regularHandler.execute(mg)
			if !reflect.DeepEqual(mg.result, test.want) {
				t.Errorf("got %v want %v", mg.result, test.want)
			}
		})
	}
}

func TestMangrep_CounterHandler(t *testing.T) {
	mg := &ManGrep{
		options: Options{
			isCounter: true,
		},
		result: []string{"a", "b", "c"},
	}

	mockHandler := &finisher{}
	counterHandler := &counterHandler{}
	counterHandler.setNext(mockHandler)

	counterHandler.execute(mg)

	want := []string{"Finded lines: 3"}
	if !reflect.DeepEqual(mg.result, want) {
		t.Errorf("got %v want %v", mg.result, want)
	}
}

func TestMangrep_Difference(t *testing.T) {
	testCases := []struct {
		name     string
		lhs, rhs []string
		result   []string
	}{
		{
			name:   "not empty result",
			lhs:    []string{"a", "b", "c"},
			rhs:    []string{"a", "b", "c", "d", "e"},
			result: []string{"d", "e"},
		},
		{
			name:   "both equal input",
			lhs:    []string{"a", "b", "c"},
			rhs:    []string{"c", "b", "a"},
			result: []string{},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := difference(test.lhs, test.rhs)
			if len(got) != len(test.result) {
				t.Errorf("got %v want %v", got, test.result)
			}
			if len(test.result) != 0 {
				if !reflect.DeepEqual(got, test.result) {
					t.Errorf("got %v want %v", got, test.result)
				}
			}
		})
	}
}

func TestMangrep_IndexOfLast(t *testing.T) {
	testCases := []struct {
		name   string
		sample string
		slice  []string
		resID  int
	}{
		{
			name:   "total differences in slice",
			sample: "a",
			slice:  []string{"b", "a", "c"},
			resID:  1,
		},
		{
			name:   "has duplicates",
			sample: "a",
			slice:  []string{"b", "c", "a", "a", "a", "d"},
			resID:  4,
		},
		{
			name:   "negative case",
			sample: "a",
			slice:  []string{"b", "c", "d"},
			resID:  -1,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := indexOfLast(test.sample, test.slice)
			if got != test.resID {
				t.Errorf("got %d want %d", got, test.resID)
			}
		})
	}
}

func TestMangrep_IndexOfBefore(t *testing.T) {
	testCases := []struct {
		name   string
		sample string
		slice  []string
		resID  int
	}{
		{
			name:   "total differences in slice",
			sample: "a",
			slice:  []string{"b", "a", "c"},
			resID:  1,
		},
		{
			name:   "has duplicates",
			sample: "a",
			slice:  []string{"b", "c", "a", "a", "a", "d"},
			resID:  2,
		},
		{
			name:   "negative case",
			sample: "a",
			slice:  []string{"b", "c", "d"},
			resID:  -1,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := indexOfBefore(test.sample, test.slice)
			if got != test.resID {
				t.Errorf("got %d want %d", got, test.resID)
			}
		})
	}
}
