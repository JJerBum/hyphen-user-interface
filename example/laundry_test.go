package example

import (
	"laundry"
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	laundry := laundry.New("../basket")
	log.Fatal(laundry.Run(":8080"))
}
