package main

import (
	"fmt"
)

func main() {
	client := NewClient("", nil)
	fmt.Println(client.GetBots())
}
