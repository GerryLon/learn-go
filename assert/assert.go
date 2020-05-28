package assert

import (
	"fmt"
	"reflect"
)

func Nil(a interface{}) {
	if a != nil {
		panic(fmt.Sprintf("%v should be nil", a))
	}
}

func Equal(a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		panic(fmt.Sprintf("%v and %v are not equal", a, b))
	}
}
