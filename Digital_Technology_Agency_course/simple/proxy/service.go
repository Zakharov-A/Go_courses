package proxy

type Service interface {
	GetData(user string) ([]string, error)
}
