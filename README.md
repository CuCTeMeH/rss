# RSS Reader package
RSS reader package, which parses concurrently multiple RSS/Atom feeds.
- To install package: `go get -u github.com/CuCTeMeH/rss/reader`.
- In order to run tests and install dependencies go inside package directory and run `make all`
- The package has only exportable RssItem struct and func Parse, which takes slice of url strings and returns a slice of RssItem structs.
