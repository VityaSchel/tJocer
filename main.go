package main

import (
	"fmt"
	"image"
	"image/color"
	"path"
	"strings"
	"time"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

const windowName string = "tJocer: The Joy of Creation: Reborn trainer"
const radarSize int = 1080
const windowSize int = 300

var focused bool = false
var wnd g.MasterWindow

var textures map[string]*g.Texture = make(map[string]*g.Texture)

func loop() {
	imgui.PushStyleVarFloat(imgui.StyleVarWindowBorderSize, 0)
	g.PushColorFrameBg(color.RGBA{10, 10, 10, 0})
	if focused {
		g.PushColorWindowBg(color.RGBA{50, 50, 50, 128})
	} else {
		g.PushColorWindowBg(color.RGBA{10, 10, 10, 0})
	}
	g.SingleWindow().Layout(
		g.Custom(func() {
			canvas := g.GetCanvas()

			playerSpriteSize := 16
			renderImage(
				canvas,
				textures["player"], 
				image.Pt(windowSize/2 - playerSpriteSize/2, windowSize/2 - playerSpriteSize/2), 
				image.Pt(playerSpriteSize, playerSpriteSize),
			)

			renderImage(
				canvas,
				textures["maps/bonnie"], 
				image.Pt(0, 0), 
				image.Pt(radarSize, radarSize),
			)
		}),
	)
	g.PopStyleColor()
	g.PopStyleColor()
	imgui.PopStyleVar()
	ChangeFocused(focused)
}

func refresh() {
	ticker := time.NewTicker(time.Second * 1)

	for {
		g.Update()

		<-ticker.C
	}
}

func main() {
	bindDefaultProcess()
	wnd := g.NewMasterWindow(
		windowName, windowSize, windowSize, 
		g.MasterWindowFlagsNotResizable|g.MasterWindowFlagsFloating|g.MasterWindowFlagsFrameless|g.MasterWindowFlagsTransparent,
	)
	
	wnd.SetBgColor(color.RGBA{0, 0, 0, 0})
	focused = false
	wnd.SetPos(0, 0)

	go refresh()
	go loadTextures()
	fmt.Println("tJocer started!")
	wnd.Run(loop)
	
}

var texturesPaths []string = []string{"player.png", "maps/bonnie.png"}
func loadTextures() {
	for _, textureFilePath := range texturesPaths {
		filepath := "./images/" + textureFilePath
		filepath = strings.ReplaceAll(filepath, "/",  "\\")

		img, _ := g.LoadImage(filepath)
		g.NewTextureFromRgba(img, func(tex *g.Texture) {
			texturePath := strings.TrimSuffix(textureFilePath, path.Ext(textureFilePath))
			textures[texturePath] = tex
		})
	}

}

func ChangeFocused(newValue bool) {
	focused = newValue
	ModifyWindow(windowName, focused)
	// if focused {
	// 	wnd.SetBgColor(color.NRGBA{0, 0, 0, 128})
	// } else {
	// 	wnd.SetBgColor(color.NRGBA{0, 0, 0, 0})
	// }
}