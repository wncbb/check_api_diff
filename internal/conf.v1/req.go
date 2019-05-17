package conf

import (
	"github.com/bitly/go-simplejson"
	"github.com/pkg/errors"
)

const (
	_ = iota
	ItemTypeRequest
	ItemTypeGroup
)

func Parse(sj *simplejson.Json) ([]*ReqRequest, error) {
	res := make([]*ReqRequest, 0)
	req := NewReq()
	if sjInfo, ok := sj.CheckGet("info"); ok {
		req.Info.Name = sjInfo.Get("name")
		req.Info.PostmanID = sjInfo.Get("_postman_id")
		req.Info.Description = sjInfo.Get("description")
		req.Info.Schema = sjInfo.Get("schema")
	}
	if itemList, ok := sj.CheckGet("item"); ok {
		arr, err := itemList.Array()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for i := 0; i < len(arr); i = i + 1 {
			curSJItem := itemList.GetIndex(i)
			if children, ok := curSJItem.CheckGet("item"); ok {
				// curSJItem is GroupItem
				curGroupItem := NewGroupItem()
			} else {
				// curSJItem is RequestItem
				curRequestItem := NewRequestItem()
				curRequestItem.Name, err = curSJItem.Get("name").String()
				if err != nil {
					// TODO: log
					err = nil
					continue
				}
				curRequestItem.Request
			}
		}
	}
}

func ItemTraversal([]Item) []*RequestItem {
	res := make([]*RequestItem, 0)
	for _, v := range Item {
	}

}

type Item interface {
	GetTyp() int
}

type RequestItem struct {
	Name     string        `json:"name"`
	Event    []*ReqEvent   `json:"event"`
	Request  *ReqRequest   `json:"request"`
	Response []interface{} `json:"response"`
	Level    []string      `json:"level"`
}

func NewRequestItem() *RequestItem {
	return &RequestItem{
		Event:    make([]*ReqEvent, 0),
		Request:  NewReqRequest(),
		Response: make([]interface{}),
		Level:    make([]string, 0),
	}
}

func (r *RequestItem) GetTyp() int {
	return ItemTypeRequest
}

func (r *RequestItem) GetName() string {
	return r.Name
}

func (r *RequestItem) GetEvent() []*ReqEvent {
	return r.Event
}

func (r *RequestItem) GetRequest() *ReqRequest {
	return r.Request
}

func (r *RequestItem) GetResponse() []interface{} {
	return r.Response
}

type GroupItem struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Item        []Item `json:"item"`
}

func NewGroupItem() *GroupItem {
	return &GroupItem{
		Item: make([]Item, 0),
	}
}

func (g *GroupItem) GetTyp() int {
	return ItemTypeGroup
}

func (g *GroupItem) GetName() int {
	return g.Name
}

func (g *GroupItem) GetDescription() string {
	return g.Description
}

func (g *GroupItem) GetItem() []Item {
	return g.Item
}

type Req struct {
	Variables []interface{} `json:"variables"`
	Info      *ReqInfo      `json:"info"`
	Item      []Item        `json:"item"`
}

func NewReq() *Req {
	return &Req{
		Variables: make([]interface{}, 0),
		Info:      NewReqInfo(),
		Item:      make([]Item, 0),
	}
}

type ReqInfo struct {
	Name        string `json:"name"`
	PostmanID   string `json:"_postman_id"`
	Description string `json:"description"`
	Schema      string `json:"schema"`
}

func NewReqInfo() *ReqInfo {
	return &ReqInfo{}
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
	URL         string          `json:"url"`
	Method      string          `json:"method"`
	Header      []*ReqHeaderRow `json:"header"`
	Body        *ReqBody        `json:"body"`
	Description string          `json:"description"`
}

func NewReqRequest() *ReqRequest {
	return &ReqRequest{
		Header: make([]*ReqHeaderRow),
		Body:   NewReqBody(),
	}
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

func NewReqBody() *ReqBody {
	return &ReqBody{}
}
