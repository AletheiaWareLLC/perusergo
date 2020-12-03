/*
 * Copyright 2020 Aletheia Ware LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gif

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"github.com/AletheiaWareLLC/perusergo"
	"github.com/AletheiaWareLLC/perusergo/media"
	"image"
	"image/draw"
	"image/gif"
	"io"
	"time"
)

func init() {
	media.Register("image/gif", GIFDecoder)
}

func GIFDecoder(peruser perusergo.Peruser, name string, source io.ReadCloser) (fyne.CanvasObject, error) {
	defer source.Close()
	gif, err := gif.DecodeAll(source)
	if err != nil {
		return nil, err
	}
	size := gif.Image[0].Bounds().Size()
	overpaintImage := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))
	output := &canvas.Image{
		Image:    overpaintImage,
		FillMode: canvas.ImageFillOriginal,
	}
	go func() {
		draw.Draw(overpaintImage, overpaintImage.Bounds(), gif.Image[0], image.ZP, draw.Src)
		loop := func() {
			for c, srcImg := range gif.Image {
				draw.Draw(overpaintImage, overpaintImage.Bounds(), srcImg, image.ZP, draw.Over)
				canvas.Refresh(output)

				time.Sleep(time.Millisecond * time.Duration(gif.Delay[c]) * 10)
			}
		}
		switch gif.LoopCount {
		case 0:
			// Loop forever
			for {
				loop()
			}
		case -1:
			// Loop once
			loop()
		default:
			// Loop LoopCount+1 times
			for i := 0; i <= gif.LoopCount; i++ {
				loop()
			}
		}
	}()
	return output, nil
}
