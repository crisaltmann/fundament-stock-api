package infrastructure

import "os"

const resultQueueName = "local-fs-quarterly-result"
const orderQueueName = "local-fs-order"
const holdingResultQueueName = "local-fs-holding-result"

func GetResultQueueName() string {
	queue := os.Getenv("RESULT_QUEUE_NAME")
	if queue == "" {
		return resultQueueName
	}
	return queue
}

func GetOrderQueueName() string {
	queue := os.Getenv("ORDER_QUEUE_NAME")
	if queue == "" {
		return orderQueueName
	}
	return queue
}

func GetHoldingResultQueueName() string {
	queue := os.Getenv("HOLDING_RESULT_QUEUE_NAME")
	if queue == "" {
		return holdingResultQueueName
	}
	return queue
}