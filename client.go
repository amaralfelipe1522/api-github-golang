package client

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	defaultBasePath = "https://api.github.com/"
	mediaType       = "application/vnd.github.v3+json"
)

// Client se refere a estrutura que irá gerenciar a comunicação com a API
type Client struct {
	client  *http.Client
	BaseURL *url.URL
}

// Response é uma resposta da API que envolve a http.Response retornada da requisição
type Response struct {
	*http.Response
}

// ErrorResponse resporta um possível erro causado na requisição
type ErrorResponse struct {
	// http.Response que causou o erro
	Response *http.Response
	// Mensagem de erro
	Message string
}

// Error implementa a interface error
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}

// retorna uma nova Response dado uma http.Response
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}
