package main

import (
	"fmt"
	"github.com/pkg/errors"
	error1 "golang/Week2"
)

func main() {

	name, err := error1.SelectRowById(2)
	if err != nil {
		fmt.Printf("%+v\n", errors.Cause(err))
	}
	fmt.Printf("name = %v", name)

}
