package account

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type HttpMethod string

const (
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
	GET    HttpMethod = "GET"
)

type RequestBuilder struct {
	client    http.Client
	url       string
	method    HttpMethod
	timeout   int64
	data      *bytes.Buffer
	headerMap map[string]string
	urlSuffix string
	request   *http.Request
}

type RequestExecutionBuilder struct {
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

func (a RequestBuilder) withHeader(key string, value string) RequestBuilder {
	if a.headerMap == nil {
		a.headerMap = make(map[string]string)
	}
	a.headerMap[key] = value
	return a
}

func (a RequestBuilder) withUrlSuffix(urlSuffix string) RequestBuilder {
	a.urlSuffix = urlSuffix
	return a
}

func (a RequestBuilder) build() RequestExecutionBuilder {
	method := ""
	switch a.method {
	case POST:
		method = http.MethodPost
	case GET:
		method = http.MethodGet
	case PUT:
		method = http.MethodPut
	case DELETE:
		method = http.MethodDelete
	default:
		log.Fatal(fmt.Sprintf("Method %s is not supported for request builder", a.method))
	}

	var request *http.Request
	if a.data == nil {
		newRequest, _ := http.NewRequest(method, fmt.Sprintf("%s%s", a.url, a.urlSuffix), nil)
		request = newRequest
	} else {
		newRequest, _ := http.NewRequest(method, fmt.Sprintf("%s%s", a.url, a.urlSuffix), a.data)
		request = newRequest
	}

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

func (a RequestBuilder) withParam(key string, value string) RequestBuilder {
	suffix := a.urlSuffix
	finalSuffix := strings.Replace(suffix, fmt.Sprintf("{%s}", key), value, 1)
	a.urlSuffix = finalSuffix
	return a
}

func (a RequestExecutionBuilder) handle() *http.Response {
	response, err := a.requestBuilder.client.Do(a.requestBuilder.request)
	if err != nil {
		log.Fatal(err)
	}
	return response
}
