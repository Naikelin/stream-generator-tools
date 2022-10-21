package gendata

import (
	"fmt"
	"generate-stream-tools/req"
	"generate-stream-tools/save"
	"log"
	"strconv"
	"time"
)

func DataWaze(times int) {

	initialDate := strconv.FormatInt(time.Now().Unix(), 10)
	i := 0
	for i < times {
		resp, err := req.Waze()
		if err != nil {
			panic(err)
		}
		date := time.Now().Format(time.RFC850)
		save.SaveWazeOutput(resp, date, i, initialDate)
		fmt.Printf("[%d] Saved at %s\n", i, date)
		i++
		time.Sleep(5 * time.Minute)
	}

	log.Printf("Data recolected. You can find it at output/dataset%s", initialDate)
}
