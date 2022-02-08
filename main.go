package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	ValCurs []Valute `xml:"Valute"`
}

type Valute struct {
	Text     string `xml:",chardata"`
	ID       string `xml:"ID,attr"`
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  string `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

func getXml(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("unable to read body '%s': %s", url, err)
	}
	var v ValCurs
	xml.Unmarshal(body, &v)

	result, err := FindVal(v, "R01820")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("Японский иен равен: %s\n", result.Value)

}

func FindVal(v ValCurs, id string) (*Valute, error) {
	for i := 0; i < len(v.ValCurs); i++ {
		if v.ValCurs[i].ID == id {
			return &v.ValCurs[i], nil
		}
	}
	return nil, errors.New("Не удалось найти валюту!")
}

func main() {
	URL := "https://www.cbr-xml-daily.ru/daily_utf8.xml"
	getXml(URL)
}
