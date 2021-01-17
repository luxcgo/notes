package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	fmt.Println("hello")
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	if 1 != 2 {
		// return nil
		return errors.New("random error")
	}
	return nil
}
