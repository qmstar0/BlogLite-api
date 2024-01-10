package initCmd

import (
	"blog/apps/commandResult"
	"fmt"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
	"github.com/ThreeDotsLabs/watermill/message"
	"time"
)

var (
	backendDuration = time.Second * 10
)

func NewPubSubBackend(pubsub PubSub) requestreply.Backend[commandResult.StateCode] {
	backend, err := requestreply.NewPubSubBackend[commandResult.StateCode](
		requestreply.PubSubBackendConfig{
			Publisher: pubsub,
			SubscriberConstructor: func(params requestreply.PubSubBackendSubscribeParams) (message.Subscriber, error) {
				return pubsub, nil
			},
			GeneratePublishTopic: func(params requestreply.PubSubBackendPublishParams) (string, error) {
				s := fmt.Sprintf("%T.%s", params.Command, params.OperationID)
				return s, nil
			},
			GenerateSubscribeTopic: func(params requestreply.PubSubBackendSubscribeParams) (string, error) {
				s := fmt.Sprintf("%T.%s", params.Command, params.OperationID)
				return s, nil
			},
			Logger:                nil,
			ListenForReplyTimeout: &backendDuration,
			ModifyNotificationMessage: func(msg *message.Message, params requestreply.PubSubBackendOnCommandProcessedParams) error {
				return nil
			},
			OnListenForReplyFinished: nil,
			AckCommandErrors:         false,
		}, requestreply.BackendPubsubJSONMarshaler[commandResult.StateCode]{})
	if err != nil {
		panic(err)
	}
	return backend
}
