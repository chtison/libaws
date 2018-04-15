#!/bin/bash

set -e

cat >> generated.go << EOF
{{- define "$1" }}
	{{- libaws.$2 . }}
{{- end }}
EOF
