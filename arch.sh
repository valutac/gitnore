#!/bin/bash

usage () {
  local arch=$(basename $0)
  cat <<EOF
  USAGE:
    $ $arch [-h|--help] COMMAND
  EXAMPLES:
    $ $arch compile     Compile binary
EOF
  exit 1;
}

compile () {
	GOOS=windows GOARCH=386 go build -o build/gitnore-windows.exe
	GOOS=windows GOARCH=amd64 go build -o build/gitnore-windows-64.exe
	GOOS=linux GOARCH=386 go build -o build/gitnore-linux-32
	GOOS=linux GOARCH=amd64 go build -o build/gitnore-linux-64
	GOOS=darwin GOARCH=386 go build -o build/gitnore-macos-32
	GOOS=darwin GOARCH=amd64 go build -o build/gitnore-macos-64
	printf 'Binary is ready\n';
}

if [ $# -ne 1 ]; then
   usage;
fi

if { [ -z "$1" ] && [ -t 0 ] ; } || [ "$1" == '-h' ] || [ "$1" == '--help' ]
then
  usage;
fi

if [ "$1" == "compile" ]; then
  compile
else
  usage;
fi
