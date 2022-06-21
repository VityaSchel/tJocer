package main

import (
	"strconv"
)

func sumHexSSS(aHex string, bHex string) string {
	aDecimal, _ := strconv.ParseInt(aHex, 16, 0)
	bDecimal, _ := strconv.ParseInt(bHex, 16, 0)
	resultDecimal := aDecimal + bDecimal
	resultHex := strconv.FormatInt(resultDecimal, 16)
	return resultHex
}

func sumHexSSI(aHex string, bHex string) int64 {
	aDecimal, _ := strconv.ParseInt(aHex, 16, 0)
	bDecimal, _ := strconv.ParseInt(bHex, 16, 0)
	resultDecimal := aDecimal + bDecimal
	return resultDecimal
}

func sumHexIIS(aDecimal int64, bDecimal int64) string {
	resultDecimal := aDecimal + bDecimal
	resultHex := strconv.FormatInt(resultDecimal, 16)
	return resultHex
}

func sumHexIII(aDecimal int64, bDecimal int64) int64 {
	return aDecimal + bDecimal
}

func sumHexSIS(aHex string, bDecimal int64) string {
	aDecimal, _ := strconv.ParseInt(aHex, 16, 0)
	resultDecimal := aDecimal + bDecimal
	resultHex := strconv.FormatInt(resultDecimal, 16)
	return resultHex
}

func sumHexISS(aDecimal int64, bHex string) string {
	return sumHexSIS(bHex, aDecimal)
}

func sumHexSII(aHex string, bDecimal int64) int64 {
	aDecimal, _ := strconv.ParseInt(aHex, 16, 0)
	resultDecimal := aDecimal + bDecimal
	return resultDecimal
}

func sumHexISI(aDecimal int64, bHex string) int64 {
	return sumHexSII(bHex, aDecimal)
}