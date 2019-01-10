package cmd

import "log"

//HandleError logs error and sends to slack
func HandleError(text string, errLog interface{}) {
	if server.NotificationURL != "" {
		ntf := Notification{
			EndpointURL: server.NotificationURL,
			Title:       text,
			Log:         "```" + errLog.(string) + "```",
			Type:        "error",
		}
		if err := ntf.Send(); err != nil {
			log.Println(ColorizeError(err))
		}
	}
	log.Println(ColorizeError(text))
	log.Println(ColorizeError(errLog.(string)))
}

//ColorizeError make error logs red
func ColorizeError(err interface{}) string {
	return "\u001b[1m\u001b[31m" + err.(string) + "\u001b[0m"
}
