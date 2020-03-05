package main

type notifier interface {
	notify()
}

type user struct {
	name string
}

func (u *user) notify()  {
	
}

func sendNotification(n notifier)  {
	n.notify()
}

func main1()  {
	
	usr := user{name: "Bot"}

	sendNotification(&usr)
}  