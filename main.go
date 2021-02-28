package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main(){
	c := colly.NewCollector()
	c.OnHTML(".img-crop .lazy[data-srcset]:first-of-type", func(e *colly.HTMLElement){
		//runs 3rd or error
		link := e.Attr("data-srcset")
		fmt.Println("Link: ",link)

		//e.Request.Visit(link)
	})

	c.OnRequest(func(r *colly.Request){
		//runs first
		fmt.Println("Visiting",r.URL)
	} )

	c.OnError(func(_ *colly.Response, err error){
		//runs 3rd or worked onHTML
		fmt.Println("Error:", err)
	})

	c.OnResponse(func(r *colly.Response){
		//runs 2nd
		fmt.Println("Visited",r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response){
		//runs last
		fmt.Println("Finished",r.Request.URL)
	})


	//runs at start
	c.Visit("https://www.awwwards.com/awwwards/collections/image-gallery-and-slideshows/")
	//attempt to scrape saved HTML doc to get my filter right
	//c.Visit("file:///Users/g3n6i/Desktop/HTML/Image-Galleries-and-SlideShows---Awwwards.html")
}

