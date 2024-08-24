package code

import (
	"html/template"
	"net/http"
	"sync"
)

// HTMLTemplate is excluded by the flag, so should be OK
var HTMLTemplate = template.Must(template.New("").Parse(""))

// sync.Once is excluded by the flag, so should be OK
var once = sync.Once{}

var HTTPClient = http.Client{} // want "HTTPClient is a global variable"
