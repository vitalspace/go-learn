package main

import "fmt"

// Module represents a module with a private message
type Module struct {
	private string
}

// sayPrivateInternal prints the module's private message
func (m Module) sayPrivateInternal() {
	fmt.Println(m.private)
}

// sayPublic calls sayPrivateInternal to print the module's private message
func (m Module) SayPublic() {
	m.sayPrivateInternal()
}

func main() {
	// Create a new module with a private message
	m := Module{private: "Hello world this is a private message"}

	// Call the module's sayPublic method to print the private message
	m.SayPublic()
}
