package main

import (
	"fmt"
	"net/url"

	"github.com/enolgor/go-transmission-lib/transmission"
)

func main() {
	client := transmission.Client{URL: &url.URL{Scheme: "http", Host: "192.168.1.2:9091", Path: "/transmission/rpc"}}
	response := &transmission.Response{Arguments: &transmission.SessionStats{}}
	err := client.Do(transmission.NewRequest("session-stats"), response)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Arguments)
	fmt.Printf("%++v", *(response.Arguments.(*transmission.SessionStats)).CumulativeStats)
	//fmt.Printf("%.2fGB uploaded\n", float64(response.Arguments.(*transmission.SessionStats).CumulativeStats.UploadedBytes)/(1024*1024*1024))
}
