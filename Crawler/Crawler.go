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
	Wait                = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section"
	Company             = "//*[@id=\"listagem\"]/div/div/div/div"
	InitXpath           = "//*[@id=\"listagem\"]/div/div/div/div["
	EndXpath            = "]/div/section/div[2]/div[2]/div/div[1]/div/div/div/div"
	EndXpathData        = "]/div/section/div[2]/div[2]/div/div[2]/div/div/div"
	closeXpath          = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[2]/i"
	sectionXpath        = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section"
	nameXpath           = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[1]/div/div[2]/div/div[1]/div/div/div/h2"
	generaXpath         = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[2]/div/div/div/div[1]/div/div/div/div"
	servicesXpath       = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[3]/div/div/div/div[3]/div/div/div/div"
	servicesInitXpath   = "//*[@id=\"jet-toggle-control-260"
	servicesEndXpath    = "\"]/div"
	founderXpath        = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[4]"
	founderInitXpath    = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[4]/div/div/div/div[3]/div/div/div/div["
	founderEndXpath     = "]/div/section/div/div[2]/div/div[1]/div/div/div/div"
	personalDataXpath   = "//*[@id=\"jet-toggle-content-2554\"]/div"
	contactXpath        = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[6]"
	contactXpathAddress = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[6]/div/div[1]/div/div[3]/div/div/div[2]/div/span"
	contactXpathSite    = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[6]/div/div[1]/div/div[4]/div/div/div[2]/div/a"
	contactXpathPhone   = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[6]/div/div[1]/div/div[5]/div/div/div[2]/div/span"
	contactXpathMail    = "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[6]/div/div[1]/div/div[6]/div/div/div[2]/div/span"
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

			continueTest, err := verifySection(driver)
			if err != nil {
				return nil, err
			}

			if continueTest {
				document, err := getPageHtml(driver)
				if err != nil {
					return nil, err
				}

				company := getData(document)
				companys = append(companys, company)
			}

			close(driver)
		}

	}

	return companys, nil
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

//Calculates loop
func totalLoop(driver selenium.WebDriver, XpathCompany string) (int, error) {
	totalLinks, err := driver.FindElements(selenium.ByXPATH, XpathCompany)
	if err != nil {
		return 0, err
	}

	return len(totalLinks), nil
}

//makes verification to avoid error in execution
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

func verifySection(driver selenium.WebDriver) (bool, error) {
	elemSection, err := driver.FindElements(selenium.ByXPATH, sectionXpath)
	if err != nil {
		return false, err
	}

	if len(elemSection) >= 6 {
		return true, nil
	} else {
		return false, nil
	}

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

//get the needed data
func getData(document *html.Node) CompanyStruct {
	var name string
	names := htmlquery.Find(document, nameXpath)
	if len(names) > 0 {
		name = htmlquery.InnerText(htmlquery.FindOne(document, nameXpath))
	}

	var general string
	generals := htmlquery.Find(document, generaXpath)
	if len(generals) > 0 {
		for i := 0; i < 4; i++ {
			general = htmlquery.InnerText(htmlquery.FindOne(document, generaXpath))
		}
	}

	var service string
	services := htmlquery.Find(document, servicesXpath)
	if len(services) > 0 {
		for i := 0; i < len(services); i++ {
			service += htmlquery.InnerText(htmlquery.FindOne(document, servicesInitXpath+strconv.Itoa(i+1)+servicesEndXpath)) + " | "
		}
	}

	var founder string
	founders := htmlquery.Find(document, founderXpath)
	if len(founders) > 0 {
		for i := 0; i < len(founders); i++ {
			founder += htmlquery.InnerText(htmlquery.FindOne(document, founderInitXpath+strconv.Itoa(i+1)+founderEndXpath)) + " | "
		}
	}

	var personalData string
	personalDatas := htmlquery.Find(document, personalDataXpath)
	if len(personalDatas) > 0 {
		personalData = htmlquery.InnerText(htmlquery.FindOne(document, personalDataXpath))
	}

	var address string
	var site string
	var phone string
	var mail string
	contacts := htmlquery.Find(document, contactXpath)
	if len(contacts) > 0 {
		address = htmlquery.InnerText(htmlquery.FindOne(document, contactXpathAddress))
		site = htmlquery.InnerText(htmlquery.FindOne(document, contactXpathSite))
		phone = htmlquery.InnerText(htmlquery.FindOne(document, contactXpathPhone))
		mail = htmlquery.InnerText(htmlquery.FindOne(document, contactXpathMail))
	}

	return CompanyStruct{
		Name:         name,
		GeneralInfo:  general,
		Service:      service,
		Founders:     founder,
		PersonalData: personalData,
		Address:      strings.TrimSpace(address),
		Site:         strings.TrimSpace(site),
		Phone:        strings.TrimSpace(phone),
		Mail:         strings.TrimSpace(mail),
	}

}
