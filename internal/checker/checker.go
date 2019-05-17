package checker

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	iconf "github.com/wncbb/check_api_diff/internal/conf"
	idiff "github.com/wncbb/check_api_diff/internal/diff"
	"github.com/wncbb/check_api_diff/internal/log"
	// istorage "github.com/wncbb/check_api_diff/internal/storage"
)

// RequestConfig [scheme:][//[userinfo@]host][/]path[?query][#fragment]
type RequestConfig struct {
	Scheme      string
	Host        string
	Path        string
	QueryValues map[string]string
}

var defaultHost string
var defaultSearchPath string
var reqConf *RequestConfig

func Init() {
	defaultHost = "10.10.53.127:8005"
	reqConf = &RequestConfig{
		Scheme: "http",
		Host:   "10.10.53.127:8005",
		// Path:   "/services/poi/search",
		Path: "/",
		QueryValues: map[string]string{
			"keyword":   "one altitude",
			"reference": "1.280714,103.848742",
		},
	}
}

/*
func SearchURL(host string) string {
	u := tdurl.NewURL().
		SetScheme(reqConf.Scheme).
		// SetHost(reqConf.Host).
		SetHost(host).
		SetPath(reqConf.Path)

	for k, v := range reqConf.QueryValues {
		u.AddQueryValue(k, v)
	}

	u.LookInside()

	return u.URL()
}
*/

func Check() {
	ParseFlag()
	log.Init(LogLevel())

	iconf.Init(EnvFile(), ReqFile())
	idiff.Init(ShowColor())
	// conf.PrintJson(conf.ReqConfByHost(onlineHost))
	reqConf := iconf.ReqConf()
	reqItems := iconf.TraversalReqItemTree(reqConf.Item, OutputDir())
	log.Default().Tracef("reqItem length: %d", len(reqItems))

	fmt.Println(StartString("DIFF RESULT", "\n", ""))
	for k, v := range reqItems {
		log.Default().Tracef("reqItem index:%d, prefix:%s, value:%#v", k, v.Prefix, v)

		diffStr, err := idiff.RunCompare(OnlinePOIHost(), DebugPOIHost(), "hostname", v)
		if err != nil {
			log.Default().Errorf("diff.RunCompare err:%#v\n", err)
			continue
		}

		diffInfo := GetDiffInfo(diffStr)

		err = WriteDiffToFile(v, diffStr, diffInfo)
		if err != nil {
			log.Default().Errorf("WriteDiffToFile err:%#v\n", err)
		}

		addNumStr := fmt.Sprintf("addNum=%-6d", diffInfo.AddNum)
		if diffInfo.AddNum > 0 {
			addNumStr = color.YellowString(addNumStr)
		}

		delNumStr := fmt.Sprintf("delNum=%-6d", diffInfo.DelNum)
		if diffInfo.DelNum > 0 {
			delNumStr = color.RedString(delNumStr)
		}

		fmt.Printf("DIFF: %s %s  ID: %s/%s\n", addNumStr, delNumStr, v.Prefix, v.Name)

	}
	return
}

type DiffInfo struct {
	AddNum int
	DelNum int
}

func GetDiffInfo(str string) *DiffInfo {
	strs := strings.Split(str, "\n")
	res := &DiffInfo{}
	for _, v := range strs {
		if len(v) == 0 {
			continue
		}

		if v[0] == '+' {
			res.AddNum += 1
		} else if v[0] == '-' {
			res.DelNum += 1
		}
	}
	return res
}
