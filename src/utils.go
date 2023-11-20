package main

import "math/rand"

// RandomBool returns a random boolean value.
func RandomBool() bool {
	return rand.Intn(2) == 0
}
