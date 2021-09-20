package mancut

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errUnparsedField      = errors.New("mancut: can not parse provided fields. use only integers separated by commas")
	errNegativeFieldValue = errors.New("mancut: field value must be positive")
	errEmptyFields        = errors.New("mancut: fields not provided")
)

// Options ...
type Options struct {
	fields    []int
	delimeter string
	separated bool
}

// GetDefaultOptions ...
func GetDefaultOptions() Options {
	return Options{
		fields:    []int{},
		delimeter: "\t",
		separated: false,
	}
}

// DefaultOptions ...
var DefaultOptions = GetDefaultOptions()

// Option ...
type Option func(*Options) error

// SetFieldsOption ...
func SetFieldsOption(fields string) Option {
	return func(o *Options) error {
		resultedFields := []int{}
		if len(fields) == 0 {
			return errEmptyFields
		}
		fieldsSlice := strings.Split(fields, ",")
		for _, f := range fieldsSlice {
			v, err := strconv.Atoi(f)
			if err != nil {
				return errUnparsedField
			}
			if v <= 0 {
				return errNegativeFieldValue
			}
			resultedFields = append(resultedFields, v-1)
		}
		o.fields = resultedFields
		return nil
	}
}

// SetDelimeterOption ...
func SetDelimeterOption(delim string) Option {
	return func(o *Options) error {
		o.delimeter = delim
		return nil
	}
}

// SetSeparatedOption ...
func SetSeparatedOption(flag bool) Option {
	return func(o *Options) error {
		o.separated = flag
		return nil
	}
}
