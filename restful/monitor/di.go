package monitor

import "github.com/google/wire"

var MonitorGraphSet = wire.NewSet(
	NewMetricServer,
)
