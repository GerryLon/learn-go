package assert

import (
	"fmt"
)

func Nil(a interface{}) {
	if a != nil {
		panic(fmt.Sprintf("%v should be nil", a))
	}
}
