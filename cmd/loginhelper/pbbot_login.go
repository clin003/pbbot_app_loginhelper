package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"pbbot_app_loginhelper/pkg/dto"

	"google.golang.org/protobuf/proto"
)

func postPbbotCreate(serverUrl string, loginReq *dto.CreateBotReq) (retText string, err error) {
	v, _ := proto.Marshal(loginReq)
	payload := bytes.NewBuffer(v)
	apiUrl := fmt.Sprintf("%s/bot/create/v1/", serverUrl)
	req, err := http.NewRequest("POST", apiUrl, payload)
	if err != nil {
		return
	}
	req.Header.Add("content-type", "application/x-protobuf")
	client := http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[状态] %s %s", resp.Status, apiUrl)
		return
	}
	retText = string(body)

	return
}
