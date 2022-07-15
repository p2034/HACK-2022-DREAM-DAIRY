package apirequest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type server_config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

var Auth_config server_config

func ParseConfig() {
	jsonFile, err := os.Open("configs/auth_server.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Auth_config)
}
