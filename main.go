package main

import (
	"fmt"
	"get_cocoa"
	"log"

	sql "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("cocoa!")
	source := "CHRIS"
	name := "ICE_CC1"
	startDate := "2017-06-01"
	endDate := "2017-06-12"
	apiKey := getAPIKey()
	quandlData := get_cocoa.Quandlv3{
		APIKey:    apiKey,
		EndDate:   endDate,
		Name:      name,
		Source:    source,
		StartDate: startDate,
	}
	data := quandlData.Data()
	// fmt.Println(data)
	// rawData := quandlData.RawData()
	// fmt.Println(rawData)
	sortedKeys := quandlData.SortedKeys(data)
	fmt.Println(sortedKeys)

	googleData := get_cocoa.GoogleFinance{
		Ticker:    "CHOC",
		StartDate: "2016-01-01",
		EndDate:   "2016-06-13",
	}
	// fmt.Println(googleData.RawData())
	// gfData := googleData.Data()
	// sortedGFKeys := googleData.SortedKeys(gfData)

	// chocPlot := plotter.XYs{}
	// for key, value := range sortedGFKeys {
	// 	chocPlot = append(chocPlot, get_cocoa.PlotStruct{
	// 		X: float64(key),
	// 		Y: gfData[value]["Close"],
	// 	})
	// }

	// fmt.Println(chocPlot)
	// get_cocoa.Plot(chocPlot)
	result(googleData)

}

func result(f get_cocoa.FinancialData) {
	fmt.Println(f)
	fmt.Println(f.RawData())
}

// getAPIKey makes a call to a mysql database I have locally to get my APIKey
func getAPIKey() string {
	db, err := sql.Open("mysql", "jint-dev:@/Quandl")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT APIKEY FROM API_KEY")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var APIKEY string
	for rows.Next() {
		if err := rows.Scan(&APIKEY); err != nil {
			log.Fatal(err)
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return APIKEY
}
