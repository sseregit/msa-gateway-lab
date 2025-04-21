package http

import "errors"

type HttpMethod string

const (
	GET    = HttpMethod("GET")
	POST   = HttpMethod("POST")
	DELETE = HttpMethod("DELETE")
	PUT    = HttpMethod("PUT")
)

func (h HttpMethod) ToString() string {
	return string(h)
}

type GetType string

const (
	QUERY = GetType("query")
	URL   = GetType("url")
)

func (g GetType) ToString() string {
	return string(g)
}

func (g GetType) CheckType() error {
	switch g {
	case QUERY, URL:
		return nil
	default:
		return errors.New("Failed to check get type")
	}
}
