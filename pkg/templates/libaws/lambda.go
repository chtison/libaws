package libaws

import (
	"fmt"
	"strings"

	"github.com/chtison/libgo/tmpl/pkg/utils"
)

//go:generate bash generate.bash "LIBAWS::Lambda" Lambda
func (l *Libaws) Lambda(m utils.Map) (string, error) {
	b := newBuilder()
	m.SetDefault("", "LogicalIdSuffix")
	m.SetDefault(m["LogicalIdSuffix"], "Function", "LogicalIdSuffix")
	if err := b.writeString(l.LambdaFunction(m["Function"].(map[interface{}]interface{}))); err != nil {
		return b.String(), err
	}
	return b.String(), nil
}

//go:generate bash generate.bash "LIBAWS::Lambda::Function" LambdaFunction
func (l *Libaws) LambdaFunction(m utils.Map) (string, error) {
	b := newBuilder()
	logicalIdSuffix := m.GetDefault("", "LogicalIdSuffix").(string)
	m.SetDefault("Function"+logicalIdSuffix, "LogicalId")
	m.SetDefault("go1.x", "Runtime")
	if len(logicalIdSuffix) > 0 {
		m.SetDefault(strings.ToLower(logicalIdSuffix), "Handler")
	} else {
		m.SetDefault("main", "Handler")
	}
	m.SetDefault(fmt.Sprintf("lambdas/%s.zip", m["Handler"]), "Code")
	if _, ok := m["Role"]; !ok {
		m.SetDefault(make(utils.Map), "RoleProperties")
		roleData := m["RoleProperties"].(utils.Map)
		roleData["LogicalIdSuffix"] = logicalIdSuffix
		if err := b.writeString(l.IamRole(roleData)); err != nil {
			return b.String(), err
		}
		m.Set(fmt.Sprintf("!GetAtt %s.Arn", roleData["LogicalId"]), "Role")
	}
	err := l.Template.ExecuteTemplate(b, "AWS::Lambda::Function", m)
	return b.String(), err
}

//go:generate bash generate.bash "LIBAWS::Lambda::Permission" LambdaPermission
func (l *Libaws) LambdaPermission(data map[interface{}]interface{}) (string, error) {
	m := utils.Map(data)
	b := newBuilder()
	m.SetDefault("lambda:InvokeFunction", "Action")
	err := l.Template.ExecuteTemplate(b, "AWS::Lambda::Permission", m)
	return b.String(), err
}
