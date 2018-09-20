package main

import (
	//"net/http"

	"github.com/gohiweeds/gocrawler"
)

func main() {
	crawler := gocrawler.NewCrawler("test.csv")

	crawler.Get("https://search.51job.com/list/040000,000000,0000,00,9,99,golang,2,1.html?lang=c&stype=&postchannel=0000&workyear=99&cotype=99&degreefrom=99&jobterm=99&companysize=99&providesalary=99&lonlat=0%2C0&radius=-1&ord_field=0&confirmdate=9&fromType=&dibiaoid=0&address=&line=&specialarea=00&from=&welfare=")
	crawler.BarChart()
	//gocrawler.ScrapeTest()
}
