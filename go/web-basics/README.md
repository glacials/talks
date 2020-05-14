# Go Web Basics talk

This is my talk that introduces you to Go from the context of a web service
architecture.

The talk is built with Go's [`present`][present-cmd] tool, with an entrypoint of
`main.slide`.

[present-cmd]: https://pkg.go.dev/golang.org/x/tools/cmd/present

## Running it yourself

The talk requires Go to run and view.

```sh
go install golang.org/x/tools/cmd/present
present
```

Then open [localhost:3999][localhost] in your browser.

[localhost]: http://localhost:3999

### Usage
Use the arrow keys, enter, or space to navigate with the keyboard. You can also
click the margins of the slides to go back and forward.

Depending on your screen width, parts of the previous and next slides may
reveal too much for your tastes. You can fix this by increasing the font size
of the page (Ctrl+ or ⌘+).

For some slides, a Run button will appear in the bottom right of the code
snippet. You can click this to run the code. For slides whose examples span
multiple files, be sure to run the file with `func main()` in it.

For some other slides, such as the `gofmt` slide, the demo takes place outside
the slideshow. In these slides, I've placed the command you should paste into
your terminal in order to execute the demonstration.

### Editing
See the documentation for the [`present`][present-pkg] package (not to be
confused with the `present` command linked above) for details on syntax of
`.slide` files.

[present-pkg]: https://pkg.go.dev/golang.org/x/tools/present
