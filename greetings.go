package greetings

import "fmt"

func Hello(Name string) string {
	message := fmt.Sprintf("!Hola, %v!, Bienvenido", Name)
	return message
}
