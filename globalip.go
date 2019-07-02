package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getGlobalIp() string {
	inetUrl := "http://inet-ip.info/ip"

	body, err := httpRequestBody(inetUrl)
	if err == nil {
		return body
	}
	log.Println(err)

	httpbinUrl := "http://httpbin.org/ip"
	body, err = httpRequestBody(httpbinUrl)

	globalIps := &struct{ Origin string }{}
	err = json.Unmarshal([]byte(body), &globalIps)

	if err != nil {
		panic(err)
	}
	return strings.Split(globalIps.Origin, ",")[0]
}

func httpRequestBody(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
