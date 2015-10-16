# GoJSON

JSON parser in Go.

```
$ echo '{ "hello": { "wow" : "value", "yo": 1 }, "arr" : [1, 2, 3, 4.423423] }' | ./go-json
2015/10/12 23:21:32 map[hello:map[wow:value yo:1] arr:[1 2 3 4.423423]]
```

## ParseJSON

You can use the `ParseJSON` function by itself:

```
impot "github.com/thejsj/go-json/parser"
parsedJSON, err := parser.ParseJSON("{ \"hello\": 1 }")
```

- Save to JSON file
- Pretty print to STDOUT
- STDERR if there's an error
- Switch away from `interface {}`
