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

package main

import (
	"flag"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/AletheiaWareLLC/perusergo/cache"
	"github.com/AletheiaWareLLC/perusergo/network"
	"github.com/AletheiaWareLLC/perusergo/ui"
	"net/url"
	"os/user"
)

var address = flag.String("address", "", "Peruse address")

func main() {
	flag.Parse()
	a := app.New()
	w := a.NewWindow("Peruser")
	c := cache.NewMemoryCache()
	n := network.NewTCPNetwork()
	b := ui.NewBrowser(a, w, c, n)
	s := *address
	if s == "" {
		if usr, err := user.Current(); err == nil {
			s = "file://" + usr.HomeDir
		} else {
			s = "https://aletheiaware.com"
		}
	}
	u, err := url.Parse(s)
	if err != nil {
		b.SetError(err)
	} else {
		b.SetAddress(u)
	}
	w.SetContent(b.CanvasObject())
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
