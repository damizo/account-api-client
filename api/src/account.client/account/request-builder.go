package account

import "net/http"

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
}

func NewRequestBuilder(client http.Client) *RequestBuilder {
	return &RequestBuilder{
		client: client,
	}
}

func (a RequestBuilder) withUrl(url string) RequestBuilder {
	a.url = url
	return a
}

func (a RequestBuilder) withMethod(method HttpMethod) RequestBuilder {
	a.method = method
	return a
}

func (a RequestBuilder) withTimeout(timeoutInMs int64) RequestBuilder {
	a.timeout = timeoutInMs
	return a
}
/*
func (a RequestBuilder) build() http.Request {
	return http.Request{
		Method:           http.MethodPost,
		URL:              ,
		Proto:            "",
		ProtoMajor:       0,
		ProtoMinor:       0,
		Header:           nil,
		Body:             nil,
		GetBody:          nil,
		ContentLength:    0,
		TransferEncoding: nil,
		Close:            false,
		Host:             "",
		Form:             nil,
		PostForm:         nil,
		MultipartForm:    nil,
		Trailer:          nil,
		RemoteAddr:       "",
		RequestURI:       "",
		TLS:              nil,
		Cancel:           nil,
		Response:         nil,
	}

}*/