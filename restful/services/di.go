package services

import "github.com/google/wire"

var ServiceGraphSet = wire.NewSet(
	NewProductCtlServices,
)
