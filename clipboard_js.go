//go:build js && wasm

// The js/wasm case, alongside clipboard_darwin.go, _unix, _windows and
// _plan9. Upstream has no branch for a browser and no fallback, so the
// package does not compile for js/wasm at all — which is what stops
// bubbles and huh compiling, through textinput.

package clipboard

import "syscall/js"

// Writing is straightforward: navigator.clipboard.writeText returns a
// promise, and nothing here needs to wait for it.
func writeAll(text string) error {
	if cb := js.Global().Get("navigator").Get("clipboard"); cb.Truthy() {
		cb.Call("writeText", text)
		return nil
	}
	// No Clipboard API — keep the text so a paste in the same page still
	// sees what was copied.
	js.Global().Set("__clipboardText", text)
	return nil
}

// Reading cannot be done here.
//
// readText is asynchronous, and a wasm module has one thread: blocking it to
// await a promise deadlocks the page, since the promise can only settle from
// the event loop that is being blocked. The page has to prime the value from
// its own paste handler instead, which is where a browser hands over the
// clipboard without asking permission.
func readAll() (string, error) {
	if v := js.Global().Get("__clipboardText"); v.Type() == js.TypeString {
		return v.String(), nil
	}
	return "", nil
}
