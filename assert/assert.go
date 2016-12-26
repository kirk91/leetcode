package assert

import (
	"log"
	"reflect"
)

func Equal(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}
	return reflect.DeepEqual(expected, actual)
}

func MustEqual(expected, actual interface{}) {
	if !Equal(expected, actual) {
		log.Fatalf("expected: %v, but received: %v", expected, actual)
	}
}
