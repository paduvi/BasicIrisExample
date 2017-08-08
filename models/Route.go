package models

import (
	"github.com/kataras/iris/context"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc context.Handler
}

type Routes []Route
