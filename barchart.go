package gocrawler

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/wcharczuk/go-chart"
)

var barData map[string]int

func (c *Crawler) BarChart() error {
	file, err := os.Open(c.filename)
	if err != nil {
		log.Printf("open csv file failed: %s\n", err.Error())
		return err
	}
	defer file.Close()

	r := csv.NewReader(file)

	barData = make(map[string]int)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(record)
		printColumn(5, record)
		//Data Analysis
	}

	for k, v := range barData {
		fmt.Printf("key=%s,%d\n", k, v)
	}

	chartBar(barData)
	return nil
}

func printColumn(col int, record []string) {
	if 0 == strings.Compare(record[col-1], "薪资") ||
		0 == strings.Compare(record[col-1], "") {

	} else {

		//		fmt.Printf("col: %d, %s\n", col, record[col-1])
		index := record[col-1]
		barData[index]++
	}
}

func chartBar(i map[string]int) {

	var values []chart.Value
	statics := 0
	for k, v := range i {
		if v < 3 {
			continue
		}
		var chartValue = chart.Value{
			Label: strings.Trim(k, "万/月"),
			Value: float64(v),
		}
		values = append(values, chartValue)
		statics = statics + v
	}
	fmt.Printf("Total: %d\n", statics)

	graph := chart.BarChart{
		Title:      "Salary",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Bars: values,
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		fmt.Println("render graph failed", err.Error())
	}

	file, err := os.Create("salary51.png")
	if err != nil {
		fmt.Println("create file failed", err.Error())
	}
	defer file.Close()

	io.Copy(file, buffer)
}
