package handler

import v1 "github.com/vulcangz/gf2-authz/api/authz/v1"

func attributesMap(attributes []*v1.Attribute) map[string]any {
	var result = map[string]any{}

	for _, attribute := range attributes {
		result[attribute.GetKey()] = attribute.GetValue()
	}

	return result
}
