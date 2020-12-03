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

package ui_test

import (
	"fmt"
	"fyne.io/fyne/test"
	"fyne.io/fyne/widget"
	"github.com/AletheiaWareLLC/perusergo/ui"
	"testing"
)

func TestPage_Empty(t *testing.T) {
	test.NewApp()
	defer test.NewApp()
	page := ui.NewPage()
	page.Refresh() // Force layout
	window := test.NewWindow(page)
	defer window.Close()
	test.AssertImageMatches(t, "page/empty.png", window.Canvas().Capture())
}

func TestPage_ShowContent(t *testing.T) {
	test.NewApp()
	defer test.NewApp()
	page := ui.NewPage()
	page.ShowContent(widget.NewLabel("This is an example"))
	window := test.NewWindow(page)
	defer window.Close()
	test.AssertImageMatches(t, "page/content.png", window.Canvas().Capture())
}

func TestPage_ShowError(t *testing.T) {
	test.NewApp()
	defer test.NewApp()
	page := ui.NewPage()
	page.ShowError(fmt.Errorf("This is an example"))
	window := test.NewWindow(page)
	defer window.Close()
	test.AssertImageMatches(t, "page/error.png", window.Canvas().Capture())
}
