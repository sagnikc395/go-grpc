
- direct client server not scalable using request response cycles, synchronousn, slow ,waits and not scalable.

- RPC:
  - Directly call functions from client to server
  - instead of JSON, we use protobuf as message interchange format , 
  - size of payload is small and that accelartes communication

-  Framework:
   -  Streaming 
   -  Continuous flow of data 
   -  async and fast 
   -  extreme scalability -> not a lot of network or any blocking requests.
   -  use cases -> micoservices !, blockchains -> peers interactions 
  
- 4 types of communication:
  - Fully Duplex -> Traditional requests and responses 
  - Server streaming -> client sends a request to the server and gets a stream back. Its a sequence of messages.
  - Client Streaming -> client sends a stream of messages to the server and the server just sends a response.
  - Bi-directional Streaming -> client and server both communicate using streams. even though its a stream , not a queue, the order of the messages is preserved. Both can do it in parallel.

