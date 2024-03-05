package rabbitmq

import (
	"github.com/kumin/GolangMaster/utils/envx"
)

var (
	RabbitMqHost   = envx.GetString("RMQ_HOST", "amqp://localhost:5672/")
	HelloQueueName = envx.GetString("QUEUE_NAME", "hello")
)
