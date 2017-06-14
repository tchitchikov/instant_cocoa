package get_cocoa

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type GoogleFinance struct {
	Ticker    string
	StartDate string
	EndDate   string
}

func (gf GoogleFinance) RawData() string {
	url := fmt.Sprintf("https://www.google.com/finance/historical?q=%s&startdate=%s&enddate=%s&output=csv",
		gf.Ticker, gf.StartDate, gf.EndDate)
	data, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer data.Body.Close()
	res, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Println(err)
	}
	return string(res)
}

func (gf GoogleFinance) Data() map[string]map[string]float64 {
	parser := csv.NewReader(strings.NewReader(gf.RawData()))
	output := make(map[string]map[string]float64)
	orderID := 0.0
	fieldValues := []string{}
	for {
		record, err := parser.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(record[0], "Date") {
			for _, fieldName := range record {
				fieldValues = append(fieldValues, fieldName)
			}
		}
		if !strings.Contains(record[0], "Date") {
			innerMap := make(map[string]float64)
			innerMap["ORDER_ID"] = orderID
			for iter, fieldName := range fieldValues {
				floatData, err := strconv.ParseFloat(record[iter], 64)
				if err != nil {
					floatData = 0.0
				}
				innerMap[fieldName] = floatData
			}
			output[record[0]] = innerMap
		}
		orderID = orderID + 1.0
	}
	return output
}
