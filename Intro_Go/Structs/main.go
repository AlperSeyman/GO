package main

type user struct {
	name   string
	number int
}

type messageTosend struct {
	message   string
	sender    user
	recipient user
}

func canSendMessage(mToSend messageTosend) bool {

	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}
	return true
}

func main() {

}
