package ziface

// Request inface
type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
}
