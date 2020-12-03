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
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	_ "github.com/AletheiaWareLLC/perusergo/media/image"
	_ "github.com/AletheiaWareLLC/perusergo/media/image/gif"
	_ "github.com/AletheiaWareLLC/perusergo/media/image/jpeg"
	_ "github.com/AletheiaWareLLC/perusergo/media/image/png"
	_ "github.com/AletheiaWareLLC/perusergo/media/image/svg"
	_ "github.com/AletheiaWareLLC/perusergo/media/text/plain"
	"image/color"
)

const pageMinSize = 32

type Page interface {
	fyne.Widget
	ShowContent(fyne.CanvasObject)
	ShowError(error)
}

type page struct {
	widget.BaseWidget
	content fyne.CanvasObject
}

func NewPage() Page {
	p := &page{}
	p.ExtendBaseWidget(p)
	return p
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer.
func (p *page) CreateRenderer() fyne.WidgetRenderer {
	p.ExtendBaseWidget(p)
	return &pageRenderer{
		page: p,
	}
}

// MinSize returns the size that this widget should not shrink below.
func (p *page) MinSize() fyne.Size {
	p.ExtendBaseWidget(p)
	return p.BaseWidget.MinSize()
}

// ShowContent sets the root content of the page.
func (p *page) ShowContent(content fyne.CanvasObject) {
	p.content = content
	p.Refresh()
}

func (p *page) ShowError(e error) {
	message := widget.NewLabel(fmt.Sprintf("Error: %s", e.Error()))
	message.Wrapping = fyne.TextWrapBreak
	p.ShowContent(container.NewVScroll(message))
}

type pageRenderer struct {
	page *page
}

func (r *pageRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (r *pageRenderer) Destroy() {
}

func (r *pageRenderer) Layout(size fyne.Size) {
	if c := r.page.content; c != nil {
		c.Resize(size)
	}
}

func (r *pageRenderer) MinSize() fyne.Size {
	min := fyne.NewSize(pageMinSize, pageMinSize)
	if c := r.page.content; c != nil {
		min = c.MinSize()
	}
	return min
}

func (r *pageRenderer) Objects() (objects []fyne.CanvasObject) {
	if c := r.page.content; c != nil {
		objects = append(objects, c)
	}
	return
}

func (r *pageRenderer) Refresh() {
	if c := r.page.content; c != nil {
		c.Refresh()
	}
	canvas.Refresh(r.page)
}
