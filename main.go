package main

import (
	"fmt"
	"image"
	"image/color"
	"path"
	"strings"
	g "github.com/AllenDang/giu"
	"time"
)

const windowName string = "tJocer: The Joy of Creation: Reborn trainer"
const radarSize int = 1080
const windowSize int = 300

type LevelUnits struct {
	pixelsIn1080Square int
	unitsInSquare int
	xoffset int
	yoffset int
}
var mapUnits = map[string]LevelUnits {
	"Freddy": LevelUnits{
		pixelsIn1080Square: 165,
		unitsInSquare: 789,
		xoffset: -395,
		yoffset: -330,
	},
	"Bonnie": LevelUnits{
		pixelsIn1080Square: 55,
		unitsInSquare: 320,
		xoffset: -430,
		yoffset: -500,
	},
	"Chica": LevelUnits{
		pixelsIn1080Square: 55,
		unitsInSquare: 320,
		xoffset: -430,
		yoffset: -500,
	},
	"Foxy": LevelUnits{
		pixelsIn1080Square: 55,
		unitsInSquare: 320,
		xoffset: -430,
		yoffset: -500,
	},
}
func getUnits(level string) (float32, float32, int, int) {
	var units = mapUnits[level]
	var pixelsInRadarSquare = units.pixelsIn1080Square*radarSize/1080
	var xUnitsInSquare = units.unitsInSquare
	var zUnitsInSquare = units.unitsInSquare
	var pixelsInUnitX float32 = float32(pixelsInRadarSquare)/float32(xUnitsInSquare)
	var pixelsInUnitZ float32 = float32(pixelsInRadarSquare)/float32(zUnitsInSquare)
	var xoffset int = units.xoffset
	var yoffset int = units.yoffset
	return pixelsInUnitX, pixelsInUnitZ, xoffset, yoffset
}

var focused bool = false
var wnd g.MasterWindow

var textures map[string]*g.Texture = make(map[string]*g.Texture)

func main() {
	pid, pidFound := bindDefaultProcess()

	if(pidFound) {
		memoryReadInit(pid)
	}

	wnd := g.NewMasterWindow(
		windowName, windowSize, windowSize, 
		g.MasterWindowFlagsNotResizable|g.MasterWindowFlagsFloating|g.MasterWindowFlagsFrameless|g.MasterWindowFlagsTransparent,
	)
	
	wnd.SetBgColor(color.RGBA{0, 0, 0, 0})
	focused = true
	wnd.SetPos(0, 0)

	if(pidFound) {
		initUI()	
	}

	go refresh()
	go loadTextures()
	ChangeFocused(focused)
	fmt.Println("tJocer started!")
	wnd.Run(loop)
	
}

var texturesPaths []string = []string{"player.png", "maps/freddy.png", "maps/bonnie.png", "maps/foxy.png", "maps/chica.png"}
func loadTextures() {
	ticker := time.NewTicker(time.Second * 1)
	<-ticker.C

	var texturesReferences map[*image.RGBA]string = make(map[*image.RGBA]string)
	for _, textureFilePath := range texturesPaths {
		filepath := "./images/" + textureFilePath
		filepath = strings.ReplaceAll(filepath, "/",  "\\")

		img, _ := g.LoadImage(filepath)
		texturesReferences[img] = textureFilePath
		g.NewTextureFromRgba(img, func(tex *g.Texture) {
			texturePath := strings.TrimSuffix(texturesReferences[img], path.Ext(texturesReferences[img]))
			textures[texturePath] = tex
		})
	}

}

func ChangeFocused(newValue bool) {
	focused = newValue
	ModifyWindow(windowName, focused)
}