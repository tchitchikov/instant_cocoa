package get_cocoa

import "time"

// FinancialData is the interface that the methods in google.go and quandl.go will use
type FinancialData interface {
	RawData() string
	Data() map[string]map[string]float64
	// SortedKeys(string, map[string]map[string]float64) []string
}

// SortedKeys implements QuickSort to order the map keys (strings that match time.Time parsing) into a list for time series work
func SortedKeys(fd FinancialData, layout string) []string {
	dateList := []time.Time{}
	for key := range fd.Data() {
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
	output = append(output, convertTimeToString(layout, quickSortFirstHalf)...)
	output = append(output, convertTimeToString(layout, quickSortSecondHalf)...)
	return output
}
