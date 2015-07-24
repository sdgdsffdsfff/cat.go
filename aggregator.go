package cat

import (
	"time"
)

var (
	aggregator_message_channel chan Message
	aggregator_max_batch_size  int
)

func cat_aggregator_init() {
	aggregator_message_channel = make(chan Message, 1<<10)
	aggregator_max_batch_size = 1 << 8
	go aggregator_run()
}

func aggregator_run() {
	for {
		if aggregator_collect() {
			time.Sleep(1 << 16 * time.Microsecond)
		}
	}
}

//False returned when it seems to be busy.
func aggregator_collect() bool {
	messages := make(chan Message, sender_max_batch_size)
	var count = 0
collect:
	for count < aggregator_max_batch_size {
		select {
		case message := <-aggregator_message_channel:
			messages <- message
			count++
		default:
			break collect
		}
	}
	close(messages)
	if count > 0 {
		aggregator_transfer(messages)
		return false
	} else {
		return true
	}
}

func aggregator_transfer(messages <-chan Message) {
	t := NewTransaction("_CatMergeTree", "_CatMergeTree", nil)
	for message := range messages {
		t.AddChild(message)
	}
	t.SetStatus("0")
	t.Complete()
	sender_transaction_channel <- t
}
