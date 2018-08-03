package main

type EventType int64

type EventHandler func(evt Event)

type GameObject interface {
	Next()
}

type Core struct {
	eq     *EventQueue
	buf    []Event
	hmap   []EventHandler
	static []GameObject
}

func (c *Core) Loop() {
	evts := c.buf[0:0]
	c.eq.Pull(evts)
	for _, evt := range evts {
		c.hmap[evt.Type()](evt)
	}
}

type Event interface {
	Type() EventType
}

type EventQueue struct {
	cap int64
	ch  chan Event
}

func NewEventQueue(cap int64) *EventQueue {
	ch := make(chan Event, cap)
	return &EventQueue{
		cap,
		ch,
	}
}

func (eq *EventQueue) Push(evts []Event) error {
	for _, evt := range evts {
		eq.ch <- evt
	}
	return nil
}
func (eq *EventQueue) Pull(buf []Event) error {

FOROUT:
	for {
		select {
		case evt := <-eq.ch:
			buf = append(buf, evt)
		default:
			break FOROUT
		}
	}
	return nil
}
