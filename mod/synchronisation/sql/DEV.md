# sql replication and sync

A CRDT architecture is what a P2P system is due tothe disconnected PUB SUB nature of it.

Materialised views are the READ side from which a SubScribe queue listens on top.
The problem is that you cant disentangle the SQL Query Predicates to know which Materialised View is effected

This lib allows that because you can SQL parse the Query and work out the Views effected.

When a mutation hits the CDC picks it up, and the Materialised VIews get updated

https://github.com/tomarrell/lbadd
