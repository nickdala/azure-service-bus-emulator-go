package cmd

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	petv1 "github.com/nickdala/azure-service-bus-emulator-go/gen/pet/v1"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

// consumeCmd represents the consume command
var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("consume called")

		connectionString := "Endpoint=sb://localhost;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=SAS_KEY_VALUE;UseDevelopmentEmulator=true;"
		client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
		if err != nil {
			fmt.Printf("Error creating service bus client: %v\n", err)
			return err
		}

		queueName := "queue.1"

		receiver, err := client.NewReceiverForQueue(queueName, &azservicebus.ReceiverOptions{
			ReceiveMode: azservicebus.ReceiveModePeekLock,
		})
		if err != nil {
			fmt.Printf("Error creating receiver: %v\n", err)
			return err
		}
		defer receiver.Close(context.TODO())

		// Check if there are messages in the queue
		peeked, err := receiver.PeekMessages(context.TODO(), 1, nil)
		if err != nil {
			fmt.Printf("Error peeking messages: %v\n", err)
			return err
		}

		fmt.Printf("The number of messages in the queue is %d\n", len(peeked))

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
	},
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
