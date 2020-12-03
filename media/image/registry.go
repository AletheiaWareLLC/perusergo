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
	"fyne.io/fyne/canvas"
	"github.com/AletheiaWareLLC/perusergo"
	"io"
)

type Decoder = func(perusergo.Peruser, string, io.ReadCloser) (*canvas.Image, error)

var decoders = map[string]Decoder{}

// Lookup returns the Decoder associated with the given Media Type.
func Lookup(mime string) (Decoder, bool) {
	d, ok := decoders[mime]
	return d, ok
}

// Register associates the given Media Type with given Decoder.
func Register(mime string, decoder Decoder) {
	decoders[mime] = decoder
}
