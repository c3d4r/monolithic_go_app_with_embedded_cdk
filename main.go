package main

import (
	"fmt"
	"os"

	"github.com/c3d4r/monolithic_go_app_with_embedded_cdk/infra"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "synth" {
		infra.Synthesize()
		return
	}

	// Normal application logic
	fmt.Println("app running")
	fmt.Println("Usage: pass 'synth' subcommand to generate CDK cloud assembly")
}
