package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	// flagPort is the open port the application listens on
	flagPort = flag.String("port", "9666", "Port to listen on")
)

type repository_and_branch struct {
	Repo   string `json:"repo"`
	Branch string `json:"branch"`
}

// curl -d '{"repo":"https://github.com/jmarrero/test-adocs.git", "branch":"master"}' -H "Content-Type: application/json" -X POST http://localhost:9666/clone
func cloneBranch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		var repo repository_and_branch
		err = json.Unmarshal(body, &repo)

		log.Println(repo.Repo)
		log.Println(repo.Branch)

		repository := repo.Repo
		branch := repo.Branch
		directory := "./" + randomAlphaNumericString()

		//start a new goroutine (lightweight thread) to handle clone/push/cleanup
		go gitClone(repository, branch, directory)

		fmt.Fprint(w, "POST done")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	flag.Parse()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/clone", cloneBranch)

	GetUploader()

	log.Printf("listening on port %s", *flagPort)
	log.Fatal(http.ListenAndServe(":"+*flagPort, mux))
}
