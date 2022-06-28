package ports

import "net/http"

type UtilsHTTPRequest interface {
	HTTPRequest(url string, method string, payload []byte) (*http.Response, error)
}
