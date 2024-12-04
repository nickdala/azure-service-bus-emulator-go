package cmd

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	petv1 "github.com/nickdala/azure-service-bus-emulator-go/gen/pet/v1"
	"github.com/nickdala/azure-service-bus-emulator-go/internal/cli"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

// consumeCmd represents the consume command
var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "consume receives messages from Azure Service Bus Emulator",
	Long: `consume receives messages from Azure Service Bus Emulator. For example:
	azure-service-bus-emulator-go consume`,
	RunE: cli.AzureServiceBusReceiverClientWrapRunE(consumerCommand),
}

func consumerCommand(cmd *cobra.Command, args []string, receiver *azservicebus.Receiver) error {
	fmt.Println("consume called")

	// Check if there are messages in the queue
	peeked, err := receiver.PeekMessages(context.TODO(), 5, nil)
	if err != nil {
		fmt.Printf("Error peeking messages: %v\n", err)
		return err
	}

	fmt.Printf("There are at least %d messages in the queue\n", len(peeked))

	if len(peeked) == 0 {
		fmt.Println("No messages in the queue. Exiting...")
		return nil
	}

	messages, err := receiver.ReceiveMessages(context.TODO(), 1, nil)
	if err != nil {
		fmt.Printf("Error receiving messages: %v\n", err)
		return err
	}

	for _, message := range messages {
		body := message.Body

		// Unmarshal the message body into Pet
		pet := &petv1.Pet{}
		err := proto.Unmarshal(body, pet)
		if err != nil {
			fmt.Printf("Error unmarshalling message body: %v\n", err)
			return err
		}

		fmt.Printf("Pet type: %s Pet name: %s\n", pet.PetType.String(), pet.Name)

		err = receiver.CompleteMessage(context.TODO(), message, nil)
		if err != nil {
			fmt.Printf("Error completing message: %v\n", err)
			return err
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(consumeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// consumeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consumeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
