package main

import (
	"github.com/IanRivas/go-middlewares/funciones"
)

func execute(name string, f funciones.MyFunction) {
	f(name)
}

func main() {
	name := "Comunidad Edteam"
	execute(name, funciones.MiddlewareLog(funciones.Saludar))
	execute(name, funciones.MiddlewareLog(funciones.Despedirse))
}
