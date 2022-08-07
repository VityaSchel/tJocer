package main

import (
	// "fmt"
	"image"
	"image/color"
	"strings"
	"time"

	// "strconv"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

var (
	selectedLevel string
)

func initUI() {
	InitAddressesManager()
}

var LevelsMaps = map[string]string {
	"Freddy": "maps/freddy",
	"Bonnie": "maps/bonnie",
	"Chica": "maps/chica",
	"Foxy": "maps/foxy",
}

var pixelsIn1080Square int32 = 55
var unitsInSquare int32 = 320
var xoffset int32 = 0
var yoffset int32 = 0

func loop() {
	imgui.PushStyleVarFloat(imgui.StyleVarWindowBorderSize, 0)
	g.PushColorFrameBg(color.RGBA{10, 10, 10, 0})
	// if focused {
	// 	g.PushColorWindowBg(color.RGBA{50, 50, 50, 128})
	// } else {
		g.PushColorWindowBg(color.RGBA{10, 10, 10, 0})
	// }

	// var pixelsInUnitX, pixelsInUnitZ, xOffset, yOffset = getUnits(selectedLevel)
	var pixelsInRadarSquare = int(pixelsIn1080Square)*radarSize/1080
	var xUnitsInSquare = unitsInSquare
	var zUnitsInSquare = unitsInSquare
	var pixelsInUnitX float32 = float32(pixelsInRadarSquare)/float32(xUnitsInSquare)
	var pixelsInUnitZ float32 = float32(pixelsInRadarSquare)/float32(zUnitsInSquare)

	g.SingleWindow().Layout(
		g.InputInt(&pixelsIn1080Square),
		g.InputInt(&unitsInSquare),
		g.InputInt(&xoffset),
		g.InputInt(&yoffset),
		g.Custom(func() {
			canvas := g.GetCanvas()

			if(!RunChecks(canvas)) {
				return
			}
			
			X, Z, _ := GetUserPosition()
			playerX := int(X*pixelsInUnitX)
			playerZ := int(Z*pixelsInUnitZ)

			playerSpriteSize := 20
			renderImage(
				canvas,
				textures["player"], 
				image.Pt(windowSize/2 - playerSpriteSize/2, windowSize/2 - playerSpriteSize/2), 
				image.Pt(playerSpriteSize, playerSpriteSize),
			)

			renderImage(
				canvas,
				textures[LevelsMaps[selectedLevel]], 
				image.Pt(-playerZ+int(xoffset), playerX+int(yoffset)), 
				image.Pt(radarSize, radarSize),
			)
		}),
	)
	g.PopStyleColor()
	g.PopStyleColor()
	imgui.PopStyleVar()
}

func RunChecks(canvas *g.Canvas) bool {
	drawText := func(text string) {
		text = "tJocer:\n" + text
		canvas.AddRectFilled(
			image.Pt(0, 0),
			image.Pt(
				windowSize, 
				len(strings.Split(text, "\n"))*18+10,
			), 
			color.RGBA{0, 0, 0, 180}, 
			1, 
			g.DrawFlagsRoundCornersAll,
		)
		canvas.AddText(image.Pt(10, 10), color.White, text)
	}

	// if(baseAddress == 0) {
	// 	drawText(t("GAME_IS_CLOSED"))
	// 	return false
	// }

	pid, success := bindDefaultProcess()
	if(!success) {	
		drawText(t("GAME_IS_CLOSED"))
		baseAddress = 0
		return false
	}

	baseAddress_, err := memoryReadInit(pid)
	if(err == "NO_HANDLE") {
		drawText(t("COULDNT_GET_HANDLE"))
		return false
	} else if (err == "BASE_ADDRESS_NOT_FOUND") {
		drawText(t("COULDNT_READ_BASE_ADDRESS"))
		return false
	} else if (err == "" && baseAddress == 0) {
		baseAddress = baseAddress_
		initUI()
	}

	val, errReading := readMemoryAtByte8(baseAddress)
	if(errReading || val == 0) {
		drawText(t("GAME_IS_CLOSED"))
		return false
	}

	selectedLevel = GetSelectedLevel()
	switch selectedLevel {
		case "":
			drawText(t("UNKNOWN_LEVEL"))
			return false

		case "Loading level":
			drawText(t("LEVEL_LOADING"))
			return false

		case "Menu":
			drawText(t("SELECT_LEVEL"))
			return false
	}

	if(LevelsMaps[selectedLevel] == "") {
		drawText(t("NO_MAP"))
		return false
	}

	return true
}

func renderImage(canvas *g.Canvas, texture *g.Texture, pos1 image.Point, pos2 image.Point) {
	if(texture != nil) {
		canvas.AddImage(texture, pos1, pos1.Add(pos2))
	}
}

func refresh() {
	ticker := time.NewTicker(time.Millisecond * 50)

	for {
		g.Update()

		<-ticker.C
	}
}