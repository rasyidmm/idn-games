package resty

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type BuilderResty struct {
	Endpoint    string
	Headers     map[string]string
	Body        interface{}
	RestyClient *resty.Client
}

func New() *BuilderResty {
	client := resty.New()
	return &BuilderResty{
		Headers:     map[string]string{},
		RestyClient: client,
	}
}

func (b *BuilderResty) SetEndpoint(Endpoint string) {
	b.Endpoint = Endpoint
}

func (b *BuilderResty) SetHeader(header map[string]string) {
	for key, val := range header {
		b.Headers[key] = val
	}
}

func (b *BuilderResty) SetBody(body interface{}) {
	b.Body = body
}

func (b *BuilderResty) SetRequest(Endpoint string, header map[string]string, body interface{}) *BuilderResty {
	b.SetEndpoint(Endpoint)
	b.SetHeader(header)
	b.SetBody(body)

	return b
}

func (b *BuilderResty) Post(response interface{}) error {
	fmt.Println("Endpoint : ", b.Endpoint)
	fmt.Println("Header : ", b.Headers)
	fmt.Println("Body : ", Stringify(b.Body))

	data, err := b.RestyClient.SetPreRequestHook(b.BeforeRequest).R().SetBody(b.Body).Post(b.Endpoint)
	if err != nil {
		return err
	}

	if data.StatusCode() != 200 {
		return errors.New(string(data.Body()))
	}

	var body = data.Body()
	err = json.Unmarshal(body, response)
	if err != nil {
		return err
	}
	return nil
}

func (b *BuilderResty) Get(response interface{}) error {
	fmt.Println("Endpoint : ", b.Endpoint)
	fmt.Println("Header : ", b.Headers)

	data, err := b.RestyClient.SetPreRequestHook(b.BeforeRequest).R().Get(b.Endpoint)
	if err != nil {
		return err
	}

	var body = data.Body()
	err = json.Unmarshal(body, response)
	if err != nil {
		return err
	}
	return nil
}

func (b *BuilderResty) Put(response interface{}) error {
	fmt.Println("Endpoint : ", b.Endpoint)
	fmt.Println("Header : ", b.Headers)
	fmt.Println("Body : ", Stringify(b.Body))

	data, err := b.RestyClient.SetPreRequestHook(b.BeforeRequest).R().SetBody(b.Body).Put(b.Endpoint)
	if err != nil {
		return err
	}

	if data.StatusCode() != 200 {
		return errors.New(string(data.Body()))
	}

	var body = data.Body()
	err = json.Unmarshal(body, response)
	if err != nil {
		return err
	}
	return nil
}

func (b *BuilderResty) BeforeRequest(r *resty.Client, h *http.Request) error {
	for k, v := range b.Headers {
		h.Header[k] = []string{v}
	}
	return nil
}

func Stringify(data interface{}) string {
	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}
