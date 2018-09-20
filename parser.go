package gocrawler

import (
	"fmt"
	"log"
	"os"

	"encoding/csv"
	"io"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
)

func ScrapeTest() error {

	file, err := os.Open("../test/welfare=")
	defer file.Close()
	if err != nil {
		log.Println("open file failed", err.Error())
		return err
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Println("newdocumentfromreader failed", err.Error())
		return err
	}

	pTitle := doc.Find("title").Text()
	var gbk = mahonia.NewDecoder("gbk")
	n := gbk.ConvertString(pTitle)
	log.Printf("title:%v\n", n)

	//	var count = 0

	saveFile, err := os.Create("../test/51.csv")
	defer saveFile.Close()

	if err != nil {
		log.Println("open file failed", err.Error())
		return err
	}
	w := csv.NewWriter(saveFile)

	//records := make([][]string, 0)
	doc.Find(".dw_table").EachWithBreak(func(i int, s *goquery.Selection) bool {
		//遍历孩子节点，需要中断跳出，所以用了EachWithBreak
		s.Children().EachWithBreak(func(j int, selection *goquery.Selection) bool {
			//fmt.Println(selection.Text())
			//获取内容
			//			str := selection.Text()
			//			fmt.Println(gbk.ConvertString(str))

			selection.Find("div.rt").EachWithBreak(func(k int, ss *goquery.Selection) bool {

				ss.Find("a").EachWithBreak(func(kk int, sss *goquery.Selection) bool {
					href, exist := sss.Find("a.dicon Dm on").Attr("href")
					if exist {
						fmt.Printf("href=%s\n", href)
					}
					return true
				})
				//				href, exist := ss.Attr("href")
				//				if exist {
				//					fmt.Printf("href=%s\n", href)
				//				}
				return true

			})

			//			title := selection.Find("p.t1").Text()
			//			company := selection.Find("span.t2").Text()
			//			location := selection.Find("span.t3").Text()
			//			salary := selection.Find("span.t4").Text()
			//			date := selection.Find("span.t5").Text()

			//			if title != "" || company != "" {
			//				title = strings.TrimSpace(title)
			//				count++
			//				//				if count == 1 {
			//				//					return true
			//				//				}
			//				fmt.Printf("%d:%s\t%s\t%s\t%s\t%s\n", count, gbk.ConvertString(title),
			//					gbk.ConvertString(company), gbk.ConvertString(location),
			//					gbk.ConvertString(salary), gbk.ConvertString(date))

			//				strslice := make([]string, 0)
			//				strslice = append(strslice, strconv.Itoa(count))
			//				strslice = append(strslice, gbk.ConvertString(title))
			//				strslice = append(strslice, gbk.ConvertString(company))
			//				strslice = append(strslice, gbk.ConvertString(location))
			//				strslice = append(strslice, gbk.ConvertString(salary))
			//				strslice = append(strslice, gbk.ConvertString(date))
			//				//records = append(records, strslice...)
			//				w.Write(strslice)
			//			}
			//fmt.Println("--------------------------")
			return true
		})
		//fmt.Println("==================")
		if i == 0 {
			return false
		}
		return true

	})
	//w.WriteAll(records)
	w.Flush()

	log.Println("END")
	return nil
}

func (c *Crawler) Scrape51JobDetail(reader io.Reader) error {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Println("newdocumentreader failed", err.Error())
		return err
	}

	var gbk = mahonia.NewDecoder("gbk")
	doc.Find("div.tCompany_main").EachWithBreak(func(i int, s *goquery.Selection) bool {
		text := s.Find("div.job_msg p").Text()
		//fmt.Printf("text : %s\n", strings.Trim(text, "\n"))
		fmt.Printf("text : %s\n", gbk.ConvertString(strings.Trim(strings.TrimSpace(text), "\n")))
		c.detail = c.detail + gbk.ConvertString(strings.Trim(strings.TrimSpace(text), " ")) + "\n"
		//		s.Children().EachWithBreak(func(j int, selection *goquery.Selection) bool {
		//			selection.Find("p").EachWithBreak(func(kk int, sss *goquery.Selection) bool {
		//				text := sss.Text()
		//				//				fmt.Printf("text : %s\n", gbk.ConvertString(strings.TrimSpace(text)))
		//				c.detail = c.detail + gbk.ConvertString(strings.TrimSpace(text)) + "\n"
		//				return true
		//			})
		//			return true
		//		})
		return true
	})
	return nil
}

func (c *Crawler) Scrape51(reader io.Reader) error {

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Println("newdocumentfromreader failed", err.Error())
		return err
	}

	//Title
	pTitle := doc.Find("title").Text()
	var gbk = mahonia.NewDecoder("gbk")
	n := gbk.ConvertString(pTitle)
	log.Printf("title:%v\n", n)

	//Content Jobs list
	doc.Find(".dw_table").EachWithBreak(func(i int, s *goquery.Selection) bool {
		//遍历孩子节点，需要中断跳出，所以用了EachWithBreak
		s.Children().EachWithBreak(func(j int, selection *goquery.Selection) bool {

			selection.Find("div.rt").EachWithBreak(func(k int, ss *goquery.Selection) bool {
				ss.Find("a").EachWithBreak(func(kk int, sss *goquery.Selection) bool {
					href, exist := sss.Attr("href")
					id, exist := sss.Attr("id")
					if exist {
						id = strings.TrimSpace(id)
						if 0 == strings.Compare(id, "rtNext") {
							c.queryUrl = strings.TrimSpace(href)
							fmt.Printf("need query: %s", c.queryUrl)
							return true
						}
					}
					return true
				})
				return true
			})

			//Find title requirements
			requirement, exist := selection.Find("p.t1").Find("a").Attr("href")
			if exist {
				//				fmt.Printf("detail: %s\n", requirement)
				//				c.RequestWithFunc(requirement, c.Scrape51JobDetail)
			}

			title := selection.Find("p.t1").Text()
			company := selection.Find("span.t2").Text()
			location := selection.Find("span.t3").Text()
			salary := selection.Find("span.t4").Text()
			date := selection.Find("span.t5").Text()

			if title != "" || company != "" {
				title = strings.TrimSpace(title)
				c.line++

				strslice := make([]string, 0)
				strslice = append(strslice, strconv.Itoa(c.line))
				strslice = append(strslice, gbk.ConvertString(title))
				strslice = append(strslice, gbk.ConvertString(company))
				strslice = append(strslice, gbk.ConvertString(location))
				strslice = append(strslice, gbk.ConvertString(salary))
				strslice = append(strslice, gbk.ConvertString(date))
				strslice = append(strslice, requirement)
				strslice = append(strslice, c.detail)
				c.detail = ""
				c.writer.Write(strslice)
			}

			return true
		})
		if i == 0 {
			return false
		}
		return true

	})
	c.writer.Flush()

	log.Println("Parse Finished!!!")

	if c.queryUrl != "" {
		c.Get(c.queryUrl)
	}
	return nil
}

//https://search.51job.com/list/040000,000000,0000,00,9,99,golang,2,2.html?lang=c&stype=1&postchannel=0000&workyear=99&cotype=99&degreefrom=99&jobterm=99&companysize=99&lonlat=0%2C0&radius=-1&ord_field=0&confirmdate=9&fromType=&dibiaoid=0&address=&line=&specialarea=00&from=&welfare=
//https://search.51job.com/list/040000,000000,0000,00,9,99,golang,2,1.html?lang=c&stype=1&postchannel=0000&workyear=99&cotype=99&degreefrom=99&jobterm=99&companysize=99&lonlat=0%2C0&radius=-1&ord_field=1&confirmdate=9&fromType=&dibiaoid=0&address=&line=&specialarea=00&from=&welfare=
//https://search.51job.com/list/040000,000000,0000,00,9,99,golang,2,1.html?lang=c&stype=1&postchannel=0000&workyear=99&cotype=99&degreefrom=99&jobterm=99&companysize=99&lonlat=0%2C0&radius=-1&ord_field=0&confirmdate=9&fromType=&dibiaoid=0&address=&line=&specialarea=00&from=&welfare=
