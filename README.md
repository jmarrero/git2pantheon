**Setting up your development environment**

1. Install golang

For fedora follow:
https://developer.fedoraproject.org/tech/languages/go/go-installation.html

For any other OS, reffer to the golang website:
https://golang.org/doc/install


2. Install go-git dependency

```
go get gopkg.in/src-d/go-git.v4
```

3. Clone this repo on your go root. Assuming the default go root... do this:

```
cd ~/go
```

```
git clone YOUR_FORK_OF_THIS_REPO
```

4. Build the project

```
cd git2pantheon
```

```
go build
```

5. run the service

```
./git2pantheon
```

6. Running go tests

```
go test
```

**Submitting a request**

The service expects only POST REST calls.

The payload must include a repo and branch to be cloned.
The repository is expected to have a pantheon2.yml file defining the pantheon enpoint and general configuration on how to handle the documentation in the repository. 

Without it no upload will happen.

An example of the payload is:

```
curl -d '{"repo":"https://github.com/jmarrero/test-adocs.git", "branch":"master"}' -H "Content-Type: application/json" -X POST http://localhost:9666/clone
```

For information about pantheon see: 
https://github.com/redhataccess/pantheon