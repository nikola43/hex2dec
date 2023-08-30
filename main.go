package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// get arguments
	args := os.Args

	// validate number of arguments
	if len(args) < 2 {
		fmt.Println("no arguments provided")
		return
	}

	if len(args) > 2 {
		fmt.Println("too many arguments")
		return
	}

	// check if argument is valid hex number
	hex := args[1]
	_, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		fmt.Println("invalid hex number")
		return
	}

	// convert hex to decimal
	dec, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		fmt.Println("invalid hex number")
		return
	}
	fmt.Println(dec)

}
