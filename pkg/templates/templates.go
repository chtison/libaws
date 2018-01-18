package templates

// CloudFormationTmplYaml ...
const CloudFormationTmplYaml = `AWSTemplateFormatVersion: 2010-09-09

Resources:
    {{- /* Add resources below */}}
`

// CloudFormationDataYaml ...
const CloudFormationDataYaml = `
`

// LibawsYaml ...
const LibawsYaml = `Templates:
    - Path: cloudformation.tmpl.yaml
`

// Templates ...
var Templates = `
{{- define "sns-topic" }}
{{- $TopicLID        := (or .TopicLID "Topic"              ) }}
{{- $SubscriptionLID := (or .SubscriptionLID "Subscription") }}

    {{- $TopicLID }}:
        Type: AWS::SNS::Topic

{{- range $i, $e := .Subscriptions }}

    {{ $SubscriptionLID }}{{ $i }}:
        Type: AWS::SNS::Subscription
        DependsOn: {{ $TopicLID }}
        Properties:
            TopicArn: !Ref {{ $TopicLID }}
            Protocol: '{{ index $e 0 }}'
            Endpoint: '{{ index $e 1 }}'
{{- end }}
{{- end }}
{{- define "lambda" }}
{{- $FunctionName   := (or .FunctionName   (printf "Function%s"   (or .DefaultName ""))) }}
{{- $RoleName       := (or .RoleName       (printf "Role%s"       (or .DefaultName ""))) }}
{{- $LogGroupName   := (or .LogGroupName   (printf "LogGroup%s"   (or .DefaultName ""))) }}
{{- $PolicyName     := (or .PolicyName     (printf "Policy%s"     (or .DefaultName ""))) }}
{{- $PermissionName := (or .PermissionName (printf "Permission%s" (or .DefaultName ""))) }}
{{- $EventName      := (or .EventName      (printf "Event%s"      (or .DefaultName ""))) }}
{{- $S3Bucket       := (or .S3Bucket       (libaws.S3Bucket ))                           }}
{{- $ZipFile        := (printf "_out/lambda/%s.zip" .Path) }}

    {{- $FunctionName }}:
        Type: AWS::Lambda::Function
        DependsOn: {{ $RoleName }}
        Properties:
            Runtime: {{ with .Runtime }}{{ . }}{{ else }}python3.6{{ end }}
            Handler: {{ with .Handler }}{{ . }}{{ else }}lambda.handler{{ end }}
            Code:
{{- if .Deploy }}
                S3Bucket: {{ $S3Bucket }}
                S3Key: {{ call libaws.Sha256File $ZipFile }}
{{- else }} {{ $ZipFile }} {{- end }}
            Role: !GetAtt {{ $RoleName }}.Arn
{{- with .Timeout }}
            Timeout: {{ . }}
{{- end }}
{{- with .Environment }}
            Environment:
                Variables:
{{- range $k, $v := . }}
                    {{ $k }}: {{ $v }}
{{- end }}{{ end }}

    {{ $RoleName }}:
        Type: AWS::IAM::Role
        Properties:
            AssumeRolePolicyDocument:
                Version: 2012-10-17
                Statement:
                    Effect: Allow
                    Action: sts:AssumeRole
                    Principal:
                        Service: lambda.amazonaws.com

    {{ $LogGroupName }}:
        Type: AWS::Logs::LogGroup
        DependsOn: {{ $FunctionName }}
        Properties:
            LogGroupName: !Sub '/aws/lambda/${ {{- $FunctionName -}} }'
            RetentionInDays: {{ with .LogGroupRetentionInDays }}{{ . }}{{ else }}7{{ end }}

    {{ $PolicyName }}:
        Type: AWS::IAM::Policy
        DependsOn: [{{ $RoleName }}, {{ $LogGroupName }}]
        Properties:
            PolicyName: Default
            Roles: [!Ref {{ $RoleName }}]
            PolicyDocument:
                Version: 2012-10-17
                Statement:
                    - Effect: Allow
                      Action:
                          - logs:CreateLogStream
                          - logs:PutLogEvents
                      Resource: !GetAtt {{ $LogGroupName }}.Arn
{{- range .Policies }}
                    - Effect: {{ with .Effect }}{{ . }}{{ else }}Allow{{ end }}
                      Action:
{{- range .Action }}
                          - {{ . }}
{{- end }}
                      Resource: {{ .Resource }}
{{- end }}

{{- range $i, $e := .Permissions }}{{ $PermissionName := (printf "%s%d" $PermissionName $i) }}

    {{ $PermissionName }}:
        Type: AWS::Lambda::Permission
        DependsOn: {{ $FunctionName }}
        Properties:
            Action: {{ with $e.Action }}{{ . }}{{ else }}lambda:InvokeFunction{{ end }}
            FunctionName: !GetAtt {{ $FunctionName }}.Arn
            Principal: {{ $e.Principal }}
{{- with $e.SourceAccount }}
            SourceAccount: {{ . }}
{{- end }}
{{- with $e.SourceArn }}
            SourceArn: {{ . }}
{{- end }}
{{- end }}

{{- range $i, $e := .Schedules }}{{ $EventName := (printf "%s%d" $EventName $i) }}

    {{ $EventName }}:
        Type: AWS::Events::Rule
        DependsOn: {{ $FunctionName }}
        Properties:
            Targets:
                - Arn: !GetAtt {{ $FunctionName }}.Arn
                  Id: !Ref {{ $FunctionName }}
            ScheduleExpression: '{{ $e }}'

    FunctionPermissionFor{{ $EventName }}:
        Type: AWS::Lambda::Permission
        DependsOn: [{{ $FunctionName }}, {{ $EventName }}]
        Properties:
            Action: lambda:InvokeFunction
            FunctionName: !GetAtt {{ $FunctionName }}.Arn
            Principal: events.amazonaws.com
            SourceArn: !GetAtt {{ $EventName }}.Arn
{{- end }}
{{- end }}
`
