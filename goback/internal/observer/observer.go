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
	log            *zerolog.Logger
	PublisherChans []<-chan Event
	Handlers       []handler
	TopicsHandler  map[int][]Topic
}

func NewObserver(l *zerolog.Logger) *Observer {
	o := Observer{}
	o.log = l
	o.TopicsHandler = make(map[int][]Topic)
	return &o
}

func (o *Observer) AddPublisherChan(ch <-chan Event) {
	o.PublisherChans = append(o.PublisherChans, ch)
}

func (o *Observer) AddTopicsHandler(topics []Topic, handler func(event Event)) {
	o.Handlers = append(o.Handlers, handler)
	o.TopicsHandler[len(o.Handlers)-1] = topics
}

func merge(cs ...<-chan Event) <-chan Event {
	if len(cs) == 1 {
		return cs[0]
	}
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
	o.log.Info().Msg("Observer/merge publishers channel")
	pubChan := merge(o.PublisherChans...)

	o.log.Info().Msg("Observer/handlers: ")
	var met map[chan Event][]Topic
	for n, topics := range o.TopicsHandler {
		h := o.Handlers[n]
		ch := make(chan Event)
		met[ch] = topics
		o.log.Info().Msgf("Observer/go handler for topics: %s", topics)
		go func(handler, chan Event) {
			for {
				event := <-ch
				h(event)
			}
		}(h, ch)
	}

	o.log.Info().Msg("Observer/go events manager")
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
