package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

func main() {
	server.StartServer(New, Version, Priority)
}

var Version = "0.2"
var Priority = 1
var (
	AuthURL        = os.Getenv("AUTH_URL")
	AuthPORT       = os.Getenv("AUTH_PORT")
	ExpirationTime = os.Getenv("EXPIRATION_TIME")
)

type Config struct {
	HeaderKey string
}

func New() interface{} {
	return &Config{}
}

func (conf Config) Access(kong *pdk.PDK) {

	kong.Log.Info("auth url = ", AuthURL, " auth port = ", AuthPORT, " expiration time = ", ExpirationTime)
	path, _ := kong.Request.GetPath()
	kong.Log.Info("path ", path)
	if strings.Contains(path, "signup") || strings.Contains(path, "login") {
		kong.Log.Info(fmt.Sprintf("Header %s valid :)", conf.HeaderKey))
		kong.ServiceRequest.SetHeader("is_valid", "true")
		return
	}

	header, err := kong.Request.GetHeader(conf.HeaderKey)
	if err != nil {
		kong.Log.Err(fmt.Sprintf("Error reading '%s' header: %s", conf.HeaderKey, err.Error()))
	}

	if header == "" {
		kong.Log.Err(fmt.Sprintf("Header %s is empty", conf.HeaderKey))
		kong.Response.Exit(401, []byte{}, nil)
		return
	}

	res, err := resolveRequest(kong, header)

	if err != nil {
		kong.Log.Err(err)
		kong.Log.Err(fmt.Sprintf("errorr 1 %s ", err.Error()))
		pl := map[string]string{
			"msg": "not authorized",
		}
		byteData, _ := json.Marshal(pl)
		if err != nil {
			kong.Response.Exit(401, []byte{}, nil)
		}
		kong.Response.Exit(401, byteData, nil)
		return
	}
	kong.Log.Info("id ", res.ID)

	if res.ID == 0 {
		kong.ServiceRequest.SetHeader("is_valid", "false")
		kong.ServiceRequest.SetHeader("user_id", strconv.Itoa(int(res.ID)))
		pl := map[string]string{
			"msg": "not authorized",
		}
		byteData, _ := json.Marshal(pl)
		if err != nil {
			kong.Response.Exit(401, []byte{}, nil)
		}
		kong.Response.Exit(401, byteData, nil)
		kong.Response.Exit(401, []byte{}, nil)

		return
	}

	kong.Log.Info(fmt.Sprintf("Header %s valid :)", conf.HeaderKey))
	kong.ServiceRequest.SetHeader("is_valid", "true")
	kong.ServiceRequest.SetHeader("user_id", strconv.Itoa(int(res.ID)))

}

type RequestResponse struct {
	ID      int32  `json:"id"`
	Message string `json:"message"`
}

func resolveRequest(kong *pdk.PDK, token string) (*RequestResponse, error) {
	kong.Log.Err("1")

	url := "http://" + AuthURL + ":" + AuthPORT + "/auth/student/get"
	kong.Log.Err(url)

	type Body struct {
		Token          string `json:"token"`
		ExpirationTime int32  `json:"expiration_time"`
	}

	method := "GET"
	exprTime, _ := strconv.Atoi(ExpirationTime)
	pl := &Body{
		Token:          token,
		ExpirationTime: int32(exprTime),
	}
	kong.Log.Err("2")

	byteData, err := json.Marshal(pl)
	if err != nil {
		kong.Log.Err(err)
		return nil, err
	}
	kong.Log.Err("3")

	payload := strings.NewReader(string(byteData))
	// payload := strings.NewReader(`{
	// 	"token":"b5fbc438-d306-466d-af8f-f90f3521589b",
	// 	"expiration_time":10
	// }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	kong.Log.Err("4")
	if err != nil {
		kong.Log.Err(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	kong.Log.Err("5")
	res, err := client.Do(req)
	kong.Log.Err("9")

	if err != nil {
		kong.Log.Err("10", err.Error(), "dfgdfg")

		kong.Log.Err(fmt.Sprintf("error 1 %s ", err.Error()))
		return nil, err
	}
	defer res.Body.Close()
	kong.Log.Err("6")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		kong.Log.Err(fmt.Sprintf("error 2 %s ", err.Error()))

		kong.Log.Err(err)
		return nil, err
	}
	kong.Log.Err(string(body))
	kong.Log.Err(fmt.Sprintf("body %s ", string(body)))

	kong.Log.Err("7")
	response := &RequestResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		kong.Log.Err(fmt.Sprintf("error 3 %s ", err.Error()))

		return nil, err
	}
	kong.Log.Err("8")
	kong.Log.Err(string(body))

	return response, nil
}
