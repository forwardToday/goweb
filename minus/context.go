package minus

import (
	"net/http"
	"net/url"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Multipart      bool
	Form           url.Values
}
