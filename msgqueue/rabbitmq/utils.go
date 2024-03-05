package rabbitmq

import "log"

func FailOnError(err error) {
	if err != nil {
		log.Panicf("%s", err)
	}
}
