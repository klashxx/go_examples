package main

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
