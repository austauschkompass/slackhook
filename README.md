# SlackHook

use slack-webhooks

# usage

`go get github.com/austauschkompass/slackhook`

main.go:
```go
package main

import (
	"fmt"

	"github.com/austauschkompass/slackhook"
)

func main() {

	url := "https://hooks.slack.com/services/xxx/yyy/zzzz"

	msg := []byte("here we go again <https://example.com|go>.")

  /* create new slackhook with  webhook-url, channel and username */

	slack := slackhook.NewMessage(url, "#service", "slackhook")

  /* pass msg as []byte */

	err := slack.Send(msg)

	if err != nil {
		fmt.Println(err.Error())
	}
}
```



