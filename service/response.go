package service

import "fmt"

func response(msg string, isError bool) {
	if isError {
		fmt.Printf("yatt[error]: %s\n", msg)
	} else {
		fmt.Printf("yatt: %s\n", msg)
	}
}
