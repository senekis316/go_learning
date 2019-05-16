package panic_recover_test

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestPanicRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
}

func TestPanicExit(t *testing.T) {
	defer func() {
		fmt.Println("finally")
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
}

func TestOSExit(t *testing.T) {
	defer func() {
		fmt.Println("finally")
	}()
	fmt.Println("Start")
	os.Exit(-1)
}
