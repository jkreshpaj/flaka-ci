package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Notification struct {
	EndpointURL string
	Message     string
	Type        string
}

type Message struct {
	Text string `json:"text"`
}

func (n *Notification) Send(message []byte) error {
	res, err := http.Post(n.EndpointURL, "application/json", bytes.NewBuffer(message))
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func (n *Notification) Parse() ([]byte, error) {
	m := Message{
		Text: n.Message,
	}

	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
