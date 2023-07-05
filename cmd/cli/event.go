package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
	"github.com/rubenvanstaden/crypto"
	"github.com/rubenvanstaden/nostr/core"
)

func NewEvent(cc *websocket.Conn) *Event {

	gc := &Event{
		fs: flag.NewFlagSet("event", flag.ContinueOnError),
        cc: cc,
	}

	gc.fs.StringVar(&gc.note, "note", "", "event text note of Kind 1")
	gc.fs.StringVar(&gc.metadata, "metadata", "", "event text note of Kind 0")
	gc.fs.StringVar(&gc.recommend, "recommend", "", "event text note of Kind 2")

	return gc
}

type Event struct {
	fs *flag.FlagSet
    cc *websocket.Conn

	note string
    metadata string
    recommend string
}

func (g *Event) Name() string {
	return g.fs.Name()
}

func (g *Event) Init(args []string) error {
	return g.fs.Parse(args)
}

func (s *Event) Run() error {

    if s.note != "" {
        s.publish(s.note)
    }

    if s.metadata != "" {
        log.Fatalln("[metadata] not implemented")
    }

    if s.recommend != "" {
        log.Fatalln("[recommend] not implemented")
    }

	return nil
}

func (s *Event) publish(content string) error {
    var msgEvent core.MessageEvent

    msgEvent.Kind = 1

    msgEvent.Tags = nil

    // The note is created now.
    msgEvent.CreatedAt = core.Now()

    // The user note that should be trimmed properly.
    msgEvent.Content = content

    // Apply NIP-19 to decode user-friendly secrets.
    var sk string
    if _, s, e := crypto.DecodeBech32(PRIVATE_KEY); e == nil {
        sk = s.(string)
    }
    if pub, e := crypto.GetPublicKey(sk); e == nil {
        // Set public with which the event wat pushed.
        msgEvent.PubKey = pub
        if npub, e := crypto.EncodePublicKey(pub); e == nil {
            fmt.Fprintln(os.Stderr, "using:", npub)
        }
    }

    // We have to sign last, since the signature is dependent on the event content.
    msgEvent.Sign(sk)

    // Marshal the signed event to a slice of bytes ready for transmission.
    msg, err := json.Marshal(msgEvent)
    if err != nil {
        log.Fatalln("unable to marchal incoming event")
    }

    log.Printf("[\033[32m*\033[0m] Client - Event published (id: %s...)", msgEvent.Id[:10])

    // Transmit event message to the spoke that connects to the relays.
    err = s.cc.WriteMessage(websocket.TextMessage, msg)
    if err != nil {
        return err
    }

    return nil
}