// reference: https://medium.com/@BastianRob/implementing-reduce-in-go-4a3e6e3affc

package util

import (
	"errors"
	"reflect"
)

// Reducer Error Collection
var (
	ErrSourceNotArray = errors.New("source value is not an array")
	ErrReducerNil     = errors.New("reducer function cannot be nil")
	ErrReducerNotFunc = errors.New("reducer argument must be a function")
)

// Reduce an array of something into another thing
func Reduce(source, initialValue, reducer interface{}) (interface{}, error) {
	srcV := reflect.ValueOf(source)
	kind := srcV.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, ErrSourceNotArray
	}
	if reducer == nil {
		return nil, ErrReducerNil
	}
	rv := reflect.ValueOf(reducer)
	if rv.Kind() != reflect.Func {
		return nil, ErrReducerNotFunc
	}
	// copy initial value as accumulator, and get the reflection value
	accumulator := initialValue
	accV := reflect.ValueOf(accumulator)
	for i := 0; i < srcV.Len(); i++ {
		entry := srcV.Index(i)
		// call reducer via reflection
		reduceResults := rv.Call([]reflect.Value{
			accV,               // send accumulator value
			entry,              // send current source entry
			reflect.ValueOf(i), // send current loop index
		})
		accV = reduceResults[0]
	}
	return accV.Interface(), nil
}
