

#!/bin/bash

set -e

WorkPath=$(pwd)

# export tools bin to PATH
PATH="${WorkPath}"/tools:"${PATH}"

if [ ! -f "$WorkPath"/tools/goimports-reviser-darwin ]; then
  {
    mkdir -p tools
    wget -P "${WorkPath}"/tools https://resource.gocloudcoder.com/goimports-reviser-darwin
    chmod +x "${WorkPath}"/tools/goimports-reviser-darwin
  }
fi

for i in $(find . -name "*.go" | grep -v ".pb.go" | grep -v ".pb.gw.go") ; do
    "$WorkPath"/tools/goimports-reviser-darwin -rm-unused -set-alias -format -file-path "$i"
done