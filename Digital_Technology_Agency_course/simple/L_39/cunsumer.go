package L_39

type Consumer interface {
	Update(pubName string)
	GetName() string
}
