package api

import "net/http"

// Middleware ...
type Middleware func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)

// Middlewares ...
var Middlewares = []Middleware{}
