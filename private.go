package private_go_module

import "fmt"

func TestPrivateGoModule() string {
	return "Test"
}

func PrintTestGoModule() {
	fmt.Println(TestPrivateGoModule())
}
