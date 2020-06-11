package transmission

import (
	"bytes"
	"encoding/json"
	"io"
)

type Request struct {
	Method    string                 `json:"method"`
	Arguments map[string]interface{} `json:"arguments,omitempty"`
	Tag       *int                   `json:"tag,omitempty"`
}

func (r *Request) AsBody() io.Reader {
	b, _ := json.Marshal(r)
	return bytes.NewBuffer(b)
}

func NewRequest(args ...interface{}) *Request {
	l := len(args)
	r := &Request{}
	if l >= 1 {
		r.Method = args[0].(string)
	}
	if l >= 2 && args[1] != nil {
		r.Arguments = args[1].(map[string]interface{})
	}
	if l >= 3 {
		i := args[2].(int)
		r.Tag = &i
	}
	return r
}
