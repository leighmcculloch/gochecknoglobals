package code

import (
	"net/http"
)

var HTTPClient = http.Client{} // want "HTTPClient is a global variable"
