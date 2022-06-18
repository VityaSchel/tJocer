package main

import (
	// "fmt"
	"image"
	"image/color"
	"time"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

func loop() {
	imgui.PushStyleVarFloat(imgui.StyleVarWindowBorderSize, 0)
	g.PushColorFrameBg(color.RGBA{10, 10, 10, 0})
	// if focused {
	// 	g.PushColorWindowBg(color.RGBA{50, 50, 50, 128})
	// } else {
		g.PushColorWindowBg(color.RGBA{10, 10, 10, 0})
	// }
	g.SingleWindow().Layout(
		g.InputInt(&xoffset),
		g.InputInt(&yoffset),
		g.Custom(func() {
			canvas := g.GetCanvas()
			
			playerX := int(readMemoryAt(0x1FF7A9497C4)*pixelsInUnitX)
			// fmt.Println(readMemoryAtByte8(0x1FF7A778278))
			GetAddresses()
			playerZ := int(readMemoryAt(0x1C8180A9948)*pixelsInUnitZ)

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

func refresh() {
	ticker := time.NewTicker(time.Millisecond * 50)

	for {
		g.Update()

		<-ticker.C
	}
}