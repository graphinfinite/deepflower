package observer

import (
	"github.com/rs/zerolog"
)

type Topic string

type Event struct {
	Topic   Topic
	Payload interface{}
}

type handler func(event Event)

type Observer struct {
	Log            *zerolog.Logger
	PublisherChans []<-chan Event
	Handlers       []handler
	TopicsHandler  map[int][]Topic
}

func NewObserver(l *zerolog.Logger) *Observer {
	return &Observer{Log: l}
}

func (o *Observer) AddPublisherChain(ch <-chan Event) {
	o.PublisherChans = append(o.PublisherChans, ch)
}

func (o *Observer) AddTopicsHandler(topics []Topic, handler func(event Event)) {
	o.Handlers = append(o.Handlers, handler)
	o.TopicsHandler[len(o.Handlers)-1] = topics
}

func merge(cs ...<-chan Event) <-chan Event {
	out := make(chan Event)
	for _, c := range cs {
		go func(c <-chan Event) {
			for v := range c {
				out <- v
			}
		}(c)
	}
	return out
}

func (o *Observer) Start() {
	o.Log.Info().Msg("Observer/merge publishers channel")
	pubChan := merge(o.PublisherChans...)

	o.Log.Info().Msg("Observer/handlers: ")
	var met map[chan Event][]Topic
	for n, topics := range o.TopicsHandler {
		h := o.Handlers[n]
		ch := make(chan Event)
		met[ch] = topics
		o.Log.Info().Msgf("Observer/go handler for topics: %s", topics)
		go func(handler, chan Event) {
			for {
				event := <-ch
				h(event)
			}
		}(h, ch)
	}

	o.Log.Info().Msg("Observer/go events manager")
	go func(map[chan Event][]Topic, <-chan Event) {
		for event := range pubChan {
			for ch, topics := range met {
				for _, topic := range topics {
					if event.Topic == topic {
						ch <- event
					}
				}
			}
		}
	}(met, pubChan)
}
