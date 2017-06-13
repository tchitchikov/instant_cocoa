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
	"time"
)

type Quandlv3 struct {
	APIKey    string
	EndDate   string
	Name      string
	Source    string
	StartDate string
}

// func (q Quandlv3) Data()

func (q Quandlv3) RawData() string {
	url := fmt.Sprintf("https://www.quandl.com/api/v3/datasets/%s/%s/data.csv?start_date=%s&end_date=%s&api_key=%s",
		q.Source, q.Name, q.StartDate, q.EndDate, q.APIKey)
	data, err := http.Get(url)
	if err != nil {
		log.Print("%v", err)
	}
	defer data.Body.Close()
	res, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Println(err)
	}
	return string(res)
}

// Data returns you a map of maps ending in float64 numbers and string dates as keys
func (q Quandlv3) Data() map[string]map[string]float64 {
	parser := csv.NewReader(strings.NewReader(q.RawData()))
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
		if record[0] == "Date" {
			for _, fieldName := range record {
				fieldValues = append(fieldValues, fieldName)
			}
		}
		if record[0] != "Date" {
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

// SortedKeys implements QuickSort to order the map keys (strings that match time.Time parsing) into a list for time series work
func (q Quandlv3) SortedKeys(data map[string]map[string]float64) []string {
	layout := "2006-01-02"
	dateList := []time.Time{}
	for key := range data {
		newVal, _ := time.Parse(layout, key)
		dateList = append(dateList, newVal)
	}

	firstHalf := []time.Time{}
	secondHalf := []time.Time{}
	initialVal := dateList[0]
	for _, val := range dateList {
		if val.Before(initialVal) {
			firstHalf = append(firstHalf, val)
		} else {
			secondHalf = append(secondHalf, val)
		}
	}
	output := []string{}
	quickSortFirstHalf := dateListSort(firstHalf)
	quickSortSecondHalf := dateListSort(secondHalf)
	output = append(output, convertTimeToString(quickSortFirstHalf)...)
	output = append(output, convertTimeToString(quickSortSecondHalf)...)
	return output

}

func convertTimeToString(timeList []time.Time) []string {
	layout := "2006-01-02"
	output := []string{}
	for _, val := range timeList {
		convertToString := val.Format(layout)
		output = append(output, convertToString)
	}
	return output
}

//dateListSort does a QuickSort on the keys using time.Time comparisons
func dateListSort(dateList []time.Time) []time.Time {
	outputList := []time.Time{}
	outputList = append(outputList, dateList[0])
	var newDateList []time.Time
	newDateList = dateList[1:]
	for _, outputVal := range newDateList {
		for iter, inputVal := range outputList {
			if outputVal.Before(inputVal) || outputVal.Equal(inputVal) {
				firstHalf := make([]time.Time, len(outputList[:iter]))
				secondHalf := make([]time.Time, len(outputList[iter:]))
				copy(firstHalf, outputList[:iter])
				copy(secondHalf, outputList[iter:])
				outputList = nil
				outputList = append(outputList, firstHalf...)
				outputList = append(outputList, outputVal)
				outputList = append(outputList, secondHalf...)
				firstHalf = nil
				secondHalf = nil
				break
			} else if iter == len(outputList)-1 {
				outputList = append(outputList, outputVal)
			}
		}
	}
	return outputList
}
