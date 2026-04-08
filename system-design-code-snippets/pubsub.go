package system_design

import "sync"

type Topic struct {
    mu          sync.RWMutex
    subscribers map[string]chan Message
}

type Broker struct {
    topics map[string]*Topic
    mu     sync.RWMutex
}

type Message struct {
    Data any
}

func (b *Broker) Publish(topic string, msg Message) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if _, ok := b.topics[topic]; !ok {
		return
	}
	for _, ch := range b.topics[topic].subscribers {
		select {
			case ch <- msg:
		default:
		}
	}
}

func (b *Broker) Subscribe (topic, subscriber string, ch chan Message) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if _, ok := b.topics[topic]; !ok {
		b.topics[topic] = &Topic{
			subscribers: make(map[string]chan Message),
		}
	}
	b.topics[topic].subscribers[subscriber] = ch
}

func (b *Broker) Unsubscribe(topic, subscriber string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if _, ok := b.topics[topic]; !ok {
		return
	}
	delete(b.topics[topic].subscribers, subscriber)
}
