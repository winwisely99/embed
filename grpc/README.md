# GRPC

There are ways to use golang and GRPC for IPC, Server and other topologies.

And we need GRPC to make the Clients agnostic !

So it can be used to make using a P2P network much easier.


https://github.com/fullstorydev/grpchan

&

https://github.com/jhump/goprotoc

Developer Dashboard
https://github.com/fullstorydev/grpcui
- Allows you to quickly see the effects and play with an API
- This will allow Flutter Web GRPC to work as an alternative to GRPC gateway

Examples: https://godoc.org/github.com/jhump/protoreflect/grpcreflect?importers

yarpc - message passing platform
https://github.com/yarpc/yarpc-go/tree/dev/peer/peerlist
- DOCS: https://godoc.org/go.uber.org/yarpc
- TOOLS: https://github.com/yarpc/yab
	- so you can call them from the command line using protoreflect

go-libp2p-grpc
https://github.com/paralin/go-libp2p-grpc
- nicely wraps libp2p
- Example:
	- https://github.com/RTradeLtd/go-libp2p-pubsub-grpc
	- https://github.com/RTradeLtd/grpc
