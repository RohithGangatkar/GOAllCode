package main

import (
	"fmt"
	"net/http"
)

func TestingReq(w http.ResponseWriter, r *http.Request) {
	data := []byte("V1 of student's called")
	w.Write(data)
}

type anySt struct {
}

func (a anySt) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := []byte("V2 of student's called")
	w.Write(data)
}

func main() {

	//-----------------------------------Method 1----------------------------------------------
	// valueSt := anySt{}

	// mux := http.NewServeMux()

	// mux.HandleFunc("/v1/api/test", TestingReq)

	// mux.Handle("/v1/api/test2", valueSt)

	// s := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: mux,
	// }
	// s.ListenAndServe()
	fmt.Println("server started")
	//-----------------------------------Another Method----------------------------------------------
	valueSt := anySt{}

	http.HandleFunc("/v1/test", TestingReq)
	http.Handle("/v1/testing", valueSt)
	http.ListenAndServe(":8080", nil)

}
