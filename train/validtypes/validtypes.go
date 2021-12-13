package main

import (
	"errors"
	"fmt"
)

type OrderStatus string

type Order struct {
	Status OrderStatus
}

func (s OrderStatus) Validate() error {
	switch s {
	case "sell":
		fallthrough
	case "in_progress":
		return nil
	default:
		return errors.New("invalid status")
	}
}

func main() {

	for _, s := range []string{"asd", "sell"} {

		var s OrderStatus = OrderStatus(s)

		err := s.Validate()
		if err != nil {

			fmt.Println(err.Error())
		}
	}

}
