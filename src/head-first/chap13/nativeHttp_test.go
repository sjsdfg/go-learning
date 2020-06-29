package chap13

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttpGet(t *testing.T) {
	response, err := http.Get("https://www.examples.com/")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(body))
}
