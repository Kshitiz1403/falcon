package main

import (
	"fmt"

	"github.com/Kshitiz1403/falcon/falcon"
)

func main() {
	// Use v1
	messageV1 := falcon.SayHello()
	fmt.Println("V1:", messageV1)

	// Use v2
	// messageV2 := v2.SayHello()
	// fmt.Println("V2:", messageV2)
}
