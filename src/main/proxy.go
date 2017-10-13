package main

import (
	"conf"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const (
	isFromWeixinParam = "__is_from_weixin__"
	isFromWeixinValue = "true"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	finalUrl := ""

	isFromWeixin := r.Form[isFromWeixinParam]
	if len(isFromWeixin) > 0 && contains(isFromWeixin, isFromWeixinValue) {
		if key := r.Form[conf.Conf.KeyParam]; len(key) > 0 {
			if redirectUrl := conf.Conf.RedirectUrls[key[0]]; len(redirectUrl) > 0 {

				u, _ := url.Parse(redirectUrl)

				q := r.URL.Query()

				q.Del(isFromWeixinParam)
				q.Del(conf.Conf.KeyParam)
				for key, value := range u.Query() {
					q.Add(key, value[0])
				}

				u.RawQuery = q.Encode()

				finalUrl = u.String()

			}
		}
	} else {
		state := "STATE"
		if passedState := r.Form["state"]; len(passedState) > 0 {
			state = passedState[0]
		}

		scope := "snsapi_base"
		if passedScope := r.Form["scope"]; len(passedScope) > 0 {
			scope = passedScope[0]
		}

		passedQuery := r.URL.Query()
		passedQuery.Del("state")
		passedQuery.Add(isFromWeixinParam, isFromWeixinValue)
		r.URL.RawQuery = passedQuery.Encode()
		redirectUri := conf.Conf.Scheme + "://" + r.Host + r.URL.String()

		authUrl, _ := url.Parse("https://open.weixin.qq.com/connect/oauth2/authorize")

		authQuery := authUrl.Query()

		authQuery.Set("appid", conf.Conf.AppId)
		authQuery.Set("redirect_uri", redirectUri)
		authQuery.Set("response_type", "code")
		authQuery.Set("scope", scope)
		authQuery.Set("state", state+"#wechat_redirect")

		authUrl.RawQuery = authQuery.Encode()

		finalUrl = authUrl.String()
	}

	if len(finalUrl) > 0 {
		fmt.Println(finalUrl)
		http.Redirect(w, r, finalUrl, http.StatusFound)
	}

}

func main() {
	cfgFile := flag.String("c", "config.yml", "configuration file")

	// parse config
	conf.ParseConfig(*cfgFile)

	log.Println("proxy starts")

	http.HandleFunc("/", redirect)                    //设置访问的路由
	err := http.ListenAndServe(conf.Conf.Listen, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
