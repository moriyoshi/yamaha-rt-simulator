package server

func IsTimeout(e error) bool {
	t, ok := e.(interface {
		Timeout() bool
	})
	if ok {
		return t.Timeout()
	}
	return false
}
