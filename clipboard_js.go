// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js

package clipboard

import "syscall/js"

// writeAll hands the text to navigator.clipboard.writeText without waiting for
// its promise. Without the Clipboard API it keeps the text for readAll.
func writeAll(text string) error {
	if cb := js.Global().Get("navigator").Get("clipboard"); cb.Truthy() {
		cb.Call("writeText", text)
		return nil
	}
	js.Global().Set("__clipboardText", text)
	return nil
}

// readAll cannot await readText on the single wasm thread, so it returns the
// last text written here, or what the page's paste handler stored.
func readAll() (string, error) {
	if v := js.Global().Get("__clipboardText"); v.Type() == js.TypeString {
		return v.String(), nil
	}
	return "", nil
}
