package ziface

type IRouter interface {
	// before connection
	PreHandle(request IRequest)
	// handleing connection
	Handle(request IRequest)
	// after connection
	PostHandle(request IRequest)
}
