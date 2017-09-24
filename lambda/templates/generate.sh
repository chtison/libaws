#!/bin/sh -e --

cd -P -- `dirname -- $0`

cat > templates.go << EOF
package templates

// ...
const (
	Python36 = \``cat lambda.py3`\`
)
EOF

echo "`dirname -- $0`/templates.go successfully created"
