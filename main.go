package main

import (
	"fmt"

	"github.com/noxterr/terraform-discordbot/src"
)

func main() {
	token := src.Initialize()
	fmt.Println(token)
}
