package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"

	. "github.com/akfaew/aeutils"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

var (
	App1 = &url.URL{
		Scheme: "https",
		Host:   "app1.example.com",
	}
	App2 = &url.URL{
		Scheme: "https",
		Host:   "app2.example.com",
	}
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// It's complicated because we base our request on Form data. See:
	// https://stackoverflow.com/questions/49745252/reverseproxy-depending-on-the-request-body-in-golang
	director := func(r *http.Request) {
		var target *url.URL

		if r.Body == nil {
			LogErrorfd(ctx, "body=nil")
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			LogErrorfd(ctx, "err=%v", err)
			return
		}

		// Reassign the now empty body
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		planID := r.FormValue("subscription_plan_id")
		LogInfofd(ctx, "subscription_plan_id=%v", planID)
		switch planID {
		case "1000":
			target = App1
		case "1001":
			target = App1

		case "2000":
			target = App2

		default:
			target = App1
		}
		LogInfofd(ctx, "target=%+v", target)

		// Reassign the body again
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		r.URL.Scheme = target.Scheme
		r.URL.Host = target.Host
	}

	rp := &httputil.ReverseProxy{
		Director:  director,
		Transport: &urlfetch.Transport{Context: ctx},
	}

	rp.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", handler)

	appengine.Main()
}
