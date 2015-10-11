package http

import (
	"io/ioutil"
	nhttp "net/http"
)

func Get(url string, query string) ([]byte, error) {
	res, err := nhttp.Get(url + "?" + query)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}
