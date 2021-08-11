package slackhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func NewMessage(url string, channel string, user string) Message {

	return Message{WebHookUrl: url, Channel: channel, Username: user}

}

type Message struct {
	Username   string `json:"username"`
	WebHookUrl string `json:"-"`
	Channel    string `json:"channel"`
	Text       string `json:"text"`
	Payload    []byte `json:"-"`
}

func (a *Message) Send(message []byte) error {
	a.Text = string(message)
	a.Payload = message

	fmt.Printf("Message to: '%v' with a message of: '%v'\n", a.WebHookUrl, string(a.Text))

	reqBody, err := json.Marshal(a)

	if err != nil {
		return err
	}

	resp, err := http.Post(a.WebHookUrl, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errmsg := fmt.Sprintf("Post to '%s' - Expected StatusCode '200', got: '%d'", a.WebHookUrl, resp.StatusCode)
		return errors.New(errmsg)
	}

	return nil
}
