package Crawler

import (
	"github.com/tebeka/selenium"
	"time"
)

func waitXpath(driver selenium.WebDriver, XPATH string) {
	start := time.Now()
	for {
		ts := time.Since(start)
		elem, _ := driver.FindElements(selenium.ByXPATH, XPATH)
		if len(elem) > 0 || ts.Minutes() >= 1 {
			break
		}
	}
}
