package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
We need a way for clients to join and leave rooms in order to ensure
that the c.room.forward <- msg in client code forwards the message
to all the clients.  To ensure that we are not trying to access the
same data at the same time, a sensible approach is to use two channels:
one that will add a client to the room and another that will remove it.
*/

type room struct {
	// forward is a channel that holds incoming messages
	// that should be forwarded to the other clients.
	forward chan []byte
	// join is a channel for clients wishing to join the room.
	join chan *client
	// leave is a channel for clients wishing to leave the room.
	leave chan *client
	// clients holds all current clients in this room.
	clients map[*client]bool
}

// newRoom makes a new room that is ready to go.
func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

/*
select statements whenever we need to synchronize or modify shared memory,
or take different actions depending on the various activities within our channels.
*/

func (r *room) run() {
	for { // runs forever
		//  Go routine, will run in the background,
		// which won't block the rest of the application.
		select {
		// the select statement will run the code for a particular case.
		// it will only run one block of case code at a time.
		//  ensure that our r.clients map is only ever modified by one thing at a time.
		case client := <-r.join:
			// joining
			r.clients[client] = true
		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward message to all clients
			for client := range r.clients {
				select {
				case client.send <- msg:
					// send the message
				default:
					// failed to send
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}

// Turning a room into an HTTP handler

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

// In order to use web sockets, we must upgrade the HTTP connection
//using the websocket.Upgrader type, which is reusable so we need only create one.

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

//  ServeHTTP method means a room can now act as a handler

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// we get the socket by calling the upgrader.Upgrade metho
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client

	// will ensure everything is tidied up after a user goes away.
	defer func() { r.leave <- client }()
	// The write method for the client is then called as a Go routine
	go client.write()
	// Finally, we call the read method in the main thread, which will block operations
	// (keeping the connection alive) until it's time to close it.
	client.read()
}
