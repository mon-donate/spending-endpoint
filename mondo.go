package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type CharitySpending struct {
	name   string
	amount string
}

type HTTPResponse struct {
	path string
	w    http.ResponseWriter
	r    *http.Request
	data CharitySpending
}

func (h *HTTPResponse) serveCSS() {
	h.w.Header().Set("Content-Type", "text/css; charset=utf-8")
	http.ServeFile(h.w, h.r, "css/main.css")
}

func (h *HTTPResponse) serveMarkup() {
	spendingMarkup, err := ioutil.ReadFile("spending.template.html")
	if err != nil {
		return
	}

	breakdownTemplate, err := ioutil.ReadFile("breakdown.template.html")
	if err != nil {
		return
	}
	breakdownTemplateMarkup := string(breakdownTemplate)

	markup := string(spendingMarkup)
	markup = strings.Replace(markup, "{{amount}}", h.data.amount, -1)
	markup = strings.Replace(markup, "{{name}}", h.data.name, -1)

	nets := 5
	var breakdownMarkupBuffer bytes.Buffer

	for i := 0; i < nets; i++ {
		breakdownMarkupBuffer.WriteString(breakdownTemplateMarkup)
	}

	markup = strings.Replace(markup, "{{breakdown}}", breakdownMarkupBuffer.String(), -1)

	h.w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(h.w, markup, "")
}

func handler(w http.ResponseWriter, r *http.Request) {

	data := CharitySpending{r.URL.Query().Get("name"), r.URL.Query().Get("amount")}

	response := HTTPResponse{r.URL.Path[1:], w, r, data}

	if response.path == "css" {
		response.serveCSS()
	} else if response.path == "" {
		response.serveMarkup()
	}

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
