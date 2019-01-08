package cmd

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var colorMap = map[string]string{
	"success": "#2eb886",
	"error":   "#b82e2e",
	"warn":    "#edad17",
	"info":    "#0075a8",
}

//Notification type contains webhook url message type of notifications
type Notification struct {
	EndpointURL string
	Title       string
	Log         string
	Type        string
}

//Message use to marshal json
type Message struct {
	Attachments Attachments `json:"attachments"`
}

//Attachments array of message settings
type Attachments []Settings

//Settings message options
type Settings struct {
	Title string `json:"title"`
	Color string `json:"color"`
	Log   string `json:"text"`
}

//Send post request to webhook url
func (n *Notification) Send() error {
	message, err := n.Parse()
	if err != nil {
		return err
	}
	_, err = http.Post(n.EndpointURL, "application/json", bytes.NewBuffer(message))
	if err != nil {
		return err
	}
	return nil
}

//Parse string message to json
func (n *Notification) Parse() ([]byte, error) {
	m := Message{}
	attach := Attachments{}
	settings := Settings{
		Title: n.Title,
		Color: colorMap[n.Type],
		Log:   n.Log,
	}
	attach = append(attach, settings)
	m.Attachments = attach

	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
