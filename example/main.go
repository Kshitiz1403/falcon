package main

import (
	"fmt"

	v1 "github.com/Kshitiz1403/falcon/v1"
	v2 "github.com/Kshitiz1403/falcon/v2"
)

func main() {
	// Use v1
	messageV1 := v1.SayHello()
	fmt.Println("V1:", messageV1)

	// Use v2
	messageV2 := v2.SayHello()
	fmt.Println("V2:", messageV2)
}
