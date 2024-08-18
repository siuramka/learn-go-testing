package helloworld

import "fmt"

const (
	englishHello    = "Hello "
	spanishHello    = "Hola "
	lithuanianHello = "Labas "
)

const (
	english    = "english"
	spanish    = "spanish"
	lithuanian = "lithuanian"
)

func Hello(name string, language string) string {
	const defaultName = "World"
	if name == "" {
		name = defaultName
	}

	switch language {
	case spanish:
		return spanishHello + name
	case english:
		return englishHello + name
	case lithuanian:
		return lithuanianHello + name

	}

	panic("unknown language")
}

func main() {
	fmt.Println(Hello("", "english"))
}
