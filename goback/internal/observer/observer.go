package observer

type Topic string

type Event struct {
	Topic   Topic
	Payload interface{}
}

type Observer interface {
	AddPublisherChan(ch chan Event)
	GetChanByTopics(topics []Topic) <-chan Event
}

type Obs struct {
	PublisherChans []<-chan Event
	ChanTopics     map[chan Event][]Topic
}

func NewObs() *Obs {
	return &Obs{}
}

func (o *Obs) AddPublisherChain(ch <-chan Event) {
	o.PublisherChans = append(o.PublisherChans, ch)
}

func (o *Obs) GetChanByTopics(topics []Topic, buffer int) <-chan Event {
	ch := make(chan Event, buffer)
	o.ChanTopics[ch] = topics
	return ch
}

func merge(buffer int, cs ...<-chan Event) <-chan Event {
	out := make(chan Event, buffer)
	for _, c := range cs {
		go func(c <-chan Event) {
			for v := range c {
				out <- v
			}
		}(c)
	}
	return out
}

func (o *Obs) Start() {
	buf := 0
	for c := range o.ChanTopics {
		buf += cap(c)
	}
	pubChan := merge(buf, o.PublisherChans...)
	for ch, topics := range o.ChanTopics {
		go func(topics []Topic, ch chan Event) {
			for {
				event := <-pubChan
				for _, topic := range topics {
					if event.Topic == topic {
						ch <- event
					}
				}
			}
		}(topics, ch)
	}
}
