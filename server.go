package main

import (
  "fmt"
  "log"
  "net/http"
  "net/http/httputil"
  "io/ioutil"
  "github.com/k0kubun/pp"
)

func handlerDigest(w http.ResponseWriter, r *http.Request) {
	pp.Println("URL: %s\n", r.URL.String())
	pp.Println("Query: %s\n", r.URL.Query())
	pp.Println("Proto: %s\n", r.Proto)
	pp.Println("Method: %s\n", r.Method)
	pp.Println("Header: %s\n", r.Header)

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("--body--\n%s\n", string(body))
	if _, ok := r.Header["Authorization"]; !ok {
		w.Header().Add("WWW-Authenticate", `Digest realm="Secret Zone", nonce="TGLc25U2BQA=f510a2780473e18e6587be702c2e67fe2b04afd", algorithm=MD5, qop="auth"`)
		w.WriteHeader(http.StatusUnauthorized)
	}else{
	  fmt.Fprintf(w, "<html><body>secret page</body></html>\n")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
  dump, err := httputil.DumpRequest(r, true)
  if err != nil {
    http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
    return
  }
  fmt.Println(string(dump))
  fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func main() {
  var httpServer http.Server
  http.HandleFunc("/", handler)
  http.HandleFunc("/digest", handlerDigest)
  log.Println("start http listening :18888")
  httpServer.Addr = ":18888"
  log.Println(httpServer.ListenAndServe())
}
