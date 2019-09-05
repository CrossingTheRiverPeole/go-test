package http_test
import (
	"fmt"
	"net/http"
	"testing"
)

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "hello world")
}

func TestHttp(t *testing.T)  {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8000",nil)
}