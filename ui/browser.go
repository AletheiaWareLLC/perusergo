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
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/storage"
	"fyne.io/fyne/theme"
	"github.com/AletheiaWareLLC/perusergo"
	"github.com/AletheiaWareLLC/perusergo/media"
	"io"
	"log"
	"net/url"
)

var _ (perusergo.Peruser) = (*Browser)(nil)

type Browser struct {
	app        fyne.App
	window     fyne.Window
	cache      perusergo.Cache
	network    perusergo.Network
	back       IconButton
	next       IconButton
	refresh    IconButton
	settings   IconButton
	addressbar AddressBar
	page       Page
	address    *url.URL
	history    []*url.URL
	index      int
}

func NewBrowser(a fyne.App, w fyne.Window, c perusergo.Cache, n perusergo.Network) *Browser {
	b := &Browser{
		app:     a,
		window:  w,
		cache:   c,
		network: n,
	}
	b.back = newIconButton(theme.NavigateBackIcon(), b.IsBackAvailable, b.Back)
	b.next = newIconButton(theme.NavigateNextIcon(), b.IsNextAvailable, b.Next)
	b.refresh = newIconButton(theme.ViewRefreshIcon(), nil, b.Refresh)
	b.settings = newIconButton(theme.SettingsIcon(), nil, b.Settings)
	b.addressbar = NewAddressBar(b)
	b.page = NewPage()
	return b
}

func (b *Browser) Address() *url.URL {
	return b.address
}

func (b *Browser) Back() {
	b.index--
	a := b.history[b.index-1]
	b.setAddress(a)
}

func (b *Browser) CanvasObject() fyne.CanvasObject {
	return container.NewBorder(
		container.NewBorder(
			nil,
			nil, // TODO progress bar?
			container.NewHBox(
				b.back,
				b.next,
				b.refresh,
			),
			container.NewHBox(
				b.settings,
			),
			b.addressbar,
		),
		nil, // TODO status bar?
		nil, // TODO page structure?
		nil, // TODO developer/debugger tools?
		b.page)
}

func (b *Browser) Get(address *url.URL) (string, io.ReadCloser, error) {
	// TODO retrieve from cache, or network
	return b.network.Get(address.String())
}

func (b *Browser) IsBackAvailable() bool {
	return b.index > 1
}

func (b *Browser) IsNextAvailable() bool {
	return b.index < len(b.history)
}

func (b *Browser) Next() {
	a := b.history[b.index]
	b.index++
	b.setAddress(a)
}

func (b *Browser) Refresh() {
	// Refresh by resetting address
	a := b.address
	b.address = nil
	b.SetAddress(a)
}

func (b *Browser) SetAddress(u *url.URL) {
	log.Println("Opening:", u)
	b.pushHistory(u)
	b.setAddress(u)
}

func (b *Browser) SetError(e error) {
	log.Println("Error:", e)
	b.page.ShowError(e)
}

func (b *Browser) SetTitle(t string) {
	log.Println("Browser.SetTitle:", t)
	if t != "" {
		b.window.SetTitle("Peruser - " + t)
	} else {
		b.window.SetTitle("Peruser")
	}
}

func (b *Browser) Settings() {
	// TODO
	log.Println("Settings not yet implemented")
}

func (b *Browser) pushHistory(address *url.URL) {
	if b.index > 0 && b.history[b.index-1] == address {
		return
	}
	b.history = append(b.history[:b.index], address)
	b.index++
}

func (b *Browser) setAddress(address *url.URL) {
	if address == nil {
		return
	}
	if address == b.address {
		// No change
		return
	}
	b.address = address
	a := address.String()

	// Show progress dialog
	// TODO add stop button
	progress := dialog.NewProgressInfinite("Loading", a, b.window)
	progress.Show()
	defer progress.Hide()

	// Update buttons
	b.back.Refresh()
	b.next.Refresh()
	b.refresh.Refresh()
	b.settings.Refresh()

	// Update address bar
	b.addressbar.SetAddress(a)

	switch address.Scheme {
	case "file":
		u := storage.NewURI(a)
		b.SetTitle(u.Name())
		if l, err := storage.ListerForURI(u); err == nil {
			b.page.ShowContent(media.DirectoryDecoder(b, l))
		} else {
			mime := u.MimeType()
			decoder, ok := media.Lookup(mime)
			if !ok {
				b.SetError(fmt.Errorf("Unrecognized Media Type: %s", mime))
				return
			}
			source, err := fyne.CurrentApp().Driver().FileReaderForURI(u)
			if err != nil {
				b.SetError(err)
				return
			}
			root, err := decoder(b, a, source)
			if err != nil {
				b.SetError(err)
				return
			}
			b.page.ShowContent(root)
		}
	case "http", "https":
		b.SetTitle(a)
		mime, source, err := b.Get(address)
		if err != nil {
			b.SetError(err)
			return
		}
		decoder, ok := media.Lookup(mime)
		if !ok {
			b.SetError(fmt.Errorf("Unrecognized Media Type: %s", mime))
			return
		}
		root, err := decoder(b, a, source)
		if err != nil {
			b.SetError(err)
			return
		}
		b.page.ShowContent(root)
	case "peruser":
		// TODO handle app destinations
	case "source":
		/* TODO use PlainDecoder to show source in plain text
		p.ShowContent(plain.PlainDecoder(b, u))
		*/
	default:
		b.SetError(fmt.Errorf("Unrecognized URI Scheme: %s", address.Scheme))
	}
}
