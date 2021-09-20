package mansort

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type finisher struct {
}

func (f *finisher) execute(m *ManSort) {
}

func (f *finisher) setNext(next handler) {
}

func TestMansort_KeyColumnSorterHandler(t *testing.T) {

	ms := &ManSort{
		options: Options{
			keyColumnNum: -1,
		},
		data: []string{"bob fisher", "alex messar", "gordon friman"},
	}
	mockFinish := &finisher{}
	columnSort := &keyColumnSorter{}
	columnSort.setNext(mockFinish)

	testCases := []struct {
		name     string
		colNum   int
		expected []string
	}{
		{
			name:     "sort by 0 column",
			colNum:   0,
			expected: []string{"alex messar", "bob fisher", "gordon friman"},
		},
		{
			name:     "test deafult",
			colNum:   -1,
			expected: []string{"alex messar", "bob fisher", "gordon friman"},
		},
		{
			name:     "sort by 1 column",
			colNum:   1,
			expected: []string{"bob fisher", "gordon friman", "alex messar"},
		},
		{
			name:     "sort by not existing column",
			colNum:   10,
			expected: []string{"alex messar", "bob fisher", "gordon friman"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			ms.keyColumnSortDone = false
			ms.options.keyColumnNum = test.colNum
			columnSort.execute(ms)
			assert.Equal(t, ms.data, test.expected)
		})
	}

}
func TestMansort_ReverseDataHandler(t *testing.T) {
	ms := &ManSort{
		options: Options{
			reverseNeeded: true,
		},
		data: []string{"bob fisher", "alex messar", "gordon friman"},
	}
	mockFinish := &finisher{}
	reverseSorter := &reverseDataSorter{}
	reverseSorter.setNext(mockFinish)

	testCases := []struct {
		name     string
		expected []string
	}{
		{
			name:     "regular reversing",
			expected: []string{"gordon friman", "alex messar", "bob fisher"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			reverseSorter.execute(ms)
			assert.Equal(t, ms.data, test.expected)
		})
	}

}
func TestMansort_UniqueSanitizerHandler(t *testing.T) {
	ms := &ManSort{
		options: Options{
			onlyUnique: true,
		},
		data: []string{"bob fisher", "alex messar", "gordon friman", "bob fisher", "alex messar"},
	}
	mockFinish := &finisher{}
	uniqueHandler := &uniqueSanitizer{}
	uniqueHandler.setNext(mockFinish)

	testCases := []struct {
		name     string
		expected []string
	}{
		{
			name:     "regular sanitizing",
			expected: []string{"bob fisher", "alex messar", "gordon friman"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			uniqueHandler.execute(ms)
			assert.Equal(t, ms.data, test.expected)
		})
	}
}
func TestMansort_TailsCheckerHandler(t *testing.T) {
	ms := &ManSort{
		options: Options{
			ignoreTails: true,
		},
	}
	mockFinish := &finisher{}
	tailsHandler := &tailsChecker{}
	tailsHandler.setNext(mockFinish)

	testCases := []struct {
		name     string
		data     []string
		expected []string
	}{
		{
			name:     "with spaces",
			data:     []string{"  a", " b ", "c  "},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "without spaces",
			data:     []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			ms.ignoreTailsDone = false
			ms.data = test.data

			tailsHandler.execute(ms)
			if !reflect.DeepEqual(ms.data, test.expected) {
				t.Errorf("got %v want %v", ms.data, test.expected)
			}
		})
	}

}
func TestMansort_NumColSorterHandler(t *testing.T) {
	ms := &ManSort{}
	mockFinish := &finisher{}
	numColHandler := &numColSorter{}
	numColHandler.setNext(mockFinish)

	testCases := []struct {
		name     string
		colNum   int
		data     []string
		expected []string
	}{
		{
			name:     "all samples valid",
			colNum:   1,
			data:     []string{"a 5", "b 3", "c 1"},
			expected: []string{"c 1", "b 3", "a 5"},
		},
		{
			name:     "invalid sort",
			colNum:   0,
			data:     []string{"a 5", "b 3", "c 1"},
			expected: []string{"a 5", "b 3", "c 1"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			ms.options.numColSort = test.colNum
			ms.numColSortDone = false
			ms.data = test.data

			numColHandler.execute(ms)
			if !reflect.DeepEqual(ms.data, test.expected) {
				t.Errorf("got %v want %v", ms.data, test.expected)
			}
		})
	}
}
func TestMansort_MonthColSorterHandler(t *testing.T) {
	ms := &ManSort{}
	mockFinish := &finisher{}
	monthColHandler := &monthColSorter{}
	monthColHandler.setNext(mockFinish)

	testCases := []struct {
		name     string
		monthCol int
		data     []string
		expected []string
	}{
		{
			name:     "valid month col",
			monthCol: 1,
			data:     []string{"Вася авг", "Маша янв", "Тест дек"},
			expected: []string{"Маша янв", "Вася авг", "Тест дек"},
		},
		{
			name:     "invalid month col",
			monthCol: 0,
			data:     []string{"Вася авг", "Маша янв", "Тест дек"},
			expected: []string{"Вася авг", "Маша янв", "Тест дек"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			ms.options.monthColNum = test.monthCol
			ms.monthColSortDone = false
			ms.data = test.data

			monthColHandler.execute(ms)

			if !reflect.DeepEqual(ms.data, test.expected) {
				t.Errorf("got %v want %v", ms.data, test.expected)
			}
		})
	}
}
