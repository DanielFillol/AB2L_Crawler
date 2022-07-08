package main

import (
	"fmt"
	"github.com/Darklabel91/AB2L_Crawler/CSV"
	"github.com/Darklabel91/AB2L_Crawler/Crawler"
)

const AB2L = "https://ab2l.org.br/radar-dinamico/"

func main() {
	driver, err := Crawler.SeleniumWebDriver()
	if err != nil {
		fmt.Println(err)
	}

	defer driver.Close()

	err = driver.Get(AB2L)
	if err != nil {
		fmt.Println(err)
	}

	Crawler.LoadWebPage(driver)

	legalTechs, err := Crawler.Craw(driver)
	if err != nil {
		fmt.Println(err)
	}

	err = CSV.WriteCSV(legalTechs)
	if err != nil {
		fmt.Println(err)
	}

}
