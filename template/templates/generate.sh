#!/bin/sh -e --

cd -P -- `dirname -- $0`

cat > templates.go << EOF
package templates

import (
	"github.com/chtison/libgo/tmpl"
	"text/template"
)

const CloudFormationTmplYaml = \``cat libaws/cloudformation.tmpl.yaml`
\`
const CloudFormationDataYaml = \``cat libaws/cloudformation.data.yaml`
\`
const LibawsYaml = \``cat libaws/libaws.yaml`
\`

var Template = template.Must(template.New("libaws").Funcs(tmpl.Funcs()).Parse(\`
`cat sns-topic/sns-topic.tmpl.yaml`
`cat lambda/lambda.tmpl.yaml`
`cat cognito-userpool/cognito-userpool.tmpl.yaml`
`cat api/api.tmpl.yaml`
\`))
EOF

echo "`dirname -- $0`/templates.go successfully created"
