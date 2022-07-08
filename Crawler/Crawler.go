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
	Company           = "//*[@id=\"listagem\"]/div/div/div/div"
	InitXpath         = "//*[@id=\"listagem\"]/div/div/div/div["
	EndXpath          = "]/div/section/div[2]/div[2]/div/div[1]/div/div/div"
	closeXpath        = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[2]/i"
	nameXpath         = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[1]/div/div[2]/div/div[1]/div/div/div/h2"
	generalXpath      = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[2]"
	servicesXpath     = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[3]"
	founderXpath      = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[4]"
	personalDataXpath = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[5]"
	contactXpath      = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[6]"
)

type CompanyStruct struct {
	Name         string
	GeneralInfo  string
	Service      string
	Founders     string
	PersonalData string
	Contact      string
}

func Craw(driver selenium.WebDriver) ([]CompanyStruct, error) {
	loop, err := totalLoop(driver, Company)
	if err != nil {
		return nil, err
	}

	var companys []CompanyStruct
	for i := 1; i <= loop; i++ {
		fmt.Println("Legal Tech: " + strconv.Itoa(i) + "/" + strconv.Itoa(loop))
		openLink := InitXpath + strconv.Itoa(i) + EndXpath

		err = bttClick(driver, openLink)
		if err != nil {
			return nil, err
		}

		waitXpath(driver, nameXpath)

		document, err := getPageHtml(driver)
		if err != nil {
			return nil, err
		}

		company := getData(document)
		companys = append(companys, company)

		err = bttClick(driver, closeXpath)
		if err != nil {
			return nil, err
		}
	}

	return companys, nil
}

func totalLoop(driver selenium.WebDriver, XpathCompany string) (int, error) {
	totalLinks, err := driver.FindElements(selenium.ByXPATH, XpathCompany)
	if err != nil {
		return 0, err
	}

	return len(totalLinks), nil
}

func bttClick(driver selenium.WebDriver, Xpathlink string) error {
	btt, err := driver.FindElement(selenium.ByXPATH, Xpathlink)
	if err != nil {
		return err
	}

	btt.Click()

	return nil
}

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

func getData(document *html.Node) CompanyStruct {
	var name string
	names := htmlquery.Find(document, nameXpath)
	if len(names) > 0 {
		name = htmlquery.InnerText(htmlquery.FindOne(document, nameXpath))
	}

	var general string
	generals := htmlquery.Find(document, generalXpath)
	if len(generals) > 0 {
		general = htmlquery.InnerText(htmlquery.FindOne(document, generalXpath))
	}

	var service string
	services := htmlquery.Find(document, servicesXpath)
	if len(services) > 0 {
		service = htmlquery.InnerText(htmlquery.FindOne(document, servicesXpath))
	}

	var founder string
	founders := htmlquery.Find(document, founderXpath)
	if len(founders) > 0 {
		founder = htmlquery.InnerText(htmlquery.FindOne(document, founderXpath))
	}

	var personalData string
	personaldatas := htmlquery.Find(document, personalDataXpath)
	if len(personaldatas) > 0 {
		personalData = htmlquery.InnerText(htmlquery.FindOne(document, personalDataXpath))
	}

	var contact string
	contacts := htmlquery.Find(document, contactXpath)
	if len(contacts) > 0 {
		contact = htmlquery.InnerText(htmlquery.FindOne(document, contactXpath))
	}

	return CompanyStruct{
		Name:         name,
		GeneralInfo:  general,
		Service:      service,
		Founders:     founder,
		PersonalData: personalData,
		Contact:      contact,
	}

}
