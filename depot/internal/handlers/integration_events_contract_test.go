//go:build contract

package handlers

import (
	"context"
	"encoding/json"
	"github.com/owezzy/soko-bora-mngt-system/depot/internal/domain"
	"github.com/owezzy/soko-bora-mngt-system/internal/am"
	"github.com/owezzy/soko-bora-mngt-system/internal/ddd"
	"github.com/owezzy/soko-bora-mngt-system/internal/registry"
	"github.com/owezzy/soko-bora-mngt-system/internal/registry/serdes"
	"github.com/owezzy/soko-bora-mngt-system/stores/storespb"
	"testing"
	"time"

	"github.com/pact-foundation/pact-go/v2/matchers"
	"github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type String = matchers.String
type Map = matchers.Map

var Like = matchers.Like

type rawEvent struct {
	Name    string
	Payload json.RawMessage
}

type incomingMessage struct {
	id         string
	name       string
	subject    string
	data       []byte
	metadata   ddd.Metadata
	sentAt     time.Time
	receivedAt time.Time
}

func (m incomingMessage) ID() string             { return m.id }
func (m incomingMessage) Subject() string        { return m.subject }
func (m incomingMessage) MessageName() string    { return m.name }
func (m incomingMessage) Metadata() ddd.Metadata { return m.metadata }
func (m incomingMessage) SentAt() time.Time      { return m.sentAt }
func (m incomingMessage) ReceivedAt() time.Time  { return m.receivedAt }
func (m incomingMessage) Data() []byte           { return m.data }
func (m incomingMessage) Ack() error             { return nil }
func (m incomingMessage) NAck() error            { return nil }
func (m incomingMessage) Extend() error          { return nil }
func (m incomingMessage) Kill() error            { return nil }

func newIncomingEventMessage(event rawEvent) (am.IncomingMessage, error) {
	data, err := proto.Marshal(&am.EventMessageData{
		Payload:    event.Payload,
		OccurredAt: timestamppb.Now(),
	})
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return incomingMessage{
		id:         event.Name,
		name:       event.Name,
		subject:    "",
		data:       data,
		metadata:   make(ddd.Metadata),
		sentAt:     now,
		receivedAt: now,
	}, nil
}

func TestStoresConsumer(t *testing.T) {
	type mocks struct {
		stores   *domain.MockStoreCacheRepository
		products *domain.MockProductCacheRepository
	}

	reg := registry.New()
	err := storespb.RegistrationsWithSerde(serdes.NewJsonSerde(reg))
	if err != nil {
		t.Fatal(err)
	}

	pact, err := v4.NewAsynchronousPact(v4.Config{
		Provider: "stores-pub",
		Consumer: "depot-sub",
		PactDir:  "./pacts",
	})
	if err != nil {
		t.Fatal(err)
	}

	tests := map[string]struct {
		given   []models.ProviderState
		content Map
		on      func(m mocks)
	}{
		"a StoreCreated message": {
			content: Map{
				"Name": String(storespb.StoreCreatedEvent),
				"Payload": Like(Map{
					"id":       String("store-id"),
					"name":     String("NewStore"),
					"location": String("NewLocation"),
				}),
			},
			on: func(m mocks) {
				m.stores.On("Add", mock.Anything, "store-id", "NewStore", "NewLocation").Return(nil)
			},
		},
		"a StoreRebranded message": {
			content: Map{
				"Name": String(storespb.StoreRebrandedEvent),
				"Payload": Like(Map{
					"id":   String("store-id"),
					"name": String("RebrandedStore"),
				}),
			},
			on: func(m mocks) {
				m.stores.On("Rename", mock.Anything, "store-id", "RebrandedStore").Return(nil)
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			m := mocks{
				stores:   domain.NewMockStoreCacheRepository(t),
				products: domain.NewMockProductCacheRepository(t),
			}
			if tc.on != nil {
				tc.on(m)
			}
			handlers := NewIntegrationEventHandlers(reg, m.stores, m.products)
			msgHandlerFn := func(message v4.AsynchronousMessage) error {
				event := message.Body.(*rawEvent)
				msg, err := newIncomingEventMessage(*event)
				if err != nil {
					return err
				}

				return handlers.HandleMessage(context.Background(), msg)
			}

			message := pact.AddAsynchronousMessage()
			for _, given := range tc.given {
				message = message.GivenWithParameter(given)
			}

			assert.NoError(t, message.
				ExpectsToReceive(name).
				WithJSONContent(tc.content).
				AsType(&rawEvent{}).
				ConsumedBy(msgHandlerFn).
				Verify(t))
		})
	}
}
