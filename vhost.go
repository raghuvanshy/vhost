package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type vhost struct {
	myUrl, myHost string
}

func main() {

	vhosts := []vhost{
		vhost{
			"theitstudio.com/",
			"http://127.0.0.1:8080",
		},
	}

	u, _ := url.Parse(vhosts[0].myHost)
	http.Handle(vhosts[0].myUrl, httputil.NewSingleHostReverseProxy(u))

	http.ListenAndServe(":80", nil)
}
