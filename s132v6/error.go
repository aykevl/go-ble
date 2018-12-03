package sd

type Error uintptr

func (e Error) Error() string {
	return "SoftDevice error"
}
