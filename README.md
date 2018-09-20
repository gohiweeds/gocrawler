# gocrawler

[ Done ] Crawl 51job.com to find target job number and generate pie and bar chart.

[ TODO ] Support lagou.com/zhilian.com

[ TODO ] Fix Chinese support in go-chart

## Usage

* open `www.51job.com` in web browser
* search keyword in the textbox in the target website, like `golang`, `linux` 
* run `go run cmd/main.go`

After above operation, you will get the chart generate.

Example:
	
	Location Pie Chart and Bar Chart

![Location Bar Chart](https://github.com/gohiweeds/gocrawler/blob/master/cmd/locationbar.png)

![Location Pie Chart](https://github.com/gohiweeds/gocrawler/blob/master/cmd/locationpie.png)


	Salary Pie Chart and Bar Chart
![Salary Bar Chart](https://github.com/gohiweeds/gocrawler/blob/master/cmd/salarybar.png)


![Salary Pie Chart](https://github.com/gohiweeds/gocrawler/blob/master/cmd/salarypie.png)
