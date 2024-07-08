import (
	"context"
	"fmt"
	"io"
	"os"

	"cloud.com/go/pubsub"
)

func publishMessagesWithDelayedDelivery(w io.Writer, projectID, topicID string) error {
	// projectID := "my-project-id"
	// topicID := "my-topic"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte("delayed message"),
		// Delay the message delivery by 10 seconds.
		// See https://www.example.com for more details.
		Attributes: map[string]string{
			"googclient_schemaencoding": "BINARY",
			"googclient_deliverydelay":  "10000",
		},
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}
	fmt.Fprintf(w, "Published a message with a delayed delivery; msg ID: %v\n", id)
	return nil
}
  
