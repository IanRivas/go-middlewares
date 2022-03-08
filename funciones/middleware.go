package funciones

import (
	"fmt"
	"time"
)

func MiddlewareLog(f func(string)) func(string) {
	return func(name string) {
		fmt.Println(time.Now())
		f(name)
	}
}
