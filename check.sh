#/bin/sh

set -e
stat go.mod > /dev/null

git status
go test -count 1 ./...
staticcheck ./...
echo "check ok"
