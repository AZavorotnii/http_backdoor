package main

import (
	"net/http"
	"os/exec"
)

type ShellServer struct{}

func (*ShellServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var args []string
	for _, values := range r.URL.Query() {
		args = append(args, values...)
	}
	// ignore first slash from URL path
	p := exec.Command(r.URL.Path[1:], args...)
	p.Stdin = r.Body
	out, err := p.CombinedOutput()

	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(out)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(out)
		w.Write([]byte(err.Error()))
	}
}

func main() {
	http.ListenAndServe(":9090", &ShellServer{})
}
