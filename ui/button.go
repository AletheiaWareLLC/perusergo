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
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"image/color"
)

type IconButton interface {
	fyne.Widget
	fyne.Tappable
}

type iconButton struct {
	widget.BaseWidget
	Icon      fyne.Resource
	IsEnabled func() bool
	OnTapped  func()
	hovered   bool
}

func newIconButton(icon fyne.Resource, enabled func() bool, tapped func()) IconButton {
	b := &iconButton{
		Icon:      icon,
		IsEnabled: enabled,
		OnTapped:  tapped,
	}
	b.ExtendBaseWidget(b)
	return b
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer.
func (b *iconButton) CreateRenderer() fyne.WidgetRenderer {
	b.ExtendBaseWidget(b)
	i := &canvas.Image{
		Resource: b.Icon,
		FillMode: canvas.ImageFillContain,
	}
	r := &iconButtonRenderer{
		button:  b,
		image:   i,
		objects: []fyne.CanvasObject{container.NewPadded(i)},
	}
	r.updateImage()
	return r
}

// MinSize returns the size that this widget should not shrink below.
func (b *iconButton) MinSize() fyne.Size {
	b.ExtendBaseWidget(b)
	return b.BaseWidget.MinSize()
}

// MouseIn is called when a desktop pointer enters the widget
func (b *iconButton) MouseIn(*desktop.MouseEvent) {
	if b.enabled() {
		b.hovered = true
		b.Refresh()
	}
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (b *iconButton) MouseMoved(*desktop.MouseEvent) {
}

// MouseOut is called when a desktop pointer exits the widget
func (b *iconButton) MouseOut() {
	if b.hovered || b.enabled() {
		b.hovered = false
		b.Refresh()
	}
}

// Tapped is called when a pointer tapped event is captured and triggers any tap handler
func (b *iconButton) Tapped(*fyne.PointEvent) {
	if b.OnTapped != nil && b.enabled() {
		b.OnTapped()
	}
}

func (b *iconButton) enabled() bool {
	if b.IsEnabled != nil {
		return b.IsEnabled()
	}
	return true
}

type iconButtonRenderer struct {
	button       *iconButton
	image        *canvas.Image
	objects      []fyne.CanvasObject
	disabledIcon fyne.Resource
}

func (r *iconButtonRenderer) BackgroundColor() color.Color {
	if r.button.hovered {
		return theme.HoverColor()
	}
	return theme.BackgroundColor()
}

func (r *iconButtonRenderer) Destroy() {
}

func (r *iconButtonRenderer) Layout(size fyne.Size) {
	r.objects[0].Resize(size)
}

func (r *iconButtonRenderer) MinSize() fyne.Size {
	return r.objects[0].MinSize()
}

func (r *iconButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *iconButtonRenderer) Refresh() {
	r.updateImage()
	r.image.Refresh()
	canvas.Refresh(r.button)
}

func (r *iconButtonRenderer) updateImage() {
	res := r.button.Icon
	if !r.button.enabled() {
		if r.disabledIcon == nil {
			r.disabledIcon = theme.NewDisabledResource(res)
		}
		res = r.disabledIcon
	}
	r.image.Resource = res
	s := theme.IconInlineSize()
	r.image.SetMinSize(fyne.NewSize(s, s))
}
