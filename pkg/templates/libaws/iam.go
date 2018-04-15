package libaws

import (
	"github.com/chtison/libgo/tmpl/pkg/utils"
	"github.com/chtison/libgo/yaml"
)

//go:generate bash generate.bash "LIBAWS::IAM::Role" IamRole
func (l *Libaws) IamRole(data map[interface{}]interface{}) (string, error) {
	m := utils.Map(data)
	logicalIdSuffix := m.GetDefault("", "LogicalIdSuffix").(string)
	m.SetDefault("Role"+logicalIdSuffix, "LogicalId")
	defaultServices := []string{"lambda.amazonaws.com"}
	defaultAssumeRolePolicyDocument := l.CreateAssumeRolePolicyDocument(defaultServices...)
	m.SetDefault(defaultAssumeRolePolicyDocument, "AssumeRolePolicyDocument")
	defaultPolicies := []string{"arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"}
	m.SetDefault(defaultPolicies, "ManagedPolicyArns")
	b := newBuilder()
	err := l.Template.ExecuteTemplate(b, "AWS::IAM::Role", m)
	return b.String(), err
}

func (*Libaws) CreateAssumeRolePolicyDocument(services ...string) yaml.MapSlice {
	return yaml.MapSlice{
		{"Version", "2012-10-17"},
		{"Statement", yaml.MapSlice{
			{"Effect", "Allow"},
			{"Action", "sts:AssumeRole"},
			{"Principal", yaml.MapSlice{
				{"Service", services},
			}},
		}},
	}
}
