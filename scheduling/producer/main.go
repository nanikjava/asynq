package main

import (
	"fmt"
	"github.com/hibiken/asynq/scheduling/tasks"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

const redisAddr = "127.0.0.1:6379"

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	// ------------------------------------------------------
	// Example 1: Enqueue task to be processed immediately.
	//            Use (*Client).Enqueue method.
	// ------------------------------------------------------

	var task *asynq.Task = &asynq.Task{}

	for i := 0; i < 100000; i++ {
		task, err := tasks.NewEmailDeliveryTask(42, fmt.Sprintf("some:template:id:%d", i))
		if err != nil {
			log.Fatalf("could not create task: %v", err)
		}
		info, err := client.Enqueue(task)
		if err != nil {
			log.Fatalf("could not enqueue task: %v", err)
		}
		log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	}

	// ------------------------------------------------------------
	// Example 2: Schedule task to be processed in the future.
	//            Use ProcessIn or ProcessAt option.
	// ------------------------------------------------------------

	info, err := client.Enqueue(task, asynq.ProcessIn(5*time.Second))
	if err != nil {
		log.Fatalf("could not schedule task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	// ----------------------------------------------------------------------------
	// Example 3: Set other options to tune task processing behavior.
	//            Options include MaxRetry, Queue, Timeout, Deadline, Unique etc.
	// ----------------------------------------------------------------------------

	client.SetDefaultOptions(tasks.TypeImageResize, asynq.MaxRetry(10), asynq.Timeout(3*time.Minute))

	task, err = tasks.NewImageResizeTask("https://example.com/myassets/image.jpg")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	info, err = client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	// ---------------------------------------------------------------------------
	// Example 4: Pass options to tune task processing behavior at enqueue time.
	//            Options passed at enqueue time override default ones.
	// ---------------------------------------------------------------------------

	info, err = client.Enqueue(task, asynq.Queue("critical"), asynq.Timeout(30*time.Second))
	if err != nil {
		log.Fatal("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}
