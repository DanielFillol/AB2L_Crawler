package main

import (
	"fmt"
	"github.com/Darklabel91/AB2L_Crawler/Crawler"
)

const (
	AB2L     = "https://ab2l.org.br/radar-dinamico/"
	ShowMore = "//*[@id=\"mais\"]/div/div/a/span/span[2]"
	Company  = "//*[@id=\"listagem\"]/div/div/div/div"
	InitLink = "//*[@id=\"listagem\"]/div/div/div/div["
	EndLink  = "]/div/section/div[2]/div[2]/div/div[1]/div/div/div"
	Resume   = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[1]/div/div[2]/div/div[1]/div/div/div/h2"
)

func main() {
	driver, err := Crawler.SeleniumWebDriver()
	if err != nil {
		fmt.Println("Web driver not created")
	}

	err = driver.Get(AB2L)
	if err != nil {
		fmt.Println("ab2l web site offline")
	}

	Crawler.LoadWebPage(driver, ShowMore)

	err = Crawler.Craw(driver, Company, Resume, InitLink, EndLink)
	if err != nil {
		fmt.Println(err)
	}
}
