package godatayes

import "fmt"

type Argument struct {
	args map[string]string
	api  map[string]interface{}
}

func NewArgument(name string, vals map[string]interface{}) *Argument {
	api, ok := apis[name]
	if !ok {
		return nil
	}
	arg := &Argument{args: map[string]string{}, api: api}

	fields := api["args"].(map[string]string)
	for k, _ := range fields {
		arg.Set(k, "")
	}

	for k, v := range vals {
		arg.Set(k, v)
	}
	return arg
}

func (arg *Argument) Set(key string, val interface{}) {
	arg.args[key] = fmt.Sprint(val)
}

func (arg *Argument) URL() string {
	url := arg.api["url"].(string)
	and := "?"
	for k, v := range arg.args {
		url += and + k + "=" + v
		and = "&"
	}

	return url
}
