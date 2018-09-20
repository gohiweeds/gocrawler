package gocrawler

import (
	"encoding/csv"
	//	"io/ioutil"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Crawler struct {
	client   *http.Client
	filename string
	writer   *csv.Writer
	line     int
	urls     []string
	detail   string
	queryUrl string
}

func NewCrawler(filename string) *Crawler {
	file, err := os.Create(filename)
	if err != nil {
		log.Println("create file failed", err.Error())
		return nil
	}

	return &Crawler{
		client:   new(http.Client),    // client to request http server
		filename: filename,            // filename to save csv
		writer:   csv.NewWriter(file), // writer to save csv
		urls:     make([]string, 0),   // urls store accessed urls
		line:     0,
	}
}

//https://search.51job.com/list/040000,000000,0000,00,9,99,golang,2,1.html?lang=c&stype=&postchannel=0000&workyear=99&cotype=99&degreefrom=99&jobterm=99&companysize=99&providesalary=99&lonlat=0%2C0&radius=-1&ord_field=0&confirmdate=9&fromType=&dibiaoid=0&address=&line=&specialarea=00&from=&welfare=
func (c *Crawler) Get(url string) error {

	//c.urls = append(c.urls, url)
	c.queryUrl = ""
	if c.client == nil {
		c.client = &http.Client{}
	}

	resp, err := c.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	c.Scrape51(resp.Body)

	return nil
}

type Callback func(reader io.Reader) error

func (c *Crawler) RequestWithFunc(url string, fn Callback) error {
	if c.client == nil {
		c.client = &http.Client{}
	}

	resp, err := c.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fn(resp.Body)
	return nil
}

func extractFilename(url string) string {
	strs := strings.Split(url, "/")
	fileName := strs[len(strs)-1]

	params := strings.Split(fileName, "&")
	param := params[len(params)-1]
	return param
}
