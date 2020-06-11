package main

import (
	"fmt"
	"net/url"
	"path"

	"github.com/enolgor/go-transmission-lib/transmission"
)

type Tracker struct {
	Announce string `json:"announce"`
}

type Torrent struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	DownloadDir string    `json:"downloadDir"`
	Trackers    []Tracker `json:"trackers"`
}

func main() {
	client := transmission.Client{URL: &url.URL{Scheme: "http", Host: "192.168.1.2:9091", Path: "/transmission/rpc"}}
	torrentData := struct {
		Torrents []Torrent `json:"torrents"`
	}{}
	response := &transmission.Response{Arguments: &torrentData}
	reqArgs := make(map[string]interface{})
	reqArgs["fields"] = []string{"id", "name", "downloadDir", "trackers"}
	request := transmission.NewRequest("torrent-get", reqArgs)
	err := client.Do(request, response)
	if err != nil {
		panic(err)
	}
	for _, torrent := range torrentData.Torrents {
		if err := checkTorrent(&torrent, &client); err != nil {
			fmt.Printf("Error: %s - %s\n", err.Error(), torrent.Name)
		} else {
			fmt.Printf("OK: %s\n", torrent.Name)
		}
	}
}

func checkTorrent(torrent *Torrent, client *transmission.Client) error {
	var hostname string
	if len(torrent.Trackers) > 0 {
		tracker := torrent.Trackers[0]
		url, err := url.Parse(tracker.Announce)
		if err != nil {
			return err
		}
		hostname = url.Hostname()
	} else {
		hostname = "unknown"
	}
	if torrent.DownloadDir != "/data/completed" {
		return fmt.Errorf("Skipping")
	}
	location := path.Join(torrent.DownloadDir, hostname)
	return move(torrent.ID, location, client)
}

func move(id int, location string, client *transmission.Client) error {
	response := &transmission.Response{}
	reqArgs := make(map[string]interface{})
	reqArgs["ids"] = id
	reqArgs["location"] = location
	reqArgs["move"] = true
	request := transmission.NewRequest("torrent-set-location", reqArgs)
	err := client.Do(request, response)
	if err != nil {
		return err
	}
	return nil
}
