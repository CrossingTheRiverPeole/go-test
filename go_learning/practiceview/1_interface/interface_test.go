package __interface

import (
	"fmt"
	"net/http"
	"testing"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello world\n")
}
func TestHttp(t *testing.T)  {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}
