package conf

import (
	"fmt"
	"path"
	"strings"

	"github.com/wncbb/check_api_diff/pkg/tdurl"
)

type Req struct {
	Variables []interface{} `json:"variables"`
	Info      *ReqInfo      `json:"info"`
	Item      []*ReqItem    `json:"item"`
}

func TraversalReqItemTree(items []*ReqItem, prefix string) []*ReqItem {
	res := make([]*ReqItem, 0)
	for _, v := range items {
		if v.Item == nil {
			v.Prefix = path.Join(prefix, v.Prefix)
			res = append(res, v)
		} else {
			curRes := TraversalReqItemTree(v.Item, path.Join(prefix, v.Name))
			res = append(res, curRes...)
		}
	}
	return res
}

type ReqInfo struct {
	Name        string `json:"name"`
	PostmanID   string `json:"_postman_id"`
	Description string `json:"description"`
	Schema      string `json:"schema"`
}

type ReqItem struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Item        []*ReqItem    `json:"item"`
	Event       []*ReqEvent   `json:"event"`
	Request     *ReqRequest   `json:"request"`
	Response    []interface{} `json:"response"`
	Prefix      string        `json:"prefix"`
}

type ReqEvent struct {
	Listen string     `json:"listen"`
	Script *ReqScript `json:"script"`
}

type ReqScript struct {
	Typ  string   `json:"type"`
	Exec []string `json:"exec"`
}

type ReqRequest struct {
	URL         *ReqURL         `json:"url"`
	Method      string          `json:"method"`
	Header      []*ReqHeaderRow `json:"header"`
	Body        *ReqBody        `json:"body"`
	Description string          `json:"description"`
}

type ReqURL struct {
	Raw   string      `json:"raw"`
	Host  []string    `json:"host"`
	Path  []string    `json:"path"`
	Query []*ReqQuery `json:"query"`
}

func (r *ReqURL) URL() string {
	u := tdurl.NewURL().SetHost(r.Host[0]).SetPath("/" + strings.Join(r.Path, "/"))
	for _, v := range r.Query {
		u.AddQueryValue(v.Key, v.Value)
	}
	return u.URL()
}

func (r *ReqURL) URLWithHost(host string) string {
	u := tdurl.NewURL().SetScheme("http").SetHost(host).SetPath("/" + strings.Join(r.Path, "/"))
	for _, v := range r.Query {
		u.AddQueryValue(v.Key, v.Value)
	}
	return u.URL()
}

type ReqQuery struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (r *ReqRequest) ShowLines() []string {
	res := make([]string, 0)
	res = append(res, fmt.Sprintf("Description: %s", r.Description))
	res = append(res, fmt.Sprintf("URL: %s", r.URL))
	res = append(res, fmt.Sprintf("Method: %s", r.Method))
	res = append(res, fmt.Sprintf("Header:"))
	for _, v := range r.Header {
		res = append(res, fmt.Sprintf("  %s: %s //%s", v.Key, v.Value, v.Description))
	}
	if r.Body != nil {
		res = append(res, fmt.Sprintf("Body:"))
		res = append(res, fmt.Sprintf("  Description:%s", r.Body.Description))
		res = append(res, fmt.Sprintf("  Mode:%s", r.Body.Mode))
		res = append(res, fmt.Sprintf("  Raw:%s", r.Body.Raw))
	}
	return res
}

func (r *ReqRequest) ShowComment(prefix string) string {
	strs := r.ShowLines()
	res := ""
	for _, v := range strs {
		res += prefix + v + "\n"
	}
	return res
}

type ReqHeaderRow struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type ReqBody struct {
	Mode        string `json:"mode"`
	Raw         string `json:"raw"`
	Description string `json:"description"`
}
