package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/AB2L_Crawler/Crawler"
	"os"
	"path/filepath"
)

const (
	fileName   = "legalTechs"
	folderName = "Result"
)

func WriteCSV(legalTechs []Crawler.CompanyStruct) error {
	var rows [][]string

	rows = append(rows, generateHeaders())

	for _, company := range legalTechs {
		rows = append(rows, generateRow(company))
	}

	cf, err := createFile(folderName + "/" + fileName + ".csv")
	if err != nil {
		return err
	}

	defer cf.Close()

	w := csv.NewWriter(cf)

	err = w.WriteAll(rows)
	if err != nil {
		return err
	}

	return nil
}

// create csv file from operating system
func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

// generate the necessary headers for csv file
func generateHeaders() []string {
	return []string{
		"Nome",
		"Descritivo",
		"Serviços",
		"Fundadores",
		"Informações Legais",
		"Endereço",
		"Site",
		"Telefone",
		"Email",
	}
}

// returns []string that compose the row in the csv file
func generateRow(legalTech Crawler.CompanyStruct) []string {
	return []string{
		legalTech.Name,
		legalTech.GeneralInfo,
		legalTech.Service,
		legalTech.Founders,
		legalTech.PersonalData,
		legalTech.Address,
		legalTech.Site,
		legalTech.Phone,
		legalTech.Mail,
	}
}
