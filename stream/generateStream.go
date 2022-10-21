package stream

import (
	"context"
	"fmt"
	"generate-stream-tools/models"
	"time"
)

func ExecuteProducers(timeDelay int64, speedMult int64) {

	res := queryMongo()

	count := 0
	var last float64
	last = 0

	var initialTime int64 = 0
	for res.Next(context.Background()) {
		var event models.Event
		res.Decode(&event)
		if initialTime == 0 {
			initialTime = event.Millisecs
		}
		t1 := time.UnixMilli(initialTime)
		t2 := time.UnixMilli(event.Millisecs)

		//fmt.Printf("[%d] Event: %s, Time: %f\n", count, event.ID, t2.Sub(t1).Minutes())

		// Delay between first task and current task
		diff := t2.Sub(t1).Seconds()
		// Same but with delay
		diff = diff + float64(timeDelay)

		//fmt.Println(diff, diff+float64(timeDelay))

		// Delay between this task and last task
		actual := (diff - float64(last)) / float64(speedMult)

		fmt.Printf("Executing task %d, delay: %f, fromInit: %f\n", count, actual, diff)

		count++
		last = actual

		time.Sleep(time.Duration(actual) * time.Second)

		go produceEvent("events", 0, event)

	}
}

//1043.483333
