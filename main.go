package main

import (
	"fmt"
	"mytest/factories"
)

func main() {
	a := factories.CreateAuthServiceInterface()
	f := a.GetFacebookAuthManager()
	l := a.GetLocalAuthManager()

	fmt.Println(f.Login("test"))
	fmt.Println(l.Login("test", "test"))
}
