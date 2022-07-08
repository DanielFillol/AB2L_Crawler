package Crawler

import "github.com/tebeka/selenium"

func LoadWebPage(driver selenium.WebDriver, Xpath string) {
	for {
		elem, _ := driver.FindElements(selenium.ByXPATH, Xpath)
		elem[0].Click()
		if len(elem) == 0 {
			break
		}
	}
}
