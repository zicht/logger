#!/bin/bash
export GO_SRC=${GOPATH}src/
export PROJECT=github.com/pbergman/logger
export PROJECT_DIR=${GO_SRC}${PROJECT}
cd $PROJECT_DIR
for i in $(find $PROJECT_DIR -type f -name "*_test.go" | xargs -n1 dirname | uniq); do
    name=${i#${GO_SRC}github.com/pbergman/}
    go test -coverprofile ./coverage/${name////.}.out ${i#${GO_SRC}}
    go tool cover -html=./coverage/${name////.}.out -o ./coverage/${name////.}.html
done