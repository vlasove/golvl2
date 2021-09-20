package mansort

// Options ...
type Options struct {
	keyColumnNum  int
	numColSort    int
	reverseNeeded bool
	onlyUnique    bool
	alreadySorted bool
	ignoreTails   bool
	monthColNum   int
}

// GetDefaultOptions ...
func GetDefaultOptions() Options {
	return Options{
		monthColNum:   -1,
		keyColumnNum:  -1,
		numColSort:    -1,
		reverseNeeded: false,
		onlyUnique:    false,
		alreadySorted: false,
		ignoreTails:   false,
	}
}

// DefaultOptions ...
var DefaultOptions = GetDefaultOptions()

// Option ...
type Option func(*Options) error

// KeyColumnNum ...
func KeyColumnNum(num int) Option {
	return func(o *Options) error {
		o.keyColumnNum = num
		return nil
	}
}

// ReverseNeeded ...
func ReverseNeeded(flag bool) Option {
	return func(o *Options) error {
		o.reverseNeeded = flag
		return nil
	}
}

// OnlyUnique ...
func OnlyUnique(flag bool) Option {
	return func(o *Options) error {
		o.onlyUnique = flag
		return nil
	}
}

// AlreadySorted ...
func AlreadySorted(flag bool) Option {
	return func(o *Options) error {
		o.alreadySorted = flag
		return nil
	}
}

// IgnoreTails ...
func IgnoreTails(flag bool) Option {
	return func(o *Options) error {
		o.ignoreTails = flag
		return nil
	}
}

// NumColSort ...
func NumColSort(val int) Option {
	return func(o *Options) error {
		o.numColSort = val
		return nil
	}
}

// MonthColSort ...
func MonthColSort(val int) Option {
	return func(o *Options) error {
		o.monthColNum = val
		return nil
	}
}
