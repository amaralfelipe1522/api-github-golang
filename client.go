package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.github.com/"
	mediaType      = "application/vnd.github.v3+json"
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

// CheckResponse se responsabiliza por verificar a existência de erros de uma Response
func CheckResponse(r *http.Response) error {
	// Verifica se o status code está entre 200 e 299 que define ser sucesso
	if code := r.StatusCode; code >= 200 && code <= 299 {
		return nil
	}

	// Lê e decodifica a mensagem de erro em uma errorResponse
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err = json.Unmarshal(data, errorResponse)
		if err != nil {
			return err
		}
	}
	return errorResponse
}

// NewClient se responsabiliza por disparar as requests
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// Converte defaultBaseURL para o tipo url.URL
	baseURL, _ := url.Parse(defaultBaseURL)

	// Instância um novo Client
	c := &Client{client: httpClient, BaseURL: baseURL}

	return c
}
