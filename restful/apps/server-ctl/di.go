package apps

import (
	"github.com/google/wire"
	"github.com/kumin/GolangMaster/restful/configs"
)

var ServerGraphSet = wire.NewSet(
	configs.ConfigGraphSet,
	NewHttpServer,
)
