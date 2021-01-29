package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
)

func main() {

	var location 		string
	var category 		string
	var title           string
	var	price           string
	var	postedOn        string
	var forSaleBy       string
	var link             string
	var metaData        string

	fmt.Print("Enter Location : ")
	fmt.Scan(&location)
	fmt.Println(location+"/n")

	fmt.Print("Enter Category : ")
	fmt.Scan(&category)
	fmt.Println(category)


	c := colly.NewCollector()
	db, _ := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/scrapper2")


	c.OnHTML(".gtm-normal-ad", func(element *colly.HTMLElement) {
		title = element.ChildText(".heading--2eONR")
		price = element.ChildText(".price--3SnqI")
		link = element.ChildAttr(".card-link--3ssYv", "href")

		fmt.Println("\tTitle: ", title)
		fmt.Println("\tPrice: ", price)
		fmt.Println("\tURL: ", link)
		err := element.Request.Visit(link)
		if err != nil {
			panic(err.Error())
		}
	})

	c.OnHTML(".sub-title--37mkY", func(element *colly.HTMLElement) {
		postedOn = element.Text
		fmt.Println("\tPosted On: ", postedOn)
	})

	c.OnHTML(".contact-name--m97Sb", func(element *colly.HTMLElement) {
		forSaleBy = element.Text
		fmt.Println("\tFor Sale By: ", forSaleBy)
	})

	c.OnHTML(".ad-meta--17Bqm", func(element *colly.HTMLElement) {
		metaData = element.Text
		fmt.Println("\tMeta: ", metaData)
	})


	c.OnScraped(func(r *colly.Response) {
		insert, err := db.Query("INSERT INTO car (district, category, title, price, url, postedOn, forSaleBy, meta) VALUES (?, ?, ?, ?, ?,?, ?, ?)", location, category, title, price, link, postedOn, forSaleBy, metaData)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
	})

	_ = c.Visit("https://ikman.lk/en/ads/"+location+"/"+category)

}
