# mindmeld

A is a toy project for port-forwarding between machines on the internet.  It works by passing traffic through an intermediary server which is reachable by all parties.

It uses gRPC streaming to pass control messages and proxy traffic, which means that the intermediary server can run in GCP Cloud Run.

## Motivation

In pre-COVID times we would often port-forward between machines  to quickly test new servers as they were being developed.  When everyone was in the same room this was easy!  Working remotely has made this a lot more complex...

Although there are a lot of existing solutions out there for this, seemed like an interesting problem to tackle - especially from the point of view of trying to make it as simple as possible, and practical to run (i.e. not too expensive).

## Design

There are two main pieces which make the system work:
* `router`: the intermediary process (see `cmd/crrouter`, named as it is the Cloud Run variant of the original `mmrouter`).
* `client` the CLI tool run by users to create/use forwards (see `cmd/mmclient`).

### Creating a service

Users share a local forward on their machine by registering a `service`.

1. Client process calls `CreateService` which creates the service in `router`, and waits for forward requests to come through.
1. When a forward request arrives on the `router` it responds to the `client` via the response stream waiting on the `CreateService` request, signalling it to create a new proxying connection to handle the traffic for the new forward.  The connection is given a `token` which identifies it when it is received by the `router`.
1. When the proxying request arrives at the `router`, it matches it to the forward (using the `token`) and bridges the two connections.

Users access a `service` by:

1. Client exposes a local port, which has a TCP listener.  Connections made to this listener will be forwarded to the service.
1. For each new connection, the client process calls `ForwardToService` which checks the service still exists, and returns a `token` to identify the proxying connection.
2. Client creates a proxying connection, using the provided `token`, and begins to copy data between the local connection and the proxying connection.

## Emulating net.Conn with gRPC

The code was initially designed so that a separate TCP server would run on the router and host the proxy connections. Though easier to debug, this meant it couldn't be used in Cloud Run.

Since Cloud Run was updated to support gRPC bi-directional streaming, it was suddenly a possibility!

The proxying connections are now handled via `internal/protoproxy` which creates implementations of `net.Conn` and `net.Listener` to translate calls to `Write` into stream sends, calls to `Read` into stream receives, and `Accept` into new connections.

The tricky bit is correctly handling when a connection closes, and so there are likely some lingering bugs here.
