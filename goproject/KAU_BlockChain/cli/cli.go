package cli

import (
	"flag"
	"fmt"
	"goproject/KAU_BlockChain/explorer"
	"goproject/KAU_BlockChain/rest"
	"os"
)

// cobra cli 명령어 만들기 좋음 추후 배워보기

func usage() {
	fmt.Printf(("Welcome to KAU-Coin!\n\n"))
	fmt.Printf(("Please use the following flags:\n\n"))
	fmt.Printf(("-port=4000:		Set the PORT \n"))
	fmt.Printf(("rest:			Start the REST API (recommended)\n\n"))
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
		//start rest api
	case "html":
		explorer.Start(*port)
		//start html explorer
	default:
		usage()
	}

	fmt.Println(*port, *mode)
}
