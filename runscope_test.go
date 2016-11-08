package runscope

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	url, _ := url.Parse(server.URL)
	client = NewClient(Options{
		Token:   "",
		BaseURL: url.String(),
	})
}

func handleGet(t *testing.T, path string, code int, data string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method: %v, want %v", r.Method, "GET")
		}
		w.WriteHeader(code)
		fmt.Fprint(w, data)
	})
}

func handlePost(t *testing.T, path string, code int, data string, iface interface{}, req interface{}) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(code)
		fmt.Fprint(w, data)
	})
}

func handlePut(t *testing.T, path string, code int, data string, iface interface{}, req interface{}) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.WriteHeader(code)
		fmt.Fprint(w, data)
	})
}

func handleDelete(t *testing.T, path string, code int) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(code)
	})
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if r.Method != want {
		t.Errorf("Request method: %v, want %v", r.Method, want)
	}
}

func testResponseData(t *testing.T, result interface{}, want interface{}) {
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Response data returned %+v, want %+v", result, want)
	}
}
