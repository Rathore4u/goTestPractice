package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-resty/resty"
	"github.com/stretchr/testify/assert"
)

type LocationResponse struct {
	Country string `json:"country"`
}

func TestGetStatusCodeShouldEqual200(t *testing.T) {

	client := resty.New()

	resp, _ := client.R().Get("http://api.zippopotam.us/us/90210")
	myResponse := LocationResponse{}
	err := json.Unmarshal(resp.Body(), &myResponse)

	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, "United States", myResponse.Country)
}
