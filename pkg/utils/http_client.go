package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/xiao-cold/deeplx-load-balancer/internal/models"
)

var client = &http.Client{
	Timeout: time.Second * 10,
}

func ForwardRequest(server *models.Server, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", server.URL+"/translate", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
