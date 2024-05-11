package stringhelper

import "go-boiler-plate/enums"

func StringPointerToString(param *string) string {
	if param == nil {
		return enums.EMPTY_STRING
	}
	return *param
}
