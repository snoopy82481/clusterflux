package main

import (
	"log"
	"os"

	"github.com/snoopy82481/clusterflux/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
