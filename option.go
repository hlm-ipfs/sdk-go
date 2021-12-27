package sdk

type Options func(*Option)

type Option struct {
	Addr  string
	AppID string
	Name  string
	Pwd   string

	Err func(err error)
}
