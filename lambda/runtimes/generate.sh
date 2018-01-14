#!/bin/sh -e --

cd -P -- `dirname -- $0`

cat > runtimes.go << EOF
package lambda

// Runtime ...
type Runtime struct {
	Extension string
	Template  string
}

// Templates ...
var runtimes = map[string]*Runtime{
	"python2.7": &Runtime{
		Extension: "py",
		Template: \``cat runtimes/lambda.2.7.py`
\`,
	},
	"python3.6": &Runtime{
		Extension: "py",
		Template: \``cat runtimes/lambda.3.6.py`
\`,
	},
}
EOF

echo "runtimes.go successfully created"
