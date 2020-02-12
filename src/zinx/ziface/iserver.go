package ziface

// server inface
type IServer interface {
	// start
	Start()
	// stop
	Stop()
	// run
	Serve()
}
