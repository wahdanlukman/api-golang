package helper

import "log"

//CheckErr untuk print error
func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
