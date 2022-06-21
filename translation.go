package main

var TRANSLATION = map[string]string {
	"GAME_IS_CLOSED": "Run the game.\n\n\nIf the problem persists: \nCouldn't get the process ID. \nMake sure you're running the game. \nMake sure you're running correct version. \nAlso make sure you're running\nthis cheat on supported OS.",
	"COULDNT_READ_BASE_ADDRESS": "Couldn't read base address. \nMake sure you're running the game. \nAlso make sure you're running\nthis cheat on supported OS.",
	"COULDNT_GET_HANDLE": "Couldn't get the process handle. \nMake sure you're running the game. \nAlso make sure you're running\nthis cheat on supported OS.",
	"UNKNOWN_LEVEL": "Awaiting game.\n\n\nThis screen appears during loading,\nif you see it in game, make sure\nyou're running correct version.",
	"LEVEL_LOADING": "Loading level...",
	"SELECT_LEVEL": "Please select the level.",
}

func t(key string) string {
	return TRANSLATION[key]
}