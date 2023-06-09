# Nostr

Nostr command line tool and relay setup.

NOTE: Use the [crypto](https://github.com/rubenvanstaden/crypto) package to generate a key exchange.

## Client CLI

- Test with Damus relay

```shell
# Set as environment variables
export NSEC="nsec..."
export NPUB="npub..."
export RELAY="relay.damus.io:443"
export RELAY_CERT=cert/relay.damus.io.pem
```

```shell
# Post a note from the terminal.
./bin/ncli --relay $RELAY note "fsociety"

# Request a set of notes based on the filter configuration.
./bin/ncli --relay $RELAY req 001 ~/.config/noztr/config.json
```

## Run Relay

```shell
// Set environment variables
export RELAY="127.0.0.1:8080"
export REPOSITORY="mongodb://127.0.0.1:27017"

// Spin up docker container
make up

// Spin up the relay service using websocket.
make relay
```

