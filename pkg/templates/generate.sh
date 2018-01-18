#!/bin/sh -e --

cd -P -- `dirname -- $0`

cat > templates.go << EOF
package templates

import (
"text/template"
"github.com/chtison/libgo/tmpl"
)
// CloudFormationTmplYaml ...
const CloudFormationTmplYaml = \``cat libaws/cloudformation.tmpl.yaml`
\`
// CloudFormationDataYaml ...
const CloudFormationDataYaml = \``cat libaws/cloudformation.data.yaml`
\`
// LibawsYaml ...
const LibawsYaml = \``cat libaws/libaws.yaml`
\`
// Templates ...
var Templates =
\`
`cat sns-topic/sns-topic.tmpl.yaml`
`cat lambda/lambda.tmpl.yaml`
\`
EOF

goimports -w templates.go

echo "`dirname -- $0`/templates.go successfully created"
