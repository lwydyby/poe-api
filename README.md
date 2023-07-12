# Golang Poe API
The golang version of https://github.com/ading2210/poe-api, used for the golang project to call poe api


The latest commit that is currently compatible is https://github.com/ading2210/poe-api/tree/7470d07b989af596293f24d0bcfc6315161505e0

# Notice
As a single person, it can be challenging to keep up with the frequent API definition changes made by Poe in the ading2210/poe-api project. Therefore, 
welcome those with the necessary skills to contribute by submitting modifications directly via MR, rather than forking the code. 
This will help streamline the process and ensure that updates are efficiently implemented.

# Instructions

## install

```bash
go get github.com/lwydyby/poe-api
```

## use

```golang

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
	// 等待全部返回后 直接返回全文
	fmt.Println(poe_api.GetFinalResponse(res))
	res, err = c.SendMessage("ChatGPT", "channel是并发安全的吗", false, 30*time.Second)
	if err != nil {
		panic(err)
	}
	// 流式返回 每次返回新增的数据
	for m := range poe_api.GetTextStream(res) {
		fmt.Println(m)
	}
	// output: 
}
```