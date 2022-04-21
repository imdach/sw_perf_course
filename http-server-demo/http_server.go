package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	s, err := ioutil.ReadAll(r.Body)
	if err != nil {
		println("err: ", err)
	}

	str := string(s)
	fmt.Printf("received: %d %s\n", len(s), str)
	time.Sleep(time.Second * 3)

	fmt.Printf("process done: %d, %s\n", len(s), str)
	fmt.Fprintf(w, "process done\n")
}
func route0(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "process done\n")
	fmt.Printf("process done\n")
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func route2(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, RandStringRunes(1500))
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	} else {
		log.Println("Damn, no flush")
	}

	time.Sleep(time.Second * 2)
	//fmt.Fprintf(w, "process done\n")
	fmt.Printf("process done\n")
}

func route3(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("process begin\n")
	time.Sleep(time.Second * 70)
	fmt.Fprintf(w, "hello world\n")
	fmt.Printf("process done\n")

}
func main() {
	http.HandleFunc("/delay", indexHandler)
	http.HandleFunc("/route0", route0)
	http.HandleFunc("/route2", route2)
	http.HandleFunc("/route3", route3)
	http.ListenAndServe(":9000", nil)
}
