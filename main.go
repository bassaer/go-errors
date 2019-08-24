package main

import (
	"log"

	"github.com/pkg/errors"
)

func work() error {
	return workA()
}

func workA() error {
	if err := workB(); err != nil {
		return errors.Wrap(err, "error from workA")
	}
	return nil
}

func workB() error {
	if err := workC(); err != nil {
		return errors.Wrap(err, "error from workB")
	}
	return nil
}

func workC() error {
	return errors.New("error from workC")
}

func main() {
	if err := work(); err != nil {
		log.Fatal(err)
	}
}
