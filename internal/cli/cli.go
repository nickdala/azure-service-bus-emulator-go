package cli

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/spf13/cobra"
)

const (
	queueName = "queue.1"
)

func AzureServiceBusSenderClientWrapRunE(
	runEFunc func(cmd *cobra.Command, args []string, sender *azservicebus.Sender) error,
) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		// Silence usage so we don't print the usage when an error occurs
		cmd.SilenceUsage = true

		client, err := newAzureServiceBusClient()
		if err != nil {
			fmt.Printf("Error creating service bus client: %v\n", err)
			return err
		}

		sender, err := client.NewSender(queueName, nil)
		if err != nil {
			fmt.Printf("Error creating sender: %v\n", err)
			return err
		}
		defer sender.Close(context.TODO())

		return runEFunc(cmd, args, sender)
	}
}

func AzureServiceBusReceiverClientWrapRunE(
	runEFunc func(cmd *cobra.Command, args []string, sender *azservicebus.Receiver) error,
) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		// Silence usage so we don't print the usage when an error occurs
		cmd.SilenceUsage = true

		client, err := newAzureServiceBusClient()
		if err != nil {
			fmt.Printf("Error creating service bus client: %v\n", err)
			return err
		}

		receiver, err := client.NewReceiverForQueue(queueName, &azservicebus.ReceiverOptions{
			ReceiveMode: azservicebus.ReceiveModePeekLock,
		})
		if err != nil {
			fmt.Printf("Error creating receiver: %v\n", err)
			return err
		}
		defer receiver.Close(context.TODO())

		return runEFunc(cmd, args, receiver)
	}
}

func newAzureServiceBusClient() (*azservicebus.Client, error) {
	connectionString := "Endpoint=sb://localhost;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=SAS_KEY_VALUE;UseDevelopmentEmulator=true;"
	return azservicebus.NewClientFromConnectionString(connectionString, nil)
}
