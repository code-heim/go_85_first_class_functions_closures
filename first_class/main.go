package main

import "fmt"

func greet(name string, formatter func(string) string) {
	fmt.Println(formatter(name))
}

func formalGreeting(name string) string {
	return "Hello, Mr./Ms. " + name
}

func main() {

	casualGreeting := func(name string) string {
		return "Hey " + name + "!"
	}

	greet("Codeheim", formalGreeting)
	greet("Codeheim", casualGreeting)
}
