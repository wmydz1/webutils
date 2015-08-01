package main

import (
	"fmt"
	"github.com/logoocc/webutils"
	"log"
	"net/http"
	"strconv"
)

var one int
var all int

// url format
// http://127.0.0.1:10000/?one=4&all=10

func testpaginator(w http.ResponseWriter, req *http.Request) {

	one, all = webutils.HanleParams(w, req, "one", "all")
	// fmt.Println(one, all)

	w.Write([]byte("golang server connected...."))
	p := webutils.NewPaginator(req, one, all)
	w.Write([]byte(strconv.Itoa(p.PageNums())))

}

func main() {
	http.HandleFunc("/", testpaginator)
	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
