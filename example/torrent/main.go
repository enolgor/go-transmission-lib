package main

import (
	"encoding/json"
	"fmt"

	"github.com/enolgor/go-transmission-lib/api"
)

func main() {
	api, err := api.NewApi("http://192.168.1.2:9091/transmission/rpc")
	if err != nil {
		panic(err)
	}
	//resp, err := api.GetTorrents(1, 2)
	//resp, err := api.GetSession()
	resp, err := api.SessionStats()
	if err != nil {
		panic(err)
	}
	s, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(s))
}
