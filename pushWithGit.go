package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
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
		directory := "/home/jmarrero/g/" + randomAlphaNumericString()

		//start a new goroutine (lightweight thread) to handle clone/push/cleanup
		go gitClone(repository, branch, directory)

		fmt.Fprint(w, "POST done")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func gitClone(repository string, branch string, directory string) {
	log.Printf("git clone -b %s --single-branch %s %s", branch, repository, directory)

	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:           repository,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
		SingleBranch:  true,
		Progress:      os.Stdout,
	})

	if err == nil {
		//Uploader call still on the same gorutine.
		push2pantheon(directory)
	} else {
		log.Print(err)
	}

}

func push2pantheon(directory string) {
	//Now call python
	args := []string{"push", "--directory", directory}
	cmd := exec.Command("/home/jmarrero/g/pantheon.py", args...) //TODO
	out, err := cmd.Output()

	log.Print(err)
	log.Print(string(out))

	// Cleanup also happens in the same thread.
	cleanup(directory)
}

func cleanup(directory string) {
	//Todo
	log.Printf("Cleanup: deleting directory: %s", directory)
	os.RemoveAll(directory)
}

func randomAlphaNumericString() string {
	const chars = "0123456789abcdefghijklmnopqrstuvwxyzREDHAT"

	var seed *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	byteSlice := make([]byte, 10)
	for i := range byteSlice {
		byteSlice[i] = chars[seed.Intn(len(chars))]
	}
	return string(byteSlice)
}

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	flag.Parse()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/clone", cloneBranch)

	log.Printf("listening on port %s", *flagPort)
	log.Fatal(http.ListenAndServe(":"+*flagPort, mux))
}
