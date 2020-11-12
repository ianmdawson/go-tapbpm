# Go Tap BPM
A command-line utility to count beats per minute (BPM) by tapping keyboard keys. Created to measure the BPM of jazz songs, since it's often hard for computers to automatically determine the BPM.

# Installation
Ensure [go](https://golang.org/doc/install) is installed.

Run the following to download the source code and dependencies:
```
$ go get github.com/ianmdawson/go-tapbpm
$ cd $GOPATH/src/github.com/ianmdawson/go-tapbpm

# install dependencies and run tests
$ make test
```

Install go-tapbpm
```
make install
```

and now you should be able to run:
```
go-tapbpm
```

# Usage
- `r` resets the count
- `control+c`, `escape`, or `q` will quit
- Most other keyboard keys should add a beat count to the total, tap along with a song to determine the beats per minute.
