package weixin

type Payload interface {
	PreSignCheck() error
}
