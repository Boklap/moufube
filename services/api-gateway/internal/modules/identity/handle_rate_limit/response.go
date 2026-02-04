package handleratelimit

type Response struct {
	Count       int
	ShouldBlock bool
	Err         error
}
