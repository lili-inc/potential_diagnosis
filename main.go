package main

import (
	"lili_style_test/src"
)

func main() {
	router := src.GetRouter()
	router.Run(":8080")
}

//func main() {
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintln(w, "Hello World!")
//	})
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}
