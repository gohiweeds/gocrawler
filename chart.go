package gocrawler

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	_ "strconv"
	"strings"

	"github.com/wcharczuk/go-chart"
)

func (c *Crawler) LocationChart() error {
	file, err := os.Open(c.filename)
	if err != nil {
		log.Printf("open csv file failed: %s\n", err.Error())
		return err
	}
	defer file.Close()

	r := csv.NewReader(file)

	barData := make(map[string]int)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Salary
		//printColumn(5, record)

		// Location
		printColumn(4, record, barData)
		//Data Analysis
	}

	for k, v := range barData {
		fmt.Printf("key=%s,%d\n", k, v)
	}

	//Salary
	//	barChart(barData, "locationbar.png")
	//	pieChart(barData, "locationpie.png")
	//Location
	barChart(barData, 1, "locationbar.png")
	pieChart(barData, 1, "locationpie.png")
	return nil
}

func (c *Crawler) SalaryChart() error {
	file, err := os.Open(c.filename)
	if err != nil {
		log.Printf("open csv file failed: %s\n", err.Error())
		return err
	}
	defer file.Close()

	r := csv.NewReader(file)

	barData := make(map[string]int)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Salary
		printColumn(5, record, barData)

		// Location
		//printColumn(4, record)
		//Data Analysis
	}

	for k, v := range barData {
		fmt.Printf("key=%s,%d\n", k, v)
	}

	//Salary
	barChart(barData, 2, "salarybar.png")
	pieChart(barData, 2, "salarypie.png")
	//Location
	//	barChart(barData, 1, "locationbar.png")
	//	pieChart(barData, 1, "locationpie.png")
	return nil
}

func printColumn(col int, record []string, barData map[string]int) {
	if 0 == strings.Compare(record[col-1], "薪资") ||
		0 == strings.Compare(record[col-1], "工作地点") ||
		0 == strings.Compare(record[col-1], "") {

	} else {

		//		fmt.Printf("col: %d, %s\n", col, record[col-1])
		index := record[col-1]
		switch col {
		case 4:
			if strings.Contains(index, "深圳") {
				barData["深圳"]++
			} else if strings.Contains(index, "北京") {
				barData["北京"]++
			} else if strings.Contains(index, "上海") {
				barData["上海"]++
			} else if strings.Contains(index, "杭州") {
				barData["杭州"]++
			} else if strings.Contains(index, "广州") {
				barData["广州"]++
			} else {
				barData[index]++
			}
			break
		case 5:
			barData[index]++
			break
		default:
		}

	}
}

func barChart(i map[string]int, kind int, filename string) {

	var values []chart.Value
	switch kind {
	case 1:
		statics := 0
		for k, v := range i {
			if v < 3 { //too much, will not display x-axis label
				continue
			}
			var chartValue = chart.Value{
				Label: k,
				Value: float64(v),
			}
			values = append(values, chartValue)
			statics = statics + v
		}

		break
	case 2:
		statics := 0
		for k, v := range i {
			if v < 3 { //too much, will not display x-axis label
				continue
			}
			if strings.Contains(k, "千") || strings.Contains(k, "天") {
				continue
			}
			var chartValue = chart.Value{
				Label: strings.Trim(strings.Trim(k, "万/月"), "万/年"),
				Value: float64(v),
			}
			values = append(values, chartValue)
			statics = statics + v
		}
		break
	default:
	}

	//	fmt.Printf("Total: %d\n", statics)

	graph := chart.BarChart{
		Width:      1024,
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
			Show:                true,
			TextRotationDegrees: 90.0,
		},
		YAxis: chart.YAxis{

			NameStyle: chart.Style{
				Show: true,
			},
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

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("create file failed", err.Error())
	}
	defer file.Close()

	io.Copy(file, buffer)
}

func pieChart(i map[string]int, kind int, filename string) {
	var values []chart.Value
	switch kind {
	case 1:
		statics := 0
		for k, v := range i {
			if v < 3 { //too much, will not display x-axis label
				continue
			}

			var chartValue = chart.Value{
				Label: k,
				Value: float64(v),
			}
			values = append(values, chartValue)
			statics = statics + v
		}

		break
	case 2:
		statics := 0
		for k, v := range i {
			if v < 3 { //too much, will not display x-axis label
				continue
			}
			if strings.Contains(k, "千") || strings.Contains(k, "天") {
				continue
			}
			var chartValue = chart.Value{
				Label: strings.Trim(strings.Trim(k, "万/月"), "万/年"),
				Value: float64(v),
			}
			values = append(values, chartValue)
			statics = statics + v
		}
		break
	default:
	}

	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: values,
	}
	buffer := bytes.NewBuffer([]byte{})
	err := pie.Render(chart.PNG, buffer)
	if err != nil {
		fmt.Println("render graph failed", err.Error())
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("create file failed", err.Error())
	}
	defer file.Close()

	io.Copy(file, buffer)
}
