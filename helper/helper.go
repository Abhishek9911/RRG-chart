package helper

import (
	"RRG/model"
	"encoding/csv"
	"os"
	"strconv"
)

func FileImport(path string) ([]model.OHLC, error) {
	var data []model.OHLC
	file, err := os.Open(path)
	if err != nil {
		return data, err
	}
	defer file.Close()
	
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return data, err
	}

	for i:= len(records)-1; i>0;i-- {
		intValue := MultipleStrToFloat(records[i][1:])
		rec := model.OHLC{
			Date: records[i][0],
			Open: intValue[0],
			High: intValue[1],
			Low: intValue[2],
			Close: intValue[3],
		}
		data = append(data, rec)
	}
	// fmt.Print(data)	

	return data, nil
}

func MultipleStrToFloat(str []string) []float64 {
	var arr []float64
	for _, val := range str{
		value, err := strconv.ParseFloat(val, 64)
		if err != nil {}
		arr = append(arr, value)
	}
	return arr
}