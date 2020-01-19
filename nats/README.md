# nats

No LB. uses NATS native one
All Clients and Servers are 100% physically not connected.

Can we abstract GPRC or FBS Schema to run over NATS. Just like the GRPC over LibP2P is ?
- HUGE WIN ?
	- MicroServices are now PUB SUB against a CDC Database. Removes a huge class of code.
	- Clients are now PUB SUB for P2P and C2S. Only one way to do things and less code and cognitive overload.
	- C2S have no LB, can hit any global server ( but will be sent to the nearest )
		- Servers can be anywhere ( at home on a rasp pi)
	- Relay Server ( when P2P fails) is very simple now as its just pub sub everywhere.
		- Can use QUIC, HTTP Pull or Push, or NATS client.

Liftbridge is changing to FLatBuffers !
- We just got a free lunch on the FLutter / Golang IPC ! Mobiles will be much faster
- GRPC (protobuf) and FLatBuffers using ONYL a GRPC IDL too. SO easy peasy for us and users
- Might even make it easier to use for P2P over Webrtc / QUIC for PUB SUB of data.
- NRPC is NOW redundant ?
- FlatBuffers
	- Has concept of Table, so may well to DB concepts and Materialised Views (CDC output) concepts
- example:
	- https://github.com/tsingson/goums
		- Istio doing AAA
		- its nice...

https://github.com/nats-rpc/nrpc
- Turns a standard PB into a GRPC RPC for Bootstrap queries and then change to nrpc to catchup 
- PRO
	- Model as a Protobuf exactly like you do with normal GRPC, but get PUB SUB for free.
	- Strong typing.
	- Protobuf directly into HIVE
	- LiftBridge sort of does the same, but does NOT give strongly typed code from your own Protbuf. BIG DIFFERENCE
		- For generic PUB SUB uses by Users, LiftBridge is good though.
	- NUID gives unique IDS, globally. Great for PUB SUB correlation.
- Mobile and Desktop can embed the NATS golang so all fine
- Web will need a NATS go-client. Talers ?
	- https://github.com/c16a/nats-dart
		- NOT at all finished
	- https://github.com/kranfix/dart-nuid
		- also half arsed
	- https://github.com/chartchuo/dart-nats
		- NICE !!
		
Between FLutter and Gomobile layer we can also then use nrpc since its just simple protocol

