package Crawler

import (
	"fmt"
	"github.com/tebeka/selenium"
	"strconv"
)

func Craw(driver selenium.WebDriver, XpathCompany string, XpathResume, InitLink string, EndLink string) error {
	totalLinks, err := driver.FindElements(selenium.ByXPATH, XpathCompany)
	if err != nil {
		return err
	}

	for i := 1; i <= len(totalLinks); i++ {
		finalLink := InitLink + strconv.Itoa(i) + EndLink

		btt, err := driver.FindElement(selenium.ByXPATH, finalLink)
		if err != nil {
			fmt.Println("could not find " + finalLink)
		}
		btt.Click()

		text, err := driver.FindElement(selenium.ByXPATH, XpathResume)

		fmt.Println(text)

	}

	return nil
}
