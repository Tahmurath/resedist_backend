package reflection

import "reflect"

// متد GetFieldName برای ModelHelper
func GetFieldName(model interface{}, field string) string {
	typ := reflect.TypeOf(model)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	// جستجوی فیلد با نام مشخص‌شده
	f, ok := typ.FieldByName(field)
	if ok {
		return f.Name
	}
	return ""
}

//depModel := DepartmentModel.Department{}
//DepScopes.Preload(request.Expand, refl.GetFieldName(depModel, "DepartmentType"), refl.GetFieldName(depModel, "Parent")),
