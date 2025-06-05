package main

import (
	"errors"
	"fmt"
)

func main() {
	err := thisfunc()
	fmt.Println(err)
}

func thisfunc() error {
	var err error

	if err == nil {
		if err == nil {

			if err == nil {
				err = errors.New("test2")
				return err

			}
		}
	}
	return err
}
