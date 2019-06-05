package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)
	mux.HandleFunc("/api/almighty", handleAlmighty)
	s := &http.Server{
		Addr:           ":1220",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is version 3"))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	time.Sleep(4 * time.Second)
	w.Write([]byte("bye bye ,this is v3 httpServer"))
}

func handleAlmighty(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数
	if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		var requestData struct {
			URL    string `json:"url"`
			Method string `json:"method"`
			Params string `json:"params"`
		}
		var err = json.Unmarshal(result, &requestData)
		fmt.Printf("%s\n", result)
		fmt.Printf("url:%s,method: %s\n", requestData.URL, requestData.Method)
		client := &http.Client{}
		req, err := http.NewRequest(requestData.Method, requestData.URL, nil)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("请求错误")
		}
		data, err := ioutil.ReadAll(resp.Body)
		w.Write(data)
	}

}
