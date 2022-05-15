/*
 * Copyrighe
 *
 */
package main

import (
	"bytes"
	"fmt"
	"net/http"
	"oga/controller/pkg/kube/client"
)

const (
	listeningPort = "8090"
)

func listPods(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() // Parses the request body
	if err != nil {
		fmt.Println("Failed to parse post")
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("Failed to read namespace parameters from request")); err != nil {
			// make linter happy
		}
		return
	}
	namespace := req.Form.Get("ns")
	clientInfo := client.ListPods(namespace)

	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("List of Pod Name and IP in namespace %s\n", namespace))
	for k, v := range clientInfo {
		b.WriteString(fmt.Sprintf("%s %s\n", k, v))
	}

	if _, err := w.Write(b.Bytes()); err != nil {
		fmt.Println("Error writing response")
	}
}

func main() {
	http.HandleFunc("/list-pods", listPods)
	fmt.Printf("Server started on port [%s]", listeningPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", listeningPort), nil)
	if err != nil {
		fmt.Printf("Failed to Listen %v\n", err)
	}
}
