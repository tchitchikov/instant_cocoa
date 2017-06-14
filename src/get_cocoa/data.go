package get_cocoa

// FinancialData is the interface that the methods in google.go and quandl.go will use
type FinancialData interface {
	RawData() string
	Data() map[string]map[string]float64
	SortedKeys(map[string]map[string]float64) []string
}
