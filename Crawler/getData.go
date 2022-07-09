package Crawler

import (
	"github.com/antchfx/htmlquery"
	"github.com/tebeka/selenium"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

const (
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

// TODO: The code bellow is to messy, needs to improve reading
//get the needed data
func getData(driver selenium.WebDriver, document *html.Node) (CompanyStruct, error) {
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
	pdts, err := driver.FindElements(selenium.ByXPATH, "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[3]/div/div/div/div[2]/div/div")
	if err != nil {
		return CompanyStruct{}, err
	}
	if len(pdts) != 0 {
		title, err := pdts[0].Text()
		if err != nil {
			return CompanyStruct{}, err
		}

		if title == "Produtos & Serviços" {
			service = getServices(document)
		}

	}

	var founder string
	fdr, err := driver.FindElements(selenium.ByXPATH, "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[4]/div/div/div/div[2]/div/div")
	if err != nil {
		return CompanyStruct{}, err
	}
	if len(fdr) != 0 {
		title, err := fdr[0].Text()
		if err != nil {
			return CompanyStruct{}, err
		}

		if title == "Sócios" {
			founder = getPartners(document)
		}
	}

	var personalData string
	perData, err := driver.FindElements(selenium.ByXPATH, "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[5]/div/div/div/div[2]/div/div")
	if err != nil {
		return CompanyStruct{}, err
	}
	if len(perData) != 0 {
		title, err := perData[0].Text()
		if err != nil {
			return CompanyStruct{}, err
		}

		if title == "Informações Extras" {
			personalData = getGeneralData(document)
		}
	}

	var address string
	var site string
	var phone string
	var mail string
	contact1, err := driver.FindElements(selenium.ByXPATH, "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[6]/div/div[1]/div/div[2]/div/div")
	if err != nil {
		return CompanyStruct{}, err
	}
	contact2, err := driver.FindElements(selenium.ByXPATH, "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[3]/div/div[1]/div/div[2]/div/div")
	if err != nil {
		return CompanyStruct{}, err
	}
	if len(contact1) != 0 {
		title, err := contact1[0].Text()
		if err != nil {
			return CompanyStruct{}, err
		}
		if title == "Endereço e Contato" {
			address, site, phone, mail = getAddress(document)
		}

	}
	if len(contact2) != 0 {
		if len(contact2) != 0 {
			title, err := contact2[0].Text()
			if err != nil {
				return CompanyStruct{}, err
			}
			if title == "Endereço e Contato" {
				address, site, phone, mail = getAddress2(driver, document)
			}

		}

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
	}, nil

}

func getServices(document *html.Node) string {
	var service string
	services := htmlquery.Find(document, servicesXpath)
	if len(services) > 0 {
		for i := 0; i < len(services); i++ {
			service += htmlquery.InnerText(htmlquery.FindOne(document, servicesInitXpath+strconv.Itoa(i+1)+servicesEndXpath)) + " | "
		}
	}
	return service
}

func getPartners(document *html.Node) string {
	var founder string
	founders := htmlquery.Find(document, founderXpath)
	if len(founders) > 0 {
		for i := 0; i < len(founders); i++ {
			founder += htmlquery.InnerText(htmlquery.FindOne(document, founderInitXpath+strconv.Itoa(i+1)+founderEndXpath)) + " | "
		}
	}
	return founder
}

func getGeneralData(document *html.Node) string {
	var personalData string
	personalDatas := htmlquery.Find(document, personalDataXpath)
	if len(personalDatas) > 0 {
		personalData = htmlquery.InnerText(htmlquery.FindOne(document, personalDataXpath))
	}
	return personalData
}

func getAddress(document *html.Node) (string, string, string, string) {
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
	return address, site, phone, mail
}

func getAddress2(driver selenium.WebDriver, document *html.Node) (string, string, string, string) {
	var address string
	var site string
	var phone string
	var mail string
	contacts := htmlquery.Find(document, "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[3]/div")
	if len(contacts) > 0 {
		address = htmlquery.InnerText(htmlquery.FindOne(document, "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[3]/div/div[1]/div/div[3]/div/div/div[2]/div/span"))
		ww, _ := driver.FindElements(selenium.ByXPATH, "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[3]/div/div[1]/div/div[4]/div/div/div[2]/div/a")
		if len(ww) != 0 {
			site = htmlquery.InnerText(htmlquery.FindOne(document, "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[3]/div/div[1]/div/div[4]/div/div/div[2]/div/a"))
		}
		phone = htmlquery.InnerText(htmlquery.FindOne(document, "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[3]/div/div[1]/div/div[5]/div/div/div[2]/div/span"))
		mail = htmlquery.InnerText(htmlquery.FindOne(document, "//*[@id=\"jet-popup-5641\"]/div/div[2]/div[1]/div[2]/div/section[3]/div/div[1]/div/div[6]/div/div/div[2]/div/span"))
	}
	return address, site, phone, mail
}
