package main

type commandID int

const (
	CMD_NICK commandID = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

//example command: /nick john
/*
	nick command have id 0
	client is the client access to chat app
	args: []string{"/nick", "john"}
*/
type command struct {
	id     commandID
	client *client
	args   []string
}
