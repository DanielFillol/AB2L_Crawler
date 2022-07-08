package Crawler

import (
	"fmt"
	"github.com/tebeka/selenium"
)

const ShowMore = "//*[@id=\"mais\"]/div/div/a/span/span[2]"

func LoadWebPage(driver selenium.WebDriver) {
	fmt.Println("Loading pages")
	for {
		elem, _ := driver.FindElements(selenium.ByXPATH, ShowMore)
		elem[0].Click()
		text, _ := elem[0].Text()
		if text != "MOSTRAR MAIS EMPRESAS" {
			fmt.Println("All pages loaded")
			break
		}
	}
}
