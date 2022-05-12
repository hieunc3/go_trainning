package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

/*
	Keep information of the user such as name, tcp connection, current room
*/
const (
	cmd_log_name   = "/nick"
	cmd_join_room  = "/join"
	cmd_enter_room = "/rooms"
	cmd_send_mes   = "/msg"
	cmd_quit_chat  = "/quit"
)

type client struct {
	conn     net.Conn
	nick     string
	room     *room
	commands chan<- command
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		//remove new line in the message
		msg = strings.Trim(msg, "\r\n")

		//get list arguments input
		args := strings.Split(msg, " ")

		//get the cmd to navigate
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case cmd_log_name:
			c.commands <- command{
				id:     CMD_NICK,
				client: c,
				args:   args,
			}
		case cmd_join_room:
			c.commands <- command{
				id:     CMD_JOIN,
				client: c,
				args:   args,
			}
		case cmd_enter_room:
			c.commands <- command{
				id:     CMD_ROOMS,
				client: c,
				args:   args,
			}
		case cmd_send_mes:
			c.commands <- command{
				id:     CMD_MSG,
				client: c,
				args:   args,
			}
		case cmd_quit_chat:
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
				args:   args,
			}
		default:
			c.err(fmt.Errorf("unknown command: %s", cmd))
		}
	}
}

//send mes err to the client
func (c *client) err(err error) {
	c.conn.Write([]byte("Err: " + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	c.conn.Write([]byte(">  " + msg + "\n"))
}
