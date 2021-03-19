package try

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/try", sayhello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http serve failed,err:%v\n", err)
		return
	}
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "<h1>a new day</h1>")
}
