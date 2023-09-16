package main

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestMain(t *testing.T) {
	text, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Testing %v", text)
}
