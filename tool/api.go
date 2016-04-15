package main

import "go/format"

type apiField struct {
	name    string
	vtype   string
	empty   bool
	comment string
}

type apiMeta struct {
	name    string
	url     string
	desc    string
	indices map[string]string
	args    []apiField
	rets    []apiField
}

var apis map[string]map[string]apiMeta

func init() {
	apis = map[string]map[string]apiMeta{}
}

func (meta apiMeta) code() string {
	str := "//" + meta.desc + "\n"
	str += "\"" + meta.name + "\":map[string]interface{}{\n"

	str += "\"url\":\"" + meta.url + "\",\n"

	str += "\"args\":map[string]string{\n"
	for _, v := range meta.args {
		str += "\"" + v.name + "\":\"" + v.vtype + "\",//" + v.comment + "\n"
	}
	str += "},\n"

	str += "\"rets\":map[string]string{\n"
	for _, v := range meta.rets {
		str += "\"" + v.name + "\":\"" + v.vtype + "\",//" + v.comment + "\n"
	}
	str += "},\n"

	str += "\"indices\":map[string]string{\n"
	for k, v := range meta.indices {
		str += "\"" + k + "\":\"" + v + "\",\n"
	}
	str += "},\n"

	str += "},\n"
	return str
}

func tostring() ([]byte, error) {
	str := "package godatayes\n\n"
	str += "var apis map[string]map[string]interface{}\n\n"

	str += "func init() {\n"

	str += "apis=map[string]map[string]interface{}{ \n"
	for k, doc := range apis {
		str += "//" + k + "\n"
		for _, v := range doc {
			str += v.code()
		}
	}

	str += "}\n"

	str += "}\n"

	return format.Source([]byte(str))
}
