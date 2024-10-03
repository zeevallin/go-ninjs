# ninjs (Go)
Go library for parsing and producing [ninjs](https://iptc.org/standards/ninjs/).

A standard developed by the [IPTC (International Press Telecommunications Council)](https://iptc.org/standards/ninjs/)

> One standard, countless use cases.
>
> ninjs standardises the representation of news in JSON – a lightweight, easy-to-parse, data interchange format.
>
> Almost every media organisation today is working on a project (or ten) that involves expressing a news item in JSON. From content APIs to mobile applications to document databases, JSON has moved into a position of technical prominence across the enterprise. In 2012, we went around the table at an IPTC meeting and realised nearly all of us had our own internal approaches for serialising news data in JSON. And so a standard was born. Approved and released in October 2013, ninjs provides key properties and structures required to represent news and publishing information.


## Examples

### Parsing

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/zeevallin/go-ninjs/2.1/ninjs"
)

const payload = `
{
    "uri": "https://github.com/zeevallin/go-ninjs/2.1/examples/parse-single",
    "version": "2",
    "versioncreated": "2024-10-02T16:02:28+02:00",
    "pubstatus": "usable",
    "urgency": 1,
    "copyrightholder": "Zee Vieira",
    "copyrightnotice": "Creative Commons BY 4.0",
    "language": "en",
    "bodies": [
        {
            "role": "web",
            "contenttype": "text/html",
            "value": "<p>Today I published a ninjs library in Go. It is a simple library to work with ninjs objects in Go.</p>"
        },
        {
            "role": "terminal",
            "contenttype": "text/plain",
            "value": "Today I published a ninjs library in Go. It is a simple library to work with ninjs objects in Go."
        }
    ],
    "headlines": [
        {
            "role": "main",
            "value": "Published a ninjs library in Go"
        }
    ],
    "by": "Zee Vieira"
}
`

func main() {
	var item ninjs.Item
	err := json.Unmarshal([]byte(payload), &item)
	if err != nil {
		panic(err)
	}

	fmt.Printf("headline: %s\n", item.Headlines[0].Value)
	fmt.Printf("body: %s\n", item.Bodies[1].Value)
	fmt.Printf("author: %s\n", item.By)
	fmt.Printf("url: %s\n", item.URI)
	fmt.Printf("version: %s\n", item.Version)
}

// headline: Published a ninjs library in Go
// body: Today I published a ninjs library in Go. It is a simple library to work with ninjs objects in Go.
// author: Zee Vieira
// url: https://github.com/zeevallin/go-ninjs/2.1/examples/parse-single
// version: 2
```