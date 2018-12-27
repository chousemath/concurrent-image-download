#### Building for windows

```bash
$ GOOS=windows GOARCH=386 go build -o images.exe main.go
```

#### Windows-Related Build Errors

```bash
$ GOOS=windows GOARCH=386 go build -o images.exe main.go

# You may see this error
../../Sirupsen/logrus/terminal_windows.go:10:2: cannot find package "github.com/konsorten/go-windows-terminal-sequences" in any of:
	/usr/local/Cellar/go/1.11/libexec/src/github.com/konsorten/go-windows-terminal-sequences (from $GOROOT)
	/Users/jo/go/src/github.com/konsorten/go-windows-terminal-sequences (from $GOPATH)
You have new mail in /var/mail/jo

# To resolve the error above, run...
$ GOOS=windows go get ./...
```