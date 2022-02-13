package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"ozon/pkg/app"
	"ozon/pkg/restmodel"
	"testing"
)

const postLinkURL = "http://localhost:8080/api"
const getLinkURL = "http://localhost:8080/api"

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func newRequest(url string, jsonStr []byte, method string) (int, []byte) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, body
}

func checkLinkPost(t *testing.T, link restmodel.Link, jsonStr []byte) restmodel.Link {
	code, body := newRequest(postLinkURL, jsonStr, http.MethodPost)

	checkResponseCode(t, http.StatusOK, code)

	var link1 restmodel.Link
	err := json.Unmarshal(body, &link1)
	if err != nil {
		t.Errorf("Error Unmarshal")
	}
	if len(link1.ShortLink) == 0 {
		t.Error("Not ShortLink")
	}
	if link.Link != link1.Link {
		t.Errorf("Expected '%s'. Got '%s'", link.Link, link1.Link)
	}
	return link1
}

func TestLinkPost(t *testing.T) {
	go app.Run("../configs")

	link := restmodel.Link{Link: "ususu"}
	var jsonStr, _ = json.Marshal(link)

	// uniqueness check
	l1 := checkLinkPost(t, link, jsonStr)
	l2 := checkLinkPost(t, link, jsonStr)
	if l1 != l2 {
		t.Error("various short links ")
		t.Errorf("Expected '%s'. Got '%s'", l1.ShortLink, l2.ShortLink)
	}

	// test empty link
	link.Link = ""
	jsonStr, _ = json.Marshal(link)
	code, _ := newRequest(postLinkURL, jsonStr, http.MethodPost)
	checkResponseCode(t, http.StatusBadRequest, code)

}

func TestLinkGet(t *testing.T) {
	//go app.Run("../configs")

	link := restmodel.Link{Link: "ususu"}
	var jsonStr, _ = json.Marshal(link)
	_, body := newRequest(postLinkURL, jsonStr, http.MethodPost)
	err := json.Unmarshal(body, &link)
	if err != nil {
		t.Errorf("Error Unmarshal")
	}

	var shortLink restmodel.Link
	shortLink.ShortLink = link.ShortLink
	jsonStr, _ = json.Marshal(shortLink)

	code, body := newRequest(getLinkURL, jsonStr, http.MethodGet)
	checkResponseCode(t, http.StatusOK, code)
	var sh1 restmodel.Link
	err = json.Unmarshal(body, &sh1)
	if err != nil {
		t.Errorf("Error Unmarshal")
	}

	if len(sh1.Link) == 0 {
		t.Error("Not Link")
	}
	if link.Link != sh1.Link {
		t.Errorf("Expected '%s'. Got '%s'", link.Link, sh1.Link)
	}

	shortLink.ShortLink = ""
	jsonStr, _ = json.Marshal(shortLink)
	code, _ = newRequest(getLinkURL, jsonStr, http.MethodGet)
	checkResponseCode(t, http.StatusBadRequest, code)
}

func TestBadLinkJSON(t *testing.T) {
	//go app.Run("../configs")

	var emptyJson []byte
	code, _ := newRequest(postLinkURL, emptyJson, http.MethodPost)
	checkResponseCode(t, http.StatusBadRequest, code)
	code, _ = newRequest(getLinkURL, emptyJson, http.MethodGet)
	checkResponseCode(t, http.StatusBadRequest, code)
}

func TestLinkNotFound(t *testing.T) {
	//go app.Run("../configs")

	link := restmodel.Link{ShortLink: "dsfdg"}
	var jsonStr, _ = json.Marshal(link)
	code, _ := newRequest(postLinkURL, jsonStr, http.MethodGet)
	checkResponseCode(t, http.StatusBadRequest, code)

}
