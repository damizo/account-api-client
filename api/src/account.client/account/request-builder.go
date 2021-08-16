package account

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type HttpMethod string

const (
	POST HttpMethod = "POST"
	PUT HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
	GET HttpMethod = "GET"
)

type RequestBuilder struct  {
	client http.Client
	url string
	method HttpMethod
	timeout int64
	data *bytes.Buffer
	headerMap map [string]string
	urlSuffix string
	request *http.Request
}

type RequestExecutionBuilder struct  {
	requestBuilder RequestBuilder
}

func NewRequestExecutionBuilder(client http.Client) *RequestBuilder {
	return &RequestBuilder{
		client: client,
	}
}


func (a RequestBuilder) withMethod(method HttpMethod) RequestBuilder {
	a.method = method
	return a
}

func (a RequestBuilder) withTimeout(timeoutInMs int64) RequestBuilder {
	a.timeout = timeoutInMs
	return a
}

func (a RequestBuilder) withBody(data *bytes.Buffer) RequestBuilder {
	a.data = data
	return a
}

func (a RequestBuilder) withHeader(key string, value string) RequestBuilder{
	if a.headerMap == nil {
		a.headerMap = make(map[string]string)
	}
	a.headerMap[key] = value
	return a
}

func (a RequestBuilder) withUrlSuffix(urlSuffix string) RequestBuilder{
	a.urlSuffix = urlSuffix
	return a
}

func (a RequestBuilder) build() RequestExecutionBuilder {
	request, _ := http.NewRequest("POST", fmt.Sprintf("%s%s", a.url, a.urlSuffix), a.data)
	for key, element := range a.headerMap {
		request.Header.Set(key, element)
	}
	a.request = request
	return RequestExecutionBuilder{requestBuilder: a}
}

func (a RequestBuilder) withUrl(url string) RequestBuilder {
	a.url = url
	return a
}

func (a RequestExecutionBuilder) handle() *http.Response {
	response, err := a.requestBuilder.client.Do(a.requestBuilder.request)
	if err != nil {
		log.Fatal(err)
	}
	return response
}
