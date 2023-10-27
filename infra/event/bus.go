package event

import (
	"sync"
)

var BusMap = make(map[string]*Bus)

type DataEvent struct {
	Data  interface{}
	Topic string
}

// DataChan 是一个能接收 DataEvent 的 channel
type DataChan chan DataEvent

// DataChanList 是一个包含 DataChannels 数据的切片
type DataChanList []DataChan

// Bus 存储有关订阅者感兴趣的特定主题的信息
type Bus struct {
	subscribers map[string]DataChanList
	rm          sync.RWMutex
}

// GetBus 按id获取一个事件bus
func GetBus(id string) *Bus {
	if BusMap[id] != nil {
		return BusMap[id]
	}
	BusMap[id] = new(Bus)
	return BusMap[id]
}

func (b *Bus) Publish(topic string, data interface{}) {
	b.rm.RLock()
	if chans, found := b.subscribers[topic]; found {
		channels := append(DataChanList{}, chans...)
		go func(data DataEvent, dataChannelSlices DataChanList) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
	b.rm.RUnlock()
}

func (b *Bus) Subscribe(topic string, ch DataChan) {
	b.rm.Lock()
	if prev, found := b.subscribers[topic]; found {
		b.subscribers[topic] = append(prev, ch)
	} else {
		b.subscribers[topic] = append([]DataChan{}, ch)
	}
	b.rm.Unlock()
}
