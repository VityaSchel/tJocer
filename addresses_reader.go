package main

// import "fmt"

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
	var Levels = map[uint32]string{
		25: "Freddy",
		41: "Bonnie",
		39: "Chica",
		26: "Foxy",
		18: "Menu",
		32: "Menu",
		14: "Loading level",
		20: "Loading level",
		35: "Loading level",
		38: "Loading level",
		44: "Loading level",
		50: "Loading level",
	}
	return Levels[selectedLevelID]
}

func GetUserPosition() (float32, float32, bool) {
	X, err1 := readMemoryAtFloat32(playerXAddress)
	Z, err2 := readMemoryAtFloat32(playerZAddress)
	return X, Z, err1 || err2
}