package example

import (
	"fmt"
	"time"

	"github.com/lwydyby/poe-api"
)

func ExampleSendMessage() {
	c := poe_api.NewClient("", nil)
	res, err := c.SendMessage("ChatGPT", "一句话描述golang的channel", true, 30*time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(poe_api.GetFinalResponse(res))
	res, err = c.SendMessage("ChatGPT", "channel是并发安全的吗", false, 30*time.Second)
	if err != nil {
		panic(err)
	}
	for m := range poe_api.GetTextStream(res) {
		fmt.Println(m)
	}
	// output:
}
