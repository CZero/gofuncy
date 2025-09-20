package errors

// PanErr panics with an error, or does nothing
func PanErr(err error) {
	if err != nil {
		panic(err)
	}
}
