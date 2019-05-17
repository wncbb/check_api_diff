package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/wncbb/check_api_diff/internal/log"
)

var env = &Env{}
var req = &Req{}

func EnvConf() *Env {
	return env
}

func ReqConf() *Req {
	return req
}

func initReq(reqFile string) {
	data, err := ioutil.ReadFile(reqFile)
	if err != nil {
		log.Default().Errorf("read req file %s failed, err:%#v", reqFile, err)
		panic(err)
	}

	err = json.Unmarshal(data, req)
	if err != nil {
		log.Default().Errorf("unmarshal req conf failed, err:%#v", err)
		panic(err)
	}
	reqBytes, _ := json.MarshalIndent(req, "", "  ")
	log.Default().Debugf("req: %s\n", string(reqBytes))
}

func initEnv(envFile string) {
	/*
		dir, file := path.Split(envFile)
		ext := path.Ext(file)
		rawFile := strings.TrimSuffix(ext)
		viper.SetConfigName(rawFile) // name of config file (without extension)
		viper.AddConfigPath(dir)     // path to look for the config file in
	*/
	data, err := ioutil.ReadFile(envFile)
	if err != nil {
		log.Default().Errorf("read env file %s failed, err:%#v", envFile, err)
		panic(err)
	}

	err = json.Unmarshal(data, &env)
	if err != nil {
		log.Default().Errorf("unmarshal env conf failed, err:%#v", err)
		panic(err)
	}

	envBytes, _ := json.MarshalIndent(env, "", "  ")
	log.Default().Debugf("env: %s\n", string(envBytes))
}

func ReqConfByHost(poiHost string) *Req {
	for _, v := range req.Items {
		for _, vv := range v.Item {
			vv.Request.URL = strings.Replace(vv.Request.URL, "{{hostname-int}}", "http://"+poiHost, 1)
		}
	}
	return req
}

func PrintJson(data interface{}) {
	dataBytes, _ := json.MarshalIndent(data, "", "  ")
	fmt.Printf("data: \n%s\n", string(dataBytes))
}

func Init(envFile, reqFile string) {
	initEnv(envFile)
	initReq(reqFile)

	/*
		envFile
		viper.SetConfigName("config") // 设置配置文件名 (不带后缀)
		viper.AddConfigPath("")       // 第一个搜索路径
		err := viper.ReadInConfig()   // 读取配置数据
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		viper.Unmarshal(&config) // 将配置信息绑定到结构体上
		fmt.Println(config)
	*/
}
