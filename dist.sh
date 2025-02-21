#/bin/sh

set -e
stat go.mod > /dev/null

rm -rf _build
mkdir _build
echo build linux
CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -o _build .
echo build windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o _build .

DEST=$HOME/go/bin
if [ -d "$DEST" ]; then
    cp _build/mdnum "$DEST"
    echo copied to "$DEST"
else
    echo "$DEST" not found
fi
