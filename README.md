go-dom
------

**This package is a work in progress**

The `go-dom` package exposes the [Web API Spec](https://spec.whatwg.org/) as a Go package.

This package is mostly just wrappers around calls to [syscall/js](https://tip.golang.org/pkg/syscall/js/)
to expose the browser Web API in a Go friendly manner.

This package's code is generated from [WebIDL](https://heycam.github.io/webidl/) definitions for the various Web API Specs.

## Examples

These examples can be compiled and run using `GOARCH=wasm GOOS=js` as described at https://github.com/golang/go/wiki/WebAssembly

Find all `div` elements and print them to the console

```go
package main

import (
    "github.com/brettlangdon/go-dom/console"
    "github.com/brettlangdon/go-dom/document"
)

func main() {
    nodes := document.QuerySelectorAll("div")
    for i := 0; i < nodes.GetLength(); i++ {
        console.Dir(nodes.Item(0))
    }
}
```

Add a `click` handler to `document** which prints the event information.

```go
package main

import (
    dom "github.com/brettlangdon/go-dom"
    "github.com/brettlangdon/go-dom/console"
    "github.com/brettlangdon/go-dom/document"
)

func main() {
    loop := dom.NewLoop()

    document.AddEventListener("click", dom.NewEventHandler(func(evt dom.Event) {
        console.Dir(evt)
    }))

    loop.Loop()
}
```

This last example uses a `dom.Loop` which is used to keep the program running instead of exiting.

If the Go program exits then the callback handlers will no longer execute.

Calling `loop.Stop()` will cause the Go program to exit.
