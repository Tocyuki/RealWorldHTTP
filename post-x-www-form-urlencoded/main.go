/*
以下のコードは次のcurlコマンドと等価
$ curl -d test=value http://localhost:18888
*/
package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{
		"test": {"value"},
	}

	resp, err := http.PostForm("http://localhost:18888", values)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
