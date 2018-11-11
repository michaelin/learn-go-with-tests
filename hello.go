package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const danish = "Danish"
const helloPrefixEnglish = "Hello, "
const helloPrefixSpanish = "Hola, "
const helloPrefixFrench = "Bonjour, "
const helloPrefixDanish = "Hejsa, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return getGreetingPrefix(language) + name + "!"
}

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloPrefixSpanish
	case french:
		prefix = helloPrefixFrench
	case danish:
		prefix = helloPrefixDanish
	default:
		prefix = helloPrefixEnglish
	}

	return
}

func main()  {
	fmt.Println(Hello("world", ""))
}