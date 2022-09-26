package handler

import "github.com/google/wire"

var HandlerGraphSet = wire.NewSet(
	NewProductCtlHandler,
)
