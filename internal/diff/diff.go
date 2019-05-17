package diff

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"net/http"
	// "net/url"
	"strings"

	"github.com/pkg/errors"
	"github.com/wncbb/check_api_diff/internal/conf"
	"github.com/wncbb/check_api_diff/internal/log"
	exdiff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

var color bool

func Init(c bool) {
	color = c
}

func Run(host, key string, req *conf.ReqItem) ([]byte, error) {
	var err error
	if req == nil {
		err = errors.New("internal/conf.Run req should not be nil")
		return nil, err
	}

	//u := strings.Replace(req.Request.URL.URL(), "{{"+key+"}}", "http://"+host, 1)
	u := req.Request.URL.URLWithHost(host)
	// u = url.QueryEscape(u)
	log.Default().Debugf("URL: %s\n", u)

	var httpReq *http.Request

	log.Default().Debugf("method: %s", req.Request.Method)

	if strings.ToUpper(req.Request.Method) == "GET" {
		// 小心被覆盖
		httpReq, err = http.NewRequest("GET", u, nil)
		if err != nil {
			return nil, errors.WithMessage(err, "Run->http.NewRequest failed")
		}

		log.Default().Debugf("LINE41 httpReq: %#v u:%s\n", httpReq, u)
		log.Default().Debugf("32 http method:%s", httpReq.Method)
		log.Default().Debugf("32 http url:%s", httpReq.URL)
		for _, v := range req.Request.Header {
			httpReq.Header.Add(v.Key, v.Value)
		}
		tmp, _ := json.Marshal(httpReq)
		log.Default().Debugf("44 httpReq: %#v\n", string(tmp))
	} else if strings.ToUpper(req.Request.Method) == "POST" {
		return nil, errors.New("POST method hasn't been supported yet")
	}

	tmp, _ := json.Marshal(httpReq)
	log.Default().Debugf("httpReq: %#v\n", string(tmp))

	log.Default().Debugf("line60: http.DefaultClient:%#v httpReq;%#v\n", http.DefaultClient, httpReq)
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		log.Default().Debugf("http.DefaultClient.Do failed, err:%#v", err)
		return nil, errors.WithMessage(err, "Run->http.DefaultClient.Do")
	}
	defer httpResp.Body.Close()
	if false {
		reader, err := gzip.NewReader(httpResp.Body)
		if err != nil {
			log.Default().Debugf("gzip.NewReader failed, err:%#v", err)
			return nil, errors.WithMessage(err, "Run->gzip.NewReader")
		}
		respData, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Default().Debugf("49 ioutil.ReadAll failed, err:%#v", err)
			return nil, errors.WithMessage(err, "Run->ioutil.ReadAll->76")
		}
		return respData, nil
	}
	respData, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		log.Default().Debugf("56 ioutil.ReadAll failed, err:%#v", err)
		return nil, errors.WithMessage(err, "Run->ioutil.ReadAll->83")
	}
	if len(respData) == 0 {
		return nil, errors.New("response should not be empty")
	}
	return respData, nil

}

func RunCompare(onlineHost, debugHost, key string, req *conf.ReqItem) (string, error) {
	if req == nil {
		err := errors.New("internal/diff.RunCompare req should not be nil")
		return "", err
	}

	onlineResp, err := Run(onlineHost, key, req)
	if err != nil {
		return "", errors.WithMessage(err, "Run onlineHost err")
	}

	debugResp, err := Run(debugHost, key, req)
	if err != nil {
		return "", errors.WithMessage(err, "Run debugHost err")
	}

	log.Default().Debugf("line106 onlineResp: %s\n", string(onlineResp))
	log.Default().Debugf("line106 debugResp: %s\n", string(debugResp))

	return Diff(onlineResp, debugResp)
}

func Diff(onlineResp, debugResp []byte) (string, error) {
	log.Default().Debugf("onlineResp:%s", string(onlineResp))
	log.Default().Debugf("debugResp:%s", string(debugResp))
	differ := exdiff.New()
	d, err := differ.Compare(onlineResp, debugResp)
	if err != nil {
		return "", errors.WithMessage(err, "Diff->differ.Compare")
	}

	format := "ascii"
	var diffString string
	if format == "ascii" {
		var aJson map[string]interface{}
		json.Unmarshal(onlineResp, &aJson)

		config := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       color,
		}

		formatter := formatter.NewAsciiFormatter(aJson, config)
		diffString, err = formatter.Format(d)
		if err != nil {
			// No error can occur
		}
	} else if format == "delta" {
		formatter := formatter.NewDeltaFormatter()
		diffString, err = formatter.Format(d)
		if err != nil {
			// No error can occur
		}
	} else {
		return "", errors.Errorf("Unknown Foramt %s\n", format)
	}

	return diffString, nil
}
