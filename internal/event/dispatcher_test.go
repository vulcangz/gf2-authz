package event

import (
	"sync"
	"testing"
	lib_time "time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
)

func TestNewDispatcher(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	clock := ctime.NewMockClock(ctrl)

	// When
	dispatcherInstance := NewDispatcher(10, clock)

	// Then
	assert.IsType(t, new(dispatcher), dispatcherInstance)
	assert.Equal(t, 10, dispatcherInstance.eventChanSize)
	assert.Equal(t, clock, dispatcherInstance.clock)
	assert.Equal(t, new(sync.Map), dispatcherInstance.subscribers)
}

func TestDispatcher_Dispatch_WhenNoSubscriber(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	clock := ctime.NewMockClock(ctrl)

	data := "my-data"

	dispatcher := NewDispatcher(10, clock)

	// When
	result := dispatcher.Dispatch(EventTypePrincipal, data)

	// Then
	assert.Nil(t, result)
}

func TestDispatcher_Dispatch_WhenSubscriber(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	date := lib_time.Date(2023, lib_time.January, 1, 0, 0, 0, 0, lib_time.UTC)

	clock := ctime.NewMockClock(ctrl)
	clock.EXPECT().Now().Return(date)

	data := "my-data"

	dispatcher := NewDispatcher(10, clock)

	// When
	_ = dispatcher.Subscribe(EventTypePrincipal)
	result := dispatcher.Dispatch(EventTypePrincipal, data)

	// Then
	assert.Nil(t, result)
}

func TestDispatcher_Subscribe(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	clock := ctime.NewMockClock(ctrl)

	dispatcher := NewDispatcher(10, clock)

	// When
	eventChan := dispatcher.Subscribe(EventTypePrincipal)

	// Then
	assert.IsType(t, make(chan *Event), eventChan)
}

func TestDispatcher_Subscribe_DispatchEvent(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	date := lib_time.Date(2023, lib_time.January, 1, 0, 0, 0, 0, lib_time.UTC)

	clock := ctime.NewMockClock(ctrl)
	clock.EXPECT().Now().Return(date)

	dispatcher := NewDispatcher(10, clock)

	// When
	eventChan := dispatcher.Subscribe(EventTypePrincipal)
	err := dispatcher.Dispatch(EventTypePrincipal, "my-data")
	close(eventChan)

	value := <-eventChan

	// Then
	assert.IsType(t, make(chan *Event), eventChan)
	assert.Nil(t, err)
	assert.Equal(t, &Event{
		Timestamp: 1672531200,
		Data:      "my-data",
	}, value)
}

func TestDispatcher_Unsubscribe_WhenNoSubscriber(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	clock := ctime.NewMockClock(ctrl)

	dispatcher := NewDispatcher(10, clock)

	// When
	eventChan := make(chan *Event)
	err := dispatcher.Unsubscribe(EventTypePrincipal, eventChan)

	// Then
	assert.Equal(t, ErrNoSubscriberForEventType, err)
}

func TestDispatcher_Unsubscribe_WhenHaveSubscriber(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	clock := ctime.NewMockClock(ctrl)

	dispatcher := NewDispatcher(10, clock)

	// When
	eventChan := dispatcher.Subscribe(EventTypePrincipal)
	err := dispatcher.Unsubscribe(EventTypePrincipal, eventChan)

	// Then
	assert.Nil(t, err)
}
