/*
Parsing URLs
*/ 

package main 

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// Parsing a URL that includes a scheme, auth info, host
	// port, path, query params, and query fragment
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parse it an ensure no errors.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Access the scheme
	fmt.Println(u.Scheme)

	// User contains all auth info; call Username and
	// Password on this for individual vals
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// Host Contains bot hostname and port, if present
	// use SplitHostPort to extract them
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Extract path and the framgent after #
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// queary params in a string of k=v format needs RawQuery
	// Parsed query param maps are from strings to slices of strings.
	// index [0] for the first value
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}