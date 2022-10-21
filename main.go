package main

import (
	"generate-stream-tools/gendata"
	"generate-stream-tools/insertdata"
	"generate-stream-tools/stream"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

/* Yo debería:
- Crear un archivo con un output de un stream (Waze)
	Opciones
	- Tener output
	- Insertar en la BDD
	- Insertar en la BDD y tener output
- Insertar un output en la BDD
	Opciones
	- Input de un archivo
- Generar un stream con Kafka
	- Tiempo de delay
	- Multiplicador de tiempo (x2, x4, x8)
	- Kafka topic
	- Kafka broker
	- Kafka port
	- Kafka user
	- Kafka password
	- Kafka group
*/

func main() {
	app := &cli.App{
		Name:  "Stream Generator",
		Usage: "A simple stream generator. You could use it to generate a stream of data for testing purposes, using Kafka.",
		Commands: []*cli.Command{
			{
				Name:  "waze",
				Usage: "Scrapping Waze data and saving it to a file",
				Action: func(c *cli.Context) error {
					gendata.DataWaze(10)
					return nil
				},
			},
			{
				Name:  "dump",
				Usage: "Insert dump file data to mongoDB",
				Action: func(cCtx *cli.Context) error {
					insertdata.InsertToMongo(cCtx.String("file"))
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Aliases:  []string{"f"},
						Usage:    "File to insert",
						Required: true,
					},
				},
			},
			{
				Name:  "stream",
				Usage: "Stream data from MongoDB to Kafka",
				Action: func(c *cli.Context) error {
					stream.ExecuteProducers(c.Int64("delay"), c.Int64("multiplier"))
					return nil
				},
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "delay",
						Aliases:     []string{"d"},
						Value:       0,
						Usage:       "Delay in secs applied to every message",
						DefaultText: "0 secs",
					},
					&cli.IntFlag{
						Name:        "multiplier",
						Aliases:     []string{"m"},
						Value:       1,
						Usage:       "Multiplier of time speed (x2, x4, x8)",
						DefaultText: "x1",
					},
					&cli.StringFlag{
						Name:        "topic",
						Aliases:     []string{"t"},
						Value:       "events",
						Usage:       "Kafka topic",
						DefaultText: "events",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
