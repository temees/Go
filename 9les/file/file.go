package main

import (
	"fmt"
	"net/http"

	"readfile/conf"
)

func main() {
	a := conf.NewConfig()
	fmt.Println(a)
	runServer(a)
	//validation
}
func runServer(a *conf.Config) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//data := ""
		fmt.Fprint(writer, a)
	})
	err := http.ListenAndServe("localhost:"+a.Port, nil)
	if err != nil {
		return
	}
}
