package middleware

import "strings"

func CleanStudent(dept string, name string, mat_no string) (string, string, string) {
	dept = strings.ToLower(dept)
	name = strings.ToLower(name)
	mat_no = strings.ToLower(mat_no)

	return dept, name, mat_no
}
