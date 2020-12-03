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

package image

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"github.com/AletheiaWareLLC/perusergo"
	"github.com/AletheiaWareLLC/perusergo/media"
	"image"
	"io"
)

func init() {
	Register("image/*", CanvasImageDecoder)
	media.Register("image/*", ImageDecoder)
}

func ImageDecoder(peruser perusergo.Peruser, name string, source io.ReadCloser) (fyne.CanvasObject, error) {
	image, err := CanvasImageDecoder(peruser, name, source)
	return image, err
}

func CanvasImageDecoder(peruser perusergo.Peruser, name string, source io.ReadCloser) (*canvas.Image, error) {
	defer source.Close()
	i, _, err := image.Decode(source)
	if err != nil {
		return nil, err
	}
	return &canvas.Image{
		Image:    i,
		FillMode: canvas.ImageFillOriginal,
	}, nil
}
