package main

import (
	// "fmt"
	"image"
	"image/color"
	"strings"
	"time"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

var (
	playerXAddress int64
	playerZAddress int64
)

func initUI() {
	playerXAddress, 
	playerZAddress, 
	baseAddress = GetAddresses()
}

func loop() {
	imgui.PushStyleVarFloat(imgui.StyleVarWindowBorderSize, 0)
	g.PushColorFrameBg(color.RGBA{10, 10, 10, 0})
	// if focused {
	// 	g.PushColorWindowBg(color.RGBA{50, 50, 50, 128})
	// } else {
		g.PushColorWindowBg(color.RGBA{10, 10, 10, 0})
	// }
	g.SingleWindow().Layout(
		// g.InputInt(&xoffset),
		// g.InputInt(&yoffset),
		g.Custom(func() {
			canvas := g.GetCanvas()

			unableToRender := false
			if(baseAddress == 0) {
				unableToRender = true
			} else {
				val, err := readMemoryAtByte8(baseAddress)
				if(err || val == 0) {
					unableToRender = true
				}
			}

			if(unableToRender) {
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

				pid, success := bindDefaultProcess()
				if(!success) {
					drawText("Run the game.\n\n\n If the problem persists: \nCouldn't get the process ID. \nMake sure you're running the game. \nMake sure you're running correct version. \nAlso make sure you're running\nthis cheat on supported OS.")
					return
				}

				baseAddress_, err := memoryReadInit(pid)
				if(err == "NO_HANDLE") {
					drawText("Couldn't get the process handle. \nMake sure you're running the game. \nAlso make sure you're running\nthis cheat on supported OS.")
					return
				} else if (err == "BASE_ADDRESS_NOT_FOUND") {
					drawText("Couldn't read base address. \nMake sure you're running the game. \nAlso make sure you're running\nthis cheat on supported OS.")
					return
				} else if (err == "") {
					baseAddress = baseAddress_
					initUI()
				}
				
				return
			}
			
			XFloat, _ := readMemoryAt(playerXAddress)
			ZFloat, _ := readMemoryAt(playerZAddress)
			playerX := int(XFloat*pixelsInUnitX)
			playerZ := int(ZFloat*pixelsInUnitZ)

			playerSpriteSize := 20
			renderImage(
				canvas,
				textures["player"], 
				image.Pt(windowSize/2 - playerSpriteSize/2, windowSize/2 - playerSpriteSize/2), 
				image.Pt(playerSpriteSize, playerSpriteSize),
				// 1,
				// 1,
			)

			renderImage(
				canvas,
				textures["maps/bonnie"], 
				image.Pt(playerZ+int(xoffset), -playerX+int(yoffset)), 
				image.Pt(radarSize, radarSize),
			)
		}),
	)
	g.PopStyleColor()
	g.PopStyleColor()
	imgui.PopStyleVar()
	ChangeFocused(focused)
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