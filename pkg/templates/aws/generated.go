package aws

// Templates ...
const Templates = `{{- define "AWS::ApiGateway::Account" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::Account
{{- if .CloudWatchRoleArn }}
        Properties:
{{- if .CloudWatchRoleArn }}
            CloudWatchRoleArn: {{ .CloudWatchRoleArn }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::ApiGateway::ApiKey" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::ApiKey
{{- if .CustomerId | or .Description | or .Enabled | or .GenerateDistinctId | or .Name | or .StageKeys }}
        Properties:
{{- if .CustomerId }}
            CustomerId: {{ .CustomerId }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .Enabled }}
            Enabled: {{ .Enabled }}
{{- end }}
{{- if .GenerateDistinctId }}
            GenerateDistinctId: {{ .GenerateDistinctId }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .StageKeys }}
            StageKeys:
{{- range $_, $v := .StageKeys }}
                -
{{- if $v.RestApiId }}
                  RestApiId: {{ $v.RestApiId }}
{{- end }}
{{- if $v.StageName }}
                  StageName: {{ $v.StageName }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::ApiGateway::Authorizer" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::Authorizer
        Properties:
            RestApiId: {{ .RestApiId }}
{{- if .AuthType }}
            AuthType: {{ .AuthType }}
{{- end }}
{{- if .AuthorizerCredentials }}
            AuthorizerCredentials: {{ .AuthorizerCredentials }}
{{- end }}
{{- if .AuthorizerResultTtlInSeconds }}
            AuthorizerResultTtlInSeconds: {{ .AuthorizerResultTtlInSeconds }}
{{- end }}
{{- if .AuthorizerUri }}
            AuthorizerUri: {{ .AuthorizerUri }}
{{- end }}
{{- if .IdentitySource }}
            IdentitySource: {{ .IdentitySource }}
{{- end }}
{{- if .IdentityValidationExpression }}
            IdentityValidationExpression: {{ .IdentityValidationExpression }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .ProviderARNs }}
            ProviderARNs:
{{- range $_, $v := .ProviderARNs }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Type }}
            Type: {{ .Type }}
{{- end }}
{{- end }}
{{- define "AWS::ApiGateway::BasePathMapping" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::BasePathMapping
        Properties:
            DomainName: {{ .DomainName }}
{{- if .BasePath }}
            BasePath: {{ .BasePath }}
{{- end }}
{{- if .RestApiId }}
            RestApiId: {{ .RestApiId }}
{{- end }}
{{- if .Stage }}
            Stage: {{ .Stage }}
{{- end }}
{{- end }}
{{- define "AWS::ApiGateway::ClientCertificate" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::ClientCertificate
{{- if .Description }}
        Properties:
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::ApiGateway::Deployment" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::Deployment
        Properties:
            RestApiId: {{ .RestApiId }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .StageDescription }}
            StageDescription:
{{- if .StageDescription.CacheClusterEnabled }}
                CacheClusterEnabled: {{ .StageDescription.CacheClusterEnabled }}
{{- end }}
{{- if .StageDescription.CacheClusterSize }}
                CacheClusterSize: {{ .StageDescription.CacheClusterSize }}
{{- end }}
{{- if .StageDescription.CacheDataEncrypted }}
                CacheDataEncrypted: {{ .StageDescription.CacheDataEncrypted }}
{{- end }}
{{- if .StageDescription.CacheTtlInSeconds }}
                CacheTtlInSeconds: {{ .StageDescription.CacheTtlInSeconds }}
{{- end }}
{{- if .StageDescription.CachingEnabled }}
                CachingEnabled: {{ .StageDescription.CachingEnabled }}
{{- end }}
{{- if .StageDescription.ClientCertificateId }}
                ClientCertificateId: {{ .StageDescription.ClientCertificateId }}
{{- end }}
{{- if .StageDescription.DataTraceEnabled }}
                DataTraceEnabled: {{ .StageDescription.DataTraceEnabled }}
{{- end }}
{{- if .StageDescription.Description }}
                Description: {{ .StageDescription.Description }}
{{- end }}
{{- if .StageDescription.DocumentationVersion }}
                DocumentationVersion: {{ .StageDescription.DocumentationVersion }}
{{- end }}
{{- if .StageDescription.LoggingLevel }}
                LoggingLevel: {{ .StageDescription.LoggingLevel }}
{{- end }}
{{- if .StageDescription.MethodSettings }}
                MethodSettings:
{{- range $_, $v := .StageDescription.MethodSettings }}
                    -
{{- if $v.CacheDataEncrypted }}
                      CacheDataEncrypted: {{ $v.CacheDataEncrypted }}
{{- end }}
{{- if $v.CacheTtlInSeconds }}
                      CacheTtlInSeconds: {{ $v.CacheTtlInSeconds }}
{{- end }}
{{- if $v.CachingEnabled }}
                      CachingEnabled: {{ $v.CachingEnabled }}
{{- end }}
{{- if $v.DataTraceEnabled }}
                      DataTraceEnabled: {{ $v.DataTraceEnabled }}
{{- end }}
{{- if $v.HttpMethod }}
                      HttpMethod: {{ $v.HttpMethod }}
{{- end }}
{{- if $v.LoggingLevel }}
                      LoggingLevel: {{ $v.LoggingLevel }}
{{- end }}
{{- if $v.MetricsEnabled }}
                      MetricsEnabled: {{ $v.MetricsEnabled }}
{{- end }}
{{- if $v.ResourcePath }}
                      ResourcePath: {{ $v.ResourcePath }}
{{- end }}
{{- if $v.ThrottlingBurstLimit }}
                      ThrottlingBurstLimit: {{ $v.ThrottlingBurstLimit }}
{{- end }}
{{- if $v.ThrottlingRateLimit }}
                      ThrottlingRateLimit: {{ $v.ThrottlingRateLimit }}
{{- end }}
{{- end }}
{{- end }}
{{- if .StageDescription.MetricsEnabled }}
                MetricsEnabled: {{ .StageDescription.MetricsEnabled }}
{{- end }}
{{- if .StageDescription.ThrottlingBurstLimit }}
                ThrottlingBurstLimit: {{ .StageDescription.ThrottlingBurstLimit }}
{{- end }}
{{- if .StageDescription.ThrottlingRateLimit }}
                ThrottlingRateLimit: {{ .StageDescription.ThrottlingRateLimit }}
{{- end }}
{{- if .StageDescription.Variables }}
                Variables:
{{- range $k, $v := .StageDescription.Variables }}
                    {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if .StageName }}
            StageName: {{ .StageName }}
{{- end }}
{{- end }}
{{- define "AWS::ApiGateway::DocumentationPart" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::DocumentationPart
        Properties:
            Location:
{{- if .Location.Method }}
                Method: {{ .Location.Method }}
{{- end }}
{{- if .Location.Name }}
                Name: {{ .Location.Name }}
{{- end }}
{{- if .Location.Path }}
                Path: {{ .Location.Path }}
{{- end }}
{{- if .Location.StatusCode }}
                StatusCode: {{ .Location.StatusCode }}
{{- end }}
{{- if .Location.Type }}
                Type: {{ .Location.Type }}
{{- end }}
            Properties: {{ .Properties }}
            RestApiId: {{ .RestApiId }}
{{- end }}
{{- define "AWS::ApiGateway::DocumentationVersion" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::DocumentationVersion
        Properties:
            DocumentationVersion: {{ .DocumentationVersion }}
            RestApiId: {{ .RestApiId }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- end }}
{{- define "AWS::ApiGateway::DomainName" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::DomainName
        Properties:
            DomainName: {{ .DomainName }}
{{- if .CertificateArn }}
            CertificateArn: {{ .CertificateArn }}
{{- end }}
{{- if .EndpointConfiguration }}
            EndpointConfiguration:
{{- if .EndpointConfiguration.Types }}
                Types:
{{- range $_, $v := .EndpointConfiguration.Types }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if .RegionalCertificateArn }}
            RegionalCertificateArn: {{ .RegionalCertificateArn }}
{{- end }}
{{- end }}
{{- define "AWS::ApiGateway::GatewayResponse" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::GatewayResponse
        Properties:
            ResponseType: {{ .ResponseType }}
            RestApiId: {{ .RestApiId }}
{{- if .ResponseParameters }}
            ResponseParameters:
{{- range $k, $v := .ResponseParameters }}
                {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .ResponseTemplates }}
            ResponseTemplates:
{{- range $k, $v := .ResponseTemplates }}
                {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .StatusCode }}
            StatusCode: {{ .StatusCode }}
{{- end }}
{{- end }}
{{- define "AWS::ApiGateway::Method" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::Method
        Properties:
            HttpMethod: {{ .HttpMethod }}
            ResourceId: {{ .ResourceId }}
            RestApiId: {{ .RestApiId }}
{{- if .ApiKeyRequired }}
            ApiKeyRequired: {{ .ApiKeyRequired }}
{{- end }}
{{- if .AuthorizationType }}
            AuthorizationType: {{ .AuthorizationType }}
{{- end }}
{{- if .AuthorizerId }}
            AuthorizerId: {{ .AuthorizerId }}
{{- end }}
{{- if .Integration }}
            Integration:
{{- if .Integration.CacheKeyParameters }}
                CacheKeyParameters:
{{- range $_, $v := .Integration.CacheKeyParameters }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- if .Integration.CacheNamespace }}
                CacheNamespace: {{ .Integration.CacheNamespace }}
{{- end }}
{{- if .Integration.ContentHandling }}
                ContentHandling: {{ .Integration.ContentHandling }}
{{- end }}
{{- if .Integration.Credentials }}
                Credentials: {{ .Integration.Credentials }}
{{- end }}
{{- if .Integration.IntegrationHttpMethod }}
                IntegrationHttpMethod: {{ .Integration.IntegrationHttpMethod }}
{{- end }}
{{- if .Integration.IntegrationResponses }}
                IntegrationResponses:
{{- range $_, $v := .Integration.IntegrationResponses }}
                    -
                      StatusCode: {{ $v.StatusCode }}
{{- if $v.ContentHandling }}
                      ContentHandling: {{ $v.ContentHandling }}
{{- end }}
{{- if $v.ResponseParameters }}
                      ResponseParameters:
{{- range $k, $v := $v.ResponseParameters }}
                          {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if $v.ResponseTemplates }}
                      ResponseTemplates:
{{- range $k, $v := $v.ResponseTemplates }}
                          {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if $v.SelectionPattern }}
                      SelectionPattern: {{ $v.SelectionPattern }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Integration.PassthroughBehavior }}
                PassthroughBehavior: {{ .Integration.PassthroughBehavior }}
{{- end }}
{{- if .Integration.RequestParameters }}
                RequestParameters:
{{- range $k, $v := .Integration.RequestParameters }}
                    {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .Integration.RequestTemplates }}
                RequestTemplates:
{{- range $k, $v := .Integration.RequestTemplates }}
                    {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .Integration.Type }}
                Type: {{ .Integration.Type }}
{{- end }}
{{- if .Integration.Uri }}
                Uri: {{ .Integration.Uri }}
{{- end }}
{{- end }}
{{- if .MethodResponses }}
            MethodResponses:
{{- range $_, $v := .MethodResponses }}
                -
                  StatusCode: {{ $v.StatusCode }}
{{- if $v.ResponseModels }}
                  ResponseModels:
{{- range $k, $v := $v.ResponseModels }}
                      {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if $v.ResponseParameters }}
                  ResponseParameters:
{{- range $k, $v := $v.ResponseParameters }}
                      {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .OperationName }}
            OperationName: {{ .OperationName }}
{{- end }}
{{- if .RequestModels }}
            RequestModels:
{{- range $k, $v := .RequestModels }}
                {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .RequestParameters }}
            RequestParameters:
{{- range $k, $v := .RequestParameters }}
                {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .RequestValidatorId }}
            RequestValidatorId: {{ .RequestValidatorId }}
{{- end }}
{{- end }}
{{- define "AWS::ApiGateway::Model" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::Model
        Properties:
            RestApiId: {{ .RestApiId }}
{{- if .ContentType }}
            ContentType: {{ .ContentType }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .Schema }}
            Schema:
                {{ yaml.MarshalIndent .Schema "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- end }}
{{- define "AWS::ApiGateway::RequestValidator" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::RequestValidator
        Properties:
            RestApiId: {{ .RestApiId }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .ValidateRequestBody }}
            ValidateRequestBody: {{ .ValidateRequestBody }}
{{- end }}
{{- if .ValidateRequestParameters }}
            ValidateRequestParameters: {{ .ValidateRequestParameters }}
{{- end }}
{{- end }}
{{- define "AWS::ApiGateway::Resource" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::Resource
        Properties:
            ParentId: {{ .ParentId }}
            PathPart: {{ .PathPart }}
            RestApiId: {{ .RestApiId }}
{{- end }}
{{- define "AWS::ApiGateway::RestApi" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::RestApi
{{- if .ApiKeySourceType | or .BinaryMediaTypes | or .Body | or .BodyS3Location | or .CloneFrom | or .Description | or .EndpointConfiguration | or .FailOnWarnings | or .MinimumCompressionSize | or .Name | or .Parameters }}
        Properties:
{{- if .ApiKeySourceType }}
            ApiKeySourceType: {{ .ApiKeySourceType }}
{{- end }}
{{- if .BinaryMediaTypes }}
            BinaryMediaTypes:
{{- range $_, $v := .BinaryMediaTypes }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Body }}
            Body:
                {{ yaml.MarshalIndent .Body "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .BodyS3Location }}
            BodyS3Location:
{{- if .BodyS3Location.Bucket }}
                Bucket: {{ .BodyS3Location.Bucket }}
{{- end }}
{{- if .BodyS3Location.ETag }}
                ETag: {{ .BodyS3Location.ETag }}
{{- end }}
{{- if .BodyS3Location.Key }}
                Key: {{ .BodyS3Location.Key }}
{{- end }}
{{- if .BodyS3Location.Version }}
                Version: {{ .BodyS3Location.Version }}
{{- end }}
{{- end }}
{{- if .CloneFrom }}
            CloneFrom: {{ .CloneFrom }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .EndpointConfiguration }}
            EndpointConfiguration:
{{- if .EndpointConfiguration.Types }}
                Types:
{{- range $_, $v := .EndpointConfiguration.Types }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if .FailOnWarnings }}
            FailOnWarnings: {{ .FailOnWarnings }}
{{- end }}
{{- if .MinimumCompressionSize }}
            MinimumCompressionSize: {{ .MinimumCompressionSize }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .Parameters }}
            Parameters:
{{- range $k, $v := .Parameters }}
                {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::ApiGateway::Stage" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::Stage
        Properties:
            RestApiId: {{ .RestApiId }}
{{- if .CacheClusterEnabled }}
            CacheClusterEnabled: {{ .CacheClusterEnabled }}
{{- end }}
{{- if .CacheClusterSize }}
            CacheClusterSize: {{ .CacheClusterSize }}
{{- end }}
{{- if .ClientCertificateId }}
            ClientCertificateId: {{ .ClientCertificateId }}
{{- end }}
{{- if .DeploymentId }}
            DeploymentId: {{ .DeploymentId }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .DocumentationVersion }}
            DocumentationVersion: {{ .DocumentationVersion }}
{{- end }}
{{- if .MethodSettings }}
            MethodSettings:
{{- range $_, $v := .MethodSettings }}
                -
{{- if $v.CacheDataEncrypted }}
                  CacheDataEncrypted: {{ $v.CacheDataEncrypted }}
{{- end }}
{{- if $v.CacheTtlInSeconds }}
                  CacheTtlInSeconds: {{ $v.CacheTtlInSeconds }}
{{- end }}
{{- if $v.CachingEnabled }}
                  CachingEnabled: {{ $v.CachingEnabled }}
{{- end }}
{{- if $v.DataTraceEnabled }}
                  DataTraceEnabled: {{ $v.DataTraceEnabled }}
{{- end }}
{{- if $v.HttpMethod }}
                  HttpMethod: {{ $v.HttpMethod }}
{{- end }}
{{- if $v.LoggingLevel }}
                  LoggingLevel: {{ $v.LoggingLevel }}
{{- end }}
{{- if $v.MetricsEnabled }}
                  MetricsEnabled: {{ $v.MetricsEnabled }}
{{- end }}
{{- if $v.ResourcePath }}
                  ResourcePath: {{ $v.ResourcePath }}
{{- end }}
{{- if $v.ThrottlingBurstLimit }}
                  ThrottlingBurstLimit: {{ $v.ThrottlingBurstLimit }}
{{- end }}
{{- if $v.ThrottlingRateLimit }}
                  ThrottlingRateLimit: {{ $v.ThrottlingRateLimit }}
{{- end }}
{{- end }}
{{- end }}
{{- if .StageName }}
            StageName: {{ .StageName }}
{{- end }}
{{- if .Variables }}
            Variables:
{{- range $k, $v := .Variables }}
                {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::ApiGateway::UsagePlan" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::UsagePlan
{{- if .ApiStages | or .Description | or .Quota | or .Throttle | or .UsagePlanName }}
        Properties:
{{- if .ApiStages }}
            ApiStages:
{{- range $_, $v := .ApiStages }}
                -
{{- if $v.ApiId }}
                  ApiId: {{ $v.ApiId }}
{{- end }}
{{- if $v.Stage }}
                  Stage: {{ $v.Stage }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .Quota }}
            Quota:
{{- if .Quota.Limit }}
                Limit: {{ .Quota.Limit }}
{{- end }}
{{- if .Quota.Offset }}
                Offset: {{ .Quota.Offset }}
{{- end }}
{{- if .Quota.Period }}
                Period: {{ .Quota.Period }}
{{- end }}
{{- end }}
{{- if .Throttle }}
            Throttle:
{{- if .Throttle.BurstLimit }}
                BurstLimit: {{ .Throttle.BurstLimit }}
{{- end }}
{{- if .Throttle.RateLimit }}
                RateLimit: {{ .Throttle.RateLimit }}
{{- end }}
{{- end }}
{{- if .UsagePlanName }}
            UsagePlanName: {{ .UsagePlanName }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::ApiGateway::UsagePlanKey" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::UsagePlanKey
        Properties:
            KeyId: {{ .KeyId }}
            KeyType: {{ .KeyType }}
            UsagePlanId: {{ .UsagePlanId }}
{{- end }}
{{- define "AWS::ApiGateway::VpcLink" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApiGateway::VpcLink
        Properties:
            Name: {{ .Name }}
            TargetArns:
{{- range $_, $v := .TargetArns }}
                - {{ $v }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- end }}
{{- define "AWS::ApplicationAutoScaling::ScalableTarget" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApplicationAutoScaling::ScalableTarget
        Properties:
            MaxCapacity: {{ .MaxCapacity }}
            MinCapacity: {{ .MinCapacity }}
            ResourceId: {{ .ResourceId }}
            RoleARN: {{ .RoleARN }}
            ScalableDimension: {{ .ScalableDimension }}
            ServiceNamespace: {{ .ServiceNamespace }}
{{- if .ScheduledActions }}
            ScheduledActions:
{{- range $_, $v := .ScheduledActions }}
                -
                  Schedule: {{ $v.Schedule }}
                  ScheduledActionName: {{ $v.ScheduledActionName }}
{{- if $v.EndTime }}
                  EndTime: {{ $v.EndTime }}
{{- end }}
{{- if $v.ScalableTargetAction }}
                  ScalableTargetAction:
{{- if $v.ScalableTargetAction.MaxCapacity }}
                      MaxCapacity: {{ $v.ScalableTargetAction.MaxCapacity }}
{{- end }}
{{- if $v.ScalableTargetAction.MinCapacity }}
                      MinCapacity: {{ $v.ScalableTargetAction.MinCapacity }}
{{- end }}
{{- end }}
{{- if $v.StartTime }}
                  StartTime: {{ $v.StartTime }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::ApplicationAutoScaling::ScalingPolicy" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ApplicationAutoScaling::ScalingPolicy
        Properties:
            PolicyName: {{ .PolicyName }}
            PolicyType: {{ .PolicyType }}
{{- if .ResourceId }}
            ResourceId: {{ .ResourceId }}
{{- end }}
{{- if .ScalableDimension }}
            ScalableDimension: {{ .ScalableDimension }}
{{- end }}
{{- if .ScalingTargetId }}
            ScalingTargetId: {{ .ScalingTargetId }}
{{- end }}
{{- if .ServiceNamespace }}
            ServiceNamespace: {{ .ServiceNamespace }}
{{- end }}
{{- if .StepScalingPolicyConfiguration }}
            StepScalingPolicyConfiguration:
{{- if .StepScalingPolicyConfiguration.AdjustmentType }}
                AdjustmentType: {{ .StepScalingPolicyConfiguration.AdjustmentType }}
{{- end }}
{{- if .StepScalingPolicyConfiguration.Cooldown }}
                Cooldown: {{ .StepScalingPolicyConfiguration.Cooldown }}
{{- end }}
{{- if .StepScalingPolicyConfiguration.MetricAggregationType }}
                MetricAggregationType: {{ .StepScalingPolicyConfiguration.MetricAggregationType }}
{{- end }}
{{- if .StepScalingPolicyConfiguration.MinAdjustmentMagnitude }}
                MinAdjustmentMagnitude: {{ .StepScalingPolicyConfiguration.MinAdjustmentMagnitude }}
{{- end }}
{{- if .StepScalingPolicyConfiguration.StepAdjustments }}
                StepAdjustments:
{{- range $_, $v := .StepScalingPolicyConfiguration.StepAdjustments }}
                    -
                      ScalingAdjustment: {{ $v.ScalingAdjustment }}
{{- if $v.MetricIntervalLowerBound }}
                      MetricIntervalLowerBound: {{ $v.MetricIntervalLowerBound }}
{{- end }}
{{- if $v.MetricIntervalUpperBound }}
                      MetricIntervalUpperBound: {{ $v.MetricIntervalUpperBound }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .TargetTrackingScalingPolicyConfiguration }}
            TargetTrackingScalingPolicyConfiguration:
                TargetValue: {{ .TargetTrackingScalingPolicyConfiguration.TargetValue }}
{{- if .TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification }}
                CustomizedMetricSpecification:
                    MetricName: {{ .TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.MetricName }}
                    Namespace: {{ .TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Namespace }}
                    Statistic: {{ .TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Statistic }}
{{- if .TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Dimensions }}
                    Dimensions:
{{- range $_, $v := .TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Dimensions }}
                        -
                          Name: {{ $v.Name }}
                          Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit }}
                    Unit: {{ .TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit }}
{{- end }}
{{- end }}
{{- if .TargetTrackingScalingPolicyConfiguration.DisableScaleIn }}
                DisableScaleIn: {{ .TargetTrackingScalingPolicyConfiguration.DisableScaleIn }}
{{- end }}
{{- if .TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification }}
                PredefinedMetricSpecification:
                    PredefinedMetricType: {{ .TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.PredefinedMetricType }}
{{- if .TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel }}
                    ResourceLabel: {{ .TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel }}
{{- end }}
{{- end }}
{{- if .TargetTrackingScalingPolicyConfiguration.ScaleInCooldown }}
                ScaleInCooldown: {{ .TargetTrackingScalingPolicyConfiguration.ScaleInCooldown }}
{{- end }}
{{- if .TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown }}
                ScaleOutCooldown: {{ .TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Athena::NamedQuery" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Athena::NamedQuery
        Properties:
            Database: {{ .Database }}
            QueryString: {{ .QueryString }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- end }}
{{- define "AWS::AutoScaling::AutoScalingGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::AutoScaling::AutoScalingGroup
        Properties:
            MaxSize: {{ .MaxSize }}
            MinSize: {{ .MinSize }}
{{- if .AutoScalingGroupName }}
            AutoScalingGroupName: {{ .AutoScalingGroupName }}
{{- end }}
{{- if .AvailabilityZones }}
            AvailabilityZones:
{{- range $_, $v := .AvailabilityZones }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Cooldown }}
            Cooldown: {{ .Cooldown }}
{{- end }}
{{- if .DesiredCapacity }}
            DesiredCapacity: {{ .DesiredCapacity }}
{{- end }}
{{- if .HealthCheckGracePeriod }}
            HealthCheckGracePeriod: {{ .HealthCheckGracePeriod }}
{{- end }}
{{- if .HealthCheckType }}
            HealthCheckType: {{ .HealthCheckType }}
{{- end }}
{{- if .InstanceId }}
            InstanceId: {{ .InstanceId }}
{{- end }}
{{- if .LaunchConfigurationName }}
            LaunchConfigurationName: {{ .LaunchConfigurationName }}
{{- end }}
{{- if .LifecycleHookSpecificationList }}
            LifecycleHookSpecificationList:
{{- range $_, $v := .LifecycleHookSpecificationList }}
                -
                  LifecycleHookName: {{ $v.LifecycleHookName }}
                  LifecycleTransition: {{ $v.LifecycleTransition }}
{{- if $v.DefaultResult }}
                  DefaultResult: {{ $v.DefaultResult }}
{{- end }}
{{- if $v.HeartbeatTimeout }}
                  HeartbeatTimeout: {{ $v.HeartbeatTimeout }}
{{- end }}
{{- if $v.NotificationMetadata }}
                  NotificationMetadata: {{ $v.NotificationMetadata }}
{{- end }}
{{- if $v.NotificationTargetARN }}
                  NotificationTargetARN: {{ $v.NotificationTargetARN }}
{{- end }}
{{- if $v.RoleARN }}
                  RoleARN: {{ $v.RoleARN }}
{{- end }}
{{- end }}
{{- end }}
{{- if .LoadBalancerNames }}
            LoadBalancerNames:
{{- range $_, $v := .LoadBalancerNames }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .MetricsCollection }}
            MetricsCollection:
{{- range $_, $v := .MetricsCollection }}
                -
                  Granularity: {{ $v.Granularity }}
{{- if $v.Metrics }}
                  Metrics:
{{- range $_, $v := $v.Metrics }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .NotificationConfigurations }}
            NotificationConfigurations:
{{- range $_, $v := .NotificationConfigurations }}
                -
                  TopicARN: {{ $v.TopicARN }}
{{- if $v.NotificationTypes }}
                  NotificationTypes:
{{- range $_, $v := $v.NotificationTypes }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .PlacementGroup }}
            PlacementGroup: {{ .PlacementGroup }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  PropagateAtLaunch: {{ $v.PropagateAtLaunch }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .TargetGroupARNs }}
            TargetGroupARNs:
{{- range $_, $v := .TargetGroupARNs }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .TerminationPolicies }}
            TerminationPolicies:
{{- range $_, $v := .TerminationPolicies }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .VPCZoneIdentifier }}
            VPCZoneIdentifier:
{{- range $_, $v := .VPCZoneIdentifier }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::AutoScaling::LaunchConfiguration" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::AutoScaling::LaunchConfiguration
        Properties:
            ImageId: {{ .ImageId }}
            InstanceType: {{ .InstanceType }}
{{- if .AssociatePublicIpAddress }}
            AssociatePublicIpAddress: {{ .AssociatePublicIpAddress }}
{{- end }}
{{- if .BlockDeviceMappings }}
            BlockDeviceMappings:
{{- range $_, $v := .BlockDeviceMappings }}
                -
                  DeviceName: {{ $v.DeviceName }}
{{- if $v.Ebs }}
                  Ebs:
{{- if $v.Ebs.DeleteOnTermination }}
                      DeleteOnTermination: {{ $v.Ebs.DeleteOnTermination }}
{{- end }}
{{- if $v.Ebs.Encrypted }}
                      Encrypted: {{ $v.Ebs.Encrypted }}
{{- end }}
{{- if $v.Ebs.Iops }}
                      Iops: {{ $v.Ebs.Iops }}
{{- end }}
{{- if $v.Ebs.SnapshotId }}
                      SnapshotId: {{ $v.Ebs.SnapshotId }}
{{- end }}
{{- if $v.Ebs.VolumeSize }}
                      VolumeSize: {{ $v.Ebs.VolumeSize }}
{{- end }}
{{- if $v.Ebs.VolumeType }}
                      VolumeType: {{ $v.Ebs.VolumeType }}
{{- end }}
{{- end }}
{{- if $v.NoDevice }}
                  NoDevice: {{ $v.NoDevice }}
{{- end }}
{{- if $v.VirtualName }}
                  VirtualName: {{ $v.VirtualName }}
{{- end }}
{{- end }}
{{- end }}
{{- if .ClassicLinkVPCId }}
            ClassicLinkVPCId: {{ .ClassicLinkVPCId }}
{{- end }}
{{- if .ClassicLinkVPCSecurityGroups }}
            ClassicLinkVPCSecurityGroups:
{{- range $_, $v := .ClassicLinkVPCSecurityGroups }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .EbsOptimized }}
            EbsOptimized: {{ .EbsOptimized }}
{{- end }}
{{- if .IamInstanceProfile }}
            IamInstanceProfile: {{ .IamInstanceProfile }}
{{- end }}
{{- if .InstanceId }}
            InstanceId: {{ .InstanceId }}
{{- end }}
{{- if .InstanceMonitoring }}
            InstanceMonitoring: {{ .InstanceMonitoring }}
{{- end }}
{{- if .KernelId }}
            KernelId: {{ .KernelId }}
{{- end }}
{{- if .KeyName }}
            KeyName: {{ .KeyName }}
{{- end }}
{{- if .PlacementTenancy }}
            PlacementTenancy: {{ .PlacementTenancy }}
{{- end }}
{{- if .RamDiskId }}
            RamDiskId: {{ .RamDiskId }}
{{- end }}
{{- if .SecurityGroups }}
            SecurityGroups:
{{- range $_, $v := .SecurityGroups }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .SpotPrice }}
            SpotPrice: {{ .SpotPrice }}
{{- end }}
{{- if .UserData }}
            UserData: {{ .UserData }}
{{- end }}
{{- end }}
{{- define "AWS::AutoScaling::LifecycleHook" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::AutoScaling::LifecycleHook
        Properties:
            AutoScalingGroupName: {{ .AutoScalingGroupName }}
            LifecycleTransition: {{ .LifecycleTransition }}
{{- if .DefaultResult }}
            DefaultResult: {{ .DefaultResult }}
{{- end }}
{{- if .HeartbeatTimeout }}
            HeartbeatTimeout: {{ .HeartbeatTimeout }}
{{- end }}
{{- if .LifecycleHookName }}
            LifecycleHookName: {{ .LifecycleHookName }}
{{- end }}
{{- if .NotificationMetadata }}
            NotificationMetadata: {{ .NotificationMetadata }}
{{- end }}
{{- if .NotificationTargetARN }}
            NotificationTargetARN: {{ .NotificationTargetARN }}
{{- end }}
{{- if .RoleARN }}
            RoleARN: {{ .RoleARN }}
{{- end }}
{{- end }}
{{- define "AWS::AutoScaling::ScalingPolicy" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::AutoScaling::ScalingPolicy
        Properties:
            AutoScalingGroupName: {{ .AutoScalingGroupName }}
{{- if .AdjustmentType }}
            AdjustmentType: {{ .AdjustmentType }}
{{- end }}
{{- if .Cooldown }}
            Cooldown: {{ .Cooldown }}
{{- end }}
{{- if .EstimatedInstanceWarmup }}
            EstimatedInstanceWarmup: {{ .EstimatedInstanceWarmup }}
{{- end }}
{{- if .MetricAggregationType }}
            MetricAggregationType: {{ .MetricAggregationType }}
{{- end }}
{{- if .MinAdjustmentMagnitude }}
            MinAdjustmentMagnitude: {{ .MinAdjustmentMagnitude }}
{{- end }}
{{- if .PolicyType }}
            PolicyType: {{ .PolicyType }}
{{- end }}
{{- if .ScalingAdjustment }}
            ScalingAdjustment: {{ .ScalingAdjustment }}
{{- end }}
{{- if .StepAdjustments }}
            StepAdjustments:
{{- range $_, $v := .StepAdjustments }}
                -
                  ScalingAdjustment: {{ $v.ScalingAdjustment }}
{{- if $v.MetricIntervalLowerBound }}
                  MetricIntervalLowerBound: {{ $v.MetricIntervalLowerBound }}
{{- end }}
{{- if $v.MetricIntervalUpperBound }}
                  MetricIntervalUpperBound: {{ $v.MetricIntervalUpperBound }}
{{- end }}
{{- end }}
{{- end }}
{{- if .TargetTrackingConfiguration }}
            TargetTrackingConfiguration:
                TargetValue: {{ .TargetTrackingConfiguration.TargetValue }}
{{- if .TargetTrackingConfiguration.CustomizedMetricSpecification }}
                CustomizedMetricSpecification:
                    MetricName: {{ .TargetTrackingConfiguration.CustomizedMetricSpecification.MetricName }}
                    Namespace: {{ .TargetTrackingConfiguration.CustomizedMetricSpecification.Namespace }}
                    Statistic: {{ .TargetTrackingConfiguration.CustomizedMetricSpecification.Statistic }}
{{- if .TargetTrackingConfiguration.CustomizedMetricSpecification.Dimensions }}
                    Dimensions:
{{- range $_, $v := .TargetTrackingConfiguration.CustomizedMetricSpecification.Dimensions }}
                        -
                          Name: {{ $v.Name }}
                          Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .TargetTrackingConfiguration.CustomizedMetricSpecification.Unit }}
                    Unit: {{ .TargetTrackingConfiguration.CustomizedMetricSpecification.Unit }}
{{- end }}
{{- end }}
{{- if .TargetTrackingConfiguration.DisableScaleIn }}
                DisableScaleIn: {{ .TargetTrackingConfiguration.DisableScaleIn }}
{{- end }}
{{- if .TargetTrackingConfiguration.PredefinedMetricSpecification }}
                PredefinedMetricSpecification:
                    PredefinedMetricType: {{ .TargetTrackingConfiguration.PredefinedMetricSpecification.PredefinedMetricType }}
{{- if .TargetTrackingConfiguration.PredefinedMetricSpecification.ResourceLabel }}
                    ResourceLabel: {{ .TargetTrackingConfiguration.PredefinedMetricSpecification.ResourceLabel }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::AutoScaling::ScheduledAction" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::AutoScaling::ScheduledAction
        Properties:
            AutoScalingGroupName: {{ .AutoScalingGroupName }}
{{- if .DesiredCapacity }}
            DesiredCapacity: {{ .DesiredCapacity }}
{{- end }}
{{- if .EndTime }}
            EndTime: {{ .EndTime }}
{{- end }}
{{- if .MaxSize }}
            MaxSize: {{ .MaxSize }}
{{- end }}
{{- if .MinSize }}
            MinSize: {{ .MinSize }}
{{- end }}
{{- if .Recurrence }}
            Recurrence: {{ .Recurrence }}
{{- end }}
{{- if .StartTime }}
            StartTime: {{ .StartTime }}
{{- end }}
{{- end }}
{{- define "AWS::Batch::ComputeEnvironment" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Batch::ComputeEnvironment
        Properties:
            ComputeResources:
                InstanceRole: {{ .ComputeResources.InstanceRole }}
                InstanceTypes:
{{- range $_, $v := .ComputeResources.InstanceTypes }}
                    - {{ $v }}
{{- end }}
                MaxvCpus: {{ .ComputeResources.MaxvCpus }}
                MinvCpus: {{ .ComputeResources.MinvCpus }}
                SecurityGroupIds:
{{- range $_, $v := .ComputeResources.SecurityGroupIds }}
                    - {{ $v }}
{{- end }}
                Subnets:
{{- range $_, $v := .ComputeResources.Subnets }}
                    - {{ $v }}
{{- end }}
                Type: {{ .ComputeResources.Type }}
{{- if .ComputeResources.BidPercentage }}
                BidPercentage: {{ .ComputeResources.BidPercentage }}
{{- end }}
{{- if .ComputeResources.DesiredvCpus }}
                DesiredvCpus: {{ .ComputeResources.DesiredvCpus }}
{{- end }}
{{- if .ComputeResources.Ec2KeyPair }}
                Ec2KeyPair: {{ .ComputeResources.Ec2KeyPair }}
{{- end }}
{{- if .ComputeResources.ImageId }}
                ImageId: {{ .ComputeResources.ImageId }}
{{- end }}
{{- if .ComputeResources.SpotIamFleetRole }}
                SpotIamFleetRole: {{ .ComputeResources.SpotIamFleetRole }}
{{- end }}
{{- if .ComputeResources.Tags }}
                Tags:
                    {{ yaml.MarshalIndent .ComputeResources.Tags "                    " "    " | fmt.Sprintf "%s" }}
{{- end }}
            ServiceRole: {{ .ServiceRole }}
            Type: {{ .Type }}
{{- if .ComputeEnvironmentName }}
            ComputeEnvironmentName: {{ .ComputeEnvironmentName }}
{{- end }}
{{- if .State }}
            State: {{ .State }}
{{- end }}
{{- end }}
{{- define "AWS::Batch::JobDefinition" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Batch::JobDefinition
        Properties:
            ContainerProperties:
                Image: {{ .ContainerProperties.Image }}
                Memory: {{ .ContainerProperties.Memory }}
                Vcpus: {{ .ContainerProperties.Vcpus }}
{{- if .ContainerProperties.Command }}
                Command:
{{- range $_, $v := .ContainerProperties.Command }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- if .ContainerProperties.Environment }}
                Environment:
{{- range $_, $v := .ContainerProperties.Environment }}
                    -
{{- if $v.Name }}
                      Name: {{ $v.Name }}
{{- end }}
{{- if $v.Value }}
                      Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- if .ContainerProperties.JobRoleArn }}
                JobRoleArn: {{ .ContainerProperties.JobRoleArn }}
{{- end }}
{{- if .ContainerProperties.MountPoints }}
                MountPoints:
{{- range $_, $v := .ContainerProperties.MountPoints }}
                    -
{{- if $v.ContainerPath }}
                      ContainerPath: {{ $v.ContainerPath }}
{{- end }}
{{- if $v.ReadOnly }}
                      ReadOnly: {{ $v.ReadOnly }}
{{- end }}
{{- if $v.SourceVolume }}
                      SourceVolume: {{ $v.SourceVolume }}
{{- end }}
{{- end }}
{{- end }}
{{- if .ContainerProperties.Privileged }}
                Privileged: {{ .ContainerProperties.Privileged }}
{{- end }}
{{- if .ContainerProperties.ReadonlyRootFilesystem }}
                ReadonlyRootFilesystem: {{ .ContainerProperties.ReadonlyRootFilesystem }}
{{- end }}
{{- if .ContainerProperties.Ulimits }}
                Ulimits:
{{- range $_, $v := .ContainerProperties.Ulimits }}
                    -
                      HardLimit: {{ $v.HardLimit }}
                      Name: {{ $v.Name }}
                      SoftLimit: {{ $v.SoftLimit }}
{{- end }}
{{- end }}
{{- if .ContainerProperties.User }}
                User: {{ .ContainerProperties.User }}
{{- end }}
{{- if .ContainerProperties.Volumes }}
                Volumes:
{{- range $_, $v := .ContainerProperties.Volumes }}
                    -
{{- if $v.Host }}
                      Host:
{{- if $v.Host.SourcePath }}
                          SourcePath: {{ $v.Host.SourcePath }}
{{- end }}
{{- end }}
{{- if $v.Name }}
                      Name: {{ $v.Name }}
{{- end }}
{{- end }}
{{- end }}
            Type: {{ .Type }}
{{- if .JobDefinitionName }}
            JobDefinitionName: {{ .JobDefinitionName }}
{{- end }}
{{- if .Parameters }}
            Parameters:
                {{ yaml.MarshalIndent .Parameters "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .RetryStrategy }}
            RetryStrategy:
{{- if .RetryStrategy.Attempts }}
                Attempts: {{ .RetryStrategy.Attempts }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Batch::JobQueue" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Batch::JobQueue
        Properties:
            ComputeEnvironmentOrder:
{{- range $_, $v := .ComputeEnvironmentOrder }}
                -
                  ComputeEnvironment: {{ $v.ComputeEnvironment }}
                  Order: {{ $v.Order }}
{{- end }}
            Priority: {{ .Priority }}
{{- if .JobQueueName }}
            JobQueueName: {{ .JobQueueName }}
{{- end }}
{{- if .State }}
            State: {{ .State }}
{{- end }}
{{- end }}
{{- define "AWS::CertificateManager::Certificate" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CertificateManager::Certificate
        Properties:
            DomainName: {{ .DomainName }}
{{- if .DomainValidationOptions }}
            DomainValidationOptions:
{{- range $_, $v := .DomainValidationOptions }}
                -
                  DomainName: {{ $v.DomainName }}
                  ValidationDomain: {{ $v.ValidationDomain }}
{{- end }}
{{- end }}
{{- if .SubjectAlternativeNames }}
            SubjectAlternativeNames:
{{- range $_, $v := .SubjectAlternativeNames }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Cloud9::EnvironmentEC2" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Cloud9::EnvironmentEC2
        Properties:
            InstanceType: {{ .InstanceType }}
{{- if .AutomaticStopTimeMinutes }}
            AutomaticStopTimeMinutes: {{ .AutomaticStopTimeMinutes }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .OwnerArn }}
            OwnerArn: {{ .OwnerArn }}
{{- end }}
{{- if .Repositories }}
            Repositories:
{{- range $_, $v := .Repositories }}
                -
                  PathComponent: {{ $v.PathComponent }}
                  RepositoryUrl: {{ $v.RepositoryUrl }}
{{- end }}
{{- end }}
{{- if .SubnetId }}
            SubnetId: {{ .SubnetId }}
{{- end }}
{{- end }}
{{- define "AWS::CloudFormation::CustomResource" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CloudFormation::CustomResource
        Properties:
            ServiceToken: {{ .ServiceToken }}
{{- end }}
{{- define "AWS::CloudFormation::Stack" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CloudFormation::Stack
        Properties:
            TemplateURL:{{ if istype .TemplateURL "string" }} {{ .TemplateURL }}{{ else }} {{ .TemplateURL }}
{{- end }}
{{- if .NotificationARNs }}
            NotificationARNs:
{{- range $_, $v := .NotificationARNs }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Parameters }}
            Parameters:
{{- range $k, $v := .Parameters }}
                {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .TimeoutInMinutes }}
            TimeoutInMinutes: {{ .TimeoutInMinutes }}
{{- end }}
{{- end }}
{{- define "AWS::CloudFormation::WaitCondition" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CloudFormation::WaitCondition
        Properties:
            Handle: {{ .Handle }}
            Timeout: {{ .Timeout }}
{{- if .Count }}
            Count: {{ .Count }}
{{- end }}
{{- end }}
{{- define "AWS::CloudFormation::WaitConditionHandle" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CloudFormation::WaitConditionHandle
{{- end }}
{{- define "AWS::CloudFront::CloudFrontOriginAccessIdentity" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CloudFront::CloudFrontOriginAccessIdentity
        Properties:
            CloudFrontOriginAccessIdentityConfig:
                Comment: {{ .CloudFrontOriginAccessIdentityConfig.Comment }}
{{- end }}
{{- define "AWS::CloudFront::Distribution" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CloudFront::Distribution
        Properties:
            DistributionConfig:
                Enabled: {{ .DistributionConfig.Enabled }}
{{- if .DistributionConfig.Aliases }}
                Aliases:
{{- range $_, $v := .DistributionConfig.Aliases }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- if .DistributionConfig.CacheBehaviors }}
                CacheBehaviors:
{{- range $_, $v := .DistributionConfig.CacheBehaviors }}
                    -
                      ForwardedValues:
                          QueryString: {{ $v.ForwardedValues.QueryString }}
{{- if $v.ForwardedValues.Cookies }}
                          Cookies:
                              Forward: {{ $v.ForwardedValues.Cookies.Forward }}
{{- if $v.ForwardedValues.Cookies.WhitelistedNames }}
                              WhitelistedNames:
{{- range $_, $v := $v.ForwardedValues.Cookies.WhitelistedNames }}
                                  - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if $v.ForwardedValues.Headers }}
                          Headers:
{{- range $_, $v := $v.ForwardedValues.Headers }}
                              - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.ForwardedValues.QueryStringCacheKeys }}
                          QueryStringCacheKeys:
{{- range $_, $v := $v.ForwardedValues.QueryStringCacheKeys }}
                              - {{ $v }}
{{- end }}
{{- end }}
                      PathPattern: {{ $v.PathPattern }}
                      TargetOriginId: {{ $v.TargetOriginId }}
                      ViewerProtocolPolicy: {{ $v.ViewerProtocolPolicy }}
{{- if $v.AllowedMethods }}
                      AllowedMethods:
{{- range $_, $v := $v.AllowedMethods }}
                          - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.CachedMethods }}
                      CachedMethods:
{{- range $_, $v := $v.CachedMethods }}
                          - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.Compress }}
                      Compress: {{ $v.Compress }}
{{- end }}
{{- if $v.DefaultTTL }}
                      DefaultTTL: {{ $v.DefaultTTL }}
{{- end }}
{{- if $v.LambdaFunctionAssociations }}
                      LambdaFunctionAssociations:
{{- range $_, $v := $v.LambdaFunctionAssociations }}
                          -
{{- if $v.EventType }}
                            EventType: {{ $v.EventType }}
{{- end }}
{{- if $v.LambdaFunctionARN }}
                            LambdaFunctionARN: {{ $v.LambdaFunctionARN }}
{{- end }}
{{- end }}
{{- end }}
{{- if $v.MaxTTL }}
                      MaxTTL: {{ $v.MaxTTL }}
{{- end }}
{{- if $v.MinTTL }}
                      MinTTL: {{ $v.MinTTL }}
{{- end }}
{{- if $v.SmoothStreaming }}
                      SmoothStreaming: {{ $v.SmoothStreaming }}
{{- end }}
{{- if $v.TrustedSigners }}
                      TrustedSigners:
{{- range $_, $v := $v.TrustedSigners }}
                          - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .DistributionConfig.Comment }}
                Comment: {{ .DistributionConfig.Comment }}
{{- end }}
{{- if .DistributionConfig.CustomErrorResponses }}
                CustomErrorResponses:
{{- range $_, $v := .DistributionConfig.CustomErrorResponses }}
                    -
                      ErrorCode: {{ $v.ErrorCode }}
{{- if $v.ErrorCachingMinTTL }}
                      ErrorCachingMinTTL: {{ $v.ErrorCachingMinTTL }}
{{- end }}
{{- if $v.ResponseCode }}
                      ResponseCode: {{ $v.ResponseCode }}
{{- end }}
{{- if $v.ResponsePagePath }}
                      ResponsePagePath: {{ $v.ResponsePagePath }}
{{- end }}
{{- end }}
{{- end }}
{{- if .DistributionConfig.DefaultCacheBehavior }}
                DefaultCacheBehavior:
                    ForwardedValues:
                        QueryString: {{ .DistributionConfig.DefaultCacheBehavior.ForwardedValues.QueryString }}
{{- if .DistributionConfig.DefaultCacheBehavior.ForwardedValues.Cookies }}
                        Cookies:
                            Forward: {{ .DistributionConfig.DefaultCacheBehavior.ForwardedValues.Cookies.Forward }}
{{- if .DistributionConfig.DefaultCacheBehavior.ForwardedValues.Cookies.WhitelistedNames }}
                            WhitelistedNames:
{{- range $_, $v := .DistributionConfig.DefaultCacheBehavior.ForwardedValues.Cookies.WhitelistedNames }}
                                - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if .DistributionConfig.DefaultCacheBehavior.ForwardedValues.Headers }}
                        Headers:
{{- range $_, $v := .DistributionConfig.DefaultCacheBehavior.ForwardedValues.Headers }}
                            - {{ $v }}
{{- end }}
{{- end }}
{{- if .DistributionConfig.DefaultCacheBehavior.ForwardedValues.QueryStringCacheKeys }}
                        QueryStringCacheKeys:
{{- range $_, $v := .DistributionConfig.DefaultCacheBehavior.ForwardedValues.QueryStringCacheKeys }}
                            - {{ $v }}
{{- end }}
{{- end }}
                    TargetOriginId: {{ .DistributionConfig.DefaultCacheBehavior.TargetOriginId }}
                    ViewerProtocolPolicy: {{ .DistributionConfig.DefaultCacheBehavior.ViewerProtocolPolicy }}
{{- if .DistributionConfig.DefaultCacheBehavior.AllowedMethods }}
                    AllowedMethods:
{{- range $_, $v := .DistributionConfig.DefaultCacheBehavior.AllowedMethods }}
                        - {{ $v }}
{{- end }}
{{- end }}
{{- if .DistributionConfig.DefaultCacheBehavior.CachedMethods }}
                    CachedMethods:
{{- range $_, $v := .DistributionConfig.DefaultCacheBehavior.CachedMethods }}
                        - {{ $v }}
{{- end }}
{{- end }}
{{- if .DistributionConfig.DefaultCacheBehavior.Compress }}
                    Compress: {{ .DistributionConfig.DefaultCacheBehavior.Compress }}
{{- end }}
{{- if .DistributionConfig.DefaultCacheBehavior.DefaultTTL }}
                    DefaultTTL: {{ .DistributionConfig.DefaultCacheBehavior.DefaultTTL }}
{{- end }}
{{- if .DistributionConfig.DefaultCacheBehavior.LambdaFunctionAssociations }}
                    LambdaFunctionAssociations:
{{- range $_, $v := .DistributionConfig.DefaultCacheBehavior.LambdaFunctionAssociations }}
                        -
{{- if $v.EventType }}
                          EventType: {{ $v.EventType }}
{{- end }}
{{- if $v.LambdaFunctionARN }}
                          LambdaFunctionARN: {{ $v.LambdaFunctionARN }}
{{- end }}
{{- end }}
{{- end }}
{{- if .DistributionConfig.DefaultCacheBehavior.MaxTTL }}
                    MaxTTL: {{ .DistributionConfig.DefaultCacheBehavior.MaxTTL }}
{{- end }}
{{- if .DistributionConfig.DefaultCacheBehavior.MinTTL }}
                    MinTTL: {{ .DistributionConfig.DefaultCacheBehavior.MinTTL }}
{{- end }}
{{- if .DistributionConfig.DefaultCacheBehavior.SmoothStreaming }}
                    SmoothStreaming: {{ .DistributionConfig.DefaultCacheBehavior.SmoothStreaming }}
{{- end }}
{{- if .DistributionConfig.DefaultCacheBehavior.TrustedSigners }}
                    TrustedSigners:
{{- range $_, $v := .DistributionConfig.DefaultCacheBehavior.TrustedSigners }}
                        - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if .DistributionConfig.DefaultRootObject }}
                DefaultRootObject: {{ .DistributionConfig.DefaultRootObject }}
{{- end }}
{{- if .DistributionConfig.HttpVersion }}
                HttpVersion: {{ .DistributionConfig.HttpVersion }}
{{- end }}
{{- if .DistributionConfig.IPV6Enabled }}
                IPV6Enabled: {{ .DistributionConfig.IPV6Enabled }}
{{- end }}
{{- if .DistributionConfig.Logging }}
                Logging:
                    Bucket: {{ .DistributionConfig.Logging.Bucket }}
{{- if .DistributionConfig.Logging.IncludeCookies }}
                    IncludeCookies: {{ .DistributionConfig.Logging.IncludeCookies }}
{{- end }}
{{- if .DistributionConfig.Logging.Prefix }}
                    Prefix: {{ .DistributionConfig.Logging.Prefix }}
{{- end }}
{{- end }}
{{- if .DistributionConfig.Origins }}
                Origins:
{{- range $_, $v := .DistributionConfig.Origins }}
                    -
                      DomainName: {{ $v.DomainName }}
                      Id: {{ $v.Id }}
{{- if $v.CustomOriginConfig }}
                      CustomOriginConfig:
                          OriginProtocolPolicy: {{ $v.CustomOriginConfig.OriginProtocolPolicy }}
{{- if $v.CustomOriginConfig.HTTPPort }}
                          HTTPPort: {{ $v.CustomOriginConfig.HTTPPort }}
{{- end }}
{{- if $v.CustomOriginConfig.HTTPSPort }}
                          HTTPSPort: {{ $v.CustomOriginConfig.HTTPSPort }}
{{- end }}
{{- if $v.CustomOriginConfig.OriginKeepaliveTimeout }}
                          OriginKeepaliveTimeout: {{ $v.CustomOriginConfig.OriginKeepaliveTimeout }}
{{- end }}
{{- if $v.CustomOriginConfig.OriginReadTimeout }}
                          OriginReadTimeout: {{ $v.CustomOriginConfig.OriginReadTimeout }}
{{- end }}
{{- if $v.CustomOriginConfig.OriginSSLProtocols }}
                          OriginSSLProtocols:
{{- range $_, $v := $v.CustomOriginConfig.OriginSSLProtocols }}
                              - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if $v.OriginCustomHeaders }}
                      OriginCustomHeaders:
{{- range $_, $v := $v.OriginCustomHeaders }}
                          -
                            HeaderName: {{ $v.HeaderName }}
                            HeaderValue: {{ $v.HeaderValue }}
{{- end }}
{{- end }}
{{- if $v.OriginPath }}
                      OriginPath: {{ $v.OriginPath }}
{{- end }}
{{- if $v.S3OriginConfig }}
                      S3OriginConfig:
{{- if $v.S3OriginConfig.OriginAccessIdentity }}
                          OriginAccessIdentity: {{ $v.S3OriginConfig.OriginAccessIdentity }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .DistributionConfig.PriceClass }}
                PriceClass: {{ .DistributionConfig.PriceClass }}
{{- end }}
{{- if .DistributionConfig.Restrictions }}
                Restrictions:
                    GeoRestriction:
                        RestrictionType: {{ .DistributionConfig.Restrictions.GeoRestriction.RestrictionType }}
{{- if .DistributionConfig.Restrictions.GeoRestriction.Locations }}
                        Locations:
{{- range $_, $v := .DistributionConfig.Restrictions.GeoRestriction.Locations }}
                            - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if .DistributionConfig.ViewerCertificate }}
                ViewerCertificate:
{{- if .DistributionConfig.ViewerCertificate.AcmCertificateArn }}
                    AcmCertificateArn: {{ .DistributionConfig.ViewerCertificate.AcmCertificateArn }}
{{- end }}
{{- if .DistributionConfig.ViewerCertificate.CloudFrontDefaultCertificate }}
                    CloudFrontDefaultCertificate: {{ .DistributionConfig.ViewerCertificate.CloudFrontDefaultCertificate }}
{{- end }}
{{- if .DistributionConfig.ViewerCertificate.IamCertificateId }}
                    IamCertificateId: {{ .DistributionConfig.ViewerCertificate.IamCertificateId }}
{{- end }}
{{- if .DistributionConfig.ViewerCertificate.MinimumProtocolVersion }}
                    MinimumProtocolVersion: {{ .DistributionConfig.ViewerCertificate.MinimumProtocolVersion }}
{{- end }}
{{- if .DistributionConfig.ViewerCertificate.SslSupportMethod }}
                    SslSupportMethod: {{ .DistributionConfig.ViewerCertificate.SslSupportMethod }}
{{- end }}
{{- end }}
{{- if .DistributionConfig.WebACLId }}
                WebACLId: {{ .DistributionConfig.WebACLId }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::CloudFront::StreamingDistribution" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CloudFront::StreamingDistribution
        Properties:
            StreamingDistributionConfig:
                Comment: {{ .StreamingDistributionConfig.Comment }}
                Enabled: {{ .StreamingDistributionConfig.Enabled }}
                S3Origin:
                    DomainName: {{ .StreamingDistributionConfig.S3Origin.DomainName }}
                    OriginAccessIdentity: {{ .StreamingDistributionConfig.S3Origin.OriginAccessIdentity }}
                TrustedSigners:
                    Enabled: {{ .StreamingDistributionConfig.TrustedSigners.Enabled }}
{{- if .StreamingDistributionConfig.TrustedSigners.AwsAccountNumbers }}
                    AwsAccountNumbers:
{{- range $_, $v := .StreamingDistributionConfig.TrustedSigners.AwsAccountNumbers }}
                        - {{ $v }}
{{- end }}
{{- end }}
{{- if .StreamingDistributionConfig.Aliases }}
                Aliases:
{{- range $_, $v := .StreamingDistributionConfig.Aliases }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- if .StreamingDistributionConfig.Logging }}
                Logging:
                    Bucket: {{ .StreamingDistributionConfig.Logging.Bucket }}
                    Enabled: {{ .StreamingDistributionConfig.Logging.Enabled }}
                    Prefix: {{ .StreamingDistributionConfig.Logging.Prefix }}
{{- end }}
{{- if .StreamingDistributionConfig.PriceClass }}
                PriceClass: {{ .StreamingDistributionConfig.PriceClass }}
{{- end }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- define "AWS::CloudTrail::Trail" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CloudTrail::Trail
        Properties:
            IsLogging: {{ .IsLogging }}
            S3BucketName: {{ .S3BucketName }}
{{- if .CloudWatchLogsLogGroupArn }}
            CloudWatchLogsLogGroupArn: {{ .CloudWatchLogsLogGroupArn }}
{{- end }}
{{- if .CloudWatchLogsRoleArn }}
            CloudWatchLogsRoleArn: {{ .CloudWatchLogsRoleArn }}
{{- end }}
{{- if .EnableLogFileValidation }}
            EnableLogFileValidation: {{ .EnableLogFileValidation }}
{{- end }}
{{- if .EventSelectors }}
            EventSelectors:
{{- range $_, $v := .EventSelectors }}
                -
{{- if $v.DataResources }}
                  DataResources:
{{- range $_, $v := $v.DataResources }}
                      -
                        Type: {{ $v.Type }}
{{- if $v.Values }}
                        Values:
{{- range $_, $v := $v.Values }}
                            - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if $v.IncludeManagementEvents }}
                  IncludeManagementEvents: {{ $v.IncludeManagementEvents }}
{{- end }}
{{- if $v.ReadWriteType }}
                  ReadWriteType: {{ $v.ReadWriteType }}
{{- end }}
{{- end }}
{{- end }}
{{- if .IncludeGlobalServiceEvents }}
            IncludeGlobalServiceEvents: {{ .IncludeGlobalServiceEvents }}
{{- end }}
{{- if .IsMultiRegionTrail }}
            IsMultiRegionTrail: {{ .IsMultiRegionTrail }}
{{- end }}
{{- if .KMSKeyId }}
            KMSKeyId: {{ .KMSKeyId }}
{{- end }}
{{- if .S3KeyPrefix }}
            S3KeyPrefix: {{ .S3KeyPrefix }}
{{- end }}
{{- if .SnsTopicName }}
            SnsTopicName: {{ .SnsTopicName }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .TrailName }}
            TrailName: {{ .TrailName }}
{{- end }}
{{- end }}
{{- define "AWS::CloudWatch::Alarm" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CloudWatch::Alarm
        Properties:
            ComparisonOperator: {{ .ComparisonOperator }}
            EvaluationPeriods: {{ .EvaluationPeriods }}
            MetricName: {{ .MetricName }}
            Namespace: {{ .Namespace }}
            Period: {{ .Period }}
            Threshold: {{ .Threshold }}
{{- if .ActionsEnabled }}
            ActionsEnabled: {{ .ActionsEnabled }}
{{- end }}
{{- if .AlarmActions }}
            AlarmActions:
{{- range $_, $v := .AlarmActions }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .AlarmDescription }}
            AlarmDescription: {{ .AlarmDescription }}
{{- end }}
{{- if .AlarmName }}
            AlarmName: {{ .AlarmName }}
{{- end }}
{{- if .Dimensions }}
            Dimensions:
{{- range $_, $v := .Dimensions }}
                -
                  Name: {{ $v.Name }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .EvaluateLowSampleCountPercentile }}
            EvaluateLowSampleCountPercentile: {{ .EvaluateLowSampleCountPercentile }}
{{- end }}
{{- if .ExtendedStatistic }}
            ExtendedStatistic: {{ .ExtendedStatistic }}
{{- end }}
{{- if .InsufficientDataActions }}
            InsufficientDataActions:
{{- range $_, $v := .InsufficientDataActions }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .OKActions }}
            OKActions:
{{- range $_, $v := .OKActions }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Statistic }}
            Statistic: {{ .Statistic }}
{{- end }}
{{- if .TreatMissingData }}
            TreatMissingData: {{ .TreatMissingData }}
{{- end }}
{{- if .Unit }}
            Unit: {{ .Unit }}
{{- end }}
{{- end }}
{{- define "AWS::CloudWatch::Dashboard" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CloudWatch::Dashboard
        Properties:
            DashboardBody: {{ .DashboardBody }}
{{- if .DashboardName }}
            DashboardName: {{ .DashboardName }}
{{- end }}
{{- end }}
{{- define "AWS::CodeBuild::Project" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CodeBuild::Project
        Properties:
            Artifacts:
                Type: {{ .Artifacts.Type }}
{{- if .Artifacts.Location }}
                Location: {{ .Artifacts.Location }}
{{- end }}
{{- if .Artifacts.Name }}
                Name: {{ .Artifacts.Name }}
{{- end }}
{{- if .Artifacts.NamespaceType }}
                NamespaceType: {{ .Artifacts.NamespaceType }}
{{- end }}
{{- if .Artifacts.Packaging }}
                Packaging: {{ .Artifacts.Packaging }}
{{- end }}
{{- if .Artifacts.Path }}
                Path: {{ .Artifacts.Path }}
{{- end }}
            Environment:
                ComputeType: {{ .Environment.ComputeType }}
                Image: {{ .Environment.Image }}
                Type: {{ .Environment.Type }}
{{- if .Environment.EnvironmentVariables }}
                EnvironmentVariables:
{{- range $_, $v := .Environment.EnvironmentVariables }}
                    -
                      Name: {{ $v.Name }}
                      Value: {{ $v.Value }}
{{- if $v.Type }}
                      Type: {{ $v.Type }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Environment.PrivilegedMode }}
                PrivilegedMode: {{ .Environment.PrivilegedMode }}
{{- end }}
            ServiceRole: {{ .ServiceRole }}
            Source:
                Type: {{ .Source.Type }}
{{- if .Source.Auth }}
                Auth:
                    Type: {{ .Source.Auth.Type }}
{{- if .Source.Auth.Resource }}
                    Resource: {{ .Source.Auth.Resource }}
{{- end }}
{{- end }}
{{- if .Source.BuildSpec }}
                BuildSpec: {{ .Source.BuildSpec }}
{{- end }}
{{- if .Source.GitCloneDepth }}
                GitCloneDepth: {{ .Source.GitCloneDepth }}
{{- end }}
{{- if .Source.InsecureSsl }}
                InsecureSsl: {{ .Source.InsecureSsl }}
{{- end }}
{{- if .Source.Location }}
                Location: {{ .Source.Location }}
{{- end }}
{{- if .BadgeEnabled }}
            BadgeEnabled: {{ .BadgeEnabled }}
{{- end }}
{{- if .Cache }}
            Cache:
                Type: {{ .Cache.Type }}
{{- if .Cache.Location }}
                Location: {{ .Cache.Location }}
{{- end }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .EncryptionKey }}
            EncryptionKey: {{ .EncryptionKey }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .TimeoutInMinutes }}
            TimeoutInMinutes: {{ .TimeoutInMinutes }}
{{- end }}
{{- if .Triggers }}
            Triggers:
{{- if .Triggers.Webhook }}
                Webhook: {{ .Triggers.Webhook }}
{{- end }}
{{- end }}
{{- if .VpcConfig }}
            VpcConfig:
                SecurityGroupIds:
{{- range $_, $v := .VpcConfig.SecurityGroupIds }}
                    - {{ $v }}
{{- end }}
                Subnets:
{{- range $_, $v := .VpcConfig.Subnets }}
                    - {{ $v }}
{{- end }}
                VpcId: {{ .VpcConfig.VpcId }}
{{- end }}
{{- end }}
{{- define "AWS::CodeCommit::Repository" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CodeCommit::Repository
        Properties:
            RepositoryName: {{ .RepositoryName }}
{{- if .RepositoryDescription }}
            RepositoryDescription: {{ .RepositoryDescription }}
{{- end }}
{{- if .Triggers }}
            Triggers:
{{- range $_, $v := .Triggers }}
                -
{{- if $v.Branches }}
                  Branches:
{{- range $_, $v := $v.Branches }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.CustomData }}
                  CustomData: {{ $v.CustomData }}
{{- end }}
{{- if $v.DestinationArn }}
                  DestinationArn: {{ $v.DestinationArn }}
{{- end }}
{{- if $v.Events }}
                  Events:
{{- range $_, $v := $v.Events }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.Name }}
                  Name: {{ $v.Name }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::CodeDeploy::Application" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CodeDeploy::Application
{{- if .ApplicationName | or .ComputePlatform }}
        Properties:
{{- if .ApplicationName }}
            ApplicationName: {{ .ApplicationName }}
{{- end }}
{{- if .ComputePlatform }}
            ComputePlatform: {{ .ComputePlatform }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::CodeDeploy::DeploymentConfig" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CodeDeploy::DeploymentConfig
{{- if .DeploymentConfigName | or .MinimumHealthyHosts }}
        Properties:
{{- if .DeploymentConfigName }}
            DeploymentConfigName: {{ .DeploymentConfigName }}
{{- end }}
{{- if .MinimumHealthyHosts }}
            MinimumHealthyHosts:
                Type: {{ .MinimumHealthyHosts.Type }}
                Value: {{ .MinimumHealthyHosts.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::CodeDeploy::DeploymentGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CodeDeploy::DeploymentGroup
        Properties:
            ApplicationName: {{ .ApplicationName }}
            ServiceRoleArn: {{ .ServiceRoleArn }}
{{- if .AlarmConfiguration }}
            AlarmConfiguration:
{{- if .AlarmConfiguration.Alarms }}
                Alarms:
{{- range $_, $v := .AlarmConfiguration.Alarms }}
                    -
{{- if $v.Name }}
                      Name: {{ $v.Name }}
{{- end }}
{{- end }}
{{- end }}
{{- if .AlarmConfiguration.Enabled }}
                Enabled: {{ .AlarmConfiguration.Enabled }}
{{- end }}
{{- if .AlarmConfiguration.IgnorePollAlarmFailure }}
                IgnorePollAlarmFailure: {{ .AlarmConfiguration.IgnorePollAlarmFailure }}
{{- end }}
{{- end }}
{{- if .AutoRollbackConfiguration }}
            AutoRollbackConfiguration:
{{- if .AutoRollbackConfiguration.Enabled }}
                Enabled: {{ .AutoRollbackConfiguration.Enabled }}
{{- end }}
{{- if .AutoRollbackConfiguration.Events }}
                Events:
{{- range $_, $v := .AutoRollbackConfiguration.Events }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if .AutoScalingGroups }}
            AutoScalingGroups:
{{- range $_, $v := .AutoScalingGroups }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Deployment }}
            Deployment:
                Revision:
{{- if .Deployment.Revision.GitHubLocation }}
                    GitHubLocation:
                        CommitId: {{ .Deployment.Revision.GitHubLocation.CommitId }}
                        Repository: {{ .Deployment.Revision.GitHubLocation.Repository }}
{{- end }}
{{- if .Deployment.Revision.RevisionType }}
                    RevisionType: {{ .Deployment.Revision.RevisionType }}
{{- end }}
{{- if .Deployment.Revision.S3Location }}
                    S3Location:
                        Bucket: {{ .Deployment.Revision.S3Location.Bucket }}
                        Key: {{ .Deployment.Revision.S3Location.Key }}
{{- if .Deployment.Revision.S3Location.BundleType }}
                        BundleType: {{ .Deployment.Revision.S3Location.BundleType }}
{{- end }}
{{- if .Deployment.Revision.S3Location.ETag }}
                        ETag: {{ .Deployment.Revision.S3Location.ETag }}
{{- end }}
{{- if .Deployment.Revision.S3Location.Version }}
                        Version: {{ .Deployment.Revision.S3Location.Version }}
{{- end }}
{{- end }}
{{- if .Deployment.Description }}
                Description: {{ .Deployment.Description }}
{{- end }}
{{- if .Deployment.IgnoreApplicationStopFailures }}
                IgnoreApplicationStopFailures: {{ .Deployment.IgnoreApplicationStopFailures }}
{{- end }}
{{- end }}
{{- if .DeploymentConfigName }}
            DeploymentConfigName: {{ .DeploymentConfigName }}
{{- end }}
{{- if .DeploymentGroupName }}
            DeploymentGroupName: {{ .DeploymentGroupName }}
{{- end }}
{{- if .DeploymentStyle }}
            DeploymentStyle:
{{- if .DeploymentStyle.DeploymentOption }}
                DeploymentOption: {{ .DeploymentStyle.DeploymentOption }}
{{- end }}
{{- if .DeploymentStyle.DeploymentType }}
                DeploymentType: {{ .DeploymentStyle.DeploymentType }}
{{- end }}
{{- end }}
{{- if .Ec2TagFilters }}
            Ec2TagFilters:
{{- range $_, $v := .Ec2TagFilters }}
                -
{{- if $v.Key }}
                  Key: {{ $v.Key }}
{{- end }}
{{- if $v.Type }}
                  Type: {{ $v.Type }}
{{- end }}
{{- if $v.Value }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- if .LoadBalancerInfo }}
            LoadBalancerInfo:
{{- if .LoadBalancerInfo.ElbInfoList }}
                ElbInfoList:
{{- range $_, $v := .LoadBalancerInfo.ElbInfoList }}
                    -
{{- if $v.Name }}
                      Name: {{ $v.Name }}
{{- end }}
{{- end }}
{{- end }}
{{- if .LoadBalancerInfo.TargetGroupInfoList }}
                TargetGroupInfoList:
{{- range $_, $v := .LoadBalancerInfo.TargetGroupInfoList }}
                    -
{{- if $v.Name }}
                      Name: {{ $v.Name }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .OnPremisesInstanceTagFilters }}
            OnPremisesInstanceTagFilters:
{{- range $_, $v := .OnPremisesInstanceTagFilters }}
                -
{{- if $v.Key }}
                  Key: {{ $v.Key }}
{{- end }}
{{- if $v.Type }}
                  Type: {{ $v.Type }}
{{- end }}
{{- if $v.Value }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- if .TriggerConfigurations }}
            TriggerConfigurations:
{{- range $_, $v := .TriggerConfigurations }}
                -
{{- if $v.TriggerEvents }}
                  TriggerEvents:
{{- range $_, $v := $v.TriggerEvents }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.TriggerName }}
                  TriggerName: {{ $v.TriggerName }}
{{- end }}
{{- if $v.TriggerTargetArn }}
                  TriggerTargetArn: {{ $v.TriggerTargetArn }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::CodePipeline::CustomActionType" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CodePipeline::CustomActionType
        Properties:
            Category: {{ .Category }}
            InputArtifactDetails:
                MaximumCount: {{ .InputArtifactDetails.MaximumCount }}
                MinimumCount: {{ .InputArtifactDetails.MinimumCount }}
            OutputArtifactDetails:
                MaximumCount: {{ .OutputArtifactDetails.MaximumCount }}
                MinimumCount: {{ .OutputArtifactDetails.MinimumCount }}
            Provider: {{ .Provider }}
{{- if .ConfigurationProperties }}
            ConfigurationProperties:
{{- range $_, $v := .ConfigurationProperties }}
                -
                  Key: {{ $v.Key }}
                  Name: {{ $v.Name }}
                  Required: {{ $v.Required }}
                  Secret: {{ $v.Secret }}
{{- if $v.Description }}
                  Description: {{ $v.Description }}
{{- end }}
{{- if $v.Queryable }}
                  Queryable: {{ $v.Queryable }}
{{- end }}
{{- if $v.Type }}
                  Type: {{ $v.Type }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Settings }}
            Settings:
{{- if .Settings.EntityUrlTemplate }}
                EntityUrlTemplate: {{ .Settings.EntityUrlTemplate }}
{{- end }}
{{- if .Settings.ExecutionUrlTemplate }}
                ExecutionUrlTemplate: {{ .Settings.ExecutionUrlTemplate }}
{{- end }}
{{- if .Settings.RevisionUrlTemplate }}
                RevisionUrlTemplate: {{ .Settings.RevisionUrlTemplate }}
{{- end }}
{{- if .Settings.ThirdPartyConfigurationUrl }}
                ThirdPartyConfigurationUrl: {{ .Settings.ThirdPartyConfigurationUrl }}
{{- end }}
{{- end }}
{{- if .Version }}
            Version: {{ .Version }}
{{- end }}
{{- end }}
{{- define "AWS::CodePipeline::Pipeline" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::CodePipeline::Pipeline
        Properties:
            ArtifactStore:
                Location: {{ .ArtifactStore.Location }}
                Type: {{ .ArtifactStore.Type }}
{{- if .ArtifactStore.EncryptionKey }}
                EncryptionKey:
                    Id: {{ .ArtifactStore.EncryptionKey.Id }}
                    Type: {{ .ArtifactStore.EncryptionKey.Type }}
{{- end }}
            RoleArn: {{ .RoleArn }}
            Stages:
{{- range $_, $v := .Stages }}
                -
                  Actions:
{{- range $_, $v := $v.Actions }}
                      -
                        ActionTypeId:
                            Category: {{ $v.ActionTypeId.Category }}
                            Owner: {{ $v.ActionTypeId.Owner }}
                            Provider: {{ $v.ActionTypeId.Provider }}
                            Version: {{ $v.ActionTypeId.Version }}
                        Name: {{ $v.Name }}
{{- if $v.Configuration }}
                        Configuration:
                            {{ yaml.MarshalIndent $v.Configuration "                            " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if $v.InputArtifacts }}
                        InputArtifacts:
{{- range $_, $v := $v.InputArtifacts }}
                            -
                              Name: {{ $v.Name }}
{{- end }}
{{- end }}
{{- if $v.OutputArtifacts }}
                        OutputArtifacts:
{{- range $_, $v := $v.OutputArtifacts }}
                            -
                              Name: {{ $v.Name }}
{{- end }}
{{- end }}
{{- if $v.RoleArn }}
                        RoleArn: {{ $v.RoleArn }}
{{- end }}
{{- if $v.RunOrder }}
                        RunOrder: {{ $v.RunOrder }}
{{- end }}
{{- end }}
                  Name: {{ $v.Name }}
{{- if $v.Blockers }}
                  Blockers:
{{- range $_, $v := $v.Blockers }}
                      -
                        Name: {{ $v.Name }}
                        Type: {{ $v.Type }}
{{- end }}
{{- end }}
{{- end }}
{{- if .DisableInboundStageTransitions }}
            DisableInboundStageTransitions:
{{- range $_, $v := .DisableInboundStageTransitions }}
                -
                  Reason: {{ $v.Reason }}
                  StageName: {{ $v.StageName }}
{{- end }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .RestartExecutionOnUpdate }}
            RestartExecutionOnUpdate: {{ .RestartExecutionOnUpdate }}
{{- end }}
{{- end }}
{{- define "AWS::Cognito::IdentityPool" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Cognito::IdentityPool
        Properties:
            AllowUnauthenticatedIdentities: {{ .AllowUnauthenticatedIdentities }}
{{- if .CognitoEvents }}
            CognitoEvents:
                {{ yaml.MarshalIndent .CognitoEvents "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .CognitoIdentityProviders }}
            CognitoIdentityProviders:
{{- range $_, $v := .CognitoIdentityProviders }}
                -
{{- if $v.ClientId }}
                  ClientId: {{ $v.ClientId }}
{{- end }}
{{- if $v.ProviderName }}
                  ProviderName: {{ $v.ProviderName }}
{{- end }}
{{- if $v.ServerSideTokenCheck }}
                  ServerSideTokenCheck: {{ $v.ServerSideTokenCheck }}
{{- end }}
{{- end }}
{{- end }}
{{- if .CognitoStreams }}
            CognitoStreams:
{{- if .CognitoStreams.RoleArn }}
                RoleArn: {{ .CognitoStreams.RoleArn }}
{{- end }}
{{- if .CognitoStreams.StreamName }}
                StreamName: {{ .CognitoStreams.StreamName }}
{{- end }}
{{- if .CognitoStreams.StreamingStatus }}
                StreamingStatus: {{ .CognitoStreams.StreamingStatus }}
{{- end }}
{{- end }}
{{- if .DeveloperProviderName }}
            DeveloperProviderName: {{ .DeveloperProviderName }}
{{- end }}
{{- if .IdentityPoolName }}
            IdentityPoolName: {{ .IdentityPoolName }}
{{- end }}
{{- if .OpenIdConnectProviderARNs }}
            OpenIdConnectProviderARNs:
{{- range $_, $v := .OpenIdConnectProviderARNs }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .PushSync }}
            PushSync:
{{- if .PushSync.ApplicationArns }}
                ApplicationArns:
{{- range $_, $v := .PushSync.ApplicationArns }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- if .PushSync.RoleArn }}
                RoleArn: {{ .PushSync.RoleArn }}
{{- end }}
{{- end }}
{{- if .SamlProviderARNs }}
            SamlProviderARNs:
{{- range $_, $v := .SamlProviderARNs }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .SupportedLoginProviders }}
            SupportedLoginProviders:
                {{ yaml.MarshalIndent .SupportedLoginProviders "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- end }}
{{- define "AWS::Cognito::IdentityPoolRoleAttachment" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Cognito::IdentityPoolRoleAttachment
        Properties:
            IdentityPoolId: {{ .IdentityPoolId }}
{{- if .RoleMappings }}
            RoleMappings:
                {{ yaml.MarshalIndent .RoleMappings "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .Roles }}
            Roles:
                {{ yaml.MarshalIndent .Roles "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- end }}
{{- define "AWS::Cognito::UserPool" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Cognito::UserPool
{{- if .AdminCreateUserConfig | or .AliasAttributes | or .AutoVerifiedAttributes | or .DeviceConfiguration | or .EmailConfiguration | or .EmailVerificationMessage | or .EmailVerificationSubject | or .LambdaConfig | or .MfaConfiguration | or .Policies | or .Schema | or .SmsAuthenticationMessage | or .SmsConfiguration | or .SmsVerificationMessage | or .UserPoolName | or .UserPoolTags }}
        Properties:
{{- if .AdminCreateUserConfig }}
            AdminCreateUserConfig:
{{- if .AdminCreateUserConfig.AllowAdminCreateUserOnly }}
                AllowAdminCreateUserOnly: {{ .AdminCreateUserConfig.AllowAdminCreateUserOnly }}
{{- end }}
{{- if .AdminCreateUserConfig.InviteMessageTemplate }}
                InviteMessageTemplate:
{{- if .AdminCreateUserConfig.InviteMessageTemplate.EmailMessage }}
                    EmailMessage: {{ .AdminCreateUserConfig.InviteMessageTemplate.EmailMessage }}
{{- end }}
{{- if .AdminCreateUserConfig.InviteMessageTemplate.EmailSubject }}
                    EmailSubject: {{ .AdminCreateUserConfig.InviteMessageTemplate.EmailSubject }}
{{- end }}
{{- if .AdminCreateUserConfig.InviteMessageTemplate.SMSMessage }}
                    SMSMessage: {{ .AdminCreateUserConfig.InviteMessageTemplate.SMSMessage }}
{{- end }}
{{- end }}
{{- if .AdminCreateUserConfig.UnusedAccountValidityDays }}
                UnusedAccountValidityDays: {{ .AdminCreateUserConfig.UnusedAccountValidityDays }}
{{- end }}
{{- end }}
{{- if .AliasAttributes }}
            AliasAttributes:
{{- range $_, $v := .AliasAttributes }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .AutoVerifiedAttributes }}
            AutoVerifiedAttributes:
{{- range $_, $v := .AutoVerifiedAttributes }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .DeviceConfiguration }}
            DeviceConfiguration:
{{- if .DeviceConfiguration.ChallengeRequiredOnNewDevice }}
                ChallengeRequiredOnNewDevice: {{ .DeviceConfiguration.ChallengeRequiredOnNewDevice }}
{{- end }}
{{- if .DeviceConfiguration.DeviceOnlyRememberedOnUserPrompt }}
                DeviceOnlyRememberedOnUserPrompt: {{ .DeviceConfiguration.DeviceOnlyRememberedOnUserPrompt }}
{{- end }}
{{- end }}
{{- if .EmailConfiguration }}
            EmailConfiguration:
{{- if .EmailConfiguration.ReplyToEmailAddress }}
                ReplyToEmailAddress: {{ .EmailConfiguration.ReplyToEmailAddress }}
{{- end }}
{{- if .EmailConfiguration.SourceArn }}
                SourceArn: {{ .EmailConfiguration.SourceArn }}
{{- end }}
{{- end }}
{{- if .EmailVerificationMessage }}
            EmailVerificationMessage: {{ .EmailVerificationMessage }}
{{- end }}
{{- if .EmailVerificationSubject }}
            EmailVerificationSubject: {{ .EmailVerificationSubject }}
{{- end }}
{{- if .LambdaConfig }}
            LambdaConfig:
{{- if .LambdaConfig.CreateAuthChallenge }}
                CreateAuthChallenge: {{ .LambdaConfig.CreateAuthChallenge }}
{{- end }}
{{- if .LambdaConfig.CustomMessage }}
                CustomMessage: {{ .LambdaConfig.CustomMessage }}
{{- end }}
{{- if .LambdaConfig.DefineAuthChallenge }}
                DefineAuthChallenge: {{ .LambdaConfig.DefineAuthChallenge }}
{{- end }}
{{- if .LambdaConfig.PostAuthentication }}
                PostAuthentication: {{ .LambdaConfig.PostAuthentication }}
{{- end }}
{{- if .LambdaConfig.PostConfirmation }}
                PostConfirmation: {{ .LambdaConfig.PostConfirmation }}
{{- end }}
{{- if .LambdaConfig.PreAuthentication }}
                PreAuthentication: {{ .LambdaConfig.PreAuthentication }}
{{- end }}
{{- if .LambdaConfig.PreSignUp }}
                PreSignUp: {{ .LambdaConfig.PreSignUp }}
{{- end }}
{{- if .LambdaConfig.VerifyAuthChallengeResponse }}
                VerifyAuthChallengeResponse: {{ .LambdaConfig.VerifyAuthChallengeResponse }}
{{- end }}
{{- end }}
{{- if .MfaConfiguration }}
            MfaConfiguration: {{ .MfaConfiguration }}
{{- end }}
{{- if .Policies }}
            Policies:
{{- if .Policies.PasswordPolicy }}
                PasswordPolicy:
{{- if .Policies.PasswordPolicy.MinimumLength }}
                    MinimumLength: {{ .Policies.PasswordPolicy.MinimumLength }}
{{- end }}
{{- if .Policies.PasswordPolicy.RequireLowercase }}
                    RequireLowercase: {{ .Policies.PasswordPolicy.RequireLowercase }}
{{- end }}
{{- if .Policies.PasswordPolicy.RequireNumbers }}
                    RequireNumbers: {{ .Policies.PasswordPolicy.RequireNumbers }}
{{- end }}
{{- if .Policies.PasswordPolicy.RequireSymbols }}
                    RequireSymbols: {{ .Policies.PasswordPolicy.RequireSymbols }}
{{- end }}
{{- if .Policies.PasswordPolicy.RequireUppercase }}
                    RequireUppercase: {{ .Policies.PasswordPolicy.RequireUppercase }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Schema }}
            Schema:
{{- range $_, $v := .Schema }}
                -
{{- if $v.AttributeDataType }}
                  AttributeDataType: {{ $v.AttributeDataType }}
{{- end }}
{{- if $v.DeveloperOnlyAttribute }}
                  DeveloperOnlyAttribute: {{ $v.DeveloperOnlyAttribute }}
{{- end }}
{{- if $v.Mutable }}
                  Mutable: {{ $v.Mutable }}
{{- end }}
{{- if $v.Name }}
                  Name: {{ $v.Name }}
{{- end }}
{{- if $v.NumberAttributeConstraints }}
                  NumberAttributeConstraints:
{{- if $v.NumberAttributeConstraints.MaxValue }}
                      MaxValue: {{ $v.NumberAttributeConstraints.MaxValue }}
{{- end }}
{{- if $v.NumberAttributeConstraints.MinValue }}
                      MinValue: {{ $v.NumberAttributeConstraints.MinValue }}
{{- end }}
{{- end }}
{{- if $v.Required }}
                  Required: {{ $v.Required }}
{{- end }}
{{- if $v.StringAttributeConstraints }}
                  StringAttributeConstraints:
{{- if $v.StringAttributeConstraints.MaxLength }}
                      MaxLength: {{ $v.StringAttributeConstraints.MaxLength }}
{{- end }}
{{- if $v.StringAttributeConstraints.MinLength }}
                      MinLength: {{ $v.StringAttributeConstraints.MinLength }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .SmsAuthenticationMessage }}
            SmsAuthenticationMessage: {{ .SmsAuthenticationMessage }}
{{- end }}
{{- if .SmsConfiguration }}
            SmsConfiguration:
{{- if .SmsConfiguration.ExternalId }}
                ExternalId: {{ .SmsConfiguration.ExternalId }}
{{- end }}
{{- if .SmsConfiguration.SnsCallerArn }}
                SnsCallerArn: {{ .SmsConfiguration.SnsCallerArn }}
{{- end }}
{{- end }}
{{- if .SmsVerificationMessage }}
            SmsVerificationMessage: {{ .SmsVerificationMessage }}
{{- end }}
{{- if .UserPoolName }}
            UserPoolName: {{ .UserPoolName }}
{{- end }}
{{- if .UserPoolTags }}
            UserPoolTags:
                {{ yaml.MarshalIndent .UserPoolTags "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Cognito::UserPoolClient" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Cognito::UserPoolClient
        Properties:
            UserPoolId: {{ .UserPoolId }}
{{- if .ClientName }}
            ClientName: {{ .ClientName }}
{{- end }}
{{- if .ExplicitAuthFlows }}
            ExplicitAuthFlows:
{{- range $_, $v := .ExplicitAuthFlows }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .GenerateSecret }}
            GenerateSecret: {{ .GenerateSecret }}
{{- end }}
{{- if .ReadAttributes }}
            ReadAttributes:
{{- range $_, $v := .ReadAttributes }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .RefreshTokenValidity }}
            RefreshTokenValidity: {{ .RefreshTokenValidity }}
{{- end }}
{{- if .WriteAttributes }}
            WriteAttributes:
{{- range $_, $v := .WriteAttributes }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Cognito::UserPoolGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Cognito::UserPoolGroup
        Properties:
            UserPoolId: {{ .UserPoolId }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .GroupName }}
            GroupName: {{ .GroupName }}
{{- end }}
{{- if .Precedence }}
            Precedence: {{ .Precedence }}
{{- end }}
{{- if .RoleArn }}
            RoleArn: {{ .RoleArn }}
{{- end }}
{{- end }}
{{- define "AWS::Cognito::UserPoolUser" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Cognito::UserPoolUser
        Properties:
            UserPoolId: {{ .UserPoolId }}
{{- if .DesiredDeliveryMediums }}
            DesiredDeliveryMediums:
{{- range $_, $v := .DesiredDeliveryMediums }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .ForceAliasCreation }}
            ForceAliasCreation: {{ .ForceAliasCreation }}
{{- end }}
{{- if .MessageAction }}
            MessageAction: {{ .MessageAction }}
{{- end }}
{{- if .UserAttributes }}
            UserAttributes:
{{- range $_, $v := .UserAttributes }}
                -
{{- if $v.Name }}
                  Name: {{ $v.Name }}
{{- end }}
{{- if $v.Value }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Username }}
            Username: {{ .Username }}
{{- end }}
{{- if .ValidationData }}
            ValidationData:
{{- range $_, $v := .ValidationData }}
                -
{{- if $v.Name }}
                  Name: {{ $v.Name }}
{{- end }}
{{- if $v.Value }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Cognito::UserPoolUserToGroupAttachment" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Cognito::UserPoolUserToGroupAttachment
        Properties:
            GroupName: {{ .GroupName }}
            UserPoolId: {{ .UserPoolId }}
            Username: {{ .Username }}
{{- end }}
{{- define "AWS::Config::ConfigRule" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Config::ConfigRule
        Properties:
            Source:
                Owner: {{ .Source.Owner }}
                SourceIdentifier: {{ .Source.SourceIdentifier }}
{{- if .Source.SourceDetails }}
                SourceDetails:
{{- range $_, $v := .Source.SourceDetails }}
                    -
                      EventSource: {{ $v.EventSource }}
                      MessageType: {{ $v.MessageType }}
{{- if $v.MaximumExecutionFrequency }}
                      MaximumExecutionFrequency: {{ $v.MaximumExecutionFrequency }}
{{- end }}
{{- end }}
{{- end }}
{{- if .ConfigRuleName }}
            ConfigRuleName: {{ .ConfigRuleName }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .InputParameters }}
            InputParameters:
                {{ yaml.MarshalIndent .InputParameters "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .MaximumExecutionFrequency }}
            MaximumExecutionFrequency: {{ .MaximumExecutionFrequency }}
{{- end }}
{{- if .Scope }}
            Scope:
{{- if .Scope.ComplianceResourceId }}
                ComplianceResourceId: {{ .Scope.ComplianceResourceId }}
{{- end }}
{{- if .Scope.ComplianceResourceTypes }}
                ComplianceResourceTypes:
{{- range $_, $v := .Scope.ComplianceResourceTypes }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- if .Scope.TagKey }}
                TagKey: {{ .Scope.TagKey }}
{{- end }}
{{- if .Scope.TagValue }}
                TagValue: {{ .Scope.TagValue }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Config::ConfigurationRecorder" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Config::ConfigurationRecorder
        Properties:
            RoleARN: {{ .RoleARN }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .RecordingGroup }}
            RecordingGroup:
{{- if .RecordingGroup.AllSupported }}
                AllSupported: {{ .RecordingGroup.AllSupported }}
{{- end }}
{{- if .RecordingGroup.IncludeGlobalResourceTypes }}
                IncludeGlobalResourceTypes: {{ .RecordingGroup.IncludeGlobalResourceTypes }}
{{- end }}
{{- if .RecordingGroup.ResourceTypes }}
                ResourceTypes:
{{- range $_, $v := .RecordingGroup.ResourceTypes }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Config::DeliveryChannel" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Config::DeliveryChannel
        Properties:
            S3BucketName: {{ .S3BucketName }}
{{- if .ConfigSnapshotDeliveryProperties }}
            ConfigSnapshotDeliveryProperties:
{{- if .ConfigSnapshotDeliveryProperties.DeliveryFrequency }}
                DeliveryFrequency: {{ .ConfigSnapshotDeliveryProperties.DeliveryFrequency }}
{{- end }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .S3KeyPrefix }}
            S3KeyPrefix: {{ .S3KeyPrefix }}
{{- end }}
{{- if .SnsTopicARN }}
            SnsTopicARN: {{ .SnsTopicARN }}
{{- end }}
{{- end }}
{{- define "AWS::DAX::Cluster" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::DAX::Cluster
        Properties:
            IAMRoleARN: {{ .IAMRoleARN }}
            NodeType: {{ .NodeType }}
            ReplicationFactor: {{ .ReplicationFactor }}
{{- if .AvailabilityZones }}
            AvailabilityZones:
{{- range $_, $v := .AvailabilityZones }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .ClusterName }}
            ClusterName: {{ .ClusterName }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .NotificationTopicARN }}
            NotificationTopicARN: {{ .NotificationTopicARN }}
{{- end }}
{{- if .ParameterGroupName }}
            ParameterGroupName: {{ .ParameterGroupName }}
{{- end }}
{{- if .PreferredMaintenanceWindow }}
            PreferredMaintenanceWindow: {{ .PreferredMaintenanceWindow }}
{{- end }}
{{- if .SecurityGroupIds }}
            SecurityGroupIds:
{{- range $_, $v := .SecurityGroupIds }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .SubnetGroupName }}
            SubnetGroupName: {{ .SubnetGroupName }}
{{- end }}
{{- if .Tags }}
            Tags:
                {{ yaml.MarshalIndent .Tags "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- end }}
{{- define "AWS::DAX::ParameterGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::DAX::ParameterGroup
{{- if .Description | or .ParameterGroupName | or .ParameterNameValues }}
        Properties:
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .ParameterGroupName }}
            ParameterGroupName: {{ .ParameterGroupName }}
{{- end }}
{{- if .ParameterNameValues }}
            ParameterNameValues:
                {{ yaml.MarshalIndent .ParameterNameValues "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::DAX::SubnetGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::DAX::SubnetGroup
        Properties:
            SubnetIds:
{{- range $_, $v := .SubnetIds }}
                - {{ $v }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .SubnetGroupName }}
            SubnetGroupName: {{ .SubnetGroupName }}
{{- end }}
{{- end }}
{{- define "AWS::DMS::Certificate" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::DMS::Certificate
{{- if .CertificateIdentifier | or .CertificatePem | or .CertificateWallet }}
        Properties:
{{- if .CertificateIdentifier }}
            CertificateIdentifier: {{ .CertificateIdentifier }}
{{- end }}
{{- if .CertificatePem }}
            CertificatePem: {{ .CertificatePem }}
{{- end }}
{{- if .CertificateWallet }}
            CertificateWallet: {{ .CertificateWallet }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::DMS::Endpoint" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::DMS::Endpoint
        Properties:
            EndpointType: {{ .EndpointType }}
            EngineName: {{ .EngineName }}
{{- if .CertificateArn }}
            CertificateArn: {{ .CertificateArn }}
{{- end }}
{{- if .DatabaseName }}
            DatabaseName: {{ .DatabaseName }}
{{- end }}
{{- if .DynamoDbSettings }}
            DynamoDbSettings:
{{- if .DynamoDbSettings.ServiceAccessRoleArn }}
                ServiceAccessRoleArn: {{ .DynamoDbSettings.ServiceAccessRoleArn }}
{{- end }}
{{- end }}
{{- if .EndpointIdentifier }}
            EndpointIdentifier: {{ .EndpointIdentifier }}
{{- end }}
{{- if .ExtraConnectionAttributes }}
            ExtraConnectionAttributes: {{ .ExtraConnectionAttributes }}
{{- end }}
{{- if .KmsKeyId }}
            KmsKeyId: {{ .KmsKeyId }}
{{- end }}
{{- if .MongoDbSettings }}
            MongoDbSettings:
{{- if .MongoDbSettings.AuthMechanism }}
                AuthMechanism: {{ .MongoDbSettings.AuthMechanism }}
{{- end }}
{{- if .MongoDbSettings.AuthSource }}
                AuthSource: {{ .MongoDbSettings.AuthSource }}
{{- end }}
{{- if .MongoDbSettings.AuthType }}
                AuthType: {{ .MongoDbSettings.AuthType }}
{{- end }}
{{- if .MongoDbSettings.DatabaseName }}
                DatabaseName: {{ .MongoDbSettings.DatabaseName }}
{{- end }}
{{- if .MongoDbSettings.DocsToInvestigate }}
                DocsToInvestigate: {{ .MongoDbSettings.DocsToInvestigate }}
{{- end }}
{{- if .MongoDbSettings.ExtractDocId }}
                ExtractDocId: {{ .MongoDbSettings.ExtractDocId }}
{{- end }}
{{- if .MongoDbSettings.NestingLevel }}
                NestingLevel: {{ .MongoDbSettings.NestingLevel }}
{{- end }}
{{- if .MongoDbSettings.Password }}
                Password: {{ .MongoDbSettings.Password }}
{{- end }}
{{- if .MongoDbSettings.Port }}
                Port: {{ .MongoDbSettings.Port }}
{{- end }}
{{- if .MongoDbSettings.ServerName }}
                ServerName: {{ .MongoDbSettings.ServerName }}
{{- end }}
{{- if .MongoDbSettings.Username }}
                Username: {{ .MongoDbSettings.Username }}
{{- end }}
{{- end }}
{{- if .Password }}
            Password: {{ .Password }}
{{- end }}
{{- if .Port }}
            Port: {{ .Port }}
{{- end }}
{{- if .S3Settings }}
            S3Settings:
{{- if .S3Settings.BucketFolder }}
                BucketFolder: {{ .S3Settings.BucketFolder }}
{{- end }}
{{- if .S3Settings.BucketName }}
                BucketName: {{ .S3Settings.BucketName }}
{{- end }}
{{- if .S3Settings.CompressionType }}
                CompressionType: {{ .S3Settings.CompressionType }}
{{- end }}
{{- if .S3Settings.CsvDelimiter }}
                CsvDelimiter: {{ .S3Settings.CsvDelimiter }}
{{- end }}
{{- if .S3Settings.CsvRowDelimiter }}
                CsvRowDelimiter: {{ .S3Settings.CsvRowDelimiter }}
{{- end }}
{{- if .S3Settings.ExternalTableDefinition }}
                ExternalTableDefinition: {{ .S3Settings.ExternalTableDefinition }}
{{- end }}
{{- if .S3Settings.ServiceAccessRoleArn }}
                ServiceAccessRoleArn: {{ .S3Settings.ServiceAccessRoleArn }}
{{- end }}
{{- end }}
{{- if .ServerName }}
            ServerName: {{ .ServerName }}
{{- end }}
{{- if .SslMode }}
            SslMode: {{ .SslMode }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .Username }}
            Username: {{ .Username }}
{{- end }}
{{- end }}
{{- define "AWS::DMS::EventSubscription" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::DMS::EventSubscription
        Properties:
            SnsTopicArn: {{ .SnsTopicArn }}
{{- if .Enabled }}
            Enabled: {{ .Enabled }}
{{- end }}
{{- if .EventCategories }}
            EventCategories:
{{- range $_, $v := .EventCategories }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .SourceIds }}
            SourceIds:
{{- range $_, $v := .SourceIds }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .SourceType }}
            SourceType: {{ .SourceType }}
{{- end }}
{{- if .SubscriptionName }}
            SubscriptionName: {{ .SubscriptionName }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::DMS::ReplicationInstance" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::DMS::ReplicationInstance
        Properties:
            ReplicationInstanceClass: {{ .ReplicationInstanceClass }}
{{- if .AllocatedStorage }}
            AllocatedStorage: {{ .AllocatedStorage }}
{{- end }}
{{- if .AllowMajorVersionUpgrade }}
            AllowMajorVersionUpgrade: {{ .AllowMajorVersionUpgrade }}
{{- end }}
{{- if .AutoMinorVersionUpgrade }}
            AutoMinorVersionUpgrade: {{ .AutoMinorVersionUpgrade }}
{{- end }}
{{- if .AvailabilityZone }}
            AvailabilityZone: {{ .AvailabilityZone }}
{{- end }}
{{- if .EngineVersion }}
            EngineVersion: {{ .EngineVersion }}
{{- end }}
{{- if .KmsKeyId }}
            KmsKeyId: {{ .KmsKeyId }}
{{- end }}
{{- if .MultiAZ }}
            MultiAZ: {{ .MultiAZ }}
{{- end }}
{{- if .PreferredMaintenanceWindow }}
            PreferredMaintenanceWindow: {{ .PreferredMaintenanceWindow }}
{{- end }}
{{- if .PubliclyAccessible }}
            PubliclyAccessible: {{ .PubliclyAccessible }}
{{- end }}
{{- if .ReplicationInstanceIdentifier }}
            ReplicationInstanceIdentifier: {{ .ReplicationInstanceIdentifier }}
{{- end }}
{{- if .ReplicationSubnetGroupIdentifier }}
            ReplicationSubnetGroupIdentifier: {{ .ReplicationSubnetGroupIdentifier }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .VpcSecurityGroupIds }}
            VpcSecurityGroupIds:
{{- range $_, $v := .VpcSecurityGroupIds }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::DMS::ReplicationSubnetGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::DMS::ReplicationSubnetGroup
        Properties:
            ReplicationSubnetGroupDescription: {{ .ReplicationSubnetGroupDescription }}
            SubnetIds:
{{- range $_, $v := .SubnetIds }}
                - {{ $v }}
{{- end }}
{{- if .ReplicationSubnetGroupIdentifier }}
            ReplicationSubnetGroupIdentifier: {{ .ReplicationSubnetGroupIdentifier }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::DMS::ReplicationTask" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::DMS::ReplicationTask
        Properties:
            MigrationType: {{ .MigrationType }}
            ReplicationInstanceArn: {{ .ReplicationInstanceArn }}
            SourceEndpointArn: {{ .SourceEndpointArn }}
            TableMappings: {{ .TableMappings }}
            TargetEndpointArn: {{ .TargetEndpointArn }}
{{- if .CdcStartTime }}
            CdcStartTime: {{ .CdcStartTime }}
{{- end }}
{{- if .ReplicationTaskIdentifier }}
            ReplicationTaskIdentifier: {{ .ReplicationTaskIdentifier }}
{{- end }}
{{- if .ReplicationTaskSettings }}
            ReplicationTaskSettings: {{ .ReplicationTaskSettings }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::DataPipeline::Pipeline" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::DataPipeline::Pipeline
        Properties:
            Name: {{ .Name }}
            ParameterObjects:
{{- range $_, $v := .ParameterObjects }}
                -
                  Attributes:
{{- range $_, $v := $v.Attributes }}
                      -
                        Key: {{ $v.Key }}
                        StringValue: {{ $v.StringValue }}
{{- end }}
                  Id: {{ $v.Id }}
{{- end }}
{{- if .Activate }}
            Activate: {{ .Activate }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .ParameterValues }}
            ParameterValues:
{{- range $_, $v := .ParameterValues }}
                -
                  Id: {{ $v.Id }}
                  StringValue: {{ $v.StringValue }}
{{- end }}
{{- end }}
{{- if .PipelineObjects }}
            PipelineObjects:
{{- range $_, $v := .PipelineObjects }}
                -
                  Fields:
{{- range $_, $v := $v.Fields }}
                      -
                        Key: {{ $v.Key }}
{{- if $v.RefValue }}
                        RefValue: {{ $v.RefValue }}
{{- end }}
{{- if $v.StringValue }}
                        StringValue: {{ $v.StringValue }}
{{- end }}
{{- end }}
                  Id: {{ $v.Id }}
                  Name: {{ $v.Name }}
{{- end }}
{{- end }}
{{- if .PipelineTags }}
            PipelineTags:
{{- range $_, $v := .PipelineTags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::DirectoryService::MicrosoftAD" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::DirectoryService::MicrosoftAD
        Properties:
            Name: {{ .Name }}
            Password: {{ .Password }}
            VpcSettings:
                SubnetIds:
{{- range $_, $v := .VpcSettings.SubnetIds }}
                    - {{ $v }}
{{- end }}
                VpcId: {{ .VpcSettings.VpcId }}
{{- if .CreateAlias }}
            CreateAlias: {{ .CreateAlias }}
{{- end }}
{{- if .EnableSso }}
            EnableSso: {{ .EnableSso }}
{{- end }}
{{- if .ShortName }}
            ShortName: {{ .ShortName }}
{{- end }}
{{- end }}
{{- define "AWS::DirectoryService::SimpleAD" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::DirectoryService::SimpleAD
        Properties:
            Name: {{ .Name }}
            Password: {{ .Password }}
            Size: {{ .Size }}
            VpcSettings:
                SubnetIds:
{{- range $_, $v := .VpcSettings.SubnetIds }}
                    - {{ $v }}
{{- end }}
                VpcId: {{ .VpcSettings.VpcId }}
{{- if .CreateAlias }}
            CreateAlias: {{ .CreateAlias }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .EnableSso }}
            EnableSso: {{ .EnableSso }}
{{- end }}
{{- if .ShortName }}
            ShortName: {{ .ShortName }}
{{- end }}
{{- end }}
{{- define "AWS::DynamoDB::Table" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::DynamoDB::Table
        Properties:
            KeySchema:
{{- range $_, $v := .KeySchema }}
                -
                  AttributeName: {{ $v.AttributeName }}
                  KeyType: {{ $v.KeyType }}
{{- end }}
            ProvisionedThroughput:
                ReadCapacityUnits: {{ .ProvisionedThroughput.ReadCapacityUnits }}
                WriteCapacityUnits: {{ .ProvisionedThroughput.WriteCapacityUnits }}
{{- if .AttributeDefinitions }}
            AttributeDefinitions:
{{- range $_, $v := .AttributeDefinitions }}
                -
                  AttributeName: {{ $v.AttributeName }}
                  AttributeType: {{ $v.AttributeType }}
{{- end }}
{{- end }}
{{- if .GlobalSecondaryIndexes }}
            GlobalSecondaryIndexes:
{{- range $_, $v := .GlobalSecondaryIndexes }}
                -
                  IndexName: {{ $v.IndexName }}
                  KeySchema:
{{- range $_, $v := $v.KeySchema }}
                      -
                        AttributeName: {{ $v.AttributeName }}
                        KeyType: {{ $v.KeyType }}
{{- end }}
                  Projection:
{{- if $v.Projection.NonKeyAttributes }}
                      NonKeyAttributes:
{{- range $_, $v := $v.Projection.NonKeyAttributes }}
                          - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.Projection.ProjectionType }}
                      ProjectionType: {{ $v.Projection.ProjectionType }}
{{- end }}
                  ProvisionedThroughput:
                      ReadCapacityUnits: {{ $v.ProvisionedThroughput.ReadCapacityUnits }}
                      WriteCapacityUnits: {{ $v.ProvisionedThroughput.WriteCapacityUnits }}
{{- end }}
{{- end }}
{{- if .LocalSecondaryIndexes }}
            LocalSecondaryIndexes:
{{- range $_, $v := .LocalSecondaryIndexes }}
                -
                  IndexName: {{ $v.IndexName }}
                  KeySchema:
{{- range $_, $v := $v.KeySchema }}
                      -
                        AttributeName: {{ $v.AttributeName }}
                        KeyType: {{ $v.KeyType }}
{{- end }}
                  Projection:
{{- if $v.Projection.NonKeyAttributes }}
                      NonKeyAttributes:
{{- range $_, $v := $v.Projection.NonKeyAttributes }}
                          - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.Projection.ProjectionType }}
                      ProjectionType: {{ $v.Projection.ProjectionType }}
{{- end }}
{{- end }}
{{- end }}
{{- if .SSESpecification }}
            SSESpecification:
                SSEEnabled: {{ .SSESpecification.SSEEnabled }}
{{- end }}
{{- if .StreamSpecification }}
            StreamSpecification:
                StreamViewType: {{ .StreamSpecification.StreamViewType }}
{{- end }}
{{- if .TableName }}
            TableName: {{ .TableName }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .TimeToLiveSpecification }}
            TimeToLiveSpecification:
                AttributeName: {{ .TimeToLiveSpecification.AttributeName }}
                Enabled: {{ .TimeToLiveSpecification.Enabled }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::CustomerGateway" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::CustomerGateway
        Properties:
            BgpAsn: {{ .BgpAsn }}
            IpAddress: {{ .IpAddress }}
            Type: {{ .Type }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::DHCPOptions" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::DHCPOptions
{{- if .DomainName | or .DomainNameServers | or .NetbiosNameServers | or .NetbiosNodeType | or .NtpServers | or .Tags }}
        Properties:
{{- if .DomainName }}
            DomainName: {{ .DomainName }}
{{- end }}
{{- if .DomainNameServers }}
            DomainNameServers:
{{- range $_, $v := .DomainNameServers }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .NetbiosNameServers }}
            NetbiosNameServers:
{{- range $_, $v := .NetbiosNameServers }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .NetbiosNodeType }}
            NetbiosNodeType: {{ .NetbiosNodeType }}
{{- end }}
{{- if .NtpServers }}
            NtpServers:
{{- range $_, $v := .NtpServers }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::EIP" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::EIP
{{- if .Domain | or .InstanceId }}
        Properties:
{{- if .Domain }}
            Domain: {{ .Domain }}
{{- end }}
{{- if .InstanceId }}
            InstanceId: {{ .InstanceId }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::EIPAssociation" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::EIPAssociation
{{- if .AllocationId | or .EIP | or .InstanceId | or .NetworkInterfaceId | or .PrivateIpAddress }}
        Properties:
{{- if .AllocationId }}
            AllocationId: {{ .AllocationId }}
{{- end }}
{{- if .EIP }}
            EIP: {{ .EIP }}
{{- end }}
{{- if .InstanceId }}
            InstanceId: {{ .InstanceId }}
{{- end }}
{{- if .NetworkInterfaceId }}
            NetworkInterfaceId: {{ .NetworkInterfaceId }}
{{- end }}
{{- if .PrivateIpAddress }}
            PrivateIpAddress: {{ .PrivateIpAddress }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::EgressOnlyInternetGateway" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::EgressOnlyInternetGateway
        Properties:
            VpcId: {{ .VpcId }}
{{- end }}
{{- define "AWS::EC2::FlowLog" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::FlowLog
        Properties:
            DeliverLogsPermissionArn: {{ .DeliverLogsPermissionArn }}
            LogGroupName: {{ .LogGroupName }}
            ResourceId: {{ .ResourceId }}
            ResourceType: {{ .ResourceType }}
            TrafficType: {{ .TrafficType }}
{{- end }}
{{- define "AWS::EC2::Host" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::Host
        Properties:
            AvailabilityZone: {{ .AvailabilityZone }}
            InstanceType: {{ .InstanceType }}
{{- if .AutoPlacement }}
            AutoPlacement: {{ .AutoPlacement }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::Instance" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::Instance
        Properties:
            ImageId: {{ .ImageId }}
{{- if .AdditionalInfo }}
            AdditionalInfo: {{ .AdditionalInfo }}
{{- end }}
{{- if .Affinity }}
            Affinity: {{ .Affinity }}
{{- end }}
{{- if .AvailabilityZone }}
            AvailabilityZone: {{ .AvailabilityZone }}
{{- end }}
{{- if .BlockDeviceMappings }}
            BlockDeviceMappings:
{{- range $_, $v := .BlockDeviceMappings }}
                -
                  DeviceName: {{ $v.DeviceName }}
{{- if $v.Ebs }}
                  Ebs:
{{- if $v.Ebs.DeleteOnTermination }}
                      DeleteOnTermination: {{ $v.Ebs.DeleteOnTermination }}
{{- end }}
{{- if $v.Ebs.Encrypted }}
                      Encrypted: {{ $v.Ebs.Encrypted }}
{{- end }}
{{- if $v.Ebs.Iops }}
                      Iops: {{ $v.Ebs.Iops }}
{{- end }}
{{- if $v.Ebs.SnapshotId }}
                      SnapshotId: {{ $v.Ebs.SnapshotId }}
{{- end }}
{{- if $v.Ebs.VolumeSize }}
                      VolumeSize: {{ $v.Ebs.VolumeSize }}
{{- end }}
{{- if $v.Ebs.VolumeType }}
                      VolumeType: {{ $v.Ebs.VolumeType }}
{{- end }}
{{- end }}
{{- if $v.NoDevice }}
                  NoDevice:
{{- end }}
{{- if $v.VirtualName }}
                  VirtualName: {{ $v.VirtualName }}
{{- end }}
{{- end }}
{{- end }}
{{- if .CreditSpecification }}
            CreditSpecification:
{{- if .CreditSpecification.CPUCredits }}
                CPUCredits: {{ .CreditSpecification.CPUCredits }}
{{- end }}
{{- end }}
{{- if .DisableApiTermination }}
            DisableApiTermination: {{ .DisableApiTermination }}
{{- end }}
{{- if .EbsOptimized }}
            EbsOptimized: {{ .EbsOptimized }}
{{- end }}
{{- if .ElasticGpuSpecifications }}
            ElasticGpuSpecifications:
{{- range $_, $v := .ElasticGpuSpecifications }}
                -
                  Type: {{ $v.Type }}
{{- end }}
{{- end }}
{{- if .HostId }}
            HostId: {{ .HostId }}
{{- end }}
{{- if .IamInstanceProfile }}
            IamInstanceProfile: {{ .IamInstanceProfile }}
{{- end }}
{{- if .InstanceInitiatedShutdownBehavior }}
            InstanceInitiatedShutdownBehavior: {{ .InstanceInitiatedShutdownBehavior }}
{{- end }}
{{- if .InstanceType }}
            InstanceType: {{ .InstanceType }}
{{- end }}
{{- if .Ipv6AddressCount }}
            Ipv6AddressCount: {{ .Ipv6AddressCount }}
{{- end }}
{{- if .Ipv6Addresses }}
            Ipv6Addresses:
{{- range $_, $v := .Ipv6Addresses }}
                -
                  Ipv6Address: {{ $v.Ipv6Address }}
{{- end }}
{{- end }}
{{- if .KernelId }}
            KernelId: {{ .KernelId }}
{{- end }}
{{- if .KeyName }}
            KeyName: {{ .KeyName }}
{{- end }}
{{- if .Monitoring }}
            Monitoring: {{ .Monitoring }}
{{- end }}
{{- if .NetworkInterfaces }}
            NetworkInterfaces:
{{- range $_, $v := .NetworkInterfaces }}
                -
                  DeviceIndex: {{ $v.DeviceIndex }}
{{- if $v.AssociatePublicIpAddress }}
                  AssociatePublicIpAddress: {{ $v.AssociatePublicIpAddress }}
{{- end }}
{{- if $v.DeleteOnTermination }}
                  DeleteOnTermination: {{ $v.DeleteOnTermination }}
{{- end }}
{{- if $v.Description }}
                  Description: {{ $v.Description }}
{{- end }}
{{- if $v.GroupSet }}
                  GroupSet:
{{- range $_, $v := $v.GroupSet }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.Ipv6AddressCount }}
                  Ipv6AddressCount: {{ $v.Ipv6AddressCount }}
{{- end }}
{{- if $v.Ipv6Addresses }}
                  Ipv6Addresses:
{{- range $_, $v := $v.Ipv6Addresses }}
                      -
                        Ipv6Address: {{ $v.Ipv6Address }}
{{- end }}
{{- end }}
{{- if $v.NetworkInterfaceId }}
                  NetworkInterfaceId: {{ $v.NetworkInterfaceId }}
{{- end }}
{{- if $v.PrivateIpAddress }}
                  PrivateIpAddress: {{ $v.PrivateIpAddress }}
{{- end }}
{{- if $v.PrivateIpAddresses }}
                  PrivateIpAddresses:
{{- range $_, $v := $v.PrivateIpAddresses }}
                      -
                        Primary: {{ $v.Primary }}
                        PrivateIpAddress: {{ $v.PrivateIpAddress }}
{{- end }}
{{- end }}
{{- if $v.SecondaryPrivateIpAddressCount }}
                  SecondaryPrivateIpAddressCount: {{ $v.SecondaryPrivateIpAddressCount }}
{{- end }}
{{- if $v.SubnetId }}
                  SubnetId: {{ $v.SubnetId }}
{{- end }}
{{- end }}
{{- end }}
{{- if .PlacementGroupName }}
            PlacementGroupName: {{ .PlacementGroupName }}
{{- end }}
{{- if .PrivateIpAddress }}
            PrivateIpAddress: {{ .PrivateIpAddress }}
{{- end }}
{{- if .RamdiskId }}
            RamdiskId: {{ .RamdiskId }}
{{- end }}
{{- if .SecurityGroupIds }}
            SecurityGroupIds:
{{- range $_, $v := .SecurityGroupIds }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .SecurityGroups }}
            SecurityGroups:
{{- range $_, $v := .SecurityGroups }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .SourceDestCheck }}
            SourceDestCheck: {{ .SourceDestCheck }}
{{- end }}
{{- if .SsmAssociations }}
            SsmAssociations:
{{- range $_, $v := .SsmAssociations }}
                -
                  DocumentName: {{ $v.DocumentName }}
{{- if $v.AssociationParameters }}
                  AssociationParameters:
{{- range $_, $v := $v.AssociationParameters }}
                      -
                        Key: {{ $v.Key }}
                        Value:
{{- range $_, $v := $v.Value }}
                            - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .SubnetId }}
            SubnetId: {{ .SubnetId }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .Tenancy }}
            Tenancy: {{ .Tenancy }}
{{- end }}
{{- if .UserData }}
            UserData: {{ .UserData }}
{{- end }}
{{- if .Volumes }}
            Volumes:
{{- range $_, $v := .Volumes }}
                -
                  Device: {{ $v.Device }}
                  VolumeId: {{ $v.VolumeId }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::InternetGateway" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::InternetGateway
{{- if .Tags }}
        Properties:
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::NatGateway" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::NatGateway
        Properties:
            AllocationId: {{ .AllocationId }}
            SubnetId: {{ .SubnetId }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::NetworkAcl" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::NetworkAcl
        Properties:
            VpcId: {{ .VpcId }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::NetworkAclEntry" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::NetworkAclEntry
        Properties:
            CidrBlock: {{ .CidrBlock }}
            NetworkAclId: {{ .NetworkAclId }}
            Protocol: {{ .Protocol }}
            RuleAction: {{ .RuleAction }}
            RuleNumber: {{ .RuleNumber }}
{{- if .Egress }}
            Egress: {{ .Egress }}
{{- end }}
{{- if .Icmp }}
            Icmp:
{{- if .Icmp.Code }}
                Code: {{ .Icmp.Code }}
{{- end }}
{{- if .Icmp.Type }}
                Type: {{ .Icmp.Type }}
{{- end }}
{{- end }}
{{- if .Ipv6CidrBlock }}
            Ipv6CidrBlock: {{ .Ipv6CidrBlock }}
{{- end }}
{{- if .PortRange }}
            PortRange:
{{- if .PortRange.From }}
                From: {{ .PortRange.From }}
{{- end }}
{{- if .PortRange.To }}
                To: {{ .PortRange.To }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::NetworkInterface" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::NetworkInterface
        Properties:
            SubnetId: {{ .SubnetId }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .GroupSet }}
            GroupSet:
{{- range $_, $v := .GroupSet }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .InterfaceType }}
            InterfaceType: {{ .InterfaceType }}
{{- end }}
{{- if .Ipv6AddressCount }}
            Ipv6AddressCount: {{ .Ipv6AddressCount }}
{{- end }}
{{- if .Ipv6Addresses }}
            Ipv6Addresses:
                Ipv6Address: {{ .Ipv6Addresses.Ipv6Address }}
{{- end }}
{{- if .PrivateIpAddress }}
            PrivateIpAddress: {{ .PrivateIpAddress }}
{{- end }}
{{- if .PrivateIpAddresses }}
            PrivateIpAddresses:
{{- range $_, $v := .PrivateIpAddresses }}
                -
                  Primary: {{ $v.Primary }}
                  PrivateIpAddress: {{ $v.PrivateIpAddress }}
{{- end }}
{{- end }}
{{- if .SecondaryPrivateIpAddressCount }}
            SecondaryPrivateIpAddressCount: {{ .SecondaryPrivateIpAddressCount }}
{{- end }}
{{- if .SourceDestCheck }}
            SourceDestCheck: {{ .SourceDestCheck }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::NetworkInterfaceAttachment" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::NetworkInterfaceAttachment
        Properties:
            DeviceIndex: {{ .DeviceIndex }}
            InstanceId: {{ .InstanceId }}
            NetworkInterfaceId: {{ .NetworkInterfaceId }}
{{- if .DeleteOnTermination }}
            DeleteOnTermination: {{ .DeleteOnTermination }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::NetworkInterfacePermission" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::NetworkInterfacePermission
        Properties:
            AwsAccountId: {{ .AwsAccountId }}
            NetworkInterfaceId: {{ .NetworkInterfaceId }}
            Permission: {{ .Permission }}
{{- end }}
{{- define "AWS::EC2::PlacementGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::PlacementGroup
{{- if .Strategy }}
        Properties:
{{- if .Strategy }}
            Strategy: {{ .Strategy }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::Route" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::Route
        Properties:
            RouteTableId: {{ .RouteTableId }}
{{- if .DestinationCidrBlock }}
            DestinationCidrBlock: {{ .DestinationCidrBlock }}
{{- end }}
{{- if .DestinationIpv6CidrBlock }}
            DestinationIpv6CidrBlock: {{ .DestinationIpv6CidrBlock }}
{{- end }}
{{- if .EgressOnlyInternetGatewayId }}
            EgressOnlyInternetGatewayId: {{ .EgressOnlyInternetGatewayId }}
{{- end }}
{{- if .GatewayId }}
            GatewayId: {{ .GatewayId }}
{{- end }}
{{- if .InstanceId }}
            InstanceId: {{ .InstanceId }}
{{- end }}
{{- if .NatGatewayId }}
            NatGatewayId: {{ .NatGatewayId }}
{{- end }}
{{- if .NetworkInterfaceId }}
            NetworkInterfaceId: {{ .NetworkInterfaceId }}
{{- end }}
{{- if .VpcPeeringConnectionId }}
            VpcPeeringConnectionId: {{ .VpcPeeringConnectionId }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::RouteTable" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::RouteTable
        Properties:
            VpcId: {{ .VpcId }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::SecurityGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::SecurityGroup
        Properties:
            GroupDescription: {{ .GroupDescription }}
{{- if .GroupName }}
            GroupName: {{ .GroupName }}
{{- end }}
{{- if .SecurityGroupEgress }}
            SecurityGroupEgress:
{{- range $_, $v := .SecurityGroupEgress }}
                -
                  IpProtocol: {{ $v.IpProtocol }}
{{- if $v.CidrIp }}
                  CidrIp: {{ $v.CidrIp }}
{{- end }}
{{- if $v.CidrIpv6 }}
                  CidrIpv6: {{ $v.CidrIpv6 }}
{{- end }}
{{- if $v.Description }}
                  Description: {{ $v.Description }}
{{- end }}
{{- if $v.DestinationPrefixListId }}
                  DestinationPrefixListId: {{ $v.DestinationPrefixListId }}
{{- end }}
{{- if $v.DestinationSecurityGroupId }}
                  DestinationSecurityGroupId: {{ $v.DestinationSecurityGroupId }}
{{- end }}
{{- if $v.FromPort }}
                  FromPort: {{ $v.FromPort }}
{{- end }}
{{- if $v.ToPort }}
                  ToPort: {{ $v.ToPort }}
{{- end }}
{{- end }}
{{- end }}
{{- if .SecurityGroupIngress }}
            SecurityGroupIngress:
{{- range $_, $v := .SecurityGroupIngress }}
                -
                  IpProtocol: {{ $v.IpProtocol }}
{{- if $v.CidrIp }}
                  CidrIp: {{ $v.CidrIp }}
{{- end }}
{{- if $v.CidrIpv6 }}
                  CidrIpv6: {{ $v.CidrIpv6 }}
{{- end }}
{{- if $v.Description }}
                  Description: {{ $v.Description }}
{{- end }}
{{- if $v.FromPort }}
                  FromPort: {{ $v.FromPort }}
{{- end }}
{{- if $v.SourceSecurityGroupId }}
                  SourceSecurityGroupId: {{ $v.SourceSecurityGroupId }}
{{- end }}
{{- if $v.SourceSecurityGroupName }}
                  SourceSecurityGroupName: {{ $v.SourceSecurityGroupName }}
{{- end }}
{{- if $v.SourceSecurityGroupOwnerId }}
                  SourceSecurityGroupOwnerId: {{ $v.SourceSecurityGroupOwnerId }}
{{- end }}
{{- if $v.ToPort }}
                  ToPort: {{ $v.ToPort }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .VpcId }}
            VpcId: {{ .VpcId }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::SecurityGroupEgress" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::SecurityGroupEgress
        Properties:
            GroupId: {{ .GroupId }}
            IpProtocol: {{ .IpProtocol }}
{{- if .CidrIp }}
            CidrIp: {{ .CidrIp }}
{{- end }}
{{- if .CidrIpv6 }}
            CidrIpv6: {{ .CidrIpv6 }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .DestinationPrefixListId }}
            DestinationPrefixListId: {{ .DestinationPrefixListId }}
{{- end }}
{{- if .DestinationSecurityGroupId }}
            DestinationSecurityGroupId: {{ .DestinationSecurityGroupId }}
{{- end }}
{{- if .FromPort }}
            FromPort: {{ .FromPort }}
{{- end }}
{{- if .ToPort }}
            ToPort: {{ .ToPort }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::SecurityGroupIngress" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::SecurityGroupIngress
        Properties:
            IpProtocol: {{ .IpProtocol }}
{{- if .CidrIp }}
            CidrIp: {{ .CidrIp }}
{{- end }}
{{- if .CidrIpv6 }}
            CidrIpv6: {{ .CidrIpv6 }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .FromPort }}
            FromPort: {{ .FromPort }}
{{- end }}
{{- if .GroupId }}
            GroupId: {{ .GroupId }}
{{- end }}
{{- if .GroupName }}
            GroupName: {{ .GroupName }}
{{- end }}
{{- if .SourceSecurityGroupId }}
            SourceSecurityGroupId: {{ .SourceSecurityGroupId }}
{{- end }}
{{- if .SourceSecurityGroupName }}
            SourceSecurityGroupName: {{ .SourceSecurityGroupName }}
{{- end }}
{{- if .SourceSecurityGroupOwnerId }}
            SourceSecurityGroupOwnerId: {{ .SourceSecurityGroupOwnerId }}
{{- end }}
{{- if .ToPort }}
            ToPort: {{ .ToPort }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::SpotFleet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::SpotFleet
        Properties:
            SpotFleetRequestConfigData:
                IamFleetRole: {{ .SpotFleetRequestConfigData.IamFleetRole }}
                LaunchSpecifications:
{{- range $_, $v := .SpotFleetRequestConfigData.LaunchSpecifications }}
                    -
                      ImageId: {{ $v.ImageId }}
                      InstanceType: {{ $v.InstanceType }}
{{- if $v.BlockDeviceMappings }}
                      BlockDeviceMappings:
{{- range $_, $v := $v.BlockDeviceMappings }}
                          -
                            DeviceName: {{ $v.DeviceName }}
{{- if $v.Ebs }}
                            Ebs:
{{- if $v.Ebs.DeleteOnTermination }}
                                DeleteOnTermination: {{ $v.Ebs.DeleteOnTermination }}
{{- end }}
{{- if $v.Ebs.Encrypted }}
                                Encrypted: {{ $v.Ebs.Encrypted }}
{{- end }}
{{- if $v.Ebs.Iops }}
                                Iops: {{ $v.Ebs.Iops }}
{{- end }}
{{- if $v.Ebs.SnapshotId }}
                                SnapshotId: {{ $v.Ebs.SnapshotId }}
{{- end }}
{{- if $v.Ebs.VolumeSize }}
                                VolumeSize: {{ $v.Ebs.VolumeSize }}
{{- end }}
{{- if $v.Ebs.VolumeType }}
                                VolumeType: {{ $v.Ebs.VolumeType }}
{{- end }}
{{- end }}
{{- if $v.NoDevice }}
                            NoDevice: {{ $v.NoDevice }}
{{- end }}
{{- if $v.VirtualName }}
                            VirtualName: {{ $v.VirtualName }}
{{- end }}
{{- end }}
{{- end }}
{{- if $v.EbsOptimized }}
                      EbsOptimized: {{ $v.EbsOptimized }}
{{- end }}
{{- if $v.IamInstanceProfile }}
                      IamInstanceProfile:
{{- if $v.IamInstanceProfile.Arn }}
                          Arn: {{ $v.IamInstanceProfile.Arn }}
{{- end }}
{{- end }}
{{- if $v.KernelId }}
                      KernelId: {{ $v.KernelId }}
{{- end }}
{{- if $v.KeyName }}
                      KeyName: {{ $v.KeyName }}
{{- end }}
{{- if $v.Monitoring }}
                      Monitoring:
{{- if $v.Monitoring.Enabled }}
                          Enabled: {{ $v.Monitoring.Enabled }}
{{- end }}
{{- end }}
{{- if $v.NetworkInterfaces }}
                      NetworkInterfaces:
{{- range $_, $v := $v.NetworkInterfaces }}
                          -
{{- if $v.AssociatePublicIpAddress }}
                            AssociatePublicIpAddress: {{ $v.AssociatePublicIpAddress }}
{{- end }}
{{- if $v.DeleteOnTermination }}
                            DeleteOnTermination: {{ $v.DeleteOnTermination }}
{{- end }}
{{- if $v.Description }}
                            Description: {{ $v.Description }}
{{- end }}
{{- if $v.DeviceIndex }}
                            DeviceIndex: {{ $v.DeviceIndex }}
{{- end }}
{{- if $v.Groups }}
                            Groups:
{{- range $_, $v := $v.Groups }}
                                - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.Ipv6AddressCount }}
                            Ipv6AddressCount: {{ $v.Ipv6AddressCount }}
{{- end }}
{{- if $v.Ipv6Addresses }}
                            Ipv6Addresses:
{{- range $_, $v := $v.Ipv6Addresses }}
                                -
                                  Ipv6Address: {{ $v.Ipv6Address }}
{{- end }}
{{- end }}
{{- if $v.NetworkInterfaceId }}
                            NetworkInterfaceId: {{ $v.NetworkInterfaceId }}
{{- end }}
{{- if $v.PrivateIpAddresses }}
                            PrivateIpAddresses:
{{- range $_, $v := $v.PrivateIpAddresses }}
                                -
                                  PrivateIpAddress: {{ $v.PrivateIpAddress }}
{{- if $v.Primary }}
                                  Primary: {{ $v.Primary }}
{{- end }}
{{- end }}
{{- end }}
{{- if $v.SecondaryPrivateIpAddressCount }}
                            SecondaryPrivateIpAddressCount: {{ $v.SecondaryPrivateIpAddressCount }}
{{- end }}
{{- if $v.SubnetId }}
                            SubnetId: {{ $v.SubnetId }}
{{- end }}
{{- end }}
{{- end }}
{{- if $v.Placement }}
                      Placement:
{{- if $v.Placement.AvailabilityZone }}
                          AvailabilityZone: {{ $v.Placement.AvailabilityZone }}
{{- end }}
{{- if $v.Placement.GroupName }}
                          GroupName: {{ $v.Placement.GroupName }}
{{- end }}
{{- end }}
{{- if $v.RamdiskId }}
                      RamdiskId: {{ $v.RamdiskId }}
{{- end }}
{{- if $v.SecurityGroups }}
                      SecurityGroups:
{{- range $_, $v := $v.SecurityGroups }}
                          -
                            GroupId: {{ $v.GroupId }}
{{- end }}
{{- end }}
{{- if $v.SpotPrice }}
                      SpotPrice: {{ $v.SpotPrice }}
{{- end }}
{{- if $v.SubnetId }}
                      SubnetId: {{ $v.SubnetId }}
{{- end }}
{{- if $v.TagSpecifications }}
                      TagSpecifications:
{{- range $_, $v := $v.TagSpecifications }}
                          -
{{- if $v.ResourceType }}
                            ResourceType: {{ $v.ResourceType }}
{{- end }}
{{- end }}
{{- end }}
{{- if $v.UserData }}
                      UserData: {{ $v.UserData }}
{{- end }}
{{- if $v.WeightedCapacity }}
                      WeightedCapacity: {{ $v.WeightedCapacity }}
{{- end }}
{{- end }}
                TargetCapacity: {{ .SpotFleetRequestConfigData.TargetCapacity }}
{{- if .SpotFleetRequestConfigData.AllocationStrategy }}
                AllocationStrategy: {{ .SpotFleetRequestConfigData.AllocationStrategy }}
{{- end }}
{{- if .SpotFleetRequestConfigData.ExcessCapacityTerminationPolicy }}
                ExcessCapacityTerminationPolicy: {{ .SpotFleetRequestConfigData.ExcessCapacityTerminationPolicy }}
{{- end }}
{{- if .SpotFleetRequestConfigData.ReplaceUnhealthyInstances }}
                ReplaceUnhealthyInstances: {{ .SpotFleetRequestConfigData.ReplaceUnhealthyInstances }}
{{- end }}
{{- if .SpotFleetRequestConfigData.SpotPrice }}
                SpotPrice: {{ .SpotFleetRequestConfigData.SpotPrice }}
{{- end }}
{{- if .SpotFleetRequestConfigData.TerminateInstancesWithExpiration }}
                TerminateInstancesWithExpiration: {{ .SpotFleetRequestConfigData.TerminateInstancesWithExpiration }}
{{- end }}
{{- if .SpotFleetRequestConfigData.Type }}
                Type: {{ .SpotFleetRequestConfigData.Type }}
{{- end }}
{{- if .SpotFleetRequestConfigData.ValidFrom }}
                ValidFrom: {{ .SpotFleetRequestConfigData.ValidFrom }}
{{- end }}
{{- if .SpotFleetRequestConfigData.ValidUntil }}
                ValidUntil: {{ .SpotFleetRequestConfigData.ValidUntil }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::Subnet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::Subnet
        Properties:
            CidrBlock: {{ .CidrBlock }}
            VpcId: {{ .VpcId }}
{{- if .AssignIpv6AddressOnCreation }}
            AssignIpv6AddressOnCreation: {{ .AssignIpv6AddressOnCreation }}
{{- end }}
{{- if .AvailabilityZone }}
            AvailabilityZone: {{ .AvailabilityZone }}
{{- end }}
{{- if .Ipv6CidrBlock }}
            Ipv6CidrBlock: {{ .Ipv6CidrBlock }}
{{- end }}
{{- if .MapPublicIpOnLaunch }}
            MapPublicIpOnLaunch: {{ .MapPublicIpOnLaunch }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::SubnetCidrBlock" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::SubnetCidrBlock
        Properties:
            Ipv6CidrBlock: {{ .Ipv6CidrBlock }}
            SubnetId: {{ .SubnetId }}
{{- end }}
{{- define "AWS::EC2::SubnetNetworkAclAssociation" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::SubnetNetworkAclAssociation
        Properties:
            NetworkAclId: {{ .NetworkAclId }}
            SubnetId: {{ .SubnetId }}
{{- end }}
{{- define "AWS::EC2::SubnetRouteTableAssociation" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::SubnetRouteTableAssociation
        Properties:
            RouteTableId: {{ .RouteTableId }}
            SubnetId: {{ .SubnetId }}
{{- end }}
{{- define "AWS::EC2::TrunkInterfaceAssociation" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::TrunkInterfaceAssociation
        Properties:
            BranchInterfaceId: {{ .BranchInterfaceId }}
            TrunkInterfaceId: {{ .TrunkInterfaceId }}
{{- if .GREKey }}
            GREKey: {{ .GREKey }}
{{- end }}
{{- if .VLANId }}
            VLANId: {{ .VLANId }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::VPC" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::VPC
        Properties:
            CidrBlock: {{ .CidrBlock }}
{{- if .EnableDnsHostnames }}
            EnableDnsHostnames: {{ .EnableDnsHostnames }}
{{- end }}
{{- if .EnableDnsSupport }}
            EnableDnsSupport: {{ .EnableDnsSupport }}
{{- end }}
{{- if .InstanceTenancy }}
            InstanceTenancy: {{ .InstanceTenancy }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::VPCCidrBlock" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::VPCCidrBlock
        Properties:
            VpcId: {{ .VpcId }}
{{- if .AmazonProvidedIpv6CidrBlock }}
            AmazonProvidedIpv6CidrBlock: {{ .AmazonProvidedIpv6CidrBlock }}
{{- end }}
{{- if .CidrBlock }}
            CidrBlock: {{ .CidrBlock }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::VPCDHCPOptionsAssociation" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::VPCDHCPOptionsAssociation
        Properties:
            DhcpOptionsId: {{ .DhcpOptionsId }}
            VpcId: {{ .VpcId }}
{{- end }}
{{- define "AWS::EC2::VPCEndpoint" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::VPCEndpoint
        Properties:
            ServiceName: {{ .ServiceName }}
            VpcId: {{ .VpcId }}
{{- if .PolicyDocument }}
            PolicyDocument:
                {{ yaml.MarshalIndent .PolicyDocument "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .RouteTableIds }}
            RouteTableIds:
{{- range $_, $v := .RouteTableIds }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::VPCGatewayAttachment" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::VPCGatewayAttachment
        Properties:
            VpcId: {{ .VpcId }}
{{- if .InternetGatewayId }}
            InternetGatewayId: {{ .InternetGatewayId }}
{{- end }}
{{- if .VpnGatewayId }}
            VpnGatewayId: {{ .VpnGatewayId }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::VPCPeeringConnection" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::VPCPeeringConnection
        Properties:
            PeerVpcId: {{ .PeerVpcId }}
            VpcId: {{ .VpcId }}
{{- if .PeerOwnerId }}
            PeerOwnerId: {{ .PeerOwnerId }}
{{- end }}
{{- if .PeerRoleArn }}
            PeerRoleArn: {{ .PeerRoleArn }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::VPNConnection" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::VPNConnection
        Properties:
            CustomerGatewayId: {{ .CustomerGatewayId }}
            Type: {{ .Type }}
            VpnGatewayId: {{ .VpnGatewayId }}
{{- if .StaticRoutesOnly }}
            StaticRoutesOnly: {{ .StaticRoutesOnly }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .VpnTunnelOptionsSpecifications }}
            VpnTunnelOptionsSpecifications:
{{- range $_, $v := .VpnTunnelOptionsSpecifications }}
                -
{{- if $v.PreSharedKey }}
                  PreSharedKey: {{ $v.PreSharedKey }}
{{- end }}
{{- if $v.TunnelInsideCidr }}
                  TunnelInsideCidr: {{ $v.TunnelInsideCidr }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::VPNConnectionRoute" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::VPNConnectionRoute
        Properties:
            DestinationCidrBlock: {{ .DestinationCidrBlock }}
            VpnConnectionId: {{ .VpnConnectionId }}
{{- end }}
{{- define "AWS::EC2::VPNGateway" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::VPNGateway
        Properties:
            Type: {{ .Type }}
{{- if .AmazonSideAsn }}
            AmazonSideAsn: {{ .AmazonSideAsn }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::VPNGatewayRoutePropagation" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::VPNGatewayRoutePropagation
        Properties:
            RouteTableIds:
{{- range $_, $v := .RouteTableIds }}
                - {{ $v }}
{{- end }}
            VpnGatewayId: {{ .VpnGatewayId }}
{{- end }}
{{- define "AWS::EC2::Volume" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::Volume
        Properties:
            AvailabilityZone: {{ .AvailabilityZone }}
{{- if .AutoEnableIO }}
            AutoEnableIO: {{ .AutoEnableIO }}
{{- end }}
{{- if .Encrypted }}
            Encrypted: {{ .Encrypted }}
{{- end }}
{{- if .Iops }}
            Iops: {{ .Iops }}
{{- end }}
{{- if .KmsKeyId }}
            KmsKeyId: {{ .KmsKeyId }}
{{- end }}
{{- if .Size }}
            Size: {{ .Size }}
{{- end }}
{{- if .SnapshotId }}
            SnapshotId: {{ .SnapshotId }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .VolumeType }}
            VolumeType: {{ .VolumeType }}
{{- end }}
{{- end }}
{{- define "AWS::EC2::VolumeAttachment" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EC2::VolumeAttachment
        Properties:
            Device: {{ .Device }}
            InstanceId: {{ .InstanceId }}
            VolumeId: {{ .VolumeId }}
{{- end }}
{{- define "AWS::ECR::Repository" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ECR::Repository
{{- if .LifecyclePolicy | or .RepositoryName | or .RepositoryPolicyText }}
        Properties:
{{- if .LifecyclePolicy }}
            LifecyclePolicy:
{{- if .LifecyclePolicy.LifecyclePolicyText }}
                LifecyclePolicyText: {{ .LifecyclePolicy.LifecyclePolicyText }}
{{- end }}
{{- if .LifecyclePolicy.RegistryId }}
                RegistryId: {{ .LifecyclePolicy.RegistryId }}
{{- end }}
{{- end }}
{{- if .RepositoryName }}
            RepositoryName: {{ .RepositoryName }}
{{- end }}
{{- if .RepositoryPolicyText }}
            RepositoryPolicyText:
                {{ yaml.MarshalIndent .RepositoryPolicyText "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::ECS::Cluster" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ECS::Cluster
{{- if .ClusterName }}
        Properties:
{{- if .ClusterName }}
            ClusterName: {{ .ClusterName }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::ECS::Service" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ECS::Service
        Properties:
            TaskDefinition: {{ .TaskDefinition }}
{{- if .Cluster }}
            Cluster: {{ .Cluster }}
{{- end }}
{{- if .DeploymentConfiguration }}
            DeploymentConfiguration:
{{- if .DeploymentConfiguration.MaximumPercent }}
                MaximumPercent: {{ .DeploymentConfiguration.MaximumPercent }}
{{- end }}
{{- if .DeploymentConfiguration.MinimumHealthyPercent }}
                MinimumHealthyPercent: {{ .DeploymentConfiguration.MinimumHealthyPercent }}
{{- end }}
{{- end }}
{{- if .DesiredCount }}
            DesiredCount: {{ .DesiredCount }}
{{- end }}
{{- if .HealthCheckGracePeriodSeconds }}
            HealthCheckGracePeriodSeconds: {{ .HealthCheckGracePeriodSeconds }}
{{- end }}
{{- if .LaunchType }}
            LaunchType: {{ .LaunchType }}
{{- end }}
{{- if .LoadBalancers }}
            LoadBalancers:
{{- range $_, $v := .LoadBalancers }}
                -
                  ContainerPort: {{ $v.ContainerPort }}
{{- if $v.ContainerName }}
                  ContainerName: {{ $v.ContainerName }}
{{- end }}
{{- if $v.LoadBalancerName }}
                  LoadBalancerName: {{ $v.LoadBalancerName }}
{{- end }}
{{- if $v.TargetGroupArn }}
                  TargetGroupArn: {{ $v.TargetGroupArn }}
{{- end }}
{{- end }}
{{- end }}
{{- if .NetworkConfiguration }}
            NetworkConfiguration:
{{- if .NetworkConfiguration.AwsvpcConfiguration }}
                AwsvpcConfiguration:
                    Subnets:
{{- range $_, $v := .NetworkConfiguration.AwsvpcConfiguration.Subnets }}
                        - {{ $v }}
{{- end }}
{{- if .NetworkConfiguration.AwsvpcConfiguration.AssignPublicIp }}
                    AssignPublicIp: {{ .NetworkConfiguration.AwsvpcConfiguration.AssignPublicIp }}
{{- end }}
{{- if .NetworkConfiguration.AwsvpcConfiguration.SecurityGroups }}
                    SecurityGroups:
{{- range $_, $v := .NetworkConfiguration.AwsvpcConfiguration.SecurityGroups }}
                        - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .PlacementConstraints }}
            PlacementConstraints:
{{- range $_, $v := .PlacementConstraints }}
                -
                  Type: {{ $v.Type }}
{{- if $v.Expression }}
                  Expression: {{ $v.Expression }}
{{- end }}
{{- end }}
{{- end }}
{{- if .PlacementStrategies }}
            PlacementStrategies:
{{- range $_, $v := .PlacementStrategies }}
                -
                  Type: {{ $v.Type }}
{{- if $v.Field }}
                  Field: {{ $v.Field }}
{{- end }}
{{- end }}
{{- end }}
{{- if .PlatformVersion }}
            PlatformVersion: {{ .PlatformVersion }}
{{- end }}
{{- if .Role }}
            Role: {{ .Role }}
{{- end }}
{{- if .ServiceName }}
            ServiceName: {{ .ServiceName }}
{{- end }}
{{- end }}
{{- define "AWS::ECS::TaskDefinition" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ECS::TaskDefinition
{{- if .ContainerDefinitions | or .Cpu | or .ExecutionRoleArn | or .Family | or .Memory | or .NetworkMode | or .PlacementConstraints | or .RequiresCompatibilities | or .TaskRoleArn | or .Volumes }}
        Properties:
{{- if .ContainerDefinitions }}
            ContainerDefinitions:
{{- range $_, $v := .ContainerDefinitions }}
                -
{{- if $v.Command }}
                  Command:
{{- range $_, $v := $v.Command }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.Cpu }}
                  Cpu: {{ $v.Cpu }}
{{- end }}
{{- if $v.DisableNetworking }}
                  DisableNetworking: {{ $v.DisableNetworking }}
{{- end }}
{{- if $v.DnsSearchDomains }}
                  DnsSearchDomains:
{{- range $_, $v := $v.DnsSearchDomains }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.DnsServers }}
                  DnsServers:
{{- range $_, $v := $v.DnsServers }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.DockerLabels }}
                  DockerLabels:
{{- range $k, $v := $v.DockerLabels }}
                      {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if $v.DockerSecurityOptions }}
                  DockerSecurityOptions:
{{- range $_, $v := $v.DockerSecurityOptions }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.EntryPoint }}
                  EntryPoint:
{{- range $_, $v := $v.EntryPoint }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.Environment }}
                  Environment:
{{- range $_, $v := $v.Environment }}
                      -
{{- if $v.Name }}
                        Name: {{ $v.Name }}
{{- end }}
{{- if $v.Value }}
                        Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- if $v.Essential }}
                  Essential: {{ $v.Essential }}
{{- end }}
{{- if $v.ExtraHosts }}
                  ExtraHosts:
{{- range $_, $v := $v.ExtraHosts }}
                      -
                        Hostname: {{ $v.Hostname }}
                        IpAddress: {{ $v.IpAddress }}
{{- end }}
{{- end }}
{{- if $v.Hostname }}
                  Hostname: {{ $v.Hostname }}
{{- end }}
{{- if $v.Image }}
                  Image: {{ $v.Image }}
{{- end }}
{{- if $v.Links }}
                  Links:
{{- range $_, $v := $v.Links }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.LinuxParameters }}
                  LinuxParameters:
{{- if $v.LinuxParameters.Capabilities }}
                      Capabilities:
{{- if $v.LinuxParameters.Capabilities.Add }}
                          Add:
{{- range $_, $v := $v.LinuxParameters.Capabilities.Add }}
                              - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.LinuxParameters.Capabilities.Drop }}
                          Drop:
{{- range $_, $v := $v.LinuxParameters.Capabilities.Drop }}
                              - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if $v.LinuxParameters.Devices }}
                      Devices:
{{- range $_, $v := $v.LinuxParameters.Devices }}
                          -
                            HostPath: {{ $v.HostPath }}
{{- if $v.ContainerPath }}
                            ContainerPath: {{ $v.ContainerPath }}
{{- end }}
{{- if $v.Permissions }}
                            Permissions:
{{- range $_, $v := $v.Permissions }}
                                - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if $v.LinuxParameters.InitProcessEnabled }}
                      InitProcessEnabled: {{ $v.LinuxParameters.InitProcessEnabled }}
{{- end }}
{{- end }}
{{- if $v.LogConfiguration }}
                  LogConfiguration:
                      LogDriver: {{ $v.LogConfiguration.LogDriver }}
{{- if $v.LogConfiguration.Options }}
                      Options:
{{- range $k, $v := $v.LogConfiguration.Options }}
                          {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if $v.Memory }}
                  Memory: {{ $v.Memory }}
{{- end }}
{{- if $v.MemoryReservation }}
                  MemoryReservation: {{ $v.MemoryReservation }}
{{- end }}
{{- if $v.MountPoints }}
                  MountPoints:
{{- range $_, $v := $v.MountPoints }}
                      -
{{- if $v.ContainerPath }}
                        ContainerPath: {{ $v.ContainerPath }}
{{- end }}
{{- if $v.ReadOnly }}
                        ReadOnly: {{ $v.ReadOnly }}
{{- end }}
{{- if $v.SourceVolume }}
                        SourceVolume: {{ $v.SourceVolume }}
{{- end }}
{{- end }}
{{- end }}
{{- if $v.Name }}
                  Name: {{ $v.Name }}
{{- end }}
{{- if $v.PortMappings }}
                  PortMappings:
{{- range $_, $v := $v.PortMappings }}
                      -
{{- if $v.ContainerPort }}
                        ContainerPort: {{ $v.ContainerPort }}
{{- end }}
{{- if $v.HostPort }}
                        HostPort: {{ $v.HostPort }}
{{- end }}
{{- if $v.Protocol }}
                        Protocol: {{ $v.Protocol }}
{{- end }}
{{- end }}
{{- end }}
{{- if $v.Privileged }}
                  Privileged: {{ $v.Privileged }}
{{- end }}
{{- if $v.ReadonlyRootFilesystem }}
                  ReadonlyRootFilesystem: {{ $v.ReadonlyRootFilesystem }}
{{- end }}
{{- if $v.Ulimits }}
                  Ulimits:
{{- range $_, $v := $v.Ulimits }}
                      -
                        HardLimit: {{ $v.HardLimit }}
                        Name: {{ $v.Name }}
                        SoftLimit: {{ $v.SoftLimit }}
{{- end }}
{{- end }}
{{- if $v.User }}
                  User: {{ $v.User }}
{{- end }}
{{- if $v.VolumesFrom }}
                  VolumesFrom:
{{- range $_, $v := $v.VolumesFrom }}
                      -
{{- if $v.ReadOnly }}
                        ReadOnly: {{ $v.ReadOnly }}
{{- end }}
{{- if $v.SourceContainer }}
                        SourceContainer: {{ $v.SourceContainer }}
{{- end }}
{{- end }}
{{- end }}
{{- if $v.WorkingDirectory }}
                  WorkingDirectory: {{ $v.WorkingDirectory }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Cpu }}
            Cpu: {{ .Cpu }}
{{- end }}
{{- if .ExecutionRoleArn }}
            ExecutionRoleArn: {{ .ExecutionRoleArn }}
{{- end }}
{{- if .Family }}
            Family: {{ .Family }}
{{- end }}
{{- if .Memory }}
            Memory: {{ .Memory }}
{{- end }}
{{- if .NetworkMode }}
            NetworkMode: {{ .NetworkMode }}
{{- end }}
{{- if .PlacementConstraints }}
            PlacementConstraints:
{{- range $_, $v := .PlacementConstraints }}
                -
                  Type: {{ $v.Type }}
{{- if $v.Expression }}
                  Expression: {{ $v.Expression }}
{{- end }}
{{- end }}
{{- end }}
{{- if .RequiresCompatibilities }}
            RequiresCompatibilities:
{{- range $_, $v := .RequiresCompatibilities }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .TaskRoleArn }}
            TaskRoleArn: {{ .TaskRoleArn }}
{{- end }}
{{- if .Volumes }}
            Volumes:
{{- range $_, $v := .Volumes }}
                -
{{- if $v.Host }}
                  Host:
{{- if $v.Host.SourcePath }}
                      SourcePath: {{ $v.Host.SourcePath }}
{{- end }}
{{- end }}
{{- if $v.Name }}
                  Name: {{ $v.Name }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EFS::FileSystem" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EFS::FileSystem
{{- if .Encrypted | or .FileSystemTags | or .KmsKeyId | or .PerformanceMode }}
        Properties:
{{- if .Encrypted }}
            Encrypted: {{ .Encrypted }}
{{- end }}
{{- if .FileSystemTags }}
            FileSystemTags:
{{- range $_, $v := .FileSystemTags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .KmsKeyId }}
            KmsKeyId: {{ .KmsKeyId }}
{{- end }}
{{- if .PerformanceMode }}
            PerformanceMode: {{ .PerformanceMode }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::EFS::MountTarget" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::EFS::MountTarget
        Properties:
            FileSystemId: {{ .FileSystemId }}
            SecurityGroups:
{{- range $_, $v := .SecurityGroups }}
                - {{ $v }}
{{- end }}
            SubnetId: {{ .SubnetId }}
{{- if .IpAddress }}
            IpAddress: {{ .IpAddress }}
{{- end }}
{{- end }}
{{- define "AWS::ElastiCache::CacheCluster" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ElastiCache::CacheCluster
        Properties:
            CacheNodeType: {{ .CacheNodeType }}
            Engine: {{ .Engine }}
            NumCacheNodes: {{ .NumCacheNodes }}
{{- if .AZMode }}
            AZMode: {{ .AZMode }}
{{- end }}
{{- if .AutoMinorVersionUpgrade }}
            AutoMinorVersionUpgrade: {{ .AutoMinorVersionUpgrade }}
{{- end }}
{{- if .CacheParameterGroupName }}
            CacheParameterGroupName: {{ .CacheParameterGroupName }}
{{- end }}
{{- if .CacheSecurityGroupNames }}
            CacheSecurityGroupNames:
{{- range $_, $v := .CacheSecurityGroupNames }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .CacheSubnetGroupName }}
            CacheSubnetGroupName: {{ .CacheSubnetGroupName }}
{{- end }}
{{- if .ClusterName }}
            ClusterName: {{ .ClusterName }}
{{- end }}
{{- if .EngineVersion }}
            EngineVersion: {{ .EngineVersion }}
{{- end }}
{{- if .NotificationTopicArn }}
            NotificationTopicArn: {{ .NotificationTopicArn }}
{{- end }}
{{- if .Port }}
            Port: {{ .Port }}
{{- end }}
{{- if .PreferredAvailabilityZone }}
            PreferredAvailabilityZone: {{ .PreferredAvailabilityZone }}
{{- end }}
{{- if .PreferredAvailabilityZones }}
            PreferredAvailabilityZones:
{{- range $_, $v := .PreferredAvailabilityZones }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .PreferredMaintenanceWindow }}
            PreferredMaintenanceWindow: {{ .PreferredMaintenanceWindow }}
{{- end }}
{{- if .SnapshotArns }}
            SnapshotArns:
{{- range $_, $v := .SnapshotArns }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .SnapshotName }}
            SnapshotName: {{ .SnapshotName }}
{{- end }}
{{- if .SnapshotRetentionLimit }}
            SnapshotRetentionLimit: {{ .SnapshotRetentionLimit }}
{{- end }}
{{- if .SnapshotWindow }}
            SnapshotWindow: {{ .SnapshotWindow }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .VpcSecurityGroupIds }}
            VpcSecurityGroupIds:
{{- range $_, $v := .VpcSecurityGroupIds }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::ElastiCache::ParameterGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ElastiCache::ParameterGroup
        Properties:
            CacheParameterGroupFamily: {{ .CacheParameterGroupFamily }}
            Description: {{ .Description }}
{{- if .Properties }}
            Properties:
{{- range $k, $v := .Properties }}
                {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::ElastiCache::ReplicationGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ElastiCache::ReplicationGroup
        Properties:
            ReplicationGroupDescription: {{ .ReplicationGroupDescription }}
{{- if .AtRestEncryptionEnabled }}
            AtRestEncryptionEnabled: {{ .AtRestEncryptionEnabled }}
{{- end }}
{{- if .AuthToken }}
            AuthToken: {{ .AuthToken }}
{{- end }}
{{- if .AutoMinorVersionUpgrade }}
            AutoMinorVersionUpgrade: {{ .AutoMinorVersionUpgrade }}
{{- end }}
{{- if .AutomaticFailoverEnabled }}
            AutomaticFailoverEnabled: {{ .AutomaticFailoverEnabled }}
{{- end }}
{{- if .CacheNodeType }}
            CacheNodeType: {{ .CacheNodeType }}
{{- end }}
{{- if .CacheParameterGroupName }}
            CacheParameterGroupName: {{ .CacheParameterGroupName }}
{{- end }}
{{- if .CacheSecurityGroupNames }}
            CacheSecurityGroupNames:
{{- range $_, $v := .CacheSecurityGroupNames }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .CacheSubnetGroupName }}
            CacheSubnetGroupName: {{ .CacheSubnetGroupName }}
{{- end }}
{{- if .Engine }}
            Engine: {{ .Engine }}
{{- end }}
{{- if .EngineVersion }}
            EngineVersion: {{ .EngineVersion }}
{{- end }}
{{- if .NodeGroupConfiguration }}
            NodeGroupConfiguration:
{{- range $_, $v := .NodeGroupConfiguration }}
                -
{{- if $v.PrimaryAvailabilityZone }}
                  PrimaryAvailabilityZone: {{ $v.PrimaryAvailabilityZone }}
{{- end }}
{{- if $v.ReplicaAvailabilityZones }}
                  ReplicaAvailabilityZones:
{{- range $_, $v := $v.ReplicaAvailabilityZones }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.ReplicaCount }}
                  ReplicaCount: {{ $v.ReplicaCount }}
{{- end }}
{{- if $v.Slots }}
                  Slots: {{ $v.Slots }}
{{- end }}
{{- end }}
{{- end }}
{{- if .NotificationTopicArn }}
            NotificationTopicArn: {{ .NotificationTopicArn }}
{{- end }}
{{- if .NumCacheClusters }}
            NumCacheClusters: {{ .NumCacheClusters }}
{{- end }}
{{- if .NumNodeGroups }}
            NumNodeGroups: {{ .NumNodeGroups }}
{{- end }}
{{- if .Port }}
            Port: {{ .Port }}
{{- end }}
{{- if .PreferredCacheClusterAZs }}
            PreferredCacheClusterAZs:
{{- range $_, $v := .PreferredCacheClusterAZs }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .PreferredMaintenanceWindow }}
            PreferredMaintenanceWindow: {{ .PreferredMaintenanceWindow }}
{{- end }}
{{- if .PrimaryClusterId }}
            PrimaryClusterId: {{ .PrimaryClusterId }}
{{- end }}
{{- if .ReplicasPerNodeGroup }}
            ReplicasPerNodeGroup: {{ .ReplicasPerNodeGroup }}
{{- end }}
{{- if .ReplicationGroupId }}
            ReplicationGroupId: {{ .ReplicationGroupId }}
{{- end }}
{{- if .SecurityGroupIds }}
            SecurityGroupIds:
{{- range $_, $v := .SecurityGroupIds }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .SnapshotArns }}
            SnapshotArns:
{{- range $_, $v := .SnapshotArns }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .SnapshotName }}
            SnapshotName: {{ .SnapshotName }}
{{- end }}
{{- if .SnapshotRetentionLimit }}
            SnapshotRetentionLimit: {{ .SnapshotRetentionLimit }}
{{- end }}
{{- if .SnapshotWindow }}
            SnapshotWindow: {{ .SnapshotWindow }}
{{- end }}
{{- if .SnapshottingClusterId }}
            SnapshottingClusterId: {{ .SnapshottingClusterId }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .TransitEncryptionEnabled }}
            TransitEncryptionEnabled: {{ .TransitEncryptionEnabled }}
{{- end }}
{{- end }}
{{- define "AWS::ElastiCache::SecurityGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ElastiCache::SecurityGroup
        Properties:
            Description: {{ .Description }}
{{- end }}
{{- define "AWS::ElastiCache::SecurityGroupIngress" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ElastiCache::SecurityGroupIngress
        Properties:
            CacheSecurityGroupName: {{ .CacheSecurityGroupName }}
            EC2SecurityGroupName: {{ .EC2SecurityGroupName }}
{{- if .EC2SecurityGroupOwnerId }}
            EC2SecurityGroupOwnerId: {{ .EC2SecurityGroupOwnerId }}
{{- end }}
{{- end }}
{{- define "AWS::ElastiCache::SubnetGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ElastiCache::SubnetGroup
        Properties:
            Description: {{ .Description }}
            SubnetIds:
{{- range $_, $v := .SubnetIds }}
                - {{ $v }}
{{- end }}
{{- if .CacheSubnetGroupName }}
            CacheSubnetGroupName: {{ .CacheSubnetGroupName }}
{{- end }}
{{- end }}
{{- define "AWS::ElasticBeanstalk::Application" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ElasticBeanstalk::Application
{{- if .ApplicationName | or .Description | or .ResourceLifecycleConfig }}
        Properties:
{{- if .ApplicationName }}
            ApplicationName: {{ .ApplicationName }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .ResourceLifecycleConfig }}
            ResourceLifecycleConfig:
{{- if .ResourceLifecycleConfig.ServiceRole }}
                ServiceRole: {{ .ResourceLifecycleConfig.ServiceRole }}
{{- end }}
{{- if .ResourceLifecycleConfig.VersionLifecycleConfig }}
                VersionLifecycleConfig:
{{- if .ResourceLifecycleConfig.VersionLifecycleConfig.MaxAgeRule }}
                    MaxAgeRule:
{{- if .ResourceLifecycleConfig.VersionLifecycleConfig.MaxAgeRule.DeleteSourceFromS3 }}
                        DeleteSourceFromS3: {{ .ResourceLifecycleConfig.VersionLifecycleConfig.MaxAgeRule.DeleteSourceFromS3 }}
{{- end }}
{{- if .ResourceLifecycleConfig.VersionLifecycleConfig.MaxAgeRule.Enabled }}
                        Enabled: {{ .ResourceLifecycleConfig.VersionLifecycleConfig.MaxAgeRule.Enabled }}
{{- end }}
{{- if .ResourceLifecycleConfig.VersionLifecycleConfig.MaxAgeRule.MaxAgeInDays }}
                        MaxAgeInDays: {{ .ResourceLifecycleConfig.VersionLifecycleConfig.MaxAgeRule.MaxAgeInDays }}
{{- end }}
{{- end }}
{{- if .ResourceLifecycleConfig.VersionLifecycleConfig.MaxCountRule }}
                    MaxCountRule:
{{- if .ResourceLifecycleConfig.VersionLifecycleConfig.MaxCountRule.DeleteSourceFromS3 }}
                        DeleteSourceFromS3: {{ .ResourceLifecycleConfig.VersionLifecycleConfig.MaxCountRule.DeleteSourceFromS3 }}
{{- end }}
{{- if .ResourceLifecycleConfig.VersionLifecycleConfig.MaxCountRule.Enabled }}
                        Enabled: {{ .ResourceLifecycleConfig.VersionLifecycleConfig.MaxCountRule.Enabled }}
{{- end }}
{{- if .ResourceLifecycleConfig.VersionLifecycleConfig.MaxCountRule.MaxCount }}
                        MaxCount: {{ .ResourceLifecycleConfig.VersionLifecycleConfig.MaxCountRule.MaxCount }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::ElasticBeanstalk::ApplicationVersion" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ElasticBeanstalk::ApplicationVersion
        Properties:
            ApplicationName: {{ .ApplicationName }}
            SourceBundle:{{ if istype .SourceBundle "string" }} {{ .SourceBundle }}{{ else }}
                S3Bucket: {{ .SourceBundle.S3Bucket }}
                S3Key: {{ .SourceBundle.S3Key }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- end }}
{{- define "AWS::ElasticBeanstalk::ConfigurationTemplate" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ElasticBeanstalk::ConfigurationTemplate
        Properties:
            ApplicationName: {{ .ApplicationName }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .EnvironmentId }}
            EnvironmentId: {{ .EnvironmentId }}
{{- end }}
{{- if .OptionSettings }}
            OptionSettings:
{{- range $_, $v := .OptionSettings }}
                -
                  Namespace: {{ $v.Namespace }}
                  OptionName: {{ $v.OptionName }}
{{- if $v.ResourceName }}
                  ResourceName: {{ $v.ResourceName }}
{{- end }}
{{- if $v.Value }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- if .PlatformArn }}
            PlatformArn: {{ .PlatformArn }}
{{- end }}
{{- if .SolutionStackName }}
            SolutionStackName: {{ .SolutionStackName }}
{{- end }}
{{- if .SourceConfiguration }}
            SourceConfiguration:
                ApplicationName: {{ .SourceConfiguration.ApplicationName }}
                TemplateName: {{ .SourceConfiguration.TemplateName }}
{{- end }}
{{- end }}
{{- define "AWS::ElasticBeanstalk::Environment" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ElasticBeanstalk::Environment
        Properties:
            ApplicationName: {{ .ApplicationName }}
{{- if .CNAMEPrefix }}
            CNAMEPrefix: {{ .CNAMEPrefix }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .EnvironmentName }}
            EnvironmentName: {{ .EnvironmentName }}
{{- end }}
{{- if .OptionSettings }}
            OptionSettings:
{{- range $_, $v := .OptionSettings }}
                -
                  Namespace: {{ $v.Namespace }}
                  OptionName: {{ $v.OptionName }}
{{- if $v.ResourceName }}
                  ResourceName: {{ $v.ResourceName }}
{{- end }}
{{- if $v.Value }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- if .PlatformArn }}
            PlatformArn: {{ .PlatformArn }}
{{- end }}
{{- if .SolutionStackName }}
            SolutionStackName: {{ .SolutionStackName }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .TemplateName }}
            TemplateName: {{ .TemplateName }}
{{- end }}
{{- if .Tier }}
            Tier:
{{- if .Tier.Name }}
                Name: {{ .Tier.Name }}
{{- end }}
{{- if .Tier.Type }}
                Type: {{ .Tier.Type }}
{{- end }}
{{- if .Tier.Version }}
                Version: {{ .Tier.Version }}
{{- end }}
{{- end }}
{{- if .VersionLabel }}
            VersionLabel: {{ .VersionLabel }}
{{- end }}
{{- end }}
{{- define "AWS::ElasticLoadBalancing::LoadBalancer" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ElasticLoadBalancing::LoadBalancer
        Properties:
            Listeners:
{{- range $_, $v := .Listeners }}
                -
                  InstancePort: {{ $v.InstancePort }}
                  LoadBalancerPort: {{ $v.LoadBalancerPort }}
                  Protocol: {{ $v.Protocol }}
{{- if $v.InstanceProtocol }}
                  InstanceProtocol: {{ $v.InstanceProtocol }}
{{- end }}
{{- if $v.PolicyNames }}
                  PolicyNames:
{{- range $_, $v := $v.PolicyNames }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.SSLCertificateId }}
                  SSLCertificateId: {{ $v.SSLCertificateId }}
{{- end }}
{{- end }}
{{- if .AccessLoggingPolicy }}
            AccessLoggingPolicy:
                Enabled: {{ .AccessLoggingPolicy.Enabled }}
                S3BucketName: {{ .AccessLoggingPolicy.S3BucketName }}
{{- if .AccessLoggingPolicy.EmitInterval }}
                EmitInterval: {{ .AccessLoggingPolicy.EmitInterval }}
{{- end }}
{{- if .AccessLoggingPolicy.S3BucketPrefix }}
                S3BucketPrefix: {{ .AccessLoggingPolicy.S3BucketPrefix }}
{{- end }}
{{- end }}
{{- if .AppCookieStickinessPolicy }}
            AppCookieStickinessPolicy:
{{- range $_, $v := .AppCookieStickinessPolicy }}
                -
                  CookieName: {{ $v.CookieName }}
                  PolicyName: {{ $v.PolicyName }}
{{- end }}
{{- end }}
{{- if .AvailabilityZones }}
            AvailabilityZones:
{{- range $_, $v := .AvailabilityZones }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .ConnectionDrainingPolicy }}
            ConnectionDrainingPolicy:
                Enabled: {{ .ConnectionDrainingPolicy.Enabled }}
{{- if .ConnectionDrainingPolicy.Timeout }}
                Timeout: {{ .ConnectionDrainingPolicy.Timeout }}
{{- end }}
{{- end }}
{{- if .ConnectionSettings }}
            ConnectionSettings:
                IdleTimeout: {{ .ConnectionSettings.IdleTimeout }}
{{- end }}
{{- if .CrossZone }}
            CrossZone: {{ .CrossZone }}
{{- end }}
{{- if .HealthCheck }}
            HealthCheck:
                HealthyThreshold: {{ .HealthCheck.HealthyThreshold }}
                Interval: {{ .HealthCheck.Interval }}
                Target: {{ .HealthCheck.Target }}
                Timeout: {{ .HealthCheck.Timeout }}
                UnhealthyThreshold: {{ .HealthCheck.UnhealthyThreshold }}
{{- end }}
{{- if .Instances }}
            Instances:
{{- range $_, $v := .Instances }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .LBCookieStickinessPolicy }}
            LBCookieStickinessPolicy:
{{- range $_, $v := .LBCookieStickinessPolicy }}
                -
{{- if $v.CookieExpirationPeriod }}
                  CookieExpirationPeriod: {{ $v.CookieExpirationPeriod }}
{{- end }}
{{- if $v.PolicyName }}
                  PolicyName: {{ $v.PolicyName }}
{{- end }}
{{- end }}
{{- end }}
{{- if .LoadBalancerName }}
            LoadBalancerName: {{ .LoadBalancerName }}
{{- end }}
{{- if .Policies }}
            Policies:
{{- range $_, $v := .Policies }}
                -
                  Attributes:
{{- range $_, $v := $v.Attributes }}
                      - {{ $v }}
{{- end }}
                  PolicyName: {{ $v.PolicyName }}
                  PolicyType: {{ $v.PolicyType }}
{{- if $v.InstancePorts }}
                  InstancePorts:
{{- range $_, $v := $v.InstancePorts }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.LoadBalancerPorts }}
                  LoadBalancerPorts:
{{- range $_, $v := $v.LoadBalancerPorts }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Scheme }}
            Scheme: {{ .Scheme }}
{{- end }}
{{- if .SecurityGroups }}
            SecurityGroups:
{{- range $_, $v := .SecurityGroups }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Subnets }}
            Subnets:
{{- range $_, $v := .Subnets }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::ElasticLoadBalancingV2::Listener" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ElasticLoadBalancingV2::Listener
        Properties:
            DefaultActions:
{{- range $_, $v := .DefaultActions }}
                -
                  TargetGroupArn: {{ $v.TargetGroupArn }}
                  Type: {{ $v.Type }}
{{- end }}
            LoadBalancerArn: {{ .LoadBalancerArn }}
            Port: {{ .Port }}
            Protocol: {{ .Protocol }}
{{- if .Certificates }}
            Certificates:
{{- range $_, $v := .Certificates }}
                -
{{- if $v.CertificateArn }}
                  CertificateArn: {{ $v.CertificateArn }}
{{- end }}
{{- end }}
{{- end }}
{{- if .SslPolicy }}
            SslPolicy: {{ .SslPolicy }}
{{- end }}
{{- end }}
{{- define "AWS::ElasticLoadBalancingV2::ListenerCertificate" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ElasticLoadBalancingV2::ListenerCertificate
        Properties:
            Certificates:
{{- range $_, $v := .Certificates }}
                -
{{- if $v.CertificateArn }}
                  CertificateArn: {{ $v.CertificateArn }}
{{- end }}
{{- end }}
            ListenerArn: {{ .ListenerArn }}
{{- end }}
{{- define "AWS::ElasticLoadBalancingV2::ListenerRule" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ElasticLoadBalancingV2::ListenerRule
        Properties:
            Actions:
{{- range $_, $v := .Actions }}
                -
                  TargetGroupArn: {{ $v.TargetGroupArn }}
                  Type: {{ $v.Type }}
{{- end }}
            Conditions:
{{- range $_, $v := .Conditions }}
                -
{{- if $v.Field }}
                  Field: {{ $v.Field }}
{{- end }}
{{- if $v.Values }}
                  Values:
{{- range $_, $v := $v.Values }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
            ListenerArn: {{ .ListenerArn }}
            Priority: {{ .Priority }}
{{- end }}
{{- define "AWS::ElasticLoadBalancingV2::LoadBalancer" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ElasticLoadBalancingV2::LoadBalancer
{{- if .IpAddressType | or .LoadBalancerAttributes | or .Name | or .Scheme | or .SecurityGroups | or .SubnetMappings | or .Subnets | or .Tags | or .Type }}
        Properties:
{{- if .IpAddressType }}
            IpAddressType: {{ .IpAddressType }}
{{- end }}
{{- if .LoadBalancerAttributes }}
            LoadBalancerAttributes:
{{- range $_, $v := .LoadBalancerAttributes }}
                -
{{- if $v.Key }}
                  Key: {{ $v.Key }}
{{- end }}
{{- if $v.Value }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .Scheme }}
            Scheme: {{ .Scheme }}
{{- end }}
{{- if .SecurityGroups }}
            SecurityGroups:
{{- range $_, $v := .SecurityGroups }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .SubnetMappings }}
            SubnetMappings:
{{- range $_, $v := .SubnetMappings }}
                -
                  AllocationId: {{ $v.AllocationId }}
                  SubnetId: {{ $v.SubnetId }}
{{- end }}
{{- end }}
{{- if .Subnets }}
            Subnets:
{{- range $_, $v := .Subnets }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .Type }}
            Type: {{ .Type }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::ElasticLoadBalancingV2::TargetGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ElasticLoadBalancingV2::TargetGroup
        Properties:
            Port: {{ .Port }}
            Protocol: {{ .Protocol }}
            VpcId: {{ .VpcId }}
{{- if .HealthCheckIntervalSeconds }}
            HealthCheckIntervalSeconds: {{ .HealthCheckIntervalSeconds }}
{{- end }}
{{- if .HealthCheckPath }}
            HealthCheckPath: {{ .HealthCheckPath }}
{{- end }}
{{- if .HealthCheckPort }}
            HealthCheckPort: {{ .HealthCheckPort }}
{{- end }}
{{- if .HealthCheckProtocol }}
            HealthCheckProtocol: {{ .HealthCheckProtocol }}
{{- end }}
{{- if .HealthCheckTimeoutSeconds }}
            HealthCheckTimeoutSeconds: {{ .HealthCheckTimeoutSeconds }}
{{- end }}
{{- if .HealthyThresholdCount }}
            HealthyThresholdCount: {{ .HealthyThresholdCount }}
{{- end }}
{{- if .Matcher }}
            Matcher:
                HttpCode: {{ .Matcher.HttpCode }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .TargetGroupAttributes }}
            TargetGroupAttributes:
{{- range $_, $v := .TargetGroupAttributes }}
                -
{{- if $v.Key }}
                  Key: {{ $v.Key }}
{{- end }}
{{- if $v.Value }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- if .TargetType }}
            TargetType: {{ .TargetType }}
{{- end }}
{{- if .Targets }}
            Targets:
{{- range $_, $v := .Targets }}
                -
                  Id: {{ $v.Id }}
{{- if $v.AvailabilityZone }}
                  AvailabilityZone: {{ $v.AvailabilityZone }}
{{- end }}
{{- if $v.Port }}
                  Port: {{ $v.Port }}
{{- end }}
{{- end }}
{{- end }}
{{- if .UnhealthyThresholdCount }}
            UnhealthyThresholdCount: {{ .UnhealthyThresholdCount }}
{{- end }}
{{- end }}
{{- define "AWS::Elasticsearch::Domain" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Elasticsearch::Domain
{{- if .AccessPolicies | or .AdvancedOptions | or .DomainName | or .EBSOptions | or .ElasticsearchClusterConfig | or .ElasticsearchVersion | or .SnapshotOptions | or .Tags | or .VPCOptions }}
        Properties:
{{- if .AccessPolicies }}
            AccessPolicies:
                {{ yaml.MarshalIndent .AccessPolicies "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .AdvancedOptions }}
            AdvancedOptions:
{{- range $k, $v := .AdvancedOptions }}
                {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .DomainName }}
            DomainName: {{ .DomainName }}
{{- end }}
{{- if .EBSOptions }}
            EBSOptions:
{{- if .EBSOptions.EBSEnabled }}
                EBSEnabled: {{ .EBSOptions.EBSEnabled }}
{{- end }}
{{- if .EBSOptions.Iops }}
                Iops: {{ .EBSOptions.Iops }}
{{- end }}
{{- if .EBSOptions.VolumeSize }}
                VolumeSize: {{ .EBSOptions.VolumeSize }}
{{- end }}
{{- if .EBSOptions.VolumeType }}
                VolumeType: {{ .EBSOptions.VolumeType }}
{{- end }}
{{- end }}
{{- if .ElasticsearchClusterConfig }}
            ElasticsearchClusterConfig:
{{- if .ElasticsearchClusterConfig.DedicatedMasterCount }}
                DedicatedMasterCount: {{ .ElasticsearchClusterConfig.DedicatedMasterCount }}
{{- end }}
{{- if .ElasticsearchClusterConfig.DedicatedMasterEnabled }}
                DedicatedMasterEnabled: {{ .ElasticsearchClusterConfig.DedicatedMasterEnabled }}
{{- end }}
{{- if .ElasticsearchClusterConfig.DedicatedMasterType }}
                DedicatedMasterType: {{ .ElasticsearchClusterConfig.DedicatedMasterType }}
{{- end }}
{{- if .ElasticsearchClusterConfig.InstanceCount }}
                InstanceCount: {{ .ElasticsearchClusterConfig.InstanceCount }}
{{- end }}
{{- if .ElasticsearchClusterConfig.InstanceType }}
                InstanceType: {{ .ElasticsearchClusterConfig.InstanceType }}
{{- end }}
{{- if .ElasticsearchClusterConfig.ZoneAwarenessEnabled }}
                ZoneAwarenessEnabled: {{ .ElasticsearchClusterConfig.ZoneAwarenessEnabled }}
{{- end }}
{{- end }}
{{- if .ElasticsearchVersion }}
            ElasticsearchVersion: {{ .ElasticsearchVersion }}
{{- end }}
{{- if .SnapshotOptions }}
            SnapshotOptions:
{{- if .SnapshotOptions.AutomatedSnapshotStartHour }}
                AutomatedSnapshotStartHour: {{ .SnapshotOptions.AutomatedSnapshotStartHour }}
{{- end }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .VPCOptions }}
            VPCOptions:
{{- if .VPCOptions.SecurityGroupIds }}
                SecurityGroupIds:
{{- range $_, $v := .VPCOptions.SecurityGroupIds }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- if .VPCOptions.SubnetIds }}
                SubnetIds:
{{- range $_, $v := .VPCOptions.SubnetIds }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Events::Rule" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Events::Rule
{{- if .Description | or .EventPattern | or .Name | or .RoleArn | or .ScheduleExpression | or .State | or .Targets }}
        Properties:
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .EventPattern }}
            EventPattern:
                {{ yaml.MarshalIndent .EventPattern "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .RoleArn }}
            RoleArn: {{ .RoleArn }}
{{- end }}
{{- if .ScheduleExpression }}
            ScheduleExpression: {{ .ScheduleExpression }}
{{- end }}
{{- if .State }}
            State: {{ .State }}
{{- end }}
{{- if .Targets }}
            Targets:
{{- range $_, $v := .Targets }}
                -
                  Arn: {{ $v.Arn }}
                  Id: {{ $v.Id }}
{{- if $v.EcsParameters }}
                  EcsParameters:
                      TaskDefinitionArn: {{ $v.EcsParameters.TaskDefinitionArn }}
{{- if $v.EcsParameters.TaskCount }}
                      TaskCount: {{ $v.EcsParameters.TaskCount }}
{{- end }}
{{- end }}
{{- if $v.Input }}
                  Input: {{ $v.Input }}
{{- end }}
{{- if $v.InputPath }}
                  InputPath: {{ $v.InputPath }}
{{- end }}
{{- if $v.InputTransformer }}
                  InputTransformer:
                      InputTemplate: {{ $v.InputTransformer.InputTemplate }}
{{- if $v.InputTransformer.InputPathsMap }}
                      InputPathsMap:
{{- range $k, $v := $v.InputTransformer.InputPathsMap }}
                          {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if $v.KinesisParameters }}
                  KinesisParameters:
                      PartitionKeyPath: {{ $v.KinesisParameters.PartitionKeyPath }}
{{- end }}
{{- if $v.RoleArn }}
                  RoleArn: {{ $v.RoleArn }}
{{- end }}
{{- if $v.RunCommandParameters }}
                  RunCommandParameters:
                      RunCommandTargets:
{{- range $_, $v := $v.RunCommandParameters.RunCommandTargets }}
                          -
                            Key: {{ $v.Key }}
                            Values:
{{- range $_, $v := $v.Values }}
                                - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::GameLift::Alias" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::GameLift::Alias
        Properties:
            Name: {{ .Name }}
            RoutingStrategy:
                Type: {{ .RoutingStrategy.Type }}
{{- if .RoutingStrategy.FleetId }}
                FleetId: {{ .RoutingStrategy.FleetId }}
{{- end }}
{{- if .RoutingStrategy.Message }}
                Message: {{ .RoutingStrategy.Message }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- end }}
{{- define "AWS::GameLift::Build" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::GameLift::Build
{{- if .Name | or .StorageLocation | or .Version }}
        Properties:
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .StorageLocation }}
            StorageLocation:
                Bucket: {{ .StorageLocation.Bucket }}
                Key: {{ .StorageLocation.Key }}
                RoleArn: {{ .StorageLocation.RoleArn }}
{{- end }}
{{- if .Version }}
            Version: {{ .Version }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::GameLift::Fleet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::GameLift::Fleet
        Properties:
            BuildId: {{ .BuildId }}
            DesiredEC2Instances: {{ .DesiredEC2Instances }}
            EC2InstanceType: {{ .EC2InstanceType }}
            Name: {{ .Name }}
            ServerLaunchPath: {{ .ServerLaunchPath }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .EC2InboundPermissions }}
            EC2InboundPermissions:
{{- range $_, $v := .EC2InboundPermissions }}
                -
                  FromPort: {{ $v.FromPort }}
                  IpRange: {{ $v.IpRange }}
                  Protocol: {{ $v.Protocol }}
                  ToPort: {{ $v.ToPort }}
{{- end }}
{{- end }}
{{- if .LogPaths }}
            LogPaths:
{{- range $_, $v := .LogPaths }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .MaxSize }}
            MaxSize: {{ .MaxSize }}
{{- end }}
{{- if .MinSize }}
            MinSize: {{ .MinSize }}
{{- end }}
{{- if .ServerLaunchParameters }}
            ServerLaunchParameters: {{ .ServerLaunchParameters }}
{{- end }}
{{- end }}
{{- define "AWS::Glue::Classifier" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Glue::Classifier
{{- if .GrokClassifier }}
        Properties:
{{- if .GrokClassifier }}
            GrokClassifier:
                Classification: {{ .GrokClassifier.Classification }}
                GrokPattern: {{ .GrokClassifier.GrokPattern }}
{{- if .GrokClassifier.CustomPatterns }}
                CustomPatterns: {{ .GrokClassifier.CustomPatterns }}
{{- end }}
{{- if .GrokClassifier.Name }}
                Name: {{ .GrokClassifier.Name }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Glue::Connection" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Glue::Connection
        Properties:
            CatalogId: {{ .CatalogId }}
            ConnectionInput:
                ConnectionProperties:
                    {{ yaml.MarshalIndent .ConnectionInput.ConnectionProperties "                    " "    " | fmt.Sprintf "%s" }}
                ConnectionType: {{ .ConnectionInput.ConnectionType }}
{{- if .ConnectionInput.Description }}
                Description: {{ .ConnectionInput.Description }}
{{- end }}
{{- if .ConnectionInput.MatchCriteria }}
                MatchCriteria:
{{- range $_, $v := .ConnectionInput.MatchCriteria }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- if .ConnectionInput.Name }}
                Name: {{ .ConnectionInput.Name }}
{{- end }}
{{- if .ConnectionInput.PhysicalConnectionRequirements }}
                PhysicalConnectionRequirements:
{{- if .ConnectionInput.PhysicalConnectionRequirements.AvailabilityZone }}
                    AvailabilityZone: {{ .ConnectionInput.PhysicalConnectionRequirements.AvailabilityZone }}
{{- end }}
{{- if .ConnectionInput.PhysicalConnectionRequirements.SecurityGroupIdList }}
                    SecurityGroupIdList:
{{- range $_, $v := .ConnectionInput.PhysicalConnectionRequirements.SecurityGroupIdList }}
                        - {{ $v }}
{{- end }}
{{- end }}
{{- if .ConnectionInput.PhysicalConnectionRequirements.SubnetId }}
                    SubnetId: {{ .ConnectionInput.PhysicalConnectionRequirements.SubnetId }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Glue::Crawler" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Glue::Crawler
        Properties:
            DatabaseName: {{ .DatabaseName }}
            Role: {{ .Role }}
            Targets:
{{- if .Targets.JdbcTargets }}
                JdbcTargets:
{{- range $_, $v := .Targets.JdbcTargets }}
                    -
{{- if $v.ConnectionName }}
                      ConnectionName: {{ $v.ConnectionName }}
{{- end }}
{{- if $v.Exclusions }}
                      Exclusions:
{{- range $_, $v := $v.Exclusions }}
                          - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.Path }}
                      Path: {{ $v.Path }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Targets.S3Targets }}
                S3Targets:
{{- range $_, $v := .Targets.S3Targets }}
                    -
{{- if $v.Exclusions }}
                      Exclusions:
{{- range $_, $v := $v.Exclusions }}
                          - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.Path }}
                      Path: {{ $v.Path }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Classifiers }}
            Classifiers:
{{- range $_, $v := .Classifiers }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .Schedule }}
            Schedule:
{{- if .Schedule.ScheduleExpression }}
                ScheduleExpression: {{ .Schedule.ScheduleExpression }}
{{- end }}
{{- end }}
{{- if .SchemaChangePolicy }}
            SchemaChangePolicy:
{{- if .SchemaChangePolicy.DeleteBehavior }}
                DeleteBehavior: {{ .SchemaChangePolicy.DeleteBehavior }}
{{- end }}
{{- if .SchemaChangePolicy.UpdateBehavior }}
                UpdateBehavior: {{ .SchemaChangePolicy.UpdateBehavior }}
{{- end }}
{{- end }}
{{- if .TablePrefix }}
            TablePrefix: {{ .TablePrefix }}
{{- end }}
{{- end }}
{{- define "AWS::Glue::Database" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Glue::Database
        Properties:
            CatalogId: {{ .CatalogId }}
            DatabaseInput:
{{- if .DatabaseInput.Description }}
                Description: {{ .DatabaseInput.Description }}
{{- end }}
{{- if .DatabaseInput.LocationUri }}
                LocationUri: {{ .DatabaseInput.LocationUri }}
{{- end }}
{{- if .DatabaseInput.Name }}
                Name: {{ .DatabaseInput.Name }}
{{- end }}
{{- if .DatabaseInput.Parameters }}
                Parameters:
                    {{ yaml.MarshalIndent .DatabaseInput.Parameters "                    " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- end }}
{{- define "AWS::Glue::DevEndpoint" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Glue::DevEndpoint
        Properties:
            PublicKey: {{ .PublicKey }}
            RoleArn: {{ .RoleArn }}
{{- if .EndpointName }}
            EndpointName: {{ .EndpointName }}
{{- end }}
{{- if .ExtraJarsS3Path }}
            ExtraJarsS3Path: {{ .ExtraJarsS3Path }}
{{- end }}
{{- if .ExtraPythonLibsS3Path }}
            ExtraPythonLibsS3Path: {{ .ExtraPythonLibsS3Path }}
{{- end }}
{{- if .NumberOfNodes }}
            NumberOfNodes: {{ .NumberOfNodes }}
{{- end }}
{{- if .SecurityGroupIds }}
            SecurityGroupIds:
{{- range $_, $v := .SecurityGroupIds }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .SubnetId }}
            SubnetId: {{ .SubnetId }}
{{- end }}
{{- end }}
{{- define "AWS::Glue::Job" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Glue::Job
        Properties:
            Command:
{{- if .Command.Name }}
                Name: {{ .Command.Name }}
{{- end }}
{{- if .Command.ScriptLocation }}
                ScriptLocation: {{ .Command.ScriptLocation }}
{{- end }}
            Role: {{ .Role }}
{{- if .AllocatedCapacity }}
            AllocatedCapacity: {{ .AllocatedCapacity }}
{{- end }}
{{- if .Connections }}
            Connections:
{{- if .Connections.Connections }}
                Connections:
{{- range $_, $v := .Connections.Connections }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if .DefaultArguments }}
            DefaultArguments:
                {{ yaml.MarshalIndent .DefaultArguments "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .ExecutionProperty }}
            ExecutionProperty:
{{- if .ExecutionProperty.MaxConcurrentRuns }}
                MaxConcurrentRuns: {{ .ExecutionProperty.MaxConcurrentRuns }}
{{- end }}
{{- end }}
{{- if .LogUri }}
            LogUri: {{ .LogUri }}
{{- end }}
{{- if .MaxRetries }}
            MaxRetries: {{ .MaxRetries }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- end }}
{{- define "AWS::Glue::Partition" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Glue::Partition
        Properties:
            CatalogId: {{ .CatalogId }}
            DatabaseName: {{ .DatabaseName }}
            PartitionInput:
                Values:
{{- range $_, $v := .PartitionInput.Values }}
                    - {{ $v }}
{{- end }}
{{- if .PartitionInput.Parameters }}
                Parameters:
                    {{ yaml.MarshalIndent .PartitionInput.Parameters "                    " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .PartitionInput.StorageDescriptor }}
                StorageDescriptor:
{{- if .PartitionInput.StorageDescriptor.BucketColumns }}
                    BucketColumns:
{{- range $_, $v := .PartitionInput.StorageDescriptor.BucketColumns }}
                        - {{ $v }}
{{- end }}
{{- end }}
{{- if .PartitionInput.StorageDescriptor.Columns }}
                    Columns:
{{- range $_, $v := .PartitionInput.StorageDescriptor.Columns }}
                        -
                          Name: {{ $v.Name }}
{{- if $v.Comment }}
                          Comment: {{ $v.Comment }}
{{- end }}
{{- if $v.Type }}
                          Type: {{ $v.Type }}
{{- end }}
{{- end }}
{{- end }}
{{- if .PartitionInput.StorageDescriptor.Compressed }}
                    Compressed: {{ .PartitionInput.StorageDescriptor.Compressed }}
{{- end }}
{{- if .PartitionInput.StorageDescriptor.InputFormat }}
                    InputFormat: {{ .PartitionInput.StorageDescriptor.InputFormat }}
{{- end }}
{{- if .PartitionInput.StorageDescriptor.Location }}
                    Location: {{ .PartitionInput.StorageDescriptor.Location }}
{{- end }}
{{- if .PartitionInput.StorageDescriptor.NumberOfBuckets }}
                    NumberOfBuckets: {{ .PartitionInput.StorageDescriptor.NumberOfBuckets }}
{{- end }}
{{- if .PartitionInput.StorageDescriptor.OutputFormat }}
                    OutputFormat: {{ .PartitionInput.StorageDescriptor.OutputFormat }}
{{- end }}
{{- if .PartitionInput.StorageDescriptor.Parameters }}
                    Parameters:
                        {{ yaml.MarshalIndent .PartitionInput.StorageDescriptor.Parameters "                        " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .PartitionInput.StorageDescriptor.SerdeInfo }}
                    SerdeInfo:
{{- if .PartitionInput.StorageDescriptor.SerdeInfo.Name }}
                        Name: {{ .PartitionInput.StorageDescriptor.SerdeInfo.Name }}
{{- end }}
{{- if .PartitionInput.StorageDescriptor.SerdeInfo.Parameters }}
                        Parameters:
                            {{ yaml.MarshalIndent .PartitionInput.StorageDescriptor.SerdeInfo.Parameters "                            " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .PartitionInput.StorageDescriptor.SerdeInfo.SerializationLibrary }}
                        SerializationLibrary: {{ .PartitionInput.StorageDescriptor.SerdeInfo.SerializationLibrary }}
{{- end }}
{{- end }}
{{- if .PartitionInput.StorageDescriptor.SkewedInfo }}
                    SkewedInfo:
{{- if .PartitionInput.StorageDescriptor.SkewedInfo.SkewedColumnNames }}
                        SkewedColumnNames:
{{- range $_, $v := .PartitionInput.StorageDescriptor.SkewedInfo.SkewedColumnNames }}
                            - {{ $v }}
{{- end }}
{{- end }}
{{- if .PartitionInput.StorageDescriptor.SkewedInfo.SkewedColumnValueLocationMaps }}
                        SkewedColumnValueLocationMaps:
                            {{ yaml.MarshalIndent .PartitionInput.StorageDescriptor.SkewedInfo.SkewedColumnValueLocationMaps "                            " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .PartitionInput.StorageDescriptor.SkewedInfo.SkewedColumnValues }}
                        SkewedColumnValues:
{{- range $_, $v := .PartitionInput.StorageDescriptor.SkewedInfo.SkewedColumnValues }}
                            - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if .PartitionInput.StorageDescriptor.SortColumns }}
                    SortColumns:
{{- range $_, $v := .PartitionInput.StorageDescriptor.SortColumns }}
                        -
                          Column: {{ $v.Column }}
{{- if $v.SortOrder }}
                          SortOrder: {{ $v.SortOrder }}
{{- end }}
{{- end }}
{{- end }}
{{- if .PartitionInput.StorageDescriptor.StoredAsSubDirectories }}
                    StoredAsSubDirectories: {{ .PartitionInput.StorageDescriptor.StoredAsSubDirectories }}
{{- end }}
{{- end }}
            TableName: {{ .TableName }}
{{- end }}
{{- define "AWS::Glue::Table" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Glue::Table
        Properties:
            CatalogId: {{ .CatalogId }}
            DatabaseName: {{ .DatabaseName }}
            TableInput:
{{- if .TableInput.Description }}
                Description: {{ .TableInput.Description }}
{{- end }}
{{- if .TableInput.Name }}
                Name: {{ .TableInput.Name }}
{{- end }}
{{- if .TableInput.Owner }}
                Owner: {{ .TableInput.Owner }}
{{- end }}
{{- if .TableInput.Parameters }}
                Parameters:
                    {{ yaml.MarshalIndent .TableInput.Parameters "                    " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .TableInput.PartitionKeys }}
                PartitionKeys:
{{- range $_, $v := .TableInput.PartitionKeys }}
                    -
                      Name: {{ $v.Name }}
{{- if $v.Comment }}
                      Comment: {{ $v.Comment }}
{{- end }}
{{- if $v.Type }}
                      Type: {{ $v.Type }}
{{- end }}
{{- end }}
{{- end }}
{{- if .TableInput.Retention }}
                Retention: {{ .TableInput.Retention }}
{{- end }}
{{- if .TableInput.StorageDescriptor }}
                StorageDescriptor:
{{- if .TableInput.StorageDescriptor.BucketColumns }}
                    BucketColumns:
{{- range $_, $v := .TableInput.StorageDescriptor.BucketColumns }}
                        - {{ $v }}
{{- end }}
{{- end }}
{{- if .TableInput.StorageDescriptor.Columns }}
                    Columns:
{{- range $_, $v := .TableInput.StorageDescriptor.Columns }}
                        -
                          Name: {{ $v.Name }}
{{- if $v.Comment }}
                          Comment: {{ $v.Comment }}
{{- end }}
{{- if $v.Type }}
                          Type: {{ $v.Type }}
{{- end }}
{{- end }}
{{- end }}
{{- if .TableInput.StorageDescriptor.Compressed }}
                    Compressed: {{ .TableInput.StorageDescriptor.Compressed }}
{{- end }}
{{- if .TableInput.StorageDescriptor.InputFormat }}
                    InputFormat: {{ .TableInput.StorageDescriptor.InputFormat }}
{{- end }}
{{- if .TableInput.StorageDescriptor.Location }}
                    Location: {{ .TableInput.StorageDescriptor.Location }}
{{- end }}
{{- if .TableInput.StorageDescriptor.NumberOfBuckets }}
                    NumberOfBuckets: {{ .TableInput.StorageDescriptor.NumberOfBuckets }}
{{- end }}
{{- if .TableInput.StorageDescriptor.OutputFormat }}
                    OutputFormat: {{ .TableInput.StorageDescriptor.OutputFormat }}
{{- end }}
{{- if .TableInput.StorageDescriptor.Parameters }}
                    Parameters:
                        {{ yaml.MarshalIndent .TableInput.StorageDescriptor.Parameters "                        " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .TableInput.StorageDescriptor.SerdeInfo }}
                    SerdeInfo:
{{- if .TableInput.StorageDescriptor.SerdeInfo.Name }}
                        Name: {{ .TableInput.StorageDescriptor.SerdeInfo.Name }}
{{- end }}
{{- if .TableInput.StorageDescriptor.SerdeInfo.Parameters }}
                        Parameters:
                            {{ yaml.MarshalIndent .TableInput.StorageDescriptor.SerdeInfo.Parameters "                            " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .TableInput.StorageDescriptor.SerdeInfo.SerializationLibrary }}
                        SerializationLibrary: {{ .TableInput.StorageDescriptor.SerdeInfo.SerializationLibrary }}
{{- end }}
{{- end }}
{{- if .TableInput.StorageDescriptor.SkewedInfo }}
                    SkewedInfo:
{{- if .TableInput.StorageDescriptor.SkewedInfo.SkewedColumnNames }}
                        SkewedColumnNames:
{{- range $_, $v := .TableInput.StorageDescriptor.SkewedInfo.SkewedColumnNames }}
                            - {{ $v }}
{{- end }}
{{- end }}
{{- if .TableInput.StorageDescriptor.SkewedInfo.SkewedColumnValueLocationMaps }}
                        SkewedColumnValueLocationMaps:
                            {{ yaml.MarshalIndent .TableInput.StorageDescriptor.SkewedInfo.SkewedColumnValueLocationMaps "                            " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .TableInput.StorageDescriptor.SkewedInfo.SkewedColumnValues }}
                        SkewedColumnValues:
{{- range $_, $v := .TableInput.StorageDescriptor.SkewedInfo.SkewedColumnValues }}
                            - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if .TableInput.StorageDescriptor.SortColumns }}
                    SortColumns:
{{- range $_, $v := .TableInput.StorageDescriptor.SortColumns }}
                        -
                          Column: {{ $v.Column }}
                          SortOrder: {{ $v.SortOrder }}
{{- end }}
{{- end }}
{{- if .TableInput.StorageDescriptor.StoredAsSubDirectories }}
                    StoredAsSubDirectories: {{ .TableInput.StorageDescriptor.StoredAsSubDirectories }}
{{- end }}
{{- end }}
{{- if .TableInput.TableType }}
                TableType: {{ .TableInput.TableType }}
{{- end }}
{{- if .TableInput.ViewExpandedText }}
                ViewExpandedText: {{ .TableInput.ViewExpandedText }}
{{- end }}
{{- if .TableInput.ViewOriginalText }}
                ViewOriginalText: {{ .TableInput.ViewOriginalText }}
{{- end }}
{{- end }}
{{- define "AWS::Glue::Trigger" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Glue::Trigger
        Properties:
            Actions:
{{- range $_, $v := .Actions }}
                -
{{- if $v.Arguments }}
                  Arguments:
                      {{ yaml.MarshalIndent $v.Arguments "                      " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if $v.JobName }}
                  JobName: {{ $v.JobName }}
{{- end }}
{{- end }}
            Type: {{ .Type }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .Predicate }}
            Predicate:
{{- if .Predicate.Conditions }}
                Conditions:
{{- range $_, $v := .Predicate.Conditions }}
                    -
{{- if $v.JobName }}
                      JobName: {{ $v.JobName }}
{{- end }}
{{- if $v.LogicalOperator }}
                      LogicalOperator: {{ $v.LogicalOperator }}
{{- end }}
{{- if $v.State }}
                      State: {{ $v.State }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Predicate.Logical }}
                Logical: {{ .Predicate.Logical }}
{{- end }}
{{- end }}
{{- if .Schedule }}
            Schedule: {{ .Schedule }}
{{- end }}
{{- end }}
{{- define "AWS::GuardDuty::Detector" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::GuardDuty::Detector
        Properties:
            Enable: {{ .Enable }}
{{- end }}
{{- define "AWS::GuardDuty::IPSet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::GuardDuty::IPSet
        Properties:
            Activate: {{ .Activate }}
            DetectorId: {{ .DetectorId }}
            Format: {{ .Format }}
            Location: {{ .Location }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- end }}
{{- define "AWS::GuardDuty::Master" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::GuardDuty::Master
        Properties:
            DetectorId: {{ .DetectorId }}
            InvitationId: {{ .InvitationId }}
            MasterId: {{ .MasterId }}
{{- end }}
{{- define "AWS::GuardDuty::Member" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::GuardDuty::Member
        Properties:
            DetectorId: {{ .DetectorId }}
            Email: {{ .Email }}
            MemberId: {{ .MemberId }}
{{- if .Message }}
            Message: {{ .Message }}
{{- end }}
{{- if .Status }}
            Status: {{ .Status }}
{{- end }}
{{- end }}
{{- define "AWS::GuardDuty::ThreatIntelSet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::GuardDuty::ThreatIntelSet
        Properties:
            Activate: {{ .Activate }}
            DetectorId: {{ .DetectorId }}
            Format: {{ .Format }}
            Location: {{ .Location }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- end }}
{{- define "AWS::IAM::AccessKey" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::IAM::AccessKey
        Properties:
            UserName: {{ .UserName }}
{{- if .Serial }}
            Serial: {{ .Serial }}
{{- end }}
{{- if .Status }}
            Status: {{ .Status }}
{{- end }}
{{- end }}
{{- define "AWS::IAM::Group" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::IAM::Group
{{- if .GroupName | or .ManagedPolicyArns | or .Path | or .Policies }}
        Properties:
{{- if .GroupName }}
            GroupName: {{ .GroupName }}
{{- end }}
{{- if .ManagedPolicyArns }}
            ManagedPolicyArns:
{{- range $_, $v := .ManagedPolicyArns }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Path }}
            Path: {{ .Path }}
{{- end }}
{{- if .Policies }}
            Policies:
{{- range $_, $v := .Policies }}
                -
                  PolicyDocument:
                      {{ yaml.MarshalIndent $v.PolicyDocument "                      " "    " | fmt.Sprintf "%s" }}
                  PolicyName: {{ $v.PolicyName }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::IAM::InstanceProfile" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::IAM::InstanceProfile
        Properties:
            Roles:
{{- range $_, $v := .Roles }}
                - {{ $v }}
{{- end }}
{{- if .InstanceProfileName }}
            InstanceProfileName: {{ .InstanceProfileName }}
{{- end }}
{{- if .Path }}
            Path: {{ .Path }}
{{- end }}
{{- end }}
{{- define "AWS::IAM::ManagedPolicy" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::IAM::ManagedPolicy
        Properties:
            PolicyDocument:
                {{ yaml.MarshalIndent .PolicyDocument "                " "    " | fmt.Sprintf "%s" }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .Groups }}
            Groups:
{{- range $_, $v := .Groups }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .ManagedPolicyName }}
            ManagedPolicyName: {{ .ManagedPolicyName }}
{{- end }}
{{- if .Path }}
            Path: {{ .Path }}
{{- end }}
{{- if .Roles }}
            Roles:
{{- range $_, $v := .Roles }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Users }}
            Users:
{{- range $_, $v := .Users }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::IAM::Policy" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::IAM::Policy
        Properties:
            PolicyDocument:
                {{ yaml.MarshalIndent .PolicyDocument "                " "    " | fmt.Sprintf "%s" }}
            PolicyName: {{ .PolicyName }}
{{- if .Groups }}
            Groups:
{{- range $_, $v := .Groups }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Roles }}
            Roles:
{{- range $_, $v := .Roles }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Users }}
            Users:
{{- range $_, $v := .Users }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::IAM::Role" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::IAM::Role
        Properties:
            AssumeRolePolicyDocument:
                {{ yaml.MarshalIndent .AssumeRolePolicyDocument "                " "    " | fmt.Sprintf "%s" }}
{{- if .ManagedPolicyArns }}
            ManagedPolicyArns:
{{- range $_, $v := .ManagedPolicyArns }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Path }}
            Path: {{ .Path }}
{{- end }}
{{- if .Policies }}
            Policies:
{{- range $_, $v := .Policies }}
                -
                  PolicyDocument:
                      {{ yaml.MarshalIndent $v.PolicyDocument "                      " "    " | fmt.Sprintf "%s" }}
                  PolicyName: {{ $v.PolicyName }}
{{- end }}
{{- end }}
{{- if .RoleName }}
            RoleName: {{ .RoleName }}
{{- end }}
{{- end }}
{{- define "AWS::IAM::User" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::IAM::User
{{- if .Groups | or .LoginProfile | or .ManagedPolicyArns | or .Path | or .Policies | or .UserName }}
        Properties:
{{- if .Groups }}
            Groups:
{{- range $_, $v := .Groups }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .LoginProfile }}
            LoginProfile:
                Password: {{ .LoginProfile.Password }}
{{- if .LoginProfile.PasswordResetRequired }}
                PasswordResetRequired: {{ .LoginProfile.PasswordResetRequired }}
{{- end }}
{{- end }}
{{- if .ManagedPolicyArns }}
            ManagedPolicyArns:
{{- range $_, $v := .ManagedPolicyArns }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Path }}
            Path: {{ .Path }}
{{- end }}
{{- if .Policies }}
            Policies:
{{- range $_, $v := .Policies }}
                -
                  PolicyDocument:
                      {{ yaml.MarshalIndent $v.PolicyDocument "                      " "    " | fmt.Sprintf "%s" }}
                  PolicyName: {{ $v.PolicyName }}
{{- end }}
{{- end }}
{{- if .UserName }}
            UserName: {{ .UserName }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::IAM::UserToGroupAddition" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::IAM::UserToGroupAddition
        Properties:
            GroupName: {{ .GroupName }}
            Users:
{{- range $_, $v := .Users }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- define "AWS::Inspector::AssessmentTarget" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Inspector::AssessmentTarget
        Properties:
            ResourceGroupArn: {{ .ResourceGroupArn }}
{{- if .AssessmentTargetName }}
            AssessmentTargetName: {{ .AssessmentTargetName }}
{{- end }}
{{- end }}
{{- define "AWS::Inspector::AssessmentTemplate" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Inspector::AssessmentTemplate
        Properties:
            AssessmentTargetArn: {{ .AssessmentTargetArn }}
            DurationInSeconds: {{ .DurationInSeconds }}
            RulesPackageArns:
{{- range $_, $v := .RulesPackageArns }}
                - {{ $v }}
{{- end }}
{{- if .AssessmentTemplateName }}
            AssessmentTemplateName: {{ .AssessmentTemplateName }}
{{- end }}
{{- if .UserAttributesForFindings }}
            UserAttributesForFindings:
{{- range $_, $v := .UserAttributesForFindings }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Inspector::ResourceGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Inspector::ResourceGroup
        Properties:
            ResourceGroupTags:
{{- range $_, $v := .ResourceGroupTags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- define "AWS::IoT::Certificate" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::IoT::Certificate
        Properties:
            CertificateSigningRequest: {{ .CertificateSigningRequest }}
            Status: {{ .Status }}
{{- end }}
{{- define "AWS::IoT::Policy" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::IoT::Policy
        Properties:
            PolicyDocument:
                {{ yaml.MarshalIndent .PolicyDocument "                " "    " | fmt.Sprintf "%s" }}
{{- if .PolicyName }}
            PolicyName: {{ .PolicyName }}
{{- end }}
{{- end }}
{{- define "AWS::IoT::PolicyPrincipalAttachment" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::IoT::PolicyPrincipalAttachment
        Properties:
            PolicyName: {{ .PolicyName }}
            Principal: {{ .Principal }}
{{- end }}
{{- define "AWS::IoT::Thing" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::IoT::Thing
{{- if .AttributePayload | or .ThingName }}
        Properties:
{{- if .AttributePayload }}
            AttributePayload:
{{- if .AttributePayload.Attributes }}
                Attributes:
{{- range $k, $v := .AttributePayload.Attributes }}
                    {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if .ThingName }}
            ThingName: {{ .ThingName }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::IoT::ThingPrincipalAttachment" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::IoT::ThingPrincipalAttachment
        Properties:
            Principal: {{ .Principal }}
            ThingName: {{ .ThingName }}
{{- end }}
{{- define "AWS::IoT::TopicRule" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::IoT::TopicRule
        Properties:
            TopicRulePayload:
                Actions:
{{- range $_, $v := .TopicRulePayload.Actions }}
                    -
{{- if $v.CloudwatchAlarm }}
                      CloudwatchAlarm:
                          AlarmName: {{ $v.CloudwatchAlarm.AlarmName }}
                          RoleArn: {{ $v.CloudwatchAlarm.RoleArn }}
                          StateReason: {{ $v.CloudwatchAlarm.StateReason }}
                          StateValue: {{ $v.CloudwatchAlarm.StateValue }}
{{- end }}
{{- if $v.CloudwatchMetric }}
                      CloudwatchMetric:
                          MetricName: {{ $v.CloudwatchMetric.MetricName }}
                          MetricNamespace: {{ $v.CloudwatchMetric.MetricNamespace }}
                          MetricUnit: {{ $v.CloudwatchMetric.MetricUnit }}
                          MetricValue: {{ $v.CloudwatchMetric.MetricValue }}
                          RoleArn: {{ $v.CloudwatchMetric.RoleArn }}
{{- if $v.CloudwatchMetric.MetricTimestamp }}
                          MetricTimestamp: {{ $v.CloudwatchMetric.MetricTimestamp }}
{{- end }}
{{- end }}
{{- if $v.DynamoDB }}
                      DynamoDB:
                          HashKeyField: {{ $v.DynamoDB.HashKeyField }}
                          HashKeyValue: {{ $v.DynamoDB.HashKeyValue }}
                          RoleArn: {{ $v.DynamoDB.RoleArn }}
                          TableName: {{ $v.DynamoDB.TableName }}
{{- if $v.DynamoDB.HashKeyType }}
                          HashKeyType: {{ $v.DynamoDB.HashKeyType }}
{{- end }}
{{- if $v.DynamoDB.PayloadField }}
                          PayloadField: {{ $v.DynamoDB.PayloadField }}
{{- end }}
{{- if $v.DynamoDB.RangeKeyField }}
                          RangeKeyField: {{ $v.DynamoDB.RangeKeyField }}
{{- end }}
{{- if $v.DynamoDB.RangeKeyType }}
                          RangeKeyType: {{ $v.DynamoDB.RangeKeyType }}
{{- end }}
{{- if $v.DynamoDB.RangeKeyValue }}
                          RangeKeyValue: {{ $v.DynamoDB.RangeKeyValue }}
{{- end }}
{{- end }}
{{- if $v.DynamoDBv2 }}
                      DynamoDBv2:
{{- if $v.DynamoDBv2.PutItem }}
                          PutItem:
                              TableName: {{ $v.DynamoDBv2.PutItem.TableName }}
{{- end }}
{{- if $v.DynamoDBv2.RoleArn }}
                          RoleArn: {{ $v.DynamoDBv2.RoleArn }}
{{- end }}
{{- end }}
{{- if $v.Elasticsearch }}
                      Elasticsearch:
                          Endpoint: {{ $v.Elasticsearch.Endpoint }}
                          Id: {{ $v.Elasticsearch.Id }}
                          Index: {{ $v.Elasticsearch.Index }}
                          RoleArn: {{ $v.Elasticsearch.RoleArn }}
                          Type: {{ $v.Elasticsearch.Type }}
{{- end }}
{{- if $v.Firehose }}
                      Firehose:
                          DeliveryStreamName: {{ $v.Firehose.DeliveryStreamName }}
                          RoleArn: {{ $v.Firehose.RoleArn }}
{{- if $v.Firehose.Separator }}
                          Separator: {{ $v.Firehose.Separator }}
{{- end }}
{{- end }}
{{- if $v.Kinesis }}
                      Kinesis:
                          RoleArn: {{ $v.Kinesis.RoleArn }}
                          StreamName: {{ $v.Kinesis.StreamName }}
{{- if $v.Kinesis.PartitionKey }}
                          PartitionKey: {{ $v.Kinesis.PartitionKey }}
{{- end }}
{{- end }}
{{- if $v.Lambda }}
                      Lambda:
{{- if $v.Lambda.FunctionArn }}
                          FunctionArn: {{ $v.Lambda.FunctionArn }}
{{- end }}
{{- end }}
{{- if $v.Republish }}
                      Republish:
                          RoleArn: {{ $v.Republish.RoleArn }}
                          Topic: {{ $v.Republish.Topic }}
{{- end }}
{{- if $v.S3 }}
                      S3:
                          BucketName: {{ $v.S3.BucketName }}
                          Key: {{ $v.S3.Key }}
                          RoleArn: {{ $v.S3.RoleArn }}
{{- end }}
{{- if $v.Sns }}
                      Sns:
                          RoleArn: {{ $v.Sns.RoleArn }}
                          TargetArn: {{ $v.Sns.TargetArn }}
{{- if $v.Sns.MessageFormat }}
                          MessageFormat: {{ $v.Sns.MessageFormat }}
{{- end }}
{{- end }}
{{- if $v.Sqs }}
                      Sqs:
                          QueueUrl: {{ $v.Sqs.QueueUrl }}
                          RoleArn: {{ $v.Sqs.RoleArn }}
{{- if $v.Sqs.UseBase64 }}
                          UseBase64: {{ $v.Sqs.UseBase64 }}
{{- end }}
{{- end }}
{{- end }}
                RuleDisabled: {{ .TopicRulePayload.RuleDisabled }}
                Sql: {{ .TopicRulePayload.Sql }}
{{- if .TopicRulePayload.AwsIotSqlVersion }}
                AwsIotSqlVersion: {{ .TopicRulePayload.AwsIotSqlVersion }}
{{- end }}
{{- if .TopicRulePayload.Description }}
                Description: {{ .TopicRulePayload.Description }}
{{- end }}
{{- if .RuleName }}
            RuleName: {{ .RuleName }}
{{- end }}
{{- end }}
{{- define "AWS::KMS::Alias" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::KMS::Alias
        Properties:
            AliasName: {{ .AliasName }}
            TargetKeyId: {{ .TargetKeyId }}
{{- end }}
{{- define "AWS::KMS::Key" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::KMS::Key
        Properties:
            KeyPolicy:
                {{ yaml.MarshalIndent .KeyPolicy "                " "    " | fmt.Sprintf "%s" }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .EnableKeyRotation }}
            EnableKeyRotation: {{ .EnableKeyRotation }}
{{- end }}
{{- if .Enabled }}
            Enabled: {{ .Enabled }}
{{- end }}
{{- if .KeyUsage }}
            KeyUsage: {{ .KeyUsage }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Kinesis::Stream" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Kinesis::Stream
        Properties:
            ShardCount: {{ .ShardCount }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .RetentionPeriodHours }}
            RetentionPeriodHours: {{ .RetentionPeriodHours }}
{{- end }}
{{- if .StreamEncryption }}
            StreamEncryption:
                EncryptionType: {{ .StreamEncryption.EncryptionType }}
                KeyId: {{ .StreamEncryption.KeyId }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::KinesisAnalytics::Application" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::KinesisAnalytics::Application
        Properties:
            Inputs:
{{- range $_, $v := .Inputs }}
                -
                  InputSchema:
                      RecordColumns:
{{- range $_, $v := $v.InputSchema.RecordColumns }}
                          -
                            Name: {{ $v.Name }}
                            SqlType: {{ $v.SqlType }}
{{- if $v.Mapping }}
                            Mapping: {{ $v.Mapping }}
{{- end }}
{{- end }}
                      RecordFormat:
                          RecordFormatType: {{ $v.InputSchema.RecordFormat.RecordFormatType }}
{{- if $v.InputSchema.RecordFormat.MappingParameters }}
                          MappingParameters:
{{- if $v.InputSchema.RecordFormat.MappingParameters.CSVMappingParameters }}
                              CSVMappingParameters:
                                  RecordColumnDelimiter: {{ $v.InputSchema.RecordFormat.MappingParameters.CSVMappingParameters.RecordColumnDelimiter }}
                                  RecordRowDelimiter: {{ $v.InputSchema.RecordFormat.MappingParameters.CSVMappingParameters.RecordRowDelimiter }}
{{- end }}
{{- if $v.InputSchema.RecordFormat.MappingParameters.JSONMappingParameters }}
                              JSONMappingParameters:
                                  RecordRowPath: {{ $v.InputSchema.RecordFormat.MappingParameters.JSONMappingParameters.RecordRowPath }}
{{- end }}
{{- end }}
{{- if $v.InputSchema.RecordEncoding }}
                      RecordEncoding: {{ $v.InputSchema.RecordEncoding }}
{{- end }}
                  NamePrefix: {{ $v.NamePrefix }}
{{- if $v.InputParallelism }}
                  InputParallelism:
{{- if $v.InputParallelism.Count }}
                      Count: {{ $v.InputParallelism.Count }}
{{- end }}
{{- end }}
{{- if $v.InputProcessingConfiguration }}
                  InputProcessingConfiguration:
{{- if $v.InputProcessingConfiguration.InputLambdaProcessor }}
                      InputLambdaProcessor:
                          ResourceARN: {{ $v.InputProcessingConfiguration.InputLambdaProcessor.ResourceARN }}
                          RoleARN: {{ $v.InputProcessingConfiguration.InputLambdaProcessor.RoleARN }}
{{- end }}
{{- end }}
{{- if $v.KinesisFirehoseInput }}
                  KinesisFirehoseInput:
                      ResourceARN: {{ $v.KinesisFirehoseInput.ResourceARN }}
                      RoleARN: {{ $v.KinesisFirehoseInput.RoleARN }}
{{- end }}
{{- if $v.KinesisStreamsInput }}
                  KinesisStreamsInput:
                      ResourceARN: {{ $v.KinesisStreamsInput.ResourceARN }}
                      RoleARN: {{ $v.KinesisStreamsInput.RoleARN }}
{{- end }}
{{- end }}
{{- if .ApplicationCode }}
            ApplicationCode: {{ .ApplicationCode }}
{{- end }}
{{- if .ApplicationDescription }}
            ApplicationDescription: {{ .ApplicationDescription }}
{{- end }}
{{- if .ApplicationName }}
            ApplicationName: {{ .ApplicationName }}
{{- end }}
{{- end }}
{{- define "AWS::KinesisAnalytics::ApplicationOutput" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::KinesisAnalytics::ApplicationOutput
        Properties:
            ApplicationName: {{ .ApplicationName }}
            Output:
                DestinationSchema:
{{- if .Output.DestinationSchema.RecordFormatType }}
                    RecordFormatType: {{ .Output.DestinationSchema.RecordFormatType }}
{{- end }}
{{- if .Output.KinesisFirehoseOutput }}
                KinesisFirehoseOutput:
                    ResourceARN: {{ .Output.KinesisFirehoseOutput.ResourceARN }}
                    RoleARN: {{ .Output.KinesisFirehoseOutput.RoleARN }}
{{- end }}
{{- if .Output.KinesisStreamsOutput }}
                KinesisStreamsOutput:
                    ResourceARN: {{ .Output.KinesisStreamsOutput.ResourceARN }}
                    RoleARN: {{ .Output.KinesisStreamsOutput.RoleARN }}
{{- end }}
{{- if .Output.LambdaOutput }}
                LambdaOutput:
                    ResourceARN: {{ .Output.LambdaOutput.ResourceARN }}
                    RoleARN: {{ .Output.LambdaOutput.RoleARN }}
{{- end }}
{{- if .Output.Name }}
                Name: {{ .Output.Name }}
{{- end }}
{{- end }}
{{- define "AWS::KinesisAnalytics::ApplicationReferenceDataSource" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::KinesisAnalytics::ApplicationReferenceDataSource
        Properties:
            ApplicationName: {{ .ApplicationName }}
            ReferenceDataSource:
                ReferenceSchema:
                    RecordColumns:
{{- range $_, $v := .ReferenceDataSource.ReferenceSchema.RecordColumns }}
                        -
                          Name: {{ $v.Name }}
                          SqlType: {{ $v.SqlType }}
{{- if $v.Mapping }}
                          Mapping: {{ $v.Mapping }}
{{- end }}
{{- end }}
                    RecordFormat:
                        RecordFormatType: {{ .ReferenceDataSource.ReferenceSchema.RecordFormat.RecordFormatType }}
{{- if .ReferenceDataSource.ReferenceSchema.RecordFormat.MappingParameters }}
                        MappingParameters:
{{- if .ReferenceDataSource.ReferenceSchema.RecordFormat.MappingParameters.CSVMappingParameters }}
                            CSVMappingParameters:
                                RecordColumnDelimiter: {{ .ReferenceDataSource.ReferenceSchema.RecordFormat.MappingParameters.CSVMappingParameters.RecordColumnDelimiter }}
                                RecordRowDelimiter: {{ .ReferenceDataSource.ReferenceSchema.RecordFormat.MappingParameters.CSVMappingParameters.RecordRowDelimiter }}
{{- end }}
{{- if .ReferenceDataSource.ReferenceSchema.RecordFormat.MappingParameters.JSONMappingParameters }}
                            JSONMappingParameters:
                                RecordRowPath: {{ .ReferenceDataSource.ReferenceSchema.RecordFormat.MappingParameters.JSONMappingParameters.RecordRowPath }}
{{- end }}
{{- end }}
{{- if .ReferenceDataSource.ReferenceSchema.RecordEncoding }}
                    RecordEncoding: {{ .ReferenceDataSource.ReferenceSchema.RecordEncoding }}
{{- end }}
{{- if .ReferenceDataSource.S3ReferenceDataSource }}
                S3ReferenceDataSource:
                    BucketARN: {{ .ReferenceDataSource.S3ReferenceDataSource.BucketARN }}
                    FileKey: {{ .ReferenceDataSource.S3ReferenceDataSource.FileKey }}
                    ReferenceRoleARN: {{ .ReferenceDataSource.S3ReferenceDataSource.ReferenceRoleARN }}
{{- end }}
{{- if .ReferenceDataSource.TableName }}
                TableName: {{ .ReferenceDataSource.TableName }}
{{- end }}
{{- end }}
{{- define "AWS::KinesisFirehose::DeliveryStream" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::KinesisFirehose::DeliveryStream
{{- if .DeliveryStreamName | or .DeliveryStreamType | or .ElasticsearchDestinationConfiguration | or .ExtendedS3DestinationConfiguration | or .KinesisStreamSourceConfiguration | or .RedshiftDestinationConfiguration | or .S3DestinationConfiguration }}
        Properties:
{{- if .DeliveryStreamName }}
            DeliveryStreamName: {{ .DeliveryStreamName }}
{{- end }}
{{- if .DeliveryStreamType }}
            DeliveryStreamType: {{ .DeliveryStreamType }}
{{- end }}
{{- if .ElasticsearchDestinationConfiguration }}
            ElasticsearchDestinationConfiguration:
                BufferingHints:
                    IntervalInSeconds: {{ .ElasticsearchDestinationConfiguration.BufferingHints.IntervalInSeconds }}
                    SizeInMBs: {{ .ElasticsearchDestinationConfiguration.BufferingHints.SizeInMBs }}
                DomainARN: {{ .ElasticsearchDestinationConfiguration.DomainARN }}
                IndexName: {{ .ElasticsearchDestinationConfiguration.IndexName }}
                IndexRotationPeriod: {{ .ElasticsearchDestinationConfiguration.IndexRotationPeriod }}
                RetryOptions:
                    DurationInSeconds: {{ .ElasticsearchDestinationConfiguration.RetryOptions.DurationInSeconds }}
                RoleARN: {{ .ElasticsearchDestinationConfiguration.RoleARN }}
                S3BackupMode: {{ .ElasticsearchDestinationConfiguration.S3BackupMode }}
                S3Configuration:
                    BucketARN: {{ .ElasticsearchDestinationConfiguration.S3Configuration.BucketARN }}
                    BufferingHints:
                        IntervalInSeconds: {{ .ElasticsearchDestinationConfiguration.S3Configuration.BufferingHints.IntervalInSeconds }}
                        SizeInMBs: {{ .ElasticsearchDestinationConfiguration.S3Configuration.BufferingHints.SizeInMBs }}
                    CompressionFormat: {{ .ElasticsearchDestinationConfiguration.S3Configuration.CompressionFormat }}
                    RoleARN: {{ .ElasticsearchDestinationConfiguration.S3Configuration.RoleARN }}
{{- if .ElasticsearchDestinationConfiguration.S3Configuration.CloudWatchLoggingOptions }}
                    CloudWatchLoggingOptions:
{{- if .ElasticsearchDestinationConfiguration.S3Configuration.CloudWatchLoggingOptions.Enabled }}
                        Enabled: {{ .ElasticsearchDestinationConfiguration.S3Configuration.CloudWatchLoggingOptions.Enabled }}
{{- end }}
{{- if .ElasticsearchDestinationConfiguration.S3Configuration.CloudWatchLoggingOptions.LogGroupName }}
                        LogGroupName: {{ .ElasticsearchDestinationConfiguration.S3Configuration.CloudWatchLoggingOptions.LogGroupName }}
{{- end }}
{{- if .ElasticsearchDestinationConfiguration.S3Configuration.CloudWatchLoggingOptions.LogStreamName }}
                        LogStreamName: {{ .ElasticsearchDestinationConfiguration.S3Configuration.CloudWatchLoggingOptions.LogStreamName }}
{{- end }}
{{- end }}
{{- if .ElasticsearchDestinationConfiguration.S3Configuration.EncryptionConfiguration }}
                    EncryptionConfiguration:
{{- if .ElasticsearchDestinationConfiguration.S3Configuration.EncryptionConfiguration.KMSEncryptionConfig }}
                        KMSEncryptionConfig:
                            AWSKMSKeyARN: {{ .ElasticsearchDestinationConfiguration.S3Configuration.EncryptionConfiguration.KMSEncryptionConfig.AWSKMSKeyARN }}
{{- end }}
{{- if .ElasticsearchDestinationConfiguration.S3Configuration.EncryptionConfiguration.NoEncryptionConfig }}
                        NoEncryptionConfig: {{ .ElasticsearchDestinationConfiguration.S3Configuration.EncryptionConfiguration.NoEncryptionConfig }}
{{- end }}
{{- end }}
{{- if .ElasticsearchDestinationConfiguration.S3Configuration.Prefix }}
                    Prefix: {{ .ElasticsearchDestinationConfiguration.S3Configuration.Prefix }}
{{- end }}
                TypeName: {{ .ElasticsearchDestinationConfiguration.TypeName }}
{{- if .ElasticsearchDestinationConfiguration.CloudWatchLoggingOptions }}
                CloudWatchLoggingOptions:
{{- if .ElasticsearchDestinationConfiguration.CloudWatchLoggingOptions.Enabled }}
                    Enabled: {{ .ElasticsearchDestinationConfiguration.CloudWatchLoggingOptions.Enabled }}
{{- end }}
{{- if .ElasticsearchDestinationConfiguration.CloudWatchLoggingOptions.LogGroupName }}
                    LogGroupName: {{ .ElasticsearchDestinationConfiguration.CloudWatchLoggingOptions.LogGroupName }}
{{- end }}
{{- if .ElasticsearchDestinationConfiguration.CloudWatchLoggingOptions.LogStreamName }}
                    LogStreamName: {{ .ElasticsearchDestinationConfiguration.CloudWatchLoggingOptions.LogStreamName }}
{{- end }}
{{- end }}
{{- if .ElasticsearchDestinationConfiguration.ProcessingConfiguration }}
                ProcessingConfiguration:
                    Enabled: {{ .ElasticsearchDestinationConfiguration.ProcessingConfiguration.Enabled }}
                    Processors:
{{- range $_, $v := .ElasticsearchDestinationConfiguration.ProcessingConfiguration.Processors }}
                        -
                          Parameters:
{{- range $_, $v := $v.Parameters }}
                              -
                                ParameterName: {{ $v.ParameterName }}
                                ParameterValue: {{ $v.ParameterValue }}
{{- end }}
                          Type: {{ $v.Type }}
{{- end }}
{{- end }}
{{- end }}
{{- if .ExtendedS3DestinationConfiguration }}
            ExtendedS3DestinationConfiguration:
                BucketARN: {{ .ExtendedS3DestinationConfiguration.BucketARN }}
                BufferingHints:
                    IntervalInSeconds: {{ .ExtendedS3DestinationConfiguration.BufferingHints.IntervalInSeconds }}
                    SizeInMBs: {{ .ExtendedS3DestinationConfiguration.BufferingHints.SizeInMBs }}
                CompressionFormat: {{ .ExtendedS3DestinationConfiguration.CompressionFormat }}
                Prefix: {{ .ExtendedS3DestinationConfiguration.Prefix }}
                RoleARN: {{ .ExtendedS3DestinationConfiguration.RoleARN }}
{{- if .ExtendedS3DestinationConfiguration.CloudWatchLoggingOptions }}
                CloudWatchLoggingOptions:
{{- if .ExtendedS3DestinationConfiguration.CloudWatchLoggingOptions.Enabled }}
                    Enabled: {{ .ExtendedS3DestinationConfiguration.CloudWatchLoggingOptions.Enabled }}
{{- end }}
{{- if .ExtendedS3DestinationConfiguration.CloudWatchLoggingOptions.LogGroupName }}
                    LogGroupName: {{ .ExtendedS3DestinationConfiguration.CloudWatchLoggingOptions.LogGroupName }}
{{- end }}
{{- if .ExtendedS3DestinationConfiguration.CloudWatchLoggingOptions.LogStreamName }}
                    LogStreamName: {{ .ExtendedS3DestinationConfiguration.CloudWatchLoggingOptions.LogStreamName }}
{{- end }}
{{- end }}
{{- if .ExtendedS3DestinationConfiguration.EncryptionConfiguration }}
                EncryptionConfiguration:
{{- if .ExtendedS3DestinationConfiguration.EncryptionConfiguration.KMSEncryptionConfig }}
                    KMSEncryptionConfig:
                        AWSKMSKeyARN: {{ .ExtendedS3DestinationConfiguration.EncryptionConfiguration.KMSEncryptionConfig.AWSKMSKeyARN }}
{{- end }}
{{- if .ExtendedS3DestinationConfiguration.EncryptionConfiguration.NoEncryptionConfig }}
                    NoEncryptionConfig: {{ .ExtendedS3DestinationConfiguration.EncryptionConfiguration.NoEncryptionConfig }}
{{- end }}
{{- end }}
{{- if .ExtendedS3DestinationConfiguration.ProcessingConfiguration }}
                ProcessingConfiguration:
                    Enabled: {{ .ExtendedS3DestinationConfiguration.ProcessingConfiguration.Enabled }}
                    Processors:
{{- range $_, $v := .ExtendedS3DestinationConfiguration.ProcessingConfiguration.Processors }}
                        -
                          Parameters:
{{- range $_, $v := $v.Parameters }}
                              -
                                ParameterName: {{ $v.ParameterName }}
                                ParameterValue: {{ $v.ParameterValue }}
{{- end }}
                          Type: {{ $v.Type }}
{{- end }}
{{- end }}
{{- if .ExtendedS3DestinationConfiguration.S3BackupConfiguration }}
                S3BackupConfiguration:
                    BucketARN: {{ .ExtendedS3DestinationConfiguration.S3BackupConfiguration.BucketARN }}
                    BufferingHints:
                        IntervalInSeconds: {{ .ExtendedS3DestinationConfiguration.S3BackupConfiguration.BufferingHints.IntervalInSeconds }}
                        SizeInMBs: {{ .ExtendedS3DestinationConfiguration.S3BackupConfiguration.BufferingHints.SizeInMBs }}
                    CompressionFormat: {{ .ExtendedS3DestinationConfiguration.S3BackupConfiguration.CompressionFormat }}
                    RoleARN: {{ .ExtendedS3DestinationConfiguration.S3BackupConfiguration.RoleARN }}
{{- if .ExtendedS3DestinationConfiguration.S3BackupConfiguration.CloudWatchLoggingOptions }}
                    CloudWatchLoggingOptions:
{{- if .ExtendedS3DestinationConfiguration.S3BackupConfiguration.CloudWatchLoggingOptions.Enabled }}
                        Enabled: {{ .ExtendedS3DestinationConfiguration.S3BackupConfiguration.CloudWatchLoggingOptions.Enabled }}
{{- end }}
{{- if .ExtendedS3DestinationConfiguration.S3BackupConfiguration.CloudWatchLoggingOptions.LogGroupName }}
                        LogGroupName: {{ .ExtendedS3DestinationConfiguration.S3BackupConfiguration.CloudWatchLoggingOptions.LogGroupName }}
{{- end }}
{{- if .ExtendedS3DestinationConfiguration.S3BackupConfiguration.CloudWatchLoggingOptions.LogStreamName }}
                        LogStreamName: {{ .ExtendedS3DestinationConfiguration.S3BackupConfiguration.CloudWatchLoggingOptions.LogStreamName }}
{{- end }}
{{- end }}
{{- if .ExtendedS3DestinationConfiguration.S3BackupConfiguration.EncryptionConfiguration }}
                    EncryptionConfiguration:
{{- if .ExtendedS3DestinationConfiguration.S3BackupConfiguration.EncryptionConfiguration.KMSEncryptionConfig }}
                        KMSEncryptionConfig:
                            AWSKMSKeyARN: {{ .ExtendedS3DestinationConfiguration.S3BackupConfiguration.EncryptionConfiguration.KMSEncryptionConfig.AWSKMSKeyARN }}
{{- end }}
{{- if .ExtendedS3DestinationConfiguration.S3BackupConfiguration.EncryptionConfiguration.NoEncryptionConfig }}
                        NoEncryptionConfig: {{ .ExtendedS3DestinationConfiguration.S3BackupConfiguration.EncryptionConfiguration.NoEncryptionConfig }}
{{- end }}
{{- end }}
{{- if .ExtendedS3DestinationConfiguration.S3BackupConfiguration.Prefix }}
                    Prefix: {{ .ExtendedS3DestinationConfiguration.S3BackupConfiguration.Prefix }}
{{- end }}
{{- end }}
{{- if .ExtendedS3DestinationConfiguration.S3BackupMode }}
                S3BackupMode: {{ .ExtendedS3DestinationConfiguration.S3BackupMode }}
{{- end }}
{{- end }}
{{- if .KinesisStreamSourceConfiguration }}
            KinesisStreamSourceConfiguration:
                KinesisStreamARN: {{ .KinesisStreamSourceConfiguration.KinesisStreamARN }}
                RoleARN: {{ .KinesisStreamSourceConfiguration.RoleARN }}
{{- end }}
{{- if .RedshiftDestinationConfiguration }}
            RedshiftDestinationConfiguration:
                ClusterJDBCURL: {{ .RedshiftDestinationConfiguration.ClusterJDBCURL }}
                CopyCommand:
                    DataTableName: {{ .RedshiftDestinationConfiguration.CopyCommand.DataTableName }}
{{- if .RedshiftDestinationConfiguration.CopyCommand.CopyOptions }}
                    CopyOptions: {{ .RedshiftDestinationConfiguration.CopyCommand.CopyOptions }}
{{- end }}
{{- if .RedshiftDestinationConfiguration.CopyCommand.DataTableColumns }}
                    DataTableColumns: {{ .RedshiftDestinationConfiguration.CopyCommand.DataTableColumns }}
{{- end }}
                Password: {{ .RedshiftDestinationConfiguration.Password }}
                RoleARN: {{ .RedshiftDestinationConfiguration.RoleARN }}
                S3Configuration:
                    BucketARN: {{ .RedshiftDestinationConfiguration.S3Configuration.BucketARN }}
                    BufferingHints:
                        IntervalInSeconds: {{ .RedshiftDestinationConfiguration.S3Configuration.BufferingHints.IntervalInSeconds }}
                        SizeInMBs: {{ .RedshiftDestinationConfiguration.S3Configuration.BufferingHints.SizeInMBs }}
                    CompressionFormat: {{ .RedshiftDestinationConfiguration.S3Configuration.CompressionFormat }}
                    RoleARN: {{ .RedshiftDestinationConfiguration.S3Configuration.RoleARN }}
{{- if .RedshiftDestinationConfiguration.S3Configuration.CloudWatchLoggingOptions }}
                    CloudWatchLoggingOptions:
{{- if .RedshiftDestinationConfiguration.S3Configuration.CloudWatchLoggingOptions.Enabled }}
                        Enabled: {{ .RedshiftDestinationConfiguration.S3Configuration.CloudWatchLoggingOptions.Enabled }}
{{- end }}
{{- if .RedshiftDestinationConfiguration.S3Configuration.CloudWatchLoggingOptions.LogGroupName }}
                        LogGroupName: {{ .RedshiftDestinationConfiguration.S3Configuration.CloudWatchLoggingOptions.LogGroupName }}
{{- end }}
{{- if .RedshiftDestinationConfiguration.S3Configuration.CloudWatchLoggingOptions.LogStreamName }}
                        LogStreamName: {{ .RedshiftDestinationConfiguration.S3Configuration.CloudWatchLoggingOptions.LogStreamName }}
{{- end }}
{{- end }}
{{- if .RedshiftDestinationConfiguration.S3Configuration.EncryptionConfiguration }}
                    EncryptionConfiguration:
{{- if .RedshiftDestinationConfiguration.S3Configuration.EncryptionConfiguration.KMSEncryptionConfig }}
                        KMSEncryptionConfig:
                            AWSKMSKeyARN: {{ .RedshiftDestinationConfiguration.S3Configuration.EncryptionConfiguration.KMSEncryptionConfig.AWSKMSKeyARN }}
{{- end }}
{{- if .RedshiftDestinationConfiguration.S3Configuration.EncryptionConfiguration.NoEncryptionConfig }}
                        NoEncryptionConfig: {{ .RedshiftDestinationConfiguration.S3Configuration.EncryptionConfiguration.NoEncryptionConfig }}
{{- end }}
{{- end }}
{{- if .RedshiftDestinationConfiguration.S3Configuration.Prefix }}
                    Prefix: {{ .RedshiftDestinationConfiguration.S3Configuration.Prefix }}
{{- end }}
                Username: {{ .RedshiftDestinationConfiguration.Username }}
{{- if .RedshiftDestinationConfiguration.CloudWatchLoggingOptions }}
                CloudWatchLoggingOptions:
{{- if .RedshiftDestinationConfiguration.CloudWatchLoggingOptions.Enabled }}
                    Enabled: {{ .RedshiftDestinationConfiguration.CloudWatchLoggingOptions.Enabled }}
{{- end }}
{{- if .RedshiftDestinationConfiguration.CloudWatchLoggingOptions.LogGroupName }}
                    LogGroupName: {{ .RedshiftDestinationConfiguration.CloudWatchLoggingOptions.LogGroupName }}
{{- end }}
{{- if .RedshiftDestinationConfiguration.CloudWatchLoggingOptions.LogStreamName }}
                    LogStreamName: {{ .RedshiftDestinationConfiguration.CloudWatchLoggingOptions.LogStreamName }}
{{- end }}
{{- end }}
{{- if .RedshiftDestinationConfiguration.ProcessingConfiguration }}
                ProcessingConfiguration:
                    Enabled: {{ .RedshiftDestinationConfiguration.ProcessingConfiguration.Enabled }}
                    Processors:
{{- range $_, $v := .RedshiftDestinationConfiguration.ProcessingConfiguration.Processors }}
                        -
                          Parameters:
{{- range $_, $v := $v.Parameters }}
                              -
                                ParameterName: {{ $v.ParameterName }}
                                ParameterValue: {{ $v.ParameterValue }}
{{- end }}
                          Type: {{ $v.Type }}
{{- end }}
{{- end }}
{{- end }}
{{- if .S3DestinationConfiguration }}
            S3DestinationConfiguration:
                BucketARN: {{ .S3DestinationConfiguration.BucketARN }}
                BufferingHints:
                    IntervalInSeconds: {{ .S3DestinationConfiguration.BufferingHints.IntervalInSeconds }}
                    SizeInMBs: {{ .S3DestinationConfiguration.BufferingHints.SizeInMBs }}
                CompressionFormat: {{ .S3DestinationConfiguration.CompressionFormat }}
                RoleARN: {{ .S3DestinationConfiguration.RoleARN }}
{{- if .S3DestinationConfiguration.CloudWatchLoggingOptions }}
                CloudWatchLoggingOptions:
{{- if .S3DestinationConfiguration.CloudWatchLoggingOptions.Enabled }}
                    Enabled: {{ .S3DestinationConfiguration.CloudWatchLoggingOptions.Enabled }}
{{- end }}
{{- if .S3DestinationConfiguration.CloudWatchLoggingOptions.LogGroupName }}
                    LogGroupName: {{ .S3DestinationConfiguration.CloudWatchLoggingOptions.LogGroupName }}
{{- end }}
{{- if .S3DestinationConfiguration.CloudWatchLoggingOptions.LogStreamName }}
                    LogStreamName: {{ .S3DestinationConfiguration.CloudWatchLoggingOptions.LogStreamName }}
{{- end }}
{{- end }}
{{- if .S3DestinationConfiguration.EncryptionConfiguration }}
                EncryptionConfiguration:
{{- if .S3DestinationConfiguration.EncryptionConfiguration.KMSEncryptionConfig }}
                    KMSEncryptionConfig:
                        AWSKMSKeyARN: {{ .S3DestinationConfiguration.EncryptionConfiguration.KMSEncryptionConfig.AWSKMSKeyARN }}
{{- end }}
{{- if .S3DestinationConfiguration.EncryptionConfiguration.NoEncryptionConfig }}
                    NoEncryptionConfig: {{ .S3DestinationConfiguration.EncryptionConfiguration.NoEncryptionConfig }}
{{- end }}
{{- end }}
{{- if .S3DestinationConfiguration.Prefix }}
                Prefix: {{ .S3DestinationConfiguration.Prefix }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Lambda::Alias" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Lambda::Alias
        Properties:
            FunctionName: {{ .FunctionName }}
            FunctionVersion: {{ .FunctionVersion }}
            Name: {{ .Name }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .RoutingConfig }}
            RoutingConfig:
                AdditionalVersionWeights:
{{- range $_, $v := .RoutingConfig.AdditionalVersionWeights }}
                    -
                      FunctionVersion: {{ $v.FunctionVersion }}
                      FunctionWeight: {{ $v.FunctionWeight }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Lambda::EventSourceMapping" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Lambda::EventSourceMapping
        Properties:
            EventSourceArn: {{ .EventSourceArn }}
            FunctionName: {{ .FunctionName }}
            StartingPosition: {{ .StartingPosition }}
{{- if .BatchSize }}
            BatchSize: {{ .BatchSize }}
{{- end }}
{{- if .Enabled }}
            Enabled: {{ .Enabled }}
{{- end }}
{{- end }}
{{- define "AWS::Lambda::Function" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Lambda::Function
        Properties:
            Code:{{ if istype .Code "string" }} {{ .Code }}{{ else }}
{{- if .Code.S3Bucket }}
                S3Bucket: {{ .Code.S3Bucket }}
{{- end }}
{{- if .Code.S3Key }}
                S3Key: {{ .Code.S3Key }}
{{- end }}
{{- if .Code.S3ObjectVersion }}
                S3ObjectVersion: {{ .Code.S3ObjectVersion }}
{{- end }}
{{- if .Code.ZipFile }}
                ZipFile: {{ .Code.ZipFile }}
{{- end }}
{{- end }}
            Handler: {{ .Handler }}
            Role: {{ .Role }}
            Runtime: {{ .Runtime }}
{{- if .DeadLetterConfig }}
            DeadLetterConfig:
{{- if .DeadLetterConfig.TargetArn }}
                TargetArn: {{ .DeadLetterConfig.TargetArn }}
{{- end }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .Environment }}
            Environment:
{{- if .Environment.Variables }}
                Variables:
{{- range $k, $v := .Environment.Variables }}
                    {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if .FunctionName }}
            FunctionName: {{ .FunctionName }}
{{- end }}
{{- if .KmsKeyArn }}
            KmsKeyArn: {{ .KmsKeyArn }}
{{- end }}
{{- if .MemorySize }}
            MemorySize: {{ .MemorySize }}
{{- end }}
{{- if .ReservedConcurrentExecutions }}
            ReservedConcurrentExecutions: {{ .ReservedConcurrentExecutions }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .Timeout }}
            Timeout: {{ .Timeout }}
{{- end }}
{{- if .TracingConfig }}
            TracingConfig:
{{- if .TracingConfig.Mode }}
                Mode: {{ .TracingConfig.Mode }}
{{- end }}
{{- end }}
{{- if .VpcConfig }}
            VpcConfig:
                SecurityGroupIds:
{{- range $_, $v := .VpcConfig.SecurityGroupIds }}
                    - {{ $v }}
{{- end }}
                SubnetIds:
{{- range $_, $v := .VpcConfig.SubnetIds }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Lambda::Permission" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Lambda::Permission
        Properties:
            Action: {{ .Action }}
            FunctionName: {{ .FunctionName }}
            Principal: {{ .Principal }}
{{- if .EventSourceToken }}
            EventSourceToken: {{ .EventSourceToken }}
{{- end }}
{{- if .SourceAccount }}
            SourceAccount: {{ .SourceAccount }}
{{- end }}
{{- if .SourceArn }}
            SourceArn: {{ .SourceArn }}
{{- end }}
{{- end }}
{{- define "AWS::Lambda::Version" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Lambda::Version
        Properties:
            FunctionName: {{ .FunctionName }}
{{- if .CodeSha256 }}
            CodeSha256: {{ .CodeSha256 }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- end }}
{{- define "AWS::Logs::Destination" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Logs::Destination
        Properties:
            DestinationName: {{ .DestinationName }}
            DestinationPolicy: {{ .DestinationPolicy }}
            RoleArn: {{ .RoleArn }}
            TargetArn: {{ .TargetArn }}
{{- end }}
{{- define "AWS::Logs::LogGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Logs::LogGroup
{{- if .LogGroupName | or .RetentionInDays }}
        Properties:
{{- if .LogGroupName }}
            LogGroupName: {{ .LogGroupName }}
{{- end }}
{{- if .RetentionInDays }}
            RetentionInDays: {{ .RetentionInDays }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Logs::LogStream" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Logs::LogStream
        Properties:
            LogGroupName: {{ .LogGroupName }}
{{- if .LogStreamName }}
            LogStreamName: {{ .LogStreamName }}
{{- end }}
{{- end }}
{{- define "AWS::Logs::MetricFilter" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Logs::MetricFilter
        Properties:
            FilterPattern: {{ .FilterPattern }}
            LogGroupName: {{ .LogGroupName }}
            MetricTransformations:
{{- range $_, $v := .MetricTransformations }}
                -
                  MetricName: {{ $v.MetricName }}
                  MetricNamespace: {{ $v.MetricNamespace }}
                  MetricValue: {{ $v.MetricValue }}
{{- end }}
{{- end }}
{{- define "AWS::Logs::SubscriptionFilter" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Logs::SubscriptionFilter
        Properties:
            DestinationArn: {{ .DestinationArn }}
            FilterPattern: {{ .FilterPattern }}
            LogGroupName: {{ .LogGroupName }}
{{- if .RoleArn }}
            RoleArn: {{ .RoleArn }}
{{- end }}
{{- end }}
{{- define "AWS::OpsWorks::App" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::OpsWorks::App
        Properties:
            Name: {{ .Name }}
            StackId: {{ .StackId }}
            Type: {{ .Type }}
{{- if .AppSource }}
            AppSource:
{{- if .AppSource.Password }}
                Password: {{ .AppSource.Password }}
{{- end }}
{{- if .AppSource.Revision }}
                Revision: {{ .AppSource.Revision }}
{{- end }}
{{- if .AppSource.SshKey }}
                SshKey: {{ .AppSource.SshKey }}
{{- end }}
{{- if .AppSource.Type }}
                Type: {{ .AppSource.Type }}
{{- end }}
{{- if .AppSource.Url }}
                Url: {{ .AppSource.Url }}
{{- end }}
{{- if .AppSource.Username }}
                Username: {{ .AppSource.Username }}
{{- end }}
{{- end }}
{{- if .Attributes }}
            Attributes:
{{- range $k, $v := .Attributes }}
                {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .DataSources }}
            DataSources:
{{- range $_, $v := .DataSources }}
                -
{{- if $v.Arn }}
                  Arn: {{ $v.Arn }}
{{- end }}
{{- if $v.DatabaseName }}
                  DatabaseName: {{ $v.DatabaseName }}
{{- end }}
{{- if $v.Type }}
                  Type: {{ $v.Type }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .Domains }}
            Domains:
{{- range $_, $v := .Domains }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .EnableSsl }}
            EnableSsl: {{ .EnableSsl }}
{{- end }}
{{- if .Environment }}
            Environment:
{{- range $_, $v := .Environment }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- if $v.Secure }}
                  Secure: {{ $v.Secure }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Shortname }}
            Shortname: {{ .Shortname }}
{{- end }}
{{- if .SslConfiguration }}
            SslConfiguration:
{{- if .SslConfiguration.Certificate }}
                Certificate: {{ .SslConfiguration.Certificate }}
{{- end }}
{{- if .SslConfiguration.Chain }}
                Chain: {{ .SslConfiguration.Chain }}
{{- end }}
{{- if .SslConfiguration.PrivateKey }}
                PrivateKey: {{ .SslConfiguration.PrivateKey }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::OpsWorks::ElasticLoadBalancerAttachment" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::OpsWorks::ElasticLoadBalancerAttachment
        Properties:
            ElasticLoadBalancerName: {{ .ElasticLoadBalancerName }}
            LayerId: {{ .LayerId }}
{{- end }}
{{- define "AWS::OpsWorks::Instance" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::OpsWorks::Instance
        Properties:
            InstanceType: {{ .InstanceType }}
            LayerIds:
{{- range $_, $v := .LayerIds }}
                - {{ $v }}
{{- end }}
            StackId: {{ .StackId }}
{{- if .AgentVersion }}
            AgentVersion: {{ .AgentVersion }}
{{- end }}
{{- if .AmiId }}
            AmiId: {{ .AmiId }}
{{- end }}
{{- if .Architecture }}
            Architecture: {{ .Architecture }}
{{- end }}
{{- if .AutoScalingType }}
            AutoScalingType: {{ .AutoScalingType }}
{{- end }}
{{- if .AvailabilityZone }}
            AvailabilityZone: {{ .AvailabilityZone }}
{{- end }}
{{- if .BlockDeviceMappings }}
            BlockDeviceMappings:
{{- range $_, $v := .BlockDeviceMappings }}
                -
{{- if $v.DeviceName }}
                  DeviceName: {{ $v.DeviceName }}
{{- end }}
{{- if $v.Ebs }}
                  Ebs:
{{- if $v.Ebs.DeleteOnTermination }}
                      DeleteOnTermination: {{ $v.Ebs.DeleteOnTermination }}
{{- end }}
{{- if $v.Ebs.Iops }}
                      Iops: {{ $v.Ebs.Iops }}
{{- end }}
{{- if $v.Ebs.SnapshotId }}
                      SnapshotId: {{ $v.Ebs.SnapshotId }}
{{- end }}
{{- if $v.Ebs.VolumeSize }}
                      VolumeSize: {{ $v.Ebs.VolumeSize }}
{{- end }}
{{- if $v.Ebs.VolumeType }}
                      VolumeType: {{ $v.Ebs.VolumeType }}
{{- end }}
{{- end }}
{{- if $v.NoDevice }}
                  NoDevice: {{ $v.NoDevice }}
{{- end }}
{{- if $v.VirtualName }}
                  VirtualName: {{ $v.VirtualName }}
{{- end }}
{{- end }}
{{- end }}
{{- if .EbsOptimized }}
            EbsOptimized: {{ .EbsOptimized }}
{{- end }}
{{- if .ElasticIps }}
            ElasticIps:
{{- range $_, $v := .ElasticIps }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Hostname }}
            Hostname: {{ .Hostname }}
{{- end }}
{{- if .InstallUpdatesOnBoot }}
            InstallUpdatesOnBoot: {{ .InstallUpdatesOnBoot }}
{{- end }}
{{- if .Os }}
            Os: {{ .Os }}
{{- end }}
{{- if .RootDeviceType }}
            RootDeviceType: {{ .RootDeviceType }}
{{- end }}
{{- if .SshKeyName }}
            SshKeyName: {{ .SshKeyName }}
{{- end }}
{{- if .SubnetId }}
            SubnetId: {{ .SubnetId }}
{{- end }}
{{- if .Tenancy }}
            Tenancy: {{ .Tenancy }}
{{- end }}
{{- if .TimeBasedAutoScaling }}
            TimeBasedAutoScaling:
{{- if .TimeBasedAutoScaling.Friday }}
                Friday:
{{- range $k, $v := .TimeBasedAutoScaling.Friday }}
                    {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .TimeBasedAutoScaling.Monday }}
                Monday:
{{- range $k, $v := .TimeBasedAutoScaling.Monday }}
                    {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .TimeBasedAutoScaling.Saturday }}
                Saturday:
{{- range $k, $v := .TimeBasedAutoScaling.Saturday }}
                    {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .TimeBasedAutoScaling.Sunday }}
                Sunday:
{{- range $k, $v := .TimeBasedAutoScaling.Sunday }}
                    {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .TimeBasedAutoScaling.Thursday }}
                Thursday:
{{- range $k, $v := .TimeBasedAutoScaling.Thursday }}
                    {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .TimeBasedAutoScaling.Tuesday }}
                Tuesday:
{{- range $k, $v := .TimeBasedAutoScaling.Tuesday }}
                    {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .TimeBasedAutoScaling.Wednesday }}
                Wednesday:
{{- range $k, $v := .TimeBasedAutoScaling.Wednesday }}
                    {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if .VirtualizationType }}
            VirtualizationType: {{ .VirtualizationType }}
{{- end }}
{{- if .Volumes }}
            Volumes:
{{- range $_, $v := .Volumes }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::OpsWorks::Layer" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::OpsWorks::Layer
        Properties:
            AutoAssignElasticIps: {{ .AutoAssignElasticIps }}
            AutoAssignPublicIps: {{ .AutoAssignPublicIps }}
            EnableAutoHealing: {{ .EnableAutoHealing }}
            Name: {{ .Name }}
            Shortname: {{ .Shortname }}
            StackId: {{ .StackId }}
            Type: {{ .Type }}
{{- if .Attributes }}
            Attributes:
{{- range $k, $v := .Attributes }}
                {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .CustomInstanceProfileArn }}
            CustomInstanceProfileArn: {{ .CustomInstanceProfileArn }}
{{- end }}
{{- if .CustomJson }}
            CustomJson:
                {{ yaml.MarshalIndent .CustomJson "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .CustomRecipes }}
            CustomRecipes:
{{- if .CustomRecipes.Configure }}
                Configure:
{{- range $_, $v := .CustomRecipes.Configure }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- if .CustomRecipes.Deploy }}
                Deploy:
{{- range $_, $v := .CustomRecipes.Deploy }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- if .CustomRecipes.Setup }}
                Setup:
{{- range $_, $v := .CustomRecipes.Setup }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- if .CustomRecipes.Shutdown }}
                Shutdown:
{{- range $_, $v := .CustomRecipes.Shutdown }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- if .CustomRecipes.Undeploy }}
                Undeploy:
{{- range $_, $v := .CustomRecipes.Undeploy }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if .CustomSecurityGroupIds }}
            CustomSecurityGroupIds:
{{- range $_, $v := .CustomSecurityGroupIds }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .InstallUpdatesOnBoot }}
            InstallUpdatesOnBoot: {{ .InstallUpdatesOnBoot }}
{{- end }}
{{- if .LifecycleEventConfiguration }}
            LifecycleEventConfiguration:
{{- if .LifecycleEventConfiguration.ShutdownEventConfiguration }}
                ShutdownEventConfiguration:
{{- if .LifecycleEventConfiguration.ShutdownEventConfiguration.DelayUntilElbConnectionsDrained }}
                    DelayUntilElbConnectionsDrained: {{ .LifecycleEventConfiguration.ShutdownEventConfiguration.DelayUntilElbConnectionsDrained }}
{{- end }}
{{- if .LifecycleEventConfiguration.ShutdownEventConfiguration.ExecutionTimeout }}
                    ExecutionTimeout: {{ .LifecycleEventConfiguration.ShutdownEventConfiguration.ExecutionTimeout }}
{{- end }}
{{- end }}
{{- end }}
{{- if .LoadBasedAutoScaling }}
            LoadBasedAutoScaling:
{{- if .LoadBasedAutoScaling.DownScaling }}
                DownScaling:
{{- if .LoadBasedAutoScaling.DownScaling.CpuThreshold }}
                    CpuThreshold: {{ .LoadBasedAutoScaling.DownScaling.CpuThreshold }}
{{- end }}
{{- if .LoadBasedAutoScaling.DownScaling.IgnoreMetricsTime }}
                    IgnoreMetricsTime: {{ .LoadBasedAutoScaling.DownScaling.IgnoreMetricsTime }}
{{- end }}
{{- if .LoadBasedAutoScaling.DownScaling.InstanceCount }}
                    InstanceCount: {{ .LoadBasedAutoScaling.DownScaling.InstanceCount }}
{{- end }}
{{- if .LoadBasedAutoScaling.DownScaling.LoadThreshold }}
                    LoadThreshold: {{ .LoadBasedAutoScaling.DownScaling.LoadThreshold }}
{{- end }}
{{- if .LoadBasedAutoScaling.DownScaling.MemoryThreshold }}
                    MemoryThreshold: {{ .LoadBasedAutoScaling.DownScaling.MemoryThreshold }}
{{- end }}
{{- if .LoadBasedAutoScaling.DownScaling.ThresholdsWaitTime }}
                    ThresholdsWaitTime: {{ .LoadBasedAutoScaling.DownScaling.ThresholdsWaitTime }}
{{- end }}
{{- end }}
{{- if .LoadBasedAutoScaling.Enable }}
                Enable: {{ .LoadBasedAutoScaling.Enable }}
{{- end }}
{{- if .LoadBasedAutoScaling.UpScaling }}
                UpScaling:
{{- if .LoadBasedAutoScaling.UpScaling.CpuThreshold }}
                    CpuThreshold: {{ .LoadBasedAutoScaling.UpScaling.CpuThreshold }}
{{- end }}
{{- if .LoadBasedAutoScaling.UpScaling.IgnoreMetricsTime }}
                    IgnoreMetricsTime: {{ .LoadBasedAutoScaling.UpScaling.IgnoreMetricsTime }}
{{- end }}
{{- if .LoadBasedAutoScaling.UpScaling.InstanceCount }}
                    InstanceCount: {{ .LoadBasedAutoScaling.UpScaling.InstanceCount }}
{{- end }}
{{- if .LoadBasedAutoScaling.UpScaling.LoadThreshold }}
                    LoadThreshold: {{ .LoadBasedAutoScaling.UpScaling.LoadThreshold }}
{{- end }}
{{- if .LoadBasedAutoScaling.UpScaling.MemoryThreshold }}
                    MemoryThreshold: {{ .LoadBasedAutoScaling.UpScaling.MemoryThreshold }}
{{- end }}
{{- if .LoadBasedAutoScaling.UpScaling.ThresholdsWaitTime }}
                    ThresholdsWaitTime: {{ .LoadBasedAutoScaling.UpScaling.ThresholdsWaitTime }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Packages }}
            Packages:
{{- range $_, $v := .Packages }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .UseEbsOptimizedInstances }}
            UseEbsOptimizedInstances: {{ .UseEbsOptimizedInstances }}
{{- end }}
{{- if .VolumeConfigurations }}
            VolumeConfigurations:
{{- range $_, $v := .VolumeConfigurations }}
                -
{{- if $v.Iops }}
                  Iops: {{ $v.Iops }}
{{- end }}
{{- if $v.MountPoint }}
                  MountPoint: {{ $v.MountPoint }}
{{- end }}
{{- if $v.NumberOfDisks }}
                  NumberOfDisks: {{ $v.NumberOfDisks }}
{{- end }}
{{- if $v.RaidLevel }}
                  RaidLevel: {{ $v.RaidLevel }}
{{- end }}
{{- if $v.Size }}
                  Size: {{ $v.Size }}
{{- end }}
{{- if $v.VolumeType }}
                  VolumeType: {{ $v.VolumeType }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::OpsWorks::Stack" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::OpsWorks::Stack
        Properties:
            DefaultInstanceProfileArn: {{ .DefaultInstanceProfileArn }}
            Name: {{ .Name }}
            ServiceRoleArn: {{ .ServiceRoleArn }}
{{- if .AgentVersion }}
            AgentVersion: {{ .AgentVersion }}
{{- end }}
{{- if .Attributes }}
            Attributes:
{{- range $k, $v := .Attributes }}
                {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .ChefConfiguration }}
            ChefConfiguration:
{{- if .ChefConfiguration.BerkshelfVersion }}
                BerkshelfVersion: {{ .ChefConfiguration.BerkshelfVersion }}
{{- end }}
{{- if .ChefConfiguration.ManageBerkshelf }}
                ManageBerkshelf: {{ .ChefConfiguration.ManageBerkshelf }}
{{- end }}
{{- end }}
{{- if .CloneAppIds }}
            CloneAppIds:
{{- range $_, $v := .CloneAppIds }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .ClonePermissions }}
            ClonePermissions: {{ .ClonePermissions }}
{{- end }}
{{- if .ConfigurationManager }}
            ConfigurationManager:
{{- if .ConfigurationManager.Name }}
                Name: {{ .ConfigurationManager.Name }}
{{- end }}
{{- if .ConfigurationManager.Version }}
                Version: {{ .ConfigurationManager.Version }}
{{- end }}
{{- end }}
{{- if .CustomCookbooksSource }}
            CustomCookbooksSource:
{{- if .CustomCookbooksSource.Password }}
                Password: {{ .CustomCookbooksSource.Password }}
{{- end }}
{{- if .CustomCookbooksSource.Revision }}
                Revision: {{ .CustomCookbooksSource.Revision }}
{{- end }}
{{- if .CustomCookbooksSource.SshKey }}
                SshKey: {{ .CustomCookbooksSource.SshKey }}
{{- end }}
{{- if .CustomCookbooksSource.Type }}
                Type: {{ .CustomCookbooksSource.Type }}
{{- end }}
{{- if .CustomCookbooksSource.Url }}
                Url: {{ .CustomCookbooksSource.Url }}
{{- end }}
{{- if .CustomCookbooksSource.Username }}
                Username: {{ .CustomCookbooksSource.Username }}
{{- end }}
{{- end }}
{{- if .CustomJson }}
            CustomJson:
                {{ yaml.MarshalIndent .CustomJson "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .DefaultAvailabilityZone }}
            DefaultAvailabilityZone: {{ .DefaultAvailabilityZone }}
{{- end }}
{{- if .DefaultOs }}
            DefaultOs: {{ .DefaultOs }}
{{- end }}
{{- if .DefaultRootDeviceType }}
            DefaultRootDeviceType: {{ .DefaultRootDeviceType }}
{{- end }}
{{- if .DefaultSshKeyName }}
            DefaultSshKeyName: {{ .DefaultSshKeyName }}
{{- end }}
{{- if .DefaultSubnetId }}
            DefaultSubnetId: {{ .DefaultSubnetId }}
{{- end }}
{{- if .EcsClusterArn }}
            EcsClusterArn: {{ .EcsClusterArn }}
{{- end }}
{{- if .ElasticIps }}
            ElasticIps:
{{- range $_, $v := .ElasticIps }}
                -
                  Ip: {{ $v.Ip }}
{{- if $v.Name }}
                  Name: {{ $v.Name }}
{{- end }}
{{- end }}
{{- end }}
{{- if .HostnameTheme }}
            HostnameTheme: {{ .HostnameTheme }}
{{- end }}
{{- if .RdsDbInstances }}
            RdsDbInstances:
{{- range $_, $v := .RdsDbInstances }}
                -
                  DbPassword: {{ $v.DbPassword }}
                  DbUser: {{ $v.DbUser }}
                  RdsDbInstanceArn: {{ $v.RdsDbInstanceArn }}
{{- end }}
{{- end }}
{{- if .SourceStackId }}
            SourceStackId: {{ .SourceStackId }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .UseCustomCookbooks }}
            UseCustomCookbooks: {{ .UseCustomCookbooks }}
{{- end }}
{{- if .UseOpsworksSecurityGroups }}
            UseOpsworksSecurityGroups: {{ .UseOpsworksSecurityGroups }}
{{- end }}
{{- if .VpcId }}
            VpcId: {{ .VpcId }}
{{- end }}
{{- end }}
{{- define "AWS::OpsWorks::UserProfile" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::OpsWorks::UserProfile
        Properties:
            IamUserArn: {{ .IamUserArn }}
{{- if .AllowSelfManagement }}
            AllowSelfManagement: {{ .AllowSelfManagement }}
{{- end }}
{{- if .SshPublicKey }}
            SshPublicKey: {{ .SshPublicKey }}
{{- end }}
{{- if .SshUsername }}
            SshUsername: {{ .SshUsername }}
{{- end }}
{{- end }}
{{- define "AWS::OpsWorks::Volume" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::OpsWorks::Volume
        Properties:
            Ec2VolumeId: {{ .Ec2VolumeId }}
            StackId: {{ .StackId }}
{{- if .MountPoint }}
            MountPoint: {{ .MountPoint }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- end }}
{{- define "AWS::RDS::DBCluster" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::RDS::DBCluster
        Properties:
            Engine: {{ .Engine }}
{{- if .AvailabilityZones }}
            AvailabilityZones:
{{- range $_, $v := .AvailabilityZones }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .BackupRetentionPeriod }}
            BackupRetentionPeriod: {{ .BackupRetentionPeriod }}
{{- end }}
{{- if .DBClusterIdentifier }}
            DBClusterIdentifier: {{ .DBClusterIdentifier }}
{{- end }}
{{- if .DBClusterParameterGroupName }}
            DBClusterParameterGroupName: {{ .DBClusterParameterGroupName }}
{{- end }}
{{- if .DBSubnetGroupName }}
            DBSubnetGroupName: {{ .DBSubnetGroupName }}
{{- end }}
{{- if .DatabaseName }}
            DatabaseName: {{ .DatabaseName }}
{{- end }}
{{- if .EngineVersion }}
            EngineVersion: {{ .EngineVersion }}
{{- end }}
{{- if .KmsKeyId }}
            KmsKeyId: {{ .KmsKeyId }}
{{- end }}
{{- if .MasterUserPassword }}
            MasterUserPassword: {{ .MasterUserPassword }}
{{- end }}
{{- if .MasterUsername }}
            MasterUsername: {{ .MasterUsername }}
{{- end }}
{{- if .Port }}
            Port: {{ .Port }}
{{- end }}
{{- if .PreferredBackupWindow }}
            PreferredBackupWindow: {{ .PreferredBackupWindow }}
{{- end }}
{{- if .PreferredMaintenanceWindow }}
            PreferredMaintenanceWindow: {{ .PreferredMaintenanceWindow }}
{{- end }}
{{- if .ReplicationSourceIdentifier }}
            ReplicationSourceIdentifier: {{ .ReplicationSourceIdentifier }}
{{- end }}
{{- if .SnapshotIdentifier }}
            SnapshotIdentifier: {{ .SnapshotIdentifier }}
{{- end }}
{{- if .StorageEncrypted }}
            StorageEncrypted: {{ .StorageEncrypted }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .VpcSecurityGroupIds }}
            VpcSecurityGroupIds:
{{- range $_, $v := .VpcSecurityGroupIds }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::RDS::DBClusterParameterGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::RDS::DBClusterParameterGroup
        Properties:
            Description: {{ .Description }}
            Family: {{ .Family }}
            Parameters:
                {{ yaml.MarshalIndent .Parameters "                " "    " | fmt.Sprintf "%s" }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::RDS::DBInstance" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::RDS::DBInstance
        Properties:
            DBInstanceClass: {{ .DBInstanceClass }}
{{- if .AllocatedStorage }}
            AllocatedStorage: {{ .AllocatedStorage }}
{{- end }}
{{- if .AllowMajorVersionUpgrade }}
            AllowMajorVersionUpgrade: {{ .AllowMajorVersionUpgrade }}
{{- end }}
{{- if .AutoMinorVersionUpgrade }}
            AutoMinorVersionUpgrade: {{ .AutoMinorVersionUpgrade }}
{{- end }}
{{- if .AvailabilityZone }}
            AvailabilityZone: {{ .AvailabilityZone }}
{{- end }}
{{- if .BackupRetentionPeriod }}
            BackupRetentionPeriod: {{ .BackupRetentionPeriod }}
{{- end }}
{{- if .CharacterSetName }}
            CharacterSetName: {{ .CharacterSetName }}
{{- end }}
{{- if .CopyTagsToSnapshot }}
            CopyTagsToSnapshot: {{ .CopyTagsToSnapshot }}
{{- end }}
{{- if .DBClusterIdentifier }}
            DBClusterIdentifier: {{ .DBClusterIdentifier }}
{{- end }}
{{- if .DBInstanceIdentifier }}
            DBInstanceIdentifier: {{ .DBInstanceIdentifier }}
{{- end }}
{{- if .DBName }}
            DBName: {{ .DBName }}
{{- end }}
{{- if .DBParameterGroupName }}
            DBParameterGroupName: {{ .DBParameterGroupName }}
{{- end }}
{{- if .DBSecurityGroups }}
            DBSecurityGroups:
{{- range $_, $v := .DBSecurityGroups }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .DBSnapshotIdentifier }}
            DBSnapshotIdentifier: {{ .DBSnapshotIdentifier }}
{{- end }}
{{- if .DBSubnetGroupName }}
            DBSubnetGroupName: {{ .DBSubnetGroupName }}
{{- end }}
{{- if .Domain }}
            Domain: {{ .Domain }}
{{- end }}
{{- if .DomainIAMRoleName }}
            DomainIAMRoleName: {{ .DomainIAMRoleName }}
{{- end }}
{{- if .Engine }}
            Engine: {{ .Engine }}
{{- end }}
{{- if .EngineVersion }}
            EngineVersion: {{ .EngineVersion }}
{{- end }}
{{- if .Iops }}
            Iops: {{ .Iops }}
{{- end }}
{{- if .KmsKeyId }}
            KmsKeyId: {{ .KmsKeyId }}
{{- end }}
{{- if .LicenseModel }}
            LicenseModel: {{ .LicenseModel }}
{{- end }}
{{- if .MasterUserPassword }}
            MasterUserPassword: {{ .MasterUserPassword }}
{{- end }}
{{- if .MasterUsername }}
            MasterUsername: {{ .MasterUsername }}
{{- end }}
{{- if .MonitoringInterval }}
            MonitoringInterval: {{ .MonitoringInterval }}
{{- end }}
{{- if .MonitoringRoleArn }}
            MonitoringRoleArn: {{ .MonitoringRoleArn }}
{{- end }}
{{- if .MultiAZ }}
            MultiAZ: {{ .MultiAZ }}
{{- end }}
{{- if .OptionGroupName }}
            OptionGroupName: {{ .OptionGroupName }}
{{- end }}
{{- if .Port }}
            Port: {{ .Port }}
{{- end }}
{{- if .PreferredBackupWindow }}
            PreferredBackupWindow: {{ .PreferredBackupWindow }}
{{- end }}
{{- if .PreferredMaintenanceWindow }}
            PreferredMaintenanceWindow: {{ .PreferredMaintenanceWindow }}
{{- end }}
{{- if .PubliclyAccessible }}
            PubliclyAccessible: {{ .PubliclyAccessible }}
{{- end }}
{{- if .SourceDBInstanceIdentifier }}
            SourceDBInstanceIdentifier: {{ .SourceDBInstanceIdentifier }}
{{- end }}
{{- if .SourceRegion }}
            SourceRegion: {{ .SourceRegion }}
{{- end }}
{{- if .StorageEncrypted }}
            StorageEncrypted: {{ .StorageEncrypted }}
{{- end }}
{{- if .StorageType }}
            StorageType: {{ .StorageType }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .Timezone }}
            Timezone: {{ .Timezone }}
{{- end }}
{{- if .VPCSecurityGroups }}
            VPCSecurityGroups:
{{- range $_, $v := .VPCSecurityGroups }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::RDS::DBParameterGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::RDS::DBParameterGroup
        Properties:
            Description: {{ .Description }}
            Family: {{ .Family }}
{{- if .Parameters }}
            Parameters:
{{- range $k, $v := .Parameters }}
                {{ $k }}: {{ $v }}
{{- end }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::RDS::DBSecurityGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::RDS::DBSecurityGroup
        Properties:
            DBSecurityGroupIngress:
{{- range $_, $v := .DBSecurityGroupIngress }}
                -
{{- if $v.CIDRIP }}
                  CIDRIP: {{ $v.CIDRIP }}
{{- end }}
{{- if $v.EC2SecurityGroupId }}
                  EC2SecurityGroupId: {{ $v.EC2SecurityGroupId }}
{{- end }}
{{- if $v.EC2SecurityGroupName }}
                  EC2SecurityGroupName: {{ $v.EC2SecurityGroupName }}
{{- end }}
{{- if $v.EC2SecurityGroupOwnerId }}
                  EC2SecurityGroupOwnerId: {{ $v.EC2SecurityGroupOwnerId }}
{{- end }}
{{- end }}
            GroupDescription: {{ .GroupDescription }}
{{- if .EC2VpcId }}
            EC2VpcId: {{ .EC2VpcId }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::RDS::DBSecurityGroupIngress" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::RDS::DBSecurityGroupIngress
        Properties:
            DBSecurityGroupName: {{ .DBSecurityGroupName }}
{{- if .CIDRIP }}
            CIDRIP: {{ .CIDRIP }}
{{- end }}
{{- if .EC2SecurityGroupId }}
            EC2SecurityGroupId: {{ .EC2SecurityGroupId }}
{{- end }}
{{- if .EC2SecurityGroupName }}
            EC2SecurityGroupName: {{ .EC2SecurityGroupName }}
{{- end }}
{{- if .EC2SecurityGroupOwnerId }}
            EC2SecurityGroupOwnerId: {{ .EC2SecurityGroupOwnerId }}
{{- end }}
{{- end }}
{{- define "AWS::RDS::DBSubnetGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::RDS::DBSubnetGroup
        Properties:
            DBSubnetGroupDescription: {{ .DBSubnetGroupDescription }}
            SubnetIds:
{{- range $_, $v := .SubnetIds }}
                - {{ $v }}
{{- end }}
{{- if .DBSubnetGroupName }}
            DBSubnetGroupName: {{ .DBSubnetGroupName }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::RDS::EventSubscription" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::RDS::EventSubscription
        Properties:
            SnsTopicArn: {{ .SnsTopicArn }}
{{- if .Enabled }}
            Enabled: {{ .Enabled }}
{{- end }}
{{- if .EventCategories }}
            EventCategories:
{{- range $_, $v := .EventCategories }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .SourceIds }}
            SourceIds:
{{- range $_, $v := .SourceIds }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .SourceType }}
            SourceType: {{ .SourceType }}
{{- end }}
{{- end }}
{{- define "AWS::RDS::OptionGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::RDS::OptionGroup
        Properties:
            EngineName: {{ .EngineName }}
            MajorEngineVersion: {{ .MajorEngineVersion }}
            OptionConfigurations:
{{- range $_, $v := .OptionConfigurations }}
                -
                  OptionName: {{ $v.OptionName }}
{{- if $v.DBSecurityGroupMemberships }}
                  DBSecurityGroupMemberships:
{{- range $_, $v := $v.DBSecurityGroupMemberships }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.OptionSettings }}
                  OptionSettings:
{{- if $v.OptionSettings.Name }}
                      Name: {{ $v.OptionSettings.Name }}
{{- end }}
{{- if $v.OptionSettings.Value }}
                      Value: {{ $v.OptionSettings.Value }}
{{- end }}
{{- end }}
{{- if $v.OptionVersion }}
                  OptionVersion: {{ $v.OptionVersion }}
{{- end }}
{{- if $v.Port }}
                  Port: {{ $v.Port }}
{{- end }}
{{- if $v.VpcSecurityGroupMemberships }}
                  VpcSecurityGroupMemberships:
{{- range $_, $v := $v.VpcSecurityGroupMemberships }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
            OptionGroupDescription: {{ .OptionGroupDescription }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Redshift::Cluster" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Redshift::Cluster
        Properties:
            ClusterType: {{ .ClusterType }}
            DBName: {{ .DBName }}
            MasterUserPassword: {{ .MasterUserPassword }}
            MasterUsername: {{ .MasterUsername }}
            NodeType: {{ .NodeType }}
{{- if .AllowVersionUpgrade }}
            AllowVersionUpgrade: {{ .AllowVersionUpgrade }}
{{- end }}
{{- if .AutomatedSnapshotRetentionPeriod }}
            AutomatedSnapshotRetentionPeriod: {{ .AutomatedSnapshotRetentionPeriod }}
{{- end }}
{{- if .AvailabilityZone }}
            AvailabilityZone: {{ .AvailabilityZone }}
{{- end }}
{{- if .ClusterIdentifier }}
            ClusterIdentifier: {{ .ClusterIdentifier }}
{{- end }}
{{- if .ClusterParameterGroupName }}
            ClusterParameterGroupName: {{ .ClusterParameterGroupName }}
{{- end }}
{{- if .ClusterSecurityGroups }}
            ClusterSecurityGroups:
{{- range $_, $v := .ClusterSecurityGroups }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .ClusterSubnetGroupName }}
            ClusterSubnetGroupName: {{ .ClusterSubnetGroupName }}
{{- end }}
{{- if .ClusterVersion }}
            ClusterVersion: {{ .ClusterVersion }}
{{- end }}
{{- if .ElasticIp }}
            ElasticIp: {{ .ElasticIp }}
{{- end }}
{{- if .Encrypted }}
            Encrypted: {{ .Encrypted }}
{{- end }}
{{- if .HsmClientCertificateIdentifier }}
            HsmClientCertificateIdentifier: {{ .HsmClientCertificateIdentifier }}
{{- end }}
{{- if .HsmConfigurationIdentifier }}
            HsmConfigurationIdentifier: {{ .HsmConfigurationIdentifier }}
{{- end }}
{{- if .IamRoles }}
            IamRoles:
{{- range $_, $v := .IamRoles }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .KmsKeyId }}
            KmsKeyId: {{ .KmsKeyId }}
{{- end }}
{{- if .LoggingProperties }}
            LoggingProperties:
                BucketName: {{ .LoggingProperties.BucketName }}
{{- if .LoggingProperties.S3KeyPrefix }}
                S3KeyPrefix: {{ .LoggingProperties.S3KeyPrefix }}
{{- end }}
{{- end }}
{{- if .NumberOfNodes }}
            NumberOfNodes: {{ .NumberOfNodes }}
{{- end }}
{{- if .OwnerAccount }}
            OwnerAccount: {{ .OwnerAccount }}
{{- end }}
{{- if .Port }}
            Port: {{ .Port }}
{{- end }}
{{- if .PreferredMaintenanceWindow }}
            PreferredMaintenanceWindow: {{ .PreferredMaintenanceWindow }}
{{- end }}
{{- if .PubliclyAccessible }}
            PubliclyAccessible: {{ .PubliclyAccessible }}
{{- end }}
{{- if .SnapshotClusterIdentifier }}
            SnapshotClusterIdentifier: {{ .SnapshotClusterIdentifier }}
{{- end }}
{{- if .SnapshotIdentifier }}
            SnapshotIdentifier: {{ .SnapshotIdentifier }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .VpcSecurityGroupIds }}
            VpcSecurityGroupIds:
{{- range $_, $v := .VpcSecurityGroupIds }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Redshift::ClusterParameterGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Redshift::ClusterParameterGroup
        Properties:
            Description: {{ .Description }}
            ParameterGroupFamily: {{ .ParameterGroupFamily }}
{{- if .Parameters }}
            Parameters:
{{- range $_, $v := .Parameters }}
                -
                  ParameterName: {{ $v.ParameterName }}
                  ParameterValue: {{ $v.ParameterValue }}
{{- end }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Redshift::ClusterSecurityGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Redshift::ClusterSecurityGroup
        Properties:
            Description: {{ .Description }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Redshift::ClusterSecurityGroupIngress" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Redshift::ClusterSecurityGroupIngress
        Properties:
            ClusterSecurityGroupName: {{ .ClusterSecurityGroupName }}
{{- if .CIDRIP }}
            CIDRIP: {{ .CIDRIP }}
{{- end }}
{{- if .EC2SecurityGroupName }}
            EC2SecurityGroupName: {{ .EC2SecurityGroupName }}
{{- end }}
{{- if .EC2SecurityGroupOwnerId }}
            EC2SecurityGroupOwnerId: {{ .EC2SecurityGroupOwnerId }}
{{- end }}
{{- end }}
{{- define "AWS::Redshift::ClusterSubnetGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Redshift::ClusterSubnetGroup
        Properties:
            Description: {{ .Description }}
            SubnetIds:
{{- range $_, $v := .SubnetIds }}
                - {{ $v }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Route53::HealthCheck" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Route53::HealthCheck
        Properties:
            HealthCheckConfig:
                Type: {{ .HealthCheckConfig.Type }}
{{- if .HealthCheckConfig.AlarmIdentifier }}
                AlarmIdentifier:
                    Name: {{ .HealthCheckConfig.AlarmIdentifier.Name }}
                    Region: {{ .HealthCheckConfig.AlarmIdentifier.Region }}
{{- end }}
{{- if .HealthCheckConfig.ChildHealthChecks }}
                ChildHealthChecks:
{{- range $_, $v := .HealthCheckConfig.ChildHealthChecks }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- if .HealthCheckConfig.EnableSNI }}
                EnableSNI: {{ .HealthCheckConfig.EnableSNI }}
{{- end }}
{{- if .HealthCheckConfig.FailureThreshold }}
                FailureThreshold: {{ .HealthCheckConfig.FailureThreshold }}
{{- end }}
{{- if .HealthCheckConfig.FullyQualifiedDomainName }}
                FullyQualifiedDomainName: {{ .HealthCheckConfig.FullyQualifiedDomainName }}
{{- end }}
{{- if .HealthCheckConfig.HealthThreshold }}
                HealthThreshold: {{ .HealthCheckConfig.HealthThreshold }}
{{- end }}
{{- if .HealthCheckConfig.IPAddress }}
                IPAddress: {{ .HealthCheckConfig.IPAddress }}
{{- end }}
{{- if .HealthCheckConfig.InsufficientDataHealthStatus }}
                InsufficientDataHealthStatus: {{ .HealthCheckConfig.InsufficientDataHealthStatus }}
{{- end }}
{{- if .HealthCheckConfig.Inverted }}
                Inverted: {{ .HealthCheckConfig.Inverted }}
{{- end }}
{{- if .HealthCheckConfig.MeasureLatency }}
                MeasureLatency: {{ .HealthCheckConfig.MeasureLatency }}
{{- end }}
{{- if .HealthCheckConfig.Port }}
                Port: {{ .HealthCheckConfig.Port }}
{{- end }}
{{- if .HealthCheckConfig.Regions }}
                Regions:
{{- range $_, $v := .HealthCheckConfig.Regions }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- if .HealthCheckConfig.RequestInterval }}
                RequestInterval: {{ .HealthCheckConfig.RequestInterval }}
{{- end }}
{{- if .HealthCheckConfig.ResourcePath }}
                ResourcePath: {{ .HealthCheckConfig.ResourcePath }}
{{- end }}
{{- if .HealthCheckConfig.SearchString }}
                SearchString: {{ .HealthCheckConfig.SearchString }}
{{- end }}
{{- if .HealthCheckTags }}
            HealthCheckTags:
{{- range $_, $v := .HealthCheckTags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Route53::HostedZone" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Route53::HostedZone
        Properties:
            Name: {{ .Name }}
{{- if .HostedZoneConfig }}
            HostedZoneConfig:
{{- if .HostedZoneConfig.Comment }}
                Comment: {{ .HostedZoneConfig.Comment }}
{{- end }}
{{- end }}
{{- if .HostedZoneTags }}
            HostedZoneTags:
{{- range $_, $v := .HostedZoneTags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .QueryLoggingConfig }}
            QueryLoggingConfig:
                CloudWatchLogsLogGroupArn: {{ .QueryLoggingConfig.CloudWatchLogsLogGroupArn }}
{{- end }}
{{- if .VPCs }}
            VPCs:
{{- range $_, $v := .VPCs }}
                -
                  VPCId: {{ $v.VPCId }}
                  VPCRegion: {{ $v.VPCRegion }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::Route53::RecordSet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Route53::RecordSet
        Properties:
            Name: {{ .Name }}
            Type: {{ .Type }}
{{- if .AliasTarget }}
            AliasTarget:
                DNSName: {{ .AliasTarget.DNSName }}
                HostedZoneId: {{ .AliasTarget.HostedZoneId }}
{{- if .AliasTarget.EvaluateTargetHealth }}
                EvaluateTargetHealth: {{ .AliasTarget.EvaluateTargetHealth }}
{{- end }}
{{- end }}
{{- if .Comment }}
            Comment: {{ .Comment }}
{{- end }}
{{- if .Failover }}
            Failover: {{ .Failover }}
{{- end }}
{{- if .GeoLocation }}
            GeoLocation:
{{- if .GeoLocation.ContinentCode }}
                ContinentCode: {{ .GeoLocation.ContinentCode }}
{{- end }}
{{- if .GeoLocation.CountryCode }}
                CountryCode: {{ .GeoLocation.CountryCode }}
{{- end }}
{{- if .GeoLocation.SubdivisionCode }}
                SubdivisionCode: {{ .GeoLocation.SubdivisionCode }}
{{- end }}
{{- end }}
{{- if .HealthCheckId }}
            HealthCheckId: {{ .HealthCheckId }}
{{- end }}
{{- if .HostedZoneId }}
            HostedZoneId: {{ .HostedZoneId }}
{{- end }}
{{- if .HostedZoneName }}
            HostedZoneName: {{ .HostedZoneName }}
{{- end }}
{{- if .Region }}
            Region: {{ .Region }}
{{- end }}
{{- if .ResourceRecords }}
            ResourceRecords:
{{- range $_, $v := .ResourceRecords }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .SetIdentifier }}
            SetIdentifier: {{ .SetIdentifier }}
{{- end }}
{{- if .TTL }}
            TTL: {{ .TTL }}
{{- end }}
{{- if .Weight }}
            Weight: {{ .Weight }}
{{- end }}
{{- end }}
{{- define "AWS::Route53::RecordSetGroup" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::Route53::RecordSetGroup
{{- if .Comment | or .HostedZoneId | or .HostedZoneName | or .RecordSets }}
        Properties:
{{- if .Comment }}
            Comment: {{ .Comment }}
{{- end }}
{{- if .HostedZoneId }}
            HostedZoneId: {{ .HostedZoneId }}
{{- end }}
{{- if .HostedZoneName }}
            HostedZoneName: {{ .HostedZoneName }}
{{- end }}
{{- if .RecordSets }}
            RecordSets:
{{- range $_, $v := .RecordSets }}
                -
                  Name: {{ $v.Name }}
                  Type: {{ $v.Type }}
{{- if $v.AliasTarget }}
                  AliasTarget:
                      DNSName: {{ $v.AliasTarget.DNSName }}
                      HostedZoneId: {{ $v.AliasTarget.HostedZoneId }}
{{- if $v.AliasTarget.EvaluateTargetHealth }}
                      EvaluateTargetHealth: {{ $v.AliasTarget.EvaluateTargetHealth }}
{{- end }}
{{- end }}
{{- if $v.Comment }}
                  Comment: {{ $v.Comment }}
{{- end }}
{{- if $v.Failover }}
                  Failover: {{ $v.Failover }}
{{- end }}
{{- if $v.GeoLocation }}
                  GeoLocation:
{{- if $v.GeoLocation.ContinentCode }}
                      ContinentCode: {{ $v.GeoLocation.ContinentCode }}
{{- end }}
{{- if $v.GeoLocation.CountryCode }}
                      CountryCode: {{ $v.GeoLocation.CountryCode }}
{{- end }}
{{- if $v.GeoLocation.SubdivisionCode }}
                      SubdivisionCode: {{ $v.GeoLocation.SubdivisionCode }}
{{- end }}
{{- end }}
{{- if $v.HealthCheckId }}
                  HealthCheckId: {{ $v.HealthCheckId }}
{{- end }}
{{- if $v.HostedZoneId }}
                  HostedZoneId: {{ $v.HostedZoneId }}
{{- end }}
{{- if $v.HostedZoneName }}
                  HostedZoneName: {{ $v.HostedZoneName }}
{{- end }}
{{- if $v.Region }}
                  Region: {{ $v.Region }}
{{- end }}
{{- if $v.ResourceRecords }}
                  ResourceRecords:
{{- range $_, $v := $v.ResourceRecords }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.SetIdentifier }}
                  SetIdentifier: {{ $v.SetIdentifier }}
{{- end }}
{{- if $v.TTL }}
                  TTL: {{ $v.TTL }}
{{- end }}
{{- if $v.Weight }}
                  Weight: {{ $v.Weight }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::S3::Bucket" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::S3::Bucket
{{- if .AccelerateConfiguration | or .AccessControl | or .AnalyticsConfigurations | or .BucketEncryption | or .BucketName | or .CorsConfiguration | or .InventoryConfigurations | or .LifecycleConfiguration | or .LoggingConfiguration | or .MetricsConfigurations | or .NotificationConfiguration | or .ReplicationConfiguration | or .Tags | or .VersioningConfiguration | or .WebsiteConfiguration }}
        Properties:
{{- if .AccelerateConfiguration }}
            AccelerateConfiguration:
                AccelerationStatus: {{ .AccelerateConfiguration.AccelerationStatus }}
{{- end }}
{{- if .AccessControl }}
            AccessControl: {{ .AccessControl }}
{{- end }}
{{- if .AnalyticsConfigurations }}
            AnalyticsConfigurations:
{{- range $_, $v := .AnalyticsConfigurations }}
                -
                  Id: {{ $v.Id }}
                  StorageClassAnalysis:
{{- if $v.StorageClassAnalysis.DataExport }}
                      DataExport:
                          Destination:
                              BucketArn: {{ $v.StorageClassAnalysis.DataExport.Destination.BucketArn }}
                              Format: {{ $v.StorageClassAnalysis.DataExport.Destination.Format }}
{{- if $v.StorageClassAnalysis.DataExport.Destination.BucketAccountId }}
                              BucketAccountId: {{ $v.StorageClassAnalysis.DataExport.Destination.BucketAccountId }}
{{- end }}
{{- if $v.StorageClassAnalysis.DataExport.Destination.Prefix }}
                              Prefix: {{ $v.StorageClassAnalysis.DataExport.Destination.Prefix }}
{{- end }}
                          OutputSchemaVersion: {{ $v.StorageClassAnalysis.DataExport.OutputSchemaVersion }}
{{- end }}
{{- if $v.Prefix }}
                  Prefix: {{ $v.Prefix }}
{{- end }}
{{- if $v.TagFilters }}
                  TagFilters:
{{- range $_, $v := $v.TagFilters }}
                      -
                        Key: {{ $v.Key }}
                        Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .BucketEncryption }}
            BucketEncryption:
                ServerSideEncryptionConfiguration:
{{- range $_, $v := .BucketEncryption.ServerSideEncryptionConfiguration }}
                    -
{{- if $v.ServerSideEncryptionByDefault }}
                      ServerSideEncryptionByDefault:
                          SSEAlgorithm: {{ $v.ServerSideEncryptionByDefault.SSEAlgorithm }}
{{- if $v.ServerSideEncryptionByDefault.KMSMasterKeyID }}
                          KMSMasterKeyID: {{ $v.ServerSideEncryptionByDefault.KMSMasterKeyID }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .BucketName }}
            BucketName: {{ .BucketName }}
{{- end }}
{{- if .CorsConfiguration }}
            CorsConfiguration:
                CorsRules:
{{- range $_, $v := .CorsConfiguration.CorsRules }}
                    -
                      AllowedMethods:
{{- range $_, $v := $v.AllowedMethods }}
                          - {{ $v }}
{{- end }}
                      AllowedOrigins:
{{- range $_, $v := $v.AllowedOrigins }}
                          - {{ $v }}
{{- end }}
{{- if $v.AllowedHeaders }}
                      AllowedHeaders:
{{- range $_, $v := $v.AllowedHeaders }}
                          - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.ExposedHeaders }}
                      ExposedHeaders:
{{- range $_, $v := $v.ExposedHeaders }}
                          - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.Id }}
                      Id: {{ $v.Id }}
{{- end }}
{{- if $v.MaxAge }}
                      MaxAge: {{ $v.MaxAge }}
{{- end }}
{{- end }}
{{- end }}
{{- if .InventoryConfigurations }}
            InventoryConfigurations:
{{- range $_, $v := .InventoryConfigurations }}
                -
                  Destination:
                      BucketArn: {{ $v.Destination.BucketArn }}
                      Format: {{ $v.Destination.Format }}
{{- if $v.Destination.BucketAccountId }}
                      BucketAccountId: {{ $v.Destination.BucketAccountId }}
{{- end }}
{{- if $v.Destination.Prefix }}
                      Prefix: {{ $v.Destination.Prefix }}
{{- end }}
                  Enabled: {{ $v.Enabled }}
                  Id: {{ $v.Id }}
                  IncludedObjectVersions: {{ $v.IncludedObjectVersions }}
                  ScheduleFrequency: {{ $v.ScheduleFrequency }}
{{- if $v.OptionalFields }}
                  OptionalFields:
{{- range $_, $v := $v.OptionalFields }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- if $v.Prefix }}
                  Prefix: {{ $v.Prefix }}
{{- end }}
{{- end }}
{{- end }}
{{- if .LifecycleConfiguration }}
            LifecycleConfiguration:
                Rules:
{{- range $_, $v := .LifecycleConfiguration.Rules }}
                    -
                      Status: {{ $v.Status }}
{{- if $v.AbortIncompleteMultipartUpload }}
                      AbortIncompleteMultipartUpload:
                          DaysAfterInitiation: {{ $v.AbortIncompleteMultipartUpload.DaysAfterInitiation }}
{{- end }}
{{- if $v.ExpirationDate }}
                      ExpirationDate: {{ $v.ExpirationDate }}
{{- end }}
{{- if $v.ExpirationInDays }}
                      ExpirationInDays: {{ $v.ExpirationInDays }}
{{- end }}
{{- if $v.Id }}
                      Id: {{ $v.Id }}
{{- end }}
{{- if $v.NoncurrentVersionExpirationInDays }}
                      NoncurrentVersionExpirationInDays: {{ $v.NoncurrentVersionExpirationInDays }}
{{- end }}
{{- if $v.NoncurrentVersionTransition }}
                      NoncurrentVersionTransition:
                          StorageClass: {{ $v.NoncurrentVersionTransition.StorageClass }}
                          TransitionInDays: {{ $v.NoncurrentVersionTransition.TransitionInDays }}
{{- end }}
{{- if $v.NoncurrentVersionTransitions }}
                      NoncurrentVersionTransitions:
{{- range $_, $v := $v.NoncurrentVersionTransitions }}
                          -
                            StorageClass: {{ $v.StorageClass }}
                            TransitionInDays: {{ $v.TransitionInDays }}
{{- end }}
{{- end }}
{{- if $v.Prefix }}
                      Prefix: {{ $v.Prefix }}
{{- end }}
{{- if $v.TagFilters }}
                      TagFilters:
{{- range $_, $v := $v.TagFilters }}
                          -
                            Key: {{ $v.Key }}
                            Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if $v.Transition }}
                      Transition:
                          StorageClass: {{ $v.Transition.StorageClass }}
{{- if $v.Transition.TransitionDate }}
                          TransitionDate: {{ $v.Transition.TransitionDate }}
{{- end }}
{{- if $v.Transition.TransitionInDays }}
                          TransitionInDays: {{ $v.Transition.TransitionInDays }}
{{- end }}
{{- end }}
{{- if $v.Transitions }}
                      Transitions:
{{- range $_, $v := $v.Transitions }}
                          -
                            StorageClass: {{ $v.StorageClass }}
{{- if $v.TransitionDate }}
                            TransitionDate: {{ $v.TransitionDate }}
{{- end }}
{{- if $v.TransitionInDays }}
                            TransitionInDays: {{ $v.TransitionInDays }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .LoggingConfiguration }}
            LoggingConfiguration:
{{- if .LoggingConfiguration.DestinationBucketName }}
                DestinationBucketName: {{ .LoggingConfiguration.DestinationBucketName }}
{{- end }}
{{- if .LoggingConfiguration.LogFilePrefix }}
                LogFilePrefix: {{ .LoggingConfiguration.LogFilePrefix }}
{{- end }}
{{- end }}
{{- if .MetricsConfigurations }}
            MetricsConfigurations:
{{- range $_, $v := .MetricsConfigurations }}
                -
                  Id: {{ $v.Id }}
{{- if $v.Prefix }}
                  Prefix: {{ $v.Prefix }}
{{- end }}
{{- if $v.TagFilters }}
                  TagFilters:
{{- range $_, $v := $v.TagFilters }}
                      -
                        Key: {{ $v.Key }}
                        Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .NotificationConfiguration }}
            NotificationConfiguration:
{{- if .NotificationConfiguration.LambdaConfigurations }}
                LambdaConfigurations:
{{- range $_, $v := .NotificationConfiguration.LambdaConfigurations }}
                    -
                      Event: {{ $v.Event }}
                      Function: {{ $v.Function }}
{{- if $v.Filter }}
                      Filter:
                          S3Key:
                              Rules:
{{- range $_, $v := $v.Filter.S3Key.Rules }}
                                  -
                                    Name: {{ $v.Name }}
                                    Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .NotificationConfiguration.QueueConfigurations }}
                QueueConfigurations:
{{- range $_, $v := .NotificationConfiguration.QueueConfigurations }}
                    -
                      Event: {{ $v.Event }}
                      Queue: {{ $v.Queue }}
{{- if $v.Filter }}
                      Filter:
                          S3Key:
                              Rules:
{{- range $_, $v := $v.Filter.S3Key.Rules }}
                                  -
                                    Name: {{ $v.Name }}
                                    Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .NotificationConfiguration.TopicConfigurations }}
                TopicConfigurations:
{{- range $_, $v := .NotificationConfiguration.TopicConfigurations }}
                    -
                      Event: {{ $v.Event }}
                      Topic: {{ $v.Topic }}
{{- if $v.Filter }}
                      Filter:
                          S3Key:
                              Rules:
{{- range $_, $v := $v.Filter.S3Key.Rules }}
                                  -
                                    Name: {{ $v.Name }}
                                    Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .ReplicationConfiguration }}
            ReplicationConfiguration:
                Role: {{ .ReplicationConfiguration.Role }}
                Rules:
{{- range $_, $v := .ReplicationConfiguration.Rules }}
                    -
                      Destination:
                          Bucket: {{ $v.Destination.Bucket }}
{{- if $v.Destination.AccessControlTranslation }}
                          AccessControlTranslation:
                              Owner: {{ $v.Destination.AccessControlTranslation.Owner }}
{{- end }}
{{- if $v.Destination.Account }}
                          Account: {{ $v.Destination.Account }}
{{- end }}
{{- if $v.Destination.EncryptionConfiguration }}
                          EncryptionConfiguration:
                              ReplicaKmsKeyID: {{ $v.Destination.EncryptionConfiguration.ReplicaKmsKeyID }}
{{- end }}
{{- if $v.Destination.StorageClass }}
                          StorageClass: {{ $v.Destination.StorageClass }}
{{- end }}
                      Prefix: {{ $v.Prefix }}
                      Status: {{ $v.Status }}
{{- if $v.Id }}
                      Id: {{ $v.Id }}
{{- end }}
{{- if $v.SourceSelectionCriteria }}
                      SourceSelectionCriteria:
                          SseKmsEncryptedObjects:
                              Status: {{ $v.SourceSelectionCriteria.SseKmsEncryptedObjects.Status }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- if .VersioningConfiguration }}
            VersioningConfiguration:
                Status: {{ .VersioningConfiguration.Status }}
{{- end }}
{{- if .WebsiteConfiguration }}
            WebsiteConfiguration:
{{- if .WebsiteConfiguration.ErrorDocument }}
                ErrorDocument: {{ .WebsiteConfiguration.ErrorDocument }}
{{- end }}
{{- if .WebsiteConfiguration.IndexDocument }}
                IndexDocument: {{ .WebsiteConfiguration.IndexDocument }}
{{- end }}
{{- if .WebsiteConfiguration.RedirectAllRequestsTo }}
                RedirectAllRequestsTo:
                    HostName: {{ .WebsiteConfiguration.RedirectAllRequestsTo.HostName }}
{{- if .WebsiteConfiguration.RedirectAllRequestsTo.Protocol }}
                    Protocol: {{ .WebsiteConfiguration.RedirectAllRequestsTo.Protocol }}
{{- end }}
{{- end }}
{{- if .WebsiteConfiguration.RoutingRules }}
                RoutingRules:
{{- range $_, $v := .WebsiteConfiguration.RoutingRules }}
                    -
                      RedirectRule:
{{- if $v.RedirectRule.HostName }}
                          HostName: {{ $v.RedirectRule.HostName }}
{{- end }}
{{- if $v.RedirectRule.HttpRedirectCode }}
                          HttpRedirectCode: {{ $v.RedirectRule.HttpRedirectCode }}
{{- end }}
{{- if $v.RedirectRule.Protocol }}
                          Protocol: {{ $v.RedirectRule.Protocol }}
{{- end }}
{{- if $v.RedirectRule.ReplaceKeyPrefixWith }}
                          ReplaceKeyPrefixWith: {{ $v.RedirectRule.ReplaceKeyPrefixWith }}
{{- end }}
{{- if $v.RedirectRule.ReplaceKeyWith }}
                          ReplaceKeyWith: {{ $v.RedirectRule.ReplaceKeyWith }}
{{- end }}
{{- if $v.RoutingRuleCondition }}
                      RoutingRuleCondition:
{{- if $v.RoutingRuleCondition.HttpErrorCodeReturnedEquals }}
                          HttpErrorCodeReturnedEquals: {{ $v.RoutingRuleCondition.HttpErrorCodeReturnedEquals }}
{{- end }}
{{- if $v.RoutingRuleCondition.KeyPrefixEquals }}
                          KeyPrefixEquals: {{ $v.RoutingRuleCondition.KeyPrefixEquals }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::S3::BucketPolicy" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::S3::BucketPolicy
        Properties:
            Bucket: {{ .Bucket }}
            PolicyDocument:
                {{ yaml.MarshalIndent .PolicyDocument "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- define "AWS::SDB::Domain" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::SDB::Domain
{{- if .Description }}
        Properties:
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::SES::ConfigurationSet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::SES::ConfigurationSet
{{- if .Name }}
        Properties:
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::SES::ConfigurationSetEventDestination" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::SES::ConfigurationSetEventDestination
        Properties:
            ConfigurationSetName: {{ .ConfigurationSetName }}
            EventDestination:
                MatchingEventTypes:
{{- range $_, $v := .EventDestination.MatchingEventTypes }}
                    - {{ $v }}
{{- end }}
{{- if .EventDestination.CloudWatchDestination }}
                CloudWatchDestination:
{{- if .EventDestination.CloudWatchDestination.DimensionConfigurations }}
                    DimensionConfigurations:
{{- range $_, $v := .EventDestination.CloudWatchDestination.DimensionConfigurations }}
                        -
                          DefaultDimensionValue: {{ $v.DefaultDimensionValue }}
                          DimensionName: {{ $v.DimensionName }}
                          DimensionValueSource: {{ $v.DimensionValueSource }}
{{- end }}
{{- end }}
{{- end }}
{{- if .EventDestination.Enabled }}
                Enabled: {{ .EventDestination.Enabled }}
{{- end }}
{{- if .EventDestination.KinesisFirehoseDestination }}
                KinesisFirehoseDestination:
                    DeliveryStreamARN: {{ .EventDestination.KinesisFirehoseDestination.DeliveryStreamARN }}
                    IAMRoleARN: {{ .EventDestination.KinesisFirehoseDestination.IAMRoleARN }}
{{- end }}
{{- if .EventDestination.Name }}
                Name: {{ .EventDestination.Name }}
{{- end }}
{{- end }}
{{- define "AWS::SES::ReceiptFilter" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::SES::ReceiptFilter
        Properties:
            Filter:
                IpFilter:
                    Cidr: {{ .Filter.IpFilter.Cidr }}
                    Policy: {{ .Filter.IpFilter.Policy }}
{{- if .Filter.Name }}
                Name: {{ .Filter.Name }}
{{- end }}
{{- end }}
{{- define "AWS::SES::ReceiptRule" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::SES::ReceiptRule
        Properties:
            Rule:
{{- if .Rule.Actions }}
                Actions:
{{- range $_, $v := .Rule.Actions }}
                    -
{{- if $v.AddHeaderAction }}
                      AddHeaderAction:
                          HeaderName: {{ $v.AddHeaderAction.HeaderName }}
                          HeaderValue: {{ $v.AddHeaderAction.HeaderValue }}
{{- end }}
{{- if $v.BounceAction }}
                      BounceAction:
                          Message: {{ $v.BounceAction.Message }}
                          Sender: {{ $v.BounceAction.Sender }}
                          SmtpReplyCode: {{ $v.BounceAction.SmtpReplyCode }}
{{- if $v.BounceAction.StatusCode }}
                          StatusCode: {{ $v.BounceAction.StatusCode }}
{{- end }}
{{- if $v.BounceAction.TopicArn }}
                          TopicArn: {{ $v.BounceAction.TopicArn }}
{{- end }}
{{- end }}
{{- if $v.LambdaAction }}
                      LambdaAction:
                          FunctionArn: {{ $v.LambdaAction.FunctionArn }}
{{- if $v.LambdaAction.InvocationType }}
                          InvocationType: {{ $v.LambdaAction.InvocationType }}
{{- end }}
{{- if $v.LambdaAction.TopicArn }}
                          TopicArn: {{ $v.LambdaAction.TopicArn }}
{{- end }}
{{- end }}
{{- if $v.S3Action }}
                      S3Action:
                          BucketName: {{ $v.S3Action.BucketName }}
{{- if $v.S3Action.KmsKeyArn }}
                          KmsKeyArn: {{ $v.S3Action.KmsKeyArn }}
{{- end }}
{{- if $v.S3Action.ObjectKeyPrefix }}
                          ObjectKeyPrefix: {{ $v.S3Action.ObjectKeyPrefix }}
{{- end }}
{{- if $v.S3Action.TopicArn }}
                          TopicArn: {{ $v.S3Action.TopicArn }}
{{- end }}
{{- end }}
{{- if $v.SNSAction }}
                      SNSAction:
{{- if $v.SNSAction.Encoding }}
                          Encoding: {{ $v.SNSAction.Encoding }}
{{- end }}
{{- if $v.SNSAction.TopicArn }}
                          TopicArn: {{ $v.SNSAction.TopicArn }}
{{- end }}
{{- end }}
{{- if $v.StopAction }}
                      StopAction:
                          Scope: {{ $v.StopAction.Scope }}
{{- if $v.StopAction.TopicArn }}
                          TopicArn: {{ $v.StopAction.TopicArn }}
{{- end }}
{{- end }}
{{- if $v.WorkmailAction }}
                      WorkmailAction:
                          OrganizationArn: {{ $v.WorkmailAction.OrganizationArn }}
{{- if $v.WorkmailAction.TopicArn }}
                          TopicArn: {{ $v.WorkmailAction.TopicArn }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .Rule.Enabled }}
                Enabled: {{ .Rule.Enabled }}
{{- end }}
{{- if .Rule.Name }}
                Name: {{ .Rule.Name }}
{{- end }}
{{- if .Rule.Recipients }}
                Recipients:
{{- range $_, $v := .Rule.Recipients }}
                    - {{ $v }}
{{- end }}
{{- end }}
{{- if .Rule.ScanEnabled }}
                ScanEnabled: {{ .Rule.ScanEnabled }}
{{- end }}
{{- if .Rule.TlsPolicy }}
                TlsPolicy: {{ .Rule.TlsPolicy }}
{{- end }}
            RuleSetName: {{ .RuleSetName }}
{{- if .After }}
            After: {{ .After }}
{{- end }}
{{- end }}
{{- define "AWS::SES::ReceiptRuleSet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::SES::ReceiptRuleSet
{{- if .RuleSetName }}
        Properties:
{{- if .RuleSetName }}
            RuleSetName: {{ .RuleSetName }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::SES::Template" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::SES::Template
{{- if .Template }}
        Properties:
{{- if .Template }}
            Template:
{{- if .Template.HtmlPart }}
                HtmlPart: {{ .Template.HtmlPart }}
{{- end }}
{{- if .Template.SubjectPart }}
                SubjectPart: {{ .Template.SubjectPart }}
{{- end }}
{{- if .Template.TemplateName }}
                TemplateName: {{ .Template.TemplateName }}
{{- end }}
{{- if .Template.TextPart }}
                TextPart: {{ .Template.TextPart }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::SNS::Subscription" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::SNS::Subscription
{{- if .Endpoint | or .Protocol | or .TopicArn }}
        Properties:
{{- if .Endpoint }}
            Endpoint: {{ .Endpoint }}
{{- end }}
{{- if .Protocol }}
            Protocol: {{ .Protocol }}
{{- end }}
{{- if .TopicArn }}
            TopicArn: {{ .TopicArn }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::SNS::Topic" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::SNS::Topic
{{- if .DisplayName | or .Subscription | or .TopicName }}
        Properties:
{{- if .DisplayName }}
            DisplayName: {{ .DisplayName }}
{{- end }}
{{- if .Subscription }}
            Subscription:
{{- range $_, $v := .Subscription }}
                -
                  Endpoint: {{ $v.Endpoint }}
                  Protocol: {{ $v.Protocol }}
{{- end }}
{{- end }}
{{- if .TopicName }}
            TopicName: {{ .TopicName }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::SNS::TopicPolicy" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::SNS::TopicPolicy
        Properties:
            PolicyDocument:
                {{ yaml.MarshalIndent .PolicyDocument "                " "    " | fmt.Sprintf "%s" }}
            Topics:
{{- range $_, $v := .Topics }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- define "AWS::SQS::Queue" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::SQS::Queue
{{- if .ContentBasedDeduplication | or .DelaySeconds | or .FifoQueue | or .KmsDataKeyReusePeriodSeconds | or .KmsMasterKeyId | or .MaximumMessageSize | or .MessageRetentionPeriod | or .QueueName | or .ReceiveMessageWaitTimeSeconds | or .RedrivePolicy | or .VisibilityTimeout }}
        Properties:
{{- if .ContentBasedDeduplication }}
            ContentBasedDeduplication: {{ .ContentBasedDeduplication }}
{{- end }}
{{- if .DelaySeconds }}
            DelaySeconds: {{ .DelaySeconds }}
{{- end }}
{{- if .FifoQueue }}
            FifoQueue: {{ .FifoQueue }}
{{- end }}
{{- if .KmsDataKeyReusePeriodSeconds }}
            KmsDataKeyReusePeriodSeconds: {{ .KmsDataKeyReusePeriodSeconds }}
{{- end }}
{{- if .KmsMasterKeyId }}
            KmsMasterKeyId: {{ .KmsMasterKeyId }}
{{- end }}
{{- if .MaximumMessageSize }}
            MaximumMessageSize: {{ .MaximumMessageSize }}
{{- end }}
{{- if .MessageRetentionPeriod }}
            MessageRetentionPeriod: {{ .MessageRetentionPeriod }}
{{- end }}
{{- if .QueueName }}
            QueueName: {{ .QueueName }}
{{- end }}
{{- if .ReceiveMessageWaitTimeSeconds }}
            ReceiveMessageWaitTimeSeconds: {{ .ReceiveMessageWaitTimeSeconds }}
{{- end }}
{{- if .RedrivePolicy }}
            RedrivePolicy:
                {{ yaml.MarshalIndent .RedrivePolicy "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .VisibilityTimeout }}
            VisibilityTimeout: {{ .VisibilityTimeout }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::SQS::QueuePolicy" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::SQS::QueuePolicy
        Properties:
            PolicyDocument:
                {{ yaml.MarshalIndent .PolicyDocument "                " "    " | fmt.Sprintf "%s" }}
            Queues:
{{- range $_, $v := .Queues }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- define "AWS::SSM::Association" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::SSM::Association
        Properties:
            Name: {{ .Name }}
{{- if .AssociationName }}
            AssociationName: {{ .AssociationName }}
{{- end }}
{{- if .DocumentVersion }}
            DocumentVersion: {{ .DocumentVersion }}
{{- end }}
{{- if .InstanceId }}
            InstanceId: {{ .InstanceId }}
{{- end }}
{{- if .Parameters }}
            Parameters:
{{- range $k, $v := .Parameters }}
                {{ $k }}:
                    ParameterValues:
{{- range $_, $v := $v.ParameterValues }}
                        - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- if .ScheduleExpression }}
            ScheduleExpression: {{ .ScheduleExpression }}
{{- end }}
{{- if .Targets }}
            Targets:
{{- range $_, $v := .Targets }}
                -
                  Key: {{ $v.Key }}
                  Values:
{{- range $_, $v := $v.Values }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::SSM::Document" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::SSM::Document
        Properties:
            Content:
                {{ yaml.MarshalIndent .Content "                " "    " | fmt.Sprintf "%s" }}
{{- if .DocumentType }}
            DocumentType: {{ .DocumentType }}
{{- end }}
{{- if .Tags }}
            Tags:
{{- range $_, $v := .Tags }}
                -
                  Key: {{ $v.Key }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::SSM::MaintenanceWindowTask" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::SSM::MaintenanceWindowTask
        Properties:
            MaxConcurrency: {{ .MaxConcurrency }}
            MaxErrors: {{ .MaxErrors }}
            Priority: {{ .Priority }}
            ServiceRoleArn: {{ .ServiceRoleArn }}
            Targets:
{{- range $_, $v := .Targets }}
                -
                  Key: {{ $v.Key }}
{{- if $v.Values }}
                  Values:
{{- range $_, $v := $v.Values }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
            TaskArn: {{ .TaskArn }}
            TaskType: {{ .TaskType }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .LoggingInfo }}
            LoggingInfo:
                Region: {{ .LoggingInfo.Region }}
                S3Bucket: {{ .LoggingInfo.S3Bucket }}
{{- if .LoggingInfo.S3Prefix }}
                S3Prefix: {{ .LoggingInfo.S3Prefix }}
{{- end }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- if .TaskInvocationParameters }}
            TaskInvocationParameters:
{{- if .TaskInvocationParameters.MaintenanceWindowAutomationParameters }}
                MaintenanceWindowAutomationParameters:
{{- if .TaskInvocationParameters.MaintenanceWindowAutomationParameters.DocumentVersion }}
                    DocumentVersion: {{ .TaskInvocationParameters.MaintenanceWindowAutomationParameters.DocumentVersion }}
{{- end }}
{{- if .TaskInvocationParameters.MaintenanceWindowAutomationParameters.Parameters }}
                    Parameters:
                        {{ yaml.MarshalIndent .TaskInvocationParameters.MaintenanceWindowAutomationParameters.Parameters "                        " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- end }}
{{- if .TaskInvocationParameters.MaintenanceWindowLambdaParameters }}
                MaintenanceWindowLambdaParameters:
{{- if .TaskInvocationParameters.MaintenanceWindowLambdaParameters.ClientContext }}
                    ClientContext: {{ .TaskInvocationParameters.MaintenanceWindowLambdaParameters.ClientContext }}
{{- end }}
{{- if .TaskInvocationParameters.MaintenanceWindowLambdaParameters.Payload }}
                    Payload: {{ .TaskInvocationParameters.MaintenanceWindowLambdaParameters.Payload }}
{{- end }}
{{- if .TaskInvocationParameters.MaintenanceWindowLambdaParameters.Qualifier }}
                    Qualifier: {{ .TaskInvocationParameters.MaintenanceWindowLambdaParameters.Qualifier }}
{{- end }}
{{- end }}
{{- if .TaskInvocationParameters.MaintenanceWindowRunCommandParameters }}
                MaintenanceWindowRunCommandParameters:
{{- if .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.Comment }}
                    Comment: {{ .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.Comment }}
{{- end }}
{{- if .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.DocumentHash }}
                    DocumentHash: {{ .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.DocumentHash }}
{{- end }}
{{- if .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.DocumentHashType }}
                    DocumentHashType: {{ .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.DocumentHashType }}
{{- end }}
{{- if .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.NotificationConfig }}
                    NotificationConfig:
                        NotificationArn: {{ .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.NotificationConfig.NotificationArn }}
{{- if .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.NotificationConfig.NotificationEvents }}
                        NotificationEvents:
{{- range $_, $v := .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.NotificationConfig.NotificationEvents }}
                            - {{ $v }}
{{- end }}
{{- end }}
{{- if .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.NotificationConfig.NotificationType }}
                        NotificationType: {{ .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.NotificationConfig.NotificationType }}
{{- end }}
{{- end }}
{{- if .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.OutputS3BucketName }}
                    OutputS3BucketName: {{ .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.OutputS3BucketName }}
{{- end }}
{{- if .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.OutputS3KeyPrefix }}
                    OutputS3KeyPrefix: {{ .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.OutputS3KeyPrefix }}
{{- end }}
{{- if .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.Parameters }}
                    Parameters:
                        {{ yaml.MarshalIndent .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.Parameters "                        " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.ServiceRoleArn }}
                    ServiceRoleArn: {{ .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.ServiceRoleArn }}
{{- end }}
{{- if .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.TimeoutSeconds }}
                    TimeoutSeconds: {{ .TaskInvocationParameters.MaintenanceWindowRunCommandParameters.TimeoutSeconds }}
{{- end }}
{{- end }}
{{- if .TaskInvocationParameters.MaintenanceWindowStepFunctionsParameters }}
                MaintenanceWindowStepFunctionsParameters:
{{- if .TaskInvocationParameters.MaintenanceWindowStepFunctionsParameters.Input }}
                    Input: {{ .TaskInvocationParameters.MaintenanceWindowStepFunctionsParameters.Input }}
{{- end }}
{{- if .TaskInvocationParameters.MaintenanceWindowStepFunctionsParameters.Name }}
                    Name: {{ .TaskInvocationParameters.MaintenanceWindowStepFunctionsParameters.Name }}
{{- end }}
{{- end }}
{{- end }}
{{- if .TaskParameters }}
            TaskParameters:
                {{ yaml.MarshalIndent .TaskParameters "                " "    " | fmt.Sprintf "%s" }}
{{- end }}
{{- if .WindowId }}
            WindowId: {{ .WindowId }}
{{- end }}
{{- end }}
{{- define "AWS::SSM::Parameter" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::SSM::Parameter
        Properties:
            Type: {{ .Type }}
            Value: {{ .Value }}
{{- if .AllowedPattern }}
            AllowedPattern: {{ .AllowedPattern }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- end }}
{{- define "AWS::SSM::PatchBaseline" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::SSM::PatchBaseline
        Properties:
            Name: {{ .Name }}
{{- if .ApprovalRules }}
            ApprovalRules:
{{- if .ApprovalRules.PatchRules }}
                PatchRules:
{{- range $_, $v := .ApprovalRules.PatchRules }}
                    -
{{- if $v.ApproveAfterDays }}
                      ApproveAfterDays: {{ $v.ApproveAfterDays }}
{{- end }}
{{- if $v.ComplianceLevel }}
                      ComplianceLevel: {{ $v.ComplianceLevel }}
{{- end }}
{{- if $v.EnableNonSecurity }}
                      EnableNonSecurity: {{ $v.EnableNonSecurity }}
{{- end }}
{{- if $v.PatchFilterGroup }}
                      PatchFilterGroup:
{{- if $v.PatchFilterGroup.PatchFilters }}
                          PatchFilters:
{{- range $_, $v := $v.PatchFilterGroup.PatchFilters }}
                              -
{{- if $v.Key }}
                                Key: {{ $v.Key }}
{{- end }}
{{- if $v.Values }}
                                Values:
{{- range $_, $v := $v.Values }}
                                    - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .ApprovedPatches }}
            ApprovedPatches:
{{- range $_, $v := .ApprovedPatches }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .ApprovedPatchesComplianceLevel }}
            ApprovedPatchesComplianceLevel: {{ .ApprovedPatchesComplianceLevel }}
{{- end }}
{{- if .ApprovedPatchesEnableNonSecurity }}
            ApprovedPatchesEnableNonSecurity: {{ .ApprovedPatchesEnableNonSecurity }}
{{- end }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .GlobalFilters }}
            GlobalFilters:
{{- if .GlobalFilters.PatchFilters }}
                PatchFilters:
{{- range $_, $v := .GlobalFilters.PatchFilters }}
                    -
{{- if $v.Key }}
                      Key: {{ $v.Key }}
{{- end }}
{{- if $v.Values }}
                      Values:
{{- range $_, $v := $v.Values }}
                          - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- if .OperatingSystem }}
            OperatingSystem: {{ .OperatingSystem }}
{{- end }}
{{- if .PatchGroups }}
            PatchGroups:
{{- range $_, $v := .PatchGroups }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .RejectedPatches }}
            RejectedPatches:
{{- range $_, $v := .RejectedPatches }}
                - {{ $v }}
{{- end }}
{{- end }}
{{- if .Sources }}
            Sources:
{{- range $_, $v := .Sources }}
                -
{{- if $v.Configuration }}
                  Configuration: {{ $v.Configuration }}
{{- end }}
{{- if $v.Name }}
                  Name: {{ $v.Name }}
{{- end }}
{{- if $v.Products }}
                  Products:
{{- range $_, $v := $v.Products }}
                      - {{ $v }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::ServiceDiscovery::Instance" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ServiceDiscovery::Instance
        Properties:
            InstanceAttributes:
                {{ yaml.MarshalIndent .InstanceAttributes "                " "    " | fmt.Sprintf "%s" }}
            ServiceId: {{ .ServiceId }}
{{- if .InstanceId }}
            InstanceId: {{ .InstanceId }}
{{- end }}
{{- end }}
{{- define "AWS::ServiceDiscovery::PrivateDnsNamespace" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ServiceDiscovery::PrivateDnsNamespace
        Properties:
            Name: {{ .Name }}
            Vpc: {{ .Vpc }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- end }}
{{- define "AWS::ServiceDiscovery::PublicDnsNamespace" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ServiceDiscovery::PublicDnsNamespace
        Properties:
            Name: {{ .Name }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- end }}
{{- define "AWS::ServiceDiscovery::Service" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::ServiceDiscovery::Service
        Properties:
            DnsConfig:
                DnsRecords:
{{- range $_, $v := .DnsConfig.DnsRecords }}
                    -
                      TTL: {{ $v.TTL }}
                      Type: {{ $v.Type }}
{{- end }}
                NamespaceId: {{ .DnsConfig.NamespaceId }}
{{- if .Description }}
            Description: {{ .Description }}
{{- end }}
{{- if .HealthCheckConfig }}
            HealthCheckConfig:
                Type: {{ .HealthCheckConfig.Type }}
{{- if .HealthCheckConfig.FailureThreshold }}
                FailureThreshold: {{ .HealthCheckConfig.FailureThreshold }}
{{- end }}
{{- if .HealthCheckConfig.ResourcePath }}
                ResourcePath: {{ .HealthCheckConfig.ResourcePath }}
{{- end }}
{{- end }}
{{- if .Name }}
            Name: {{ .Name }}
{{- end }}
{{- end }}
{{- define "AWS::StepFunctions::Activity" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::StepFunctions::Activity
        Properties:
            Name: {{ .Name }}
{{- end }}
{{- define "AWS::StepFunctions::StateMachine" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::StepFunctions::StateMachine
        Properties:
            DefinitionString: {{ .DefinitionString }}
            RoleArn: {{ .RoleArn }}
{{- if .StateMachineName }}
            StateMachineName: {{ .StateMachineName }}
{{- end }}
{{- end }}
{{- define "AWS::WAF::ByteMatchSet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::WAF::ByteMatchSet
        Properties:
            Name: {{ .Name }}
{{- if .ByteMatchTuples }}
            ByteMatchTuples:
{{- range $_, $v := .ByteMatchTuples }}
                -
                  FieldToMatch:
                      Type: {{ $v.FieldToMatch.Type }}
{{- if $v.FieldToMatch.Data }}
                      Data: {{ $v.FieldToMatch.Data }}
{{- end }}
                  PositionalConstraint: {{ $v.PositionalConstraint }}
                  TextTransformation: {{ $v.TextTransformation }}
{{- if $v.TargetString }}
                  TargetString: {{ $v.TargetString }}
{{- end }}
{{- if $v.TargetStringBase64 }}
                  TargetStringBase64: {{ $v.TargetStringBase64 }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::WAF::IPSet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::WAF::IPSet
        Properties:
            Name: {{ .Name }}
{{- if .IPSetDescriptors }}
            IPSetDescriptors:
{{- range $_, $v := .IPSetDescriptors }}
                -
                  Type: {{ $v.Type }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::WAF::Rule" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::WAF::Rule
        Properties:
            MetricName: {{ .MetricName }}
            Name: {{ .Name }}
{{- if .Predicates }}
            Predicates:
{{- range $_, $v := .Predicates }}
                -
                  DataId: {{ $v.DataId }}
                  Negated: {{ $v.Negated }}
                  Type: {{ $v.Type }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::WAF::SizeConstraintSet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::WAF::SizeConstraintSet
        Properties:
            Name: {{ .Name }}
            SizeConstraints:
{{- range $_, $v := .SizeConstraints }}
                -
                  ComparisonOperator: {{ $v.ComparisonOperator }}
                  FieldToMatch:
                      Type: {{ $v.FieldToMatch.Type }}
{{- if $v.FieldToMatch.Data }}
                      Data: {{ $v.FieldToMatch.Data }}
{{- end }}
                  Size: {{ $v.Size }}
                  TextTransformation: {{ $v.TextTransformation }}
{{- end }}
{{- end }}
{{- define "AWS::WAF::SqlInjectionMatchSet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::WAF::SqlInjectionMatchSet
        Properties:
            Name: {{ .Name }}
{{- if .SqlInjectionMatchTuples }}
            SqlInjectionMatchTuples:
{{- range $_, $v := .SqlInjectionMatchTuples }}
                -
                  FieldToMatch:
                      Type: {{ $v.FieldToMatch.Type }}
{{- if $v.FieldToMatch.Data }}
                      Data: {{ $v.FieldToMatch.Data }}
{{- end }}
                  TextTransformation: {{ $v.TextTransformation }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::WAF::WebACL" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::WAF::WebACL
        Properties:
            DefaultAction:
                Type: {{ .DefaultAction.Type }}
            MetricName: {{ .MetricName }}
            Name: {{ .Name }}
{{- if .Rules }}
            Rules:
{{- range $_, $v := .Rules }}
                -
                  Action:
                      Type: {{ $v.Action.Type }}
                  Priority: {{ $v.Priority }}
                  RuleId: {{ $v.RuleId }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::WAF::XssMatchSet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::WAF::XssMatchSet
        Properties:
            Name: {{ .Name }}
            XssMatchTuples:
{{- range $_, $v := .XssMatchTuples }}
                -
                  FieldToMatch:
                      Type: {{ $v.FieldToMatch.Type }}
{{- if $v.FieldToMatch.Data }}
                      Data: {{ $v.FieldToMatch.Data }}
{{- end }}
                  TextTransformation: {{ $v.TextTransformation }}
{{- end }}
{{- end }}
{{- define "AWS::WAFRegional::ByteMatchSet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::WAFRegional::ByteMatchSet
        Properties:
            Name: {{ .Name }}
{{- if .ByteMatchTuples }}
            ByteMatchTuples:
{{- range $_, $v := .ByteMatchTuples }}
                -
                  FieldToMatch:
                      Type: {{ $v.FieldToMatch.Type }}
{{- if $v.FieldToMatch.Data }}
                      Data: {{ $v.FieldToMatch.Data }}
{{- end }}
                  PositionalConstraint: {{ $v.PositionalConstraint }}
                  TextTransformation: {{ $v.TextTransformation }}
{{- if $v.TargetString }}
                  TargetString: {{ $v.TargetString }}
{{- end }}
{{- if $v.TargetStringBase64 }}
                  TargetStringBase64: {{ $v.TargetStringBase64 }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::WAFRegional::IPSet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::WAFRegional::IPSet
        Properties:
            Name: {{ .Name }}
{{- if .IPSetDescriptors }}
            IPSetDescriptors:
{{- range $_, $v := .IPSetDescriptors }}
                -
                  Type: {{ $v.Type }}
                  Value: {{ $v.Value }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::WAFRegional::Rule" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::WAFRegional::Rule
        Properties:
            MetricName: {{ .MetricName }}
            Name: {{ .Name }}
{{- if .Predicates }}
            Predicates:
{{- range $_, $v := .Predicates }}
                -
                  DataId: {{ $v.DataId }}
                  Negated: {{ $v.Negated }}
                  Type: {{ $v.Type }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::WAFRegional::SizeConstraintSet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::WAFRegional::SizeConstraintSet
        Properties:
            Name: {{ .Name }}
{{- if .SizeConstraints }}
            SizeConstraints:
{{- range $_, $v := .SizeConstraints }}
                -
                  ComparisonOperator: {{ $v.ComparisonOperator }}
                  FieldToMatch:
                      Type: {{ $v.FieldToMatch.Type }}
{{- if $v.FieldToMatch.Data }}
                      Data: {{ $v.FieldToMatch.Data }}
{{- end }}
                  Size: {{ $v.Size }}
                  TextTransformation: {{ $v.TextTransformation }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::WAFRegional::SqlInjectionMatchSet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::WAFRegional::SqlInjectionMatchSet
        Properties:
            Name: {{ .Name }}
{{- if .SqlInjectionMatchTuples }}
            SqlInjectionMatchTuples:
{{- range $_, $v := .SqlInjectionMatchTuples }}
                -
                  FieldToMatch:
                      Type: {{ $v.FieldToMatch.Type }}
{{- if $v.FieldToMatch.Data }}
                      Data: {{ $v.FieldToMatch.Data }}
{{- end }}
                  TextTransformation: {{ $v.TextTransformation }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::WAFRegional::WebACL" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::WAFRegional::WebACL
        Properties:
            DefaultAction:
                Type: {{ .DefaultAction.Type }}
            MetricName: {{ .MetricName }}
            Name: {{ .Name }}
{{- if .Rules }}
            Rules:
{{- range $_, $v := .Rules }}
                -
                  Action:
                      Type: {{ $v.Action.Type }}
                  Priority: {{ $v.Priority }}
                  RuleId: {{ $v.RuleId }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::WAFRegional::WebACLAssociation" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::WAFRegional::WebACLAssociation
        Properties:
            ResourceArn: {{ .ResourceArn }}
            WebACLId: {{ .WebACLId }}
{{- end }}
{{- define "AWS::WAFRegional::XssMatchSet" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::WAFRegional::XssMatchSet
        Properties:
            Name: {{ .Name }}
{{- if .XssMatchTuples }}
            XssMatchTuples:
{{- range $_, $v := .XssMatchTuples }}
                -
                  FieldToMatch:
                      Type: {{ $v.FieldToMatch.Type }}
{{- if $v.FieldToMatch.Data }}
                      Data: {{ $v.FieldToMatch.Data }}
{{- end }}
                  TextTransformation: {{ $v.TextTransformation }}
{{- end }}
{{- end }}
{{- end }}
{{- define "AWS::WorkSpaces::Workspace" }}
{{- with .LogicalId }}
    {{ . }}:
{{- end }}
{{- with .DependsOn }}
        DependsOn:
{{- range $_, $v := . }}
            - {{ $v }}
{{- end }}
{{- end }}
        Type: AWS::WorkSpaces::Workspace
        Properties:
            BundleId: {{ .BundleId }}
            DirectoryId: {{ .DirectoryId }}
            UserName: {{ .UserName }}
{{- if .RootVolumeEncryptionEnabled }}
            RootVolumeEncryptionEnabled: {{ .RootVolumeEncryptionEnabled }}
{{- end }}
{{- if .UserVolumeEncryptionEnabled }}
            UserVolumeEncryptionEnabled: {{ .UserVolumeEncryptionEnabled }}
{{- end }}
{{- if .VolumeEncryptionKey }}
            VolumeEncryptionKey: {{ .VolumeEncryptionKey }}
{{- end }}
{{- end }}`
