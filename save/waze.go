package save

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

func SaveWazeOutput(A string, date string, i int, initialDate string) {
	/* A es la respuesta completa del GET */
	var temp interface{}

	/* Necesitamos pasar el string a una interface */
	json.Unmarshal([]byte(A), &temp)

	/* Pasamos la interface a Map, por lo que podemos sacar las 'alerts' */
	m := temp.(map[string]interface{})

	/* Pasamos las 'alerts' a una interface para así enviar sólamente la data importante */

	alerts := m["alerts"].([]interface{})

	_ = os.Mkdir("output", 0777)

	f, err := os.OpenFile("output/dataset"+initialDate, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(f)
	w.WriteString("[" + strconv.Itoa(i) + "] " + date + "\n")
	for _, alert := range alerts {
		valStr, _ := json.Marshal(alert)
		w.WriteString(string(valStr) + "\n")
	}
	w.Flush()
}
