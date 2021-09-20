package mancut

import (
	"reflect"
	"testing"
)

func TestHandlers_FieldsHandler(t *testing.T) {
	mc := &ManCut{}
	fieldsHandler := &fieldsHandler{}

	testCases := []struct {
		name          string
		delimeter     string
		onlySeparated bool
		fields        []int
		data          []string
		want          []string
	}{
		{
			name:          "regular with space separator",
			delimeter:     " ",
			onlySeparated: false,
			fields:        []int{0, 1, 2},
			data: []string{
				`this is bob`,
				`and he has things`,
			},
			want: []string{
				`this is bob`,
				`and he has`,
			},
		},
		{
			name:          "regular with asterisc separator",
			delimeter:     "*",
			onlySeparated: false,
			fields:        []int{0, 1},
			data: []string{
				`this*is*bob`,
				`and*he*has*things`,
			},
			want: []string{
				`this*is`,
				`and*he`,
			},
		},
		{
			name:          "regular with only separated",
			delimeter:     "#",
			onlySeparated: true,
			fields:        []int{0, 1},
			data: []string{
				`this#is#bob`,
				`and*he*has*things`,
			},
			want: []string{
				`this#is`,
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			mc.result = []string{}
			mc.delimeter = test.delimeter
			mc.onlySeparated = test.onlySeparated
			mc.options.fields = test.fields
			mc.data = test.data

			mc.isFieldsHandlerDone = false

			fieldsHandler.execute(mc)

			if !reflect.DeepEqual(mc.result, test.want) {
				t.Errorf("got %v want %v", mc.result, test.want)
			}
		})
	}

}
