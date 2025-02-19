package protojson

type fieldsFilter struct {
	replacement string
	fields      map[string]struct{}
}

func FieldsFilter(fields []string, replacement string) fieldsFilter {
	fieldsMap := make(map[string]struct{}, len(fields))
	for _, f := range fields {
		fieldsMap[f] = struct{}{}
	}
	return fieldsFilter{
		replacement: replacement,
		fields:      fieldsMap,
	}
}
