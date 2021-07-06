package main

import (
	"fmt"
	"net/http"

	"flag/config"
)

func main() {
	a := config.NewConfig()
	fmt.Println(a)
	runServer(a)
	//validation
}
func runServer(a *config.Config) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//data := ""
		fmt.Fprint(writer, a)
	})
	err := http.ListenAndServe("localhost:"+a.Port, nil)
	if err != nil {
		return
	}
}
