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
	if req == nil {
		err := errors.New("internal/conf.Run req should not be nil")
		return nil, err
	}

	u := strings.Replace(req.Request.URL, "{{"+key+"}}", "http://"+host, 1)
	// u = url.QueryEscape(u)
	log.Default().Debugf("URL: %s\n", u)

	var httpReq *http.Request

	log.Default().Debugf("method: %s", req.Request.Method)

	if strings.ToUpper(req.Request.Method) == "GET" {
		httpReq, _ = http.NewRequest("GET", u, nil)
		log.Default().Debugf("32 http method:%s", httpReq.Method)
		log.Default().Debugf("32 http url:%s", httpReq.URL)
		for _, v := range req.Request.Header {
			httpReq.Header.Add(v.Key, v.Value)
			// httpReq.Header.Add("User-Agent", "PostmanRuntime/7.11.0")
			// httpReq.Header.Add("Accept", "*/*")
			// httpReq.Header.Add("Cache-Control", "no-cache")
			/*
				httpReq.Header.Add("Postman-Token", "d6f69e67-63cf-4eb8-bd1e-693a794e2734,8b109349-fb86-476d-9940-4b0e67668899")
				httpReq.Header.Add("Host", "10.10.53.127:8005")
				httpReq.Header.Add("accept-encoding", "gzip, deflate")
				httpReq.Header.Add("Connection", "keep-alive")
				httpReq.Header.Add("cache-control", "no-cache")
			*/
		}
		tmp, _ := json.Marshal(httpReq)
		log.Default().Debugf("44 httpReq: %#v\n", string(tmp))
	}

	tmp, _ := json.Marshal(httpReq)
	log.Default().Debugf("httpReq: %#v\n", string(tmp))

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		log.Default().Debugf("http.DefaultClient.Do failed, err:%#v", err)
		return nil, errors.WithStack(err)
	}
	defer httpResp.Body.Close()
	if false {
		reader, err := gzip.NewReader(httpResp.Body)
		if err != nil {
			log.Default().Debugf("gzip.NewReader failed, err:%#v", err)
			return nil, errors.WithStack(err)
		}
		respData, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Default().Debugf("49 ioutil.ReadAll failed, err:%#v", err)
			return nil, errors.WithStack(err)
		}
		return respData, nil
	}
	respData, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		log.Default().Debugf("56 ioutil.ReadAll failed, err:%#v", err)
		return nil, errors.WithStack(err)
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
		return "", errors.WithStack(err)
	}

	debugResp, err := Run(debugHost, key, req)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return Diff(onlineResp, debugResp)
}

func Diff(onlineResp, debugResp []byte) (string, error) {
	log.Default().Debugf("onlineResp:%s", string(onlineResp))
	log.Default().Debugf("debugResp:%s", string(debugResp))
	differ := exdiff.New()
	d, err := differ.Compare(onlineResp, debugResp)
	if err != nil {
		return "", errors.WithStack(err)
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
