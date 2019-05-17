package conf

type Req struct {
	Variables []interface{} `json:"variables"`
	Info      *ReqInfo      `json:"info"`
	Items     []*ReqItem    `json:"item"`
}

type ReqInfo struct {
	Name        string `json:"name"`
	PostmanID   string `json:"_postman_id"`
	Description string `json:"description"`
	Schema      string `json:"schema"`
}

type ReqItem struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Item        []*ReqItemItem `json:"item"`
}

type ReqItemItem struct {
	Name     string        `json:"name"`
	Event    []*ReqEvent   `json:"event"`
	Request  *ReqRequest   `json:"request"`
	Response []interface{} `json:"response"`
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
