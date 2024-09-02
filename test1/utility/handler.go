package utility

import "fmt"

type str string
type App struct {
	first string
	last  string
}

func ComputerClient() {

	fmt.Println("Hello, World!")
}
func (a App) ComputerServer() {
	fmt.Println(a.first)
}
