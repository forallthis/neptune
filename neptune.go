package neptune

// CheckErr 检查错误
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
