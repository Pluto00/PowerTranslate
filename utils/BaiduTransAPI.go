package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/go-ini/ini"
	"golang.org/x/time/rate"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var limiter *rate.Limiter
var appId string
var token string

func init() {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal(err)
	}
	env := cfg.Section("env")
	appIdKey, _ := env.GetKey("appid")
	appId = appIdKey.String()
	tokenKey, _ := env.GetKey("token")
	token = tokenKey.String()
	// qps < 1
	limiter = rate.NewLimiter(1, 1)
}

func BaiduTransAPI(q, from, to string) (map[string]interface{}, error) {
	for !limiter.Allow() { // wait
		time.Sleep(time.Second)
	}
	salt := strconv.Itoa(rand.Int())
	sign := fmt.Sprintf("%x", md5.Sum([]byte(appId+q+salt+token)))
	url := "http://api.fanyi.baidu.com/api/trans/vip/translate"
	method := "POST"
	payload := strings.NewReader(fmt.Sprintf("q=%s&from=%s&to=%s&appid=%s&salt=%s&sign=%s", q, from, to, appId, salt, sign))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var ret map[string]interface{}
	err = json.Unmarshal(body, &ret)
	return ret, err
}
