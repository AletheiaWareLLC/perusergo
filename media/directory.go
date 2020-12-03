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

package media

import (
	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/storage"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/AletheiaWareLLC/perusergo"
	//"sort"
	"net/url"
)

func DirectoryDecoder(peruser perusergo.Peruser, uri fyne.ListableURI) fyne.CanvasObject {
	// TODO Create new Type implementing CanvasObject
	//  - Only show parent button if uri has parent
	//  - Query uri children once per refresh, sort alphabetically with directories before files
	/* Sort URIs
	sort.Slice(uris, func(i, j int) bool {
		// TODO sort directories before files
		return uris[i].String() < uris[j].String()
	})
	*/
	parent := widget.NewButtonWithIcon("Parent Directory", theme.MoveUpIcon(), func() {
		p, err := storage.Parent(uri)
		if err != nil {
			peruser.SetError(err)
			return
		}
		u, err := url.Parse(p.String())
		if err != nil {
			peruser.SetError(err)
			return
		}
		peruser.SetAddress(u)
	})
	list := widget.NewList(
		func() int {
			list, err := uri.List()
			if err != nil {
				peruser.SetError(err)
				return 0
			}
			return len(list)
		},
		func() fyne.CanvasObject {
			label := widget.NewLabel("Template Object")
			label.Wrapping = fyne.TextTruncate
			icon := widget.NewFileIcon(nil)
			return fyne.NewContainerWithLayout(layout.NewHBoxLayout(), icon, label)
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			list, err := uri.List()
			if err != nil {
				peruser.SetError(err)
				return
			}
			if id < 0 && id >= len(list) {
				return
			}
			u := list[id]
			c := item.(*fyne.Container)
			c.Objects[0].(*widget.FileIcon).SetURI(u)
			c.Objects[1].(*widget.Label).SetText(u.Name())
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		list, err := uri.List()
		if err != nil {
			peruser.SetError(err)
			return
		}
		if id < 0 && id >= len(list) {
			return
		}
		uri := list[id]
		u, err := url.Parse(uri.String())
		if err != nil {
			peruser.SetError(err)
			return
		}
		peruser.SetAddress(u)
	}
	list.ExtendBaseWidget(list)
	return container.NewBorder(parent, nil, nil, nil, list)
}
