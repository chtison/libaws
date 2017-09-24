#!/bin/sh -e --

cd -P -- `dirname -- $0`

cat > templates.go << EOF
package templates

import (
	"github.com/chtison/libgo/tmpl"
	"text/template"
)

const LibAwsYaml = \``cat libaws.yaml`\`

var Template = template.Must(template.New("libaws").Funcs(tmpl.Funcs()).Parse(\`
`cat sns-topic/sns-topic.tmpl.yaml`
`cat lambda/lambda.tmpl.yaml`
`cat cognito-userpool/cognito-userpool.tmpl.yaml`
\`))
EOF

echo "`dirname -- $0`/templates.go successfully created"
