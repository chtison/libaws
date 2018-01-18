#!/bin/sh -e --

cd -P -- `dirname -- $0`

cat > runtimes.go << EOF
package runtimes

// Runtime ...
type Runtime struct {
	Extension string
	Template  string
}

// Runtimes ...
var Runtimes = map[string]*Runtime{
	"python2.7": &Runtime{
		Extension: "py",
		Template: \``cat lambda.2.7.py`
\`,
	},
	"python3.6": &Runtime{
		Extension: "py",
		Template: \``cat lambda.3.6.py`
\`,
	},
}
EOF

goimports -w runtimes.go

echo "`dirname -- $0`/runtimes.go successfully created"
