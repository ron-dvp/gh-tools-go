package tools

import (
	"bytes"
	"ghtools/srv"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type RequestInfo struct {
	ReqType string
	Path    string
	Body    string
}

type Response struct {
	Status       string
	ResponseInfo []byte
}

var baseDomain = "https://api.github.com"

func ClientCred() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[-] Missing env data file")
	}

	val := srv.EnvVal("GITHUB_TOKEN")
	if val == "" {
		println("[-] Error : missing github token")
		os.Exit(1)
	}

	return "Bearer " + val
}

func SendRequest(info RequestInfo) (Response, error) {

	c := &http.Client{}
	result := Response{}

	body := []byte("")
	if info.Body != "" {
		body = []byte(info.Body)
	}
	req, err := http.NewRequest(info.ReqType, baseDomain+info.Path, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("[-] Error: %s", err)
		return result, err
	}
	req.Header.Add("Authorization", ClientCred())
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalf("[-] Error: %s", err)
		return result, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("[-] Error: %s", err)
		return result, err
	}

	result.Status = resp.Status
	result.ResponseInfo = []byte(data)
	return result, nil
}
