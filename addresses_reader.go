package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

var (
	playerXAddress int64
	playerZAddress int64
)

func InitAddressesManager() {
	playerXAddress,
		playerZAddress,
		baseAddress = GetAddresses()
}

func GetSelectedLevel() string {
	selectedLevelAddress := sumHexISI(baseAddress, "257B93C")
	selectedLevelID, _ := readMemoryAtByte4(selectedLevelAddress)
	fmt.Println("selectedLevelID", selectedLevelID)
	var Levels = map[uint32]string{
		18: "Menu",
		32: "Menu",
		25: "Freddy",
		24: "Foxy",
		39: "Chica",
		41: "Bonnie",
	}
	loadedLevel := Levels[selectedLevelID]
	var LoadingLevelIDs = []uint32{
		6, 8, 11, 12, 13, 14, 16, 20, 21, 22, 23, 27, 30, 33, 35, 36, 37, 38, 42, 44, 47, 50,
	}
	
	if(loadedLevel == "" && slices.Contains(LoadingLevelIDs, selectedLevelID)) {
		return "Loading level"
	} else {
		return loadedLevel
	}
}

func GetUserPosition() (float32, float32, bool) {
	X, err1 := readMemoryAtFloat32(playerXAddress)
	Z, err2 := readMemoryAtFloat32(playerZAddress)
	return X, Z, err1 || err2
}