package gocrawler

import (
	"encoding/csv"
	"log"
	"os"
)

func saveFile(filename string, data [][]string) error {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Printf("open file failed %s\n", err.Error())
		return err
	}

	w := csv.NewWriter(file)

	w.WriteAll(data)
	w.Flush()
	if err = w.Error(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
