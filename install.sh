#!/bin/sh
# Install development libraries I use for emacs.
go get golang.org/x/tools/cmd/gorename
go get github.com/rogpeppe/godef
go get -u github.com/ptrv/goflycheck
go get -u github.com/dougm/goflymake
go get -u golang.org/x/tools/cmd/goimports
#go get -u golang.org/x/tools/cmd/godoc
go get -u github.com/golang/lint/golint
go get -u github.com/nsf/gocode
go get -u github.com/alecthomas/gometalinter
gometalinter --install --update


#Install libraries used by this library
go get -u github.com/nsf/termbox-go
go get -u github.com/JoelOtter/termloop
echo 'Done'
