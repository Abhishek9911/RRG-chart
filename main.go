package main

import (
	"fmt"
	// "os"
	"RRG/helper"
)

func main(){
	niftyPath := "data/NIFTY50.csv"
	accPath := "data/ACC.csv"
	siemensPath := "data/siemens.csv"
	suzlonPath := "data/Suzlon.csv"
	wareeEPath := "data/WareeE.csv"
	wareeTPath := "data/WareeT.csv"


	indices, err := helper.FileImport(niftyPath)
	if err != nil {
		fmt.Print("main (main.go) unable to extract data from file: ", err, "Path: ", niftyPath)
	}
	
	acc, err := helper.FileImport(accPath)
	if err != nil {
		fmt.Print("main (main.go) unable to extract data from file: ", err, "Path: ", accPath)
	}

	siemens, err := helper.FileImport(siemensPath)
	if err != nil {
		fmt.Print("main (main.go) unable to extract data from file: ", err, "Path: ", siemensPath)
	}

	suzlon, err := helper.FileImport(suzlonPath)
	if err != nil {
		fmt.Print("main (main.go) unable to extract data from file: ", err, "Path: ", suzlonPath)
	}

	wareeE, err := helper.FileImport(wareeEPath)
	if err != nil {
		fmt.Print("main (main.go) unable to extract data from file: ", err, "Path: ", wareeEPath)
	}

	wareeT, err := helper.FileImport(wareeTPath)
	if err != nil {
		fmt.Print("main (main.go) unable to extract data from file: ", err, "Path: ", wareeTPath)
	}
	fmt.Print(len(suzlon), len(wareeT), len(wareeE), len(siemens), len(acc), len(indices))
	// lets go forward tomorrow
	 
}