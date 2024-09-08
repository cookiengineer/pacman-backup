package console

func Result(result bool, message string) {

	if result == true {
		Info(message)
	} else {
		Error(message)
	}

}
