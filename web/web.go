package web

import (
	"net/http"

	"github.com/haleyrc/api"
)

type HandlerFunc func(api.ExecutionContext, http.ResponseWriter, *http.Request)
