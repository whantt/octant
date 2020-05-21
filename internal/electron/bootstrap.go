/*
 * Copyright (c) 2020 the Octant contributors. All Rights Reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package electron

import (
	"fmt"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func initWindows(a *astilectron.Astilectron, appURL string, listener MessageListener, logger astikit.SeverityLogger) ([]*astilectron.Window, error) {
	height := 750
	width := 1200
	windowMinWidth := 768
	windowMinHeight := windowMinWidth * 10 / 16 // base min height off ultra wide ratio

	window, err := a.NewWindow(appURL, &astilectron.WindowOptions{
		Center:        astikit.BoolPtr(true),
		Height:        astikit.IntPtr(height),
		Width:         astikit.IntPtr(width),
		MinWidth:      astikit.IntPtr(windowMinWidth),
		MinHeight:     astikit.IntPtr(windowMinHeight),
		TitleBarStyle: astikit.StrPtr("hiddenInset"),
	})

	if err != nil {
		return nil, fmt.Errorf("create main window: %w", err)
	}

	if listener != nil {
		window.OnMessage(handleMessage(window, listener, logger))
	}

	windows := []*astilectron.Window{
		window,
	}

	return windows, nil
}

func initMenuItems(window *astilectron.Window) ([]*astilectron.MenuItemOptions, error) {
	menuItems := []*astilectron.MenuItemOptions{
		{
			Label: astikit.StrPtr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astikit.StrPtr("Quit"),
					Role:  astilectron.MenuItemRoleQuit,
				},
			},
		},
		{
			Label: astikit.StrPtr("View"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astikit.StrPtr("Develop"),
					SubMenu: []*astilectron.MenuItemOptions{
						{
							Label:       astikit.StrPtr("Developer Tools"),
							Accelerator: astilectron.NewAccelerator("CommandOrControl", "Option", "I"),
							OnClick: func(e astilectron.Event) (deleteListener bool) {
								if err := window.OpenDevTools(); err != nil {
									// TODO: do something if this fails
								}

								return
							},
						},
					},
				},
			},
		},
	}

	return menuItems, nil
}
