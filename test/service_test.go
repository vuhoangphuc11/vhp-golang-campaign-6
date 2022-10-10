package test

import (
	"github.com/vuhoangphuc11/vhp-golang-campaign-6/pkg"
	"testing"
)

func Test_Hello_Pass(t *testing.T) {
	result := pkg.Hello("Gopher")

	if result != "Hello Gopher" {
		t.Errorf("Output expect Hello Gopher instead of %v", result)
	}
}

func Test_Hello_Empty(t *testing.T) {
	emptyNameResult := pkg.Hello("")

	if emptyNameResult != "What is your name ?" {
		t.Errorf("Output expect What is your name ? instead of %v", emptyNameResult)
	}
}

func Test_Sum_Pass(t *testing.T) {
	if pkg.Sum(5, 10) != 15 {
		t.Errorf("Result is incorrect")
	}
}
