package main

import (
	"fmt"
	"os"

	"github.com/brevis-network/brevis-sdk/examples/dummy"
	"github.com/brevis-network/brevis-sdk/sdk/prover"
)

func main() {
	fmt.Println("Hello, World!")

	// app, err := sdk.NewBrevisApp()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(app)

	proverService, err := prover.NewService(dummy.DefaultAppCircuit(), prover.ServiceConfig{
		SetupDir: "$HOME/code/axal/prover-axal/circuitOut",
		SrsDir:   "$HOME/code/axal/prover-axal/kzgsrs",
	})
	if err != nil {
		// fmt.Println(err)
		fmt.Println("Error creating prover service")
		os.Exit(1)
	}

	proverService.Serve("localhost", 33247)
}
