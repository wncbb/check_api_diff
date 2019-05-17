package tdurl

import (
	"fmt"
	"net/url"
)

type URL struct {
	// [scheme:][//[userinfo@]host][/]path[?query][#fragment]
	u      url.URL
	values url.Values
}

func NewURL() *URL {
	u := &URL{
		u:      url.URL{},
		values: make(map[string][]string),
	}
	return u
}

func (u *URL) SetScheme(scheme string) *URL {
	u.u.Scheme = scheme
	return u
}

func (u *URL) SetUserinfo(username, password string, passwordSet bool) *URL {
	if passwordSet {
		u.u.User = url.UserPassword(username, password)
	} else {
		u.u.User = url.User(username)
	}
	return u
}

func (u *URL) SetHost(host string) *URL {
	u.u.Host = host
	return u
}

func (u *URL) SetPath(path string) *URL {
	u.u.Path = path
	return u
}

func (u *URL) SetQueryValue(k, v string) *URL {
	u.values.Set(k, v)
	return u
}

func (u *URL) AddQueryValue(k, v string) *URL {
	u.values.Add(k, v)
	return u
}

func (u *URL) AddFragment(fragment string) *URL {
	u.u.Fragment = fragment
	return u
}

func (u *URL) LookInside() {
	fmt.Printf("URL: %#v\n", u)
}

func (u *URL) URL() string {
	u.u.RawQuery = u.values.Encode()
	return u.u.String()
}
