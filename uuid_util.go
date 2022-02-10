package util

import (
	uuid "github.com/iris-contrib/go.uuid"
	"strings"
)

func GetUUID() string {
	uid, err := uuid.NewV4()
	if err != nil {
		return ""
	}
	str := uid.String()
	oldStr := "-"
	newStr := ""
	str = strings.Replace(str, oldStr, newStr, -1)
	return str
}


