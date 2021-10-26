package util

import "strconv"

func ProjectIDAsString(projectID int32) string {
	return strconv.Itoa(int(projectID))
}
