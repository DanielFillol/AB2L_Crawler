package Crawler

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/tebeka/selenium"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

const (
	Company      = "//*[@id=\"listagem\"]/div/div/div/div"
	InitXpath    = "//*[@id=\"listagem\"]/div/div/div/div["
	EndXpath     = "]/div/section/div[2]/div[2]/div/div[1]/div/div/div/div"
	Wait         = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section"
	closeXpath   = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[2]/i"
	EndXpathData = "]/div/section/div[2]/div[2]/div/div[2]/div/div/div"
)

type CompanyStruct struct {
	Name         string
	GeneralInfo  string
	Service      string
	Founders     string
	PersonalData string
	Address      string
	Site         string
	Phone        string
	Mail         string
}

func Craw(driver selenium.WebDriver) ([]CompanyStruct, error) {
	loop, err := totalLoop(driver, Company)
	if err != nil {
		return nil, err
	}

	var companys []CompanyStruct
	for i := 0; i < loop; i++ {
		fmt.Println("Legal Tech: " + strconv.Itoa(i+1) + "/" + strconv.Itoa(loop))

		forward, err := verifyData(driver, i)
		if err != nil {
			return nil, err
		}

		if forward {
			openLink := InitXpath + strconv.Itoa(i+1) + EndXpath

			open(driver, openLink)

			document, err := getPageHtml(driver)
			if err != nil {
				return nil, err
			}

			company, err := getData(driver, document)
			if err != nil {
				return nil, err
			}

			companys = append(companys, company)

			close(driver)
		}

	}

	return companys, nil
}

//Calculates loop
func totalLoop(driver selenium.WebDriver, XpathCompany string) (int, error) {
	totalLinks, err := driver.FindElements(selenium.ByXPATH, XpathCompany)
	if err != nil {
		return 0, err
	}

	return len(totalLinks), nil
}

//Makes verification to avoid error in execution
func verifyData(driver selenium.WebDriver, i int) (bool, error) {
	titleXpath := InitXpath + strconv.Itoa(i+1) + EndXpathData

	elemTitle, err := driver.FindElement(selenium.ByXPATH, titleXpath)
	if err != nil {
		return false, err
	}
	title, err := elemTitle.Text()

	if title != "No data was found" {
		return true, nil
	} else {
		return false, nil
	}
}

//Interact with link to open and close
func bttClick(driver selenium.WebDriver, XpathLink string) error {
	btt, err := driver.FindElement(selenium.ByXPATH, XpathLink)
	if err != nil {
		return err
	}
	btt.Click()
	return nil
}

func open(driver selenium.WebDriver, openLink string) {
	bttClick(driver, openLink)
	bttClick(driver, openLink)
	bttClick(driver, openLink)
	waitXpath(driver, Wait)
}

func close(driver selenium.WebDriver) {
	bttClick(driver, closeXpath)
	bttClick(driver, closeXpath)
	bttClick(driver, closeXpath)
}

//get the page source
func getPageHtml(driver selenium.WebDriver) (*html.Node, error) {
	pageSourceCode, err := driver.PageSource()
	if err != nil {
		return nil, err
	}

	document, err := htmlquery.Parse(strings.NewReader(pageSourceCode))
	if err != nil {
		return nil, err
	}

	return document, nil
}
