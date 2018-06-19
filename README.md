# goroyale
[![GoDoc](https://godoc.org/github.com/Altarrel/goroyale?status.svg)](https://godoc.org/github.com/Altarrel/goroyale)
[![Go report](https://goreportcard.com/badge/github.com/Altarrel/goroyale)](https://goreportcard.com/report/github.com/Altarrel/goroyale)

A Golang wrapper for the Clash Royale API at https://royaleapi.com/.

## Installing
If you have Go installed you can run this command.
```sh
go get github.com/Altarrel/goroyale
```

## Example
```golang
package main

import (
	"fmt"

	"github.com/Altarrel/goroyale"
)

var token = "API KEY GOES HERE"

func main() {
	c, err := goroyale.New(token, 0) // 0 will use the default request timeout of 10 seconds
	if err != nil {
		fmt.Println(err)
		return
	}

	ver, err := c.GetAPIVersion()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("API Version:", ver)
	}

	params := map[string][]string{
		"exclude": {"name"},
	}
	p, err := c.GetPlayer("8L9L9GL", params)
	if err != nil {
		fmt.Println(err)
	} else {
		// will just print "Name:" as p.Name is "" because it was excluded
		// more info about this at https://docs.royaleapi.com/#/field_filter
		fmt.Println("Name:", p.Name)

		fmt.Println("Tag:", p.Tag)
		fmt.Println("Clan:", p.Clan.Name)
	}
}
```

## Ratelimits
If you hit the RoyaleAPI ratelimit, the lib will just refuse to run your request and return an error of type `goroyale.RatelimitError`.
You can handle this with a type switch on the error.
```golang
switch e := err.(type) {
case goroyale.RatelimitError:
	// RatelimitError.RetryAfter is a time.Duration to wait before sending another request to the API.
	time.Sleep(e.RetryAfter)
}
```