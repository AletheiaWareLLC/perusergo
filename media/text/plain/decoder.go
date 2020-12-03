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

package plain

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/AletheiaWareLLC/perusergo"
	"github.com/AletheiaWareLLC/perusergo/media"
	"io"
	"io/ioutil"
)

func init() {
	media.Register("text/*", PlainDecoder)
	media.Register("text/plain", PlainDecoder)
}

func PlainDecoder(peruser perusergo.Peruser, name string, source io.ReadCloser) (fyne.CanvasObject, error) {
	label := widget.NewLabel("")
	label.Wrapping = fyne.TextWrapBreak
	defer source.Close()
	bytes, err := ioutil.ReadAll(source)
	if err != nil {
		return nil, err
	}
	label.SetText(string(bytes))
	return label, nil
}
