package main
import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"net/url"
	"os"
	"strings"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat"
	"io/ioutil"
	"io"
)

func ShowHello(rw http.ResponseWriter, req *http.Request) {
	todos := Todos{
		Todo{Name: "Welcome!"},
	}

	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(todos); err != nil {
		panic(err)
	}
}

func Jssdk_api(rw http.ResponseWriter, req *http.Request) {
	_ , err1 := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err1 != nil {
		panic(err1)
	}
	if err2 := req.Body.Close(); err2 != nil {
		panic(err2)
	}

	req.ParseForm()
	getUrl, found := req.Form["url"]

	if !(found) {
		returnError(rw,"miss_params")
		return
	}

	vars := mux.Vars(req)
	getAct := vars["appid"]
	decodeUrl, _ := url.QueryUnescape(getUrl[0])

	file, _ := os.Open("./wxgoconf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("configiure_error:", err)
		returnError(rw,"jssdk_api_error(0)")
		return
	}

	if strings.Compare(getAct,configuration.Weixin.Appid) == -1 {
		returnError(rw,"jssdk_api_error(1)")
		return
	}

	memRedis := cache.NewRedis(&cache.RedisOpts{
		Host:     configuration.Redis.Host,
		Password: configuration.Redis.Password,
		Database: 0,
	})

	config := &wechat.Config{
		AppID:     configuration.Weixin.Appid,
		AppSecret: configuration.Weixin.Appsecret,
		Cache:     memRedis,
	}
	wc := wechat.NewWechat(config)

	js := wc.GetJs()
	cfg, err := js.GetConfig(decodeUrl)
	if err != nil {
		returnError(rw, "jssdk_api_error(2)")
		return
	}

	resData := Resmsg{
		Code:200,
		Data:Jsapi{cfg.AppID,cfg.Timestamp,cfg.NonceStr,cfg.Signature,decodeUrl},
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(rw).Encode(resData); err != nil {
		panic(err)
	}
}

func returnError(w http.ResponseWriter,s string){
	resData := Reserr{
		Code:400,
		Msg:s,
	}

	if err := json.NewEncoder(w).Encode(resData); err != nil {
		panic(err)
	}

	return
}
