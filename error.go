package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func runc() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func run() error {
	return fmt.Errorf("error")
}

func main() {
	if err := runc(); err != nil {
		fmt.Println(err)
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}

}
