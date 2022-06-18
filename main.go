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

const pixelsIn1080Square = 55
const pixelsInRadarSquare = pixelsIn1080Square*radarSize/1080
const xUnitsInSquare = 320
const zUnitsInSquare = 320
const pixelsInUnitX float32 = float32(pixelsInRadarSquare)/float32(xUnitsInSquare)
const pixelsInUnitZ float32 = float32(pixelsInRadarSquare)/float32(zUnitsInSquare)

var focused bool = false
var wnd g.MasterWindow
var xoffset int32 = -430
var yoffset int32 = -500

var textures map[string]*g.Texture = make(map[string]*g.Texture)

func main() {
	bindDefaultProcess()
	wnd := g.NewMasterWindow(
		windowName, windowSize, windowSize, 
		g.MasterWindowFlagsNotResizable|g.MasterWindowFlagsFloating|g.MasterWindowFlagsFrameless|g.MasterWindowFlagsTransparent,
	)
	
	wnd.SetBgColor(color.RGBA{0, 0, 0, 0})
	focused = true
	wnd.SetPos(0, 0)

	initUI()

	go refresh()
	go loadTextures()
	fmt.Println("tJocer started!")
	wnd.Run(loop)
	
}

var texturesPaths []string = []string{"player.png", "maps/bonnie.png"}
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

func renderImage(canvas *g.Canvas, texture *g.Texture, pos1 image.Point, pos2 image.Point) {
	if(texture != nil) {
		canvas.AddImage(texture, pos1, pos1.Add(pos2))
	}
}