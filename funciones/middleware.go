package funciones

import (
	"fmt"
	"time"
)

type MyFunction func(string)

func MiddlewareLog(f MyFunction) MyFunction {
	return func(name string) {
		fmt.Println(time.Now().Format("15:04:05"))
		f(name)
	}
}
