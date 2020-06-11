package transmission

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	URL           *url.URL
	sessionID     string
	rpcVersion    string
	rpcVersionMin string
	httpClient    http.Client
}

func (c *Client) Do(req *Request, resp *Response) error {
	if c.sessionID == "" {
		sessionID, err := c.getSessionID()
		if err != nil {
			return err
		}
		c.sessionID = sessionID
	}
	httpReq, err := http.NewRequest("POST", c.URL.String(), req.AsBody())
	if err != nil {
		return err
	}
	httpReq.Header.Set("X-Transmission-Session-Id", c.sessionID)
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	//
	/*
		buf := new(bytes.Buffer)
		buf.ReadFrom(httpResp.Body)
		fmt.Println(string(buf.Bytes()))
		r := bytes.NewReader(buf.Bytes())
		dec := json.NewDecoder(r)
		///
	*/
	dec := json.NewDecoder(httpResp.Body)
	if err = dec.Decode(resp); err != nil {
		return err
	}
	if req.Tag != nil && *req.Tag != resp.Tag {
		return fmt.Errorf("Tag mismatch: %d/%d", *req.Tag, resp.Tag)
	}
	return nil
}

func (c *Client) getSessionID() (string, error) {
	req, err := http.NewRequest("POST", c.URL.String(), NewRequest("session-get").AsBody())
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	sessionID := resp.Header["X-Transmission-Session-Id"][0]
	return sessionID, nil
}
