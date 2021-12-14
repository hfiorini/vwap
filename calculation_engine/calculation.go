package calculation_engine

import (
	"log"
	"sync"
)

type Subscription interface {
	AddItem(item ItemSubscription)
	PrintVWAP()
}

type subscription struct {
	locker         sync.Mutex
	Items          []ItemSubscription
	MaxSize        int
	PriceVolumeMap map[string]float64
	VolumeMap      map[string]float64
	VWAPMap        map[string]float64
}

func NewSubscription(maxSize int) Subscription {

	return &subscription{
		Items:          []ItemSubscription{},
		MaxSize:        maxSize,
		PriceVolumeMap: make(map[string]float64),
		VolumeMap:      make(map[string]float64),
		VWAPMap:        make(map[string]float64),
	}
}

func (s *subscription) AddItem(item ItemSubscription) {
	s.locker.Lock()
	defer s.locker.Unlock()

	if len(s.Items) == s.MaxSize {
		s.handleFullQueue()
	}

	productID := item.ProductID
	volume := item.Volume
	price := item.Price

	_, found := s.VWAPMap[productID]

	if !found {
		priceVolume := price * volume

		s.PriceVolumeMap[productID] = priceVolume
		s.VolumeMap[productID] = volume
		s.VWAPMap[productID] = priceVolume / volume
	} else {
		s.PriceVolumeMap[productID] = s.VWAPMap[productID] + (price * volume)
		s.VolumeMap[productID] = s.VolumeMap[productID] + volume
		s.VWAPMap[productID] = s.PriceVolumeMap[productID] / s.VolumeMap[productID]
	}

	s.Items = append(s.Items, item)
}

func (s *subscription) handleFullQueue() {
	firstItem := s.Items[0]
	s.Items = s.Items[1:]

	productID := firstItem.ProductID
	volume := firstItem.Volume
	price := firstItem.Price

	s.PriceVolumeMap[productID] = s.PriceVolumeMap[productID] - (price * volume)
	s.VolumeMap[productID] = s.VolumeMap[productID] - volume
	if s.VolumeMap[productID] != 0 {
		s.VWAPMap[productID] = s.PriceVolumeMap[productID] / s.VolumeMap[productID]
	}
}

func (s *subscription) PrintVWAP() {
	for key, value := range s.VWAPMap {
		log.Printf("%s : %v", key, value)
	}
	log.Print("*********************")
}
