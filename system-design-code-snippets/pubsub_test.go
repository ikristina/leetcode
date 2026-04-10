package system_design

import (
	"testing"
	"time"
)

func TestPubSub_Basic(t *testing.T) {
	broker := &Broker{
		topics: make(map[string]*Topic),
	}

	ch := make(chan Message, 1)
	topic := "news"
	subscriber := "sub1"

	broker.Subscribe(topic, subscriber, ch)

	msg := Message{Data: "hello"}
	broker.Publish(topic, msg)

	select {
	case received := <-ch:
		if received.Data != "hello" {
			t.Errorf("Expected message data 'hello', got %v", received.Data)
		}
	case <-time.After(time.Second):
		t.Error("Timed out waiting for message")
	}
}

func TestPubSub_NonBlockingSend(t *testing.T) {
	broker := &Broker{
		topics: make(map[string]*Topic),
	}

	// Create a channel with capacity 1
	ch := make(chan Message, 1)
	topic := "alerts"
	subscriber := "sub1"

	broker.Subscribe(topic, subscriber, ch)

	// Publish two messages. The first will fit in the channel.
	// The second would normally block, but our Publish is non-blocking and drops it.
	broker.Publish(topic, Message{Data: "first"})
	broker.Publish(topic, Message{Data: "second"})

	// Verify we got the first message
	select {
	case received := <-ch:
		if received.Data != "first" {
			t.Errorf("Expected 'first', got %v", received.Data)
		}
	default:
		t.Error("Expected a message, got none")
	}

	// The channel should now be empty because the second message was dropped
	select {
	case received := <-ch:
		t.Errorf("Did not expect second message to be queued, got %v", received.Data)
	default:
		// Success! The channel is correctly empty
	}
}

func TestPubSub_Unsubscribe(t *testing.T) {
	broker := &Broker{
		topics: make(map[string]*Topic),
	}

	ch := make(chan Message, 1)
	topic := "updates"
	subscriber := "sub1"

	broker.Subscribe(topic, subscriber, ch)
	broker.Unsubscribe(topic, subscriber)

	broker.Publish(topic, Message{Data: "hello"})

	select {
	case <-ch:
		t.Error("Received message after unsubscribing")
	default:
		// Success! We shouldn't receive anything
	}
}
