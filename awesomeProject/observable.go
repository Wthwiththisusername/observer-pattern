package awesomeProject

type Observable interface {
	subscribe(observable Observable)
	unsubscribe(observable Observable)
	sendAll()
}
