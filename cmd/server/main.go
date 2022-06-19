package main

import "fmt"

// Run - is going to be responsible for the init and startup of our application.
func Run() error {
	fmt.Println("starting up our application")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
