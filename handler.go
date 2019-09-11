package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jhillyerd/enmime"
	"go.uber.org/zap"
	"net"
	"net/http"
)

func mailHandler(origin net.Addr, from string, to []string, data []byte) {
	r :=bytes.NewReader(data)
	envelope, err := enmime.ReadEnvelope(r)
	if err != nil {
		fmt.Print(err)
		return
	}
	msg:= Message{}
	msg.Text =envelope.Text
	msg.Html = envelope.HTML
	msg.Subject = envelope.Root.Header.Get("Subject")
	msg.FromEmail = from
	msg.Mailclass = cfg.MailClass
	for _, addr := range to {
		msg.To = append(msg.To,Recipient{Email:addr})
	}
	gamsg := GAMessage{Message:msg,Username:cfg.User,Password:cfg.Password}
	jsonStr,err:=json.Marshal(gamsg)
	fmt.Println("Json string::", string(jsonStr))
	req, err := http.NewRequest("POST", cfg.Url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Failed to send http request",
			zap.String("err", err.Error()),
		)
		return
	}
	if resp.StatusCode != 200 {
		logger.Error("Failed to send http request",
			zap.String("err", err.Error()),
			zap.Int("status",resp.StatusCode),
		)
		return
	}
	defer resp.Body.Close()
	var gr  GAResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		logger.Error("Failed to decode json response",
			zap.String("err", err.Error()),
		)
		return
	}
	if gr.Error != "" {
		logger.Error("Failed to send email",
			zap.String("err", err.Error()),
		)
	}
}