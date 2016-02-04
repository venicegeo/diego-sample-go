package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var args []string
		args = append(args, "--version")
		out, err := exec.Command("gdalinfo", args...).CombinedOutput()
		fmt.Fprintf(w, string(out))
		if err != nil {
			log.Println(err.Error())
		}
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
