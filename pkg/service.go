package pkg

import (
	"fmt"
)

func Hello(name string) string {
	if name == "" {
		return fmt.Sprintf("What is your name ?")
	} else {
		return fmt.Sprintf("Hello %s", name)
	}
}

func Sum(a, b int) int {
	return a + b
}
