/*
以下のコードは次のcurlコマンドと等価
$ curl ーG --data-urlencode "query=hello world" http://localhost:18888
*/
package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{
		"query": {"hello world"},
	}

	resp, _ := http.Get("http://localhost:18888" + "?" + values.Encode())
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	log.Println(string(body))
}
