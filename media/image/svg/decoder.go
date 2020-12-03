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
	"github.com/AletheiaWareLLC/perusergo/media/image"
	"io"
	"io/ioutil"
)

func init() {
	image.Register("image/svg", svgDecoder)
	image.Register("image/svg+xml", svgDecoder)
	media.Register("image/svg", SVGDecoder)
	media.Register("image/svg+xml", SVGDecoder)
}

func SVGDecoder(peruser perusergo.Peruser, name string, source io.ReadCloser) (fyne.CanvasObject, error) {
	image, err := svgDecoder(peruser, name, source)
	return image, err
}

func svgDecoder(peruser perusergo.Peruser, name string, source io.ReadCloser) (*canvas.Image, error) {
	defer source.Close()
	bytes, err := ioutil.ReadAll(source)
	if err != nil {
		return nil, err
	}
	return &canvas.Image{
		Resource: fyne.NewStaticResource(name, bytes),
		FillMode: canvas.ImageFillOriginal,
	}, nil
}
