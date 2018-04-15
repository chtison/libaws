package libaws

import (
	"github.com/chtison/libaws/pkg/aws/apigateway"
)

type ApiGateway struct {
	RestApi apigateway.RestApi
}

func (l *Libaws) ApiGateway(ag *ApiGateway) (string, error) {

}

func (l *Libaws) ApiGatewayRestApi(restApi *apigateway.RestApi) string {
	if restApi.LogicalId == "" {
		restApi.LogicalId = fmt.Sprintf("RestApi%s", l.LogicalIdSuffix)
	}
	return ""
}

func (l *Libaws) ApiGatewayMethods(methods []*apigateway.Method) (string, error) {

}

func (l *Libaws) ApiGateway(t *template.Template, data map[interface{}]interface{}, out io.Writer) {
	data = deepcopy.Copy(data).(map[interface{}]interface{})
}

func execRestApi(t *template.Template, data map[interface{}]interface{}, out io.Writer) error {
	return t.ExecuteTemplate(out, "AWS::ApiGateway::RestApi", data["RestApi"])
}

func execMethods(t *template.Template, data map[interface{}]interface{}, out io.Writer) error {
	if _, ok := data["Methods"]; !ok {
		return nil
	}
	if !IsKind(data["Methods"], reflect.Slice) {
		return fmt.Errorf(`["Methods"]: %s`, ErrorType(data["Methods"], reflect.Slice.String()))
	}
	for i, method := range data["Methods"] {
		if !IsKind(method, reflect.Map) {
			return fmt.Errorf(`["Methods"][%d]: %s`, i, ErrorType(method, reflect.Map.String()))
		}
		resource := deepcopy.Copy(method)
	}
}

func IsType(i interface{}, typeString string) bool {
	if t := reflect.TypeOf(i); t != nil {
		return t.String() == typeString
	}
	return false
}

func IsKind(i interface{}, kind reflect.Kind) bool {
	if t := reflect.TypeOf(i); t != nil {
		return t.Kind() == kind
	}
	return false
}

func ErrorType(i interface{}, expected string) error {
	return fmt.Errorf(`is of type %T but expected is %s`, i, expected)
}
