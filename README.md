> **Fork.** This is [atotto/clipboard](https://github.com/atotto/clipboard)
> with one file added: `clipboard_js.go`, the `js && wasm` case. Upstream has
> darwin, unix, windows and plan9 branches and no fallback, so the package
> does not compile for `js/wasm` at all — which is what stops `bubbles` and
> `huh` compiling, through `textinput`.
>
> Writing goes to `navigator.clipboard`. Reading cannot: `readText` is async
> and a wasm module has one thread, so awaiting it deadlocks the page. The
> host primes the value from its own paste handler instead.
>
> The module path is `github.com/0magnet/clipboard` so it can be a `replace`
> target.

[![Build Status](https://travis-ci.com/atotto/clipboard.svg?branch=master)](https://travis-ci.com/atotto/clipboard)

[![GoDoc](https://godoc.org/github.com/atotto/clipboard?status.svg)](http://godoc.org/github.com/atotto/clipboard)

# Clipboard for Go

Provide copying and pasting to the Clipboard for Go.

Build:

    $ go get github.com/atotto/clipboard

Platforms:

* OSX
* Windows 7 (probably work on other Windows)
* Linux, Unix (requires 'xclip' or 'xsel' command to be installed)


Document: 

* http://godoc.org/github.com/atotto/clipboard

Notes:

* Text string only
* UTF-8 text encoding only (no conversion)

TODO:

* Clipboard watcher(?)

## Commands:

paste shell command:

    $ go install github.com/atotto/clipboard/cmd/gopaste@latest
    $ # example:
    $ gopaste > document.txt

copy shell command:

    $ go install github.com/atotto/clipboard/cmd/gocopy@latest
    $ # example:
    $ cat document.txt | gocopy



