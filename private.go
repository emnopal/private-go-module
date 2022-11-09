package private_go_module

import "fmt"

func TestPrivateGoModule(params string) string {
	return "Test" + params
}

func PrintTestGoModule() {
	fmt.Println(TestPrivateGoModule("Params"))
}
