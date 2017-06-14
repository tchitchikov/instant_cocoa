package get_cocoa

import "time"

func convertTimeToString(layout string, timeList []time.Time) []string {
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
