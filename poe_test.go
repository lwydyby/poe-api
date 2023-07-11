package poe_api

import (
	"fmt"
	"testing"
	"time"
)

func TestSendMessage(t *testing.T) {
	c := NewClient("", nil)
	res, err := c.SendMessage("ChatGPT", "", true, 30*time.Second)
	if err != nil {
		panic(err)
	}
	for r := range res {
		fmt.Println(r)
	}

}
