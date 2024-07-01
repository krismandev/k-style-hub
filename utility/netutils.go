package utility

import "os"

func GetHostName() string {
	name, err := os.Hostname()
	if err != nil {
		return ""
	}
	return name

}
