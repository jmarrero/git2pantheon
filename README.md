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

7. To build the Container using buildah

```
buildah -t YOURTAG bud .
```

8. To run the container using podman

```
podman run --network=host --rm -p 9666:9666 YOURTAG
```
**Notice the --network flag** this is only needed if you are running pantheon on the same localbox, that way this service can upload files to "localhost".

9. To get inside the container and debug

get the container process
```
podman ps
```

```
podman exec -it PROCESS bash
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