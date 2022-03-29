package main

import (
	"goproject/KAU_BlockChain/cli"
	"goproject/KAU_BlockChain/db"
)

func main() {

	defer db.Close()
	cli.Start()
}
