package orderbook

import "github.com/crypto-bank/proto/order"

// NewEvent - Creates new event structure containing event.
func NewEvent(ev interface{}) *Event {
	switch event := ev.(type) {
	case *order.Order:
		return &Event{Event: &Event_Order{Order: event}}
	case *order.Trade:
		return &Event{Event: &Event_Trade{Trade: event}}
	default:
		panic("invalid event type")
	}
}

// Inner - Returns inner event.
func (event *Event) Inner() interface{} {
	switch ev := event.Event.(type) {
	case *Event_Order:
		return ev.Order
	case *Event_Trade:
		return ev.Trade
	default:
		panic("invalid event type")
	}
}
