package helper

func IfError(err error) {
	if err != nil {
		panic(err)
	}
}

// var ErrValidate = errors.New("ErrorValidation")

// func ErrValidation(err error) error {
// 	if err != nil {
// 		return ErrValidate
// 	}
// 	return nil
// }
