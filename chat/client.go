package main

import (
	"github.com/gorilla/websocket"
)

/*
socket will hold a reference to the web socket that will allow us
to communicate with the client, and  send  is a buffered
channel through which received messages are queued ready to be
forwarded to the user's browser (via the socket).
The room  will keep a reference to the room that the client
is chatting so  we can forward messages to everyone in the room.
*/

// client represents a single chatting user.
type client struct {
	// socket is the web socket for this client.
	socket *websocket.Conn
	// send is a channel on which messages are sent.
	send chan []byte
	// room is the room this client is chatting in.
	room *room
}

/*
The read method allows our client to read from the socket via the ReadMessage
method, continually sending any received messages to the forward channel on
the room type. If it encounters an error (such as 'the socket has died'),
the loop will break and the socket will be closed.
*/

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

/*
Similarly, the write method continually accepts messages from the send channel
writing everything out of the socket via the WriteMessage method. If writing to
the socket fails, the for loop is broken and the socket is closed.
*/

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
