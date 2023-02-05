package response

import "fmt"

func CheckError(err error, msg string, args ...interface{}) bool {
	if err != nil {
		if len(args) > 0 {
			fmt.Println("%s:%v", fmt.Sprintf(msg, args...), err)
		} else {
			fmt.Println("%s:%v", msg, err)
		}
		return true
	}

	return false
}
