package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parse the URL and ensure there are no errors.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Accessing the scheme / protocol
	fmt.Println(u.Scheme)

	// User contains all authentication info; call Username and Password on this for individual values.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// Host contains both the hostname and the port, if present. Use SplitHostPort to extract them.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Extract the path and fragment after the #.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// Extract the query string after the ?.
	fmt.Println(u.RawQuery)
	// Parse into a map of key/value pairs.
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}
