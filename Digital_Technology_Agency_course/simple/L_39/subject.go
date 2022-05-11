package L_39

type Subject interface {
	Subscribe(Consumer)
	UnSubscribe(Consumer)
	Notify()
}
