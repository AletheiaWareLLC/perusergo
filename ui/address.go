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

package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/AletheiaWareLLC/perusergo"
	"net/url"
)

type AddressBar interface {
	fyne.CanvasObject
	SetAddress(string)
}

type addressBar struct {
	*widget.Entry // TODO SelectEntry to show suggestions and history
	peruser       perusergo.Peruser
}

func NewAddressBar(peruser perusergo.Peruser) AddressBar {
	l := &addressBar{
		Entry:   &widget.Entry{},
		peruser: peruser,
	}
	l.ExtendBaseWidget(l)
	return l
}

func (l *addressBar) SetAddress(a string) {
	l.SetText(a)
}

// TypedKey receives key presses to handle enter/return.
func (l *addressBar) TypedKey(e *fyne.KeyEvent) {
	switch e.Name {
	case fyne.KeyEnter, fyne.KeyReturn:
		u, err := url.Parse(l.Text)
		if err != nil {
			l.peruser.SetError(err)
			return
		}
		l.peruser.SetAddress(u)
	default:
		l.Entry.TypedKey(e)
	}
}
