package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	petv1 "github.com/nickdala/azure-service-bus-emulator-go/gen/pet/v1"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

const (
	CAT     = "cat"
	DOG     = "dog"
	SNAKE   = "snake"
	HAMSTER = "hamster"
)

func SupportedPets() []string {
	return []string{CAT, DOG, SNAKE, HAMSTER}
}

func IsSupportedPet(pet string) bool {
	for _, supportedPet := range SupportedPets() {
		if supportedPet == pet {
			return true
		}
	}
	return false
}

// produceCmd represents the produce command
var produceCmd = &cobra.Command{
	Use:   "produce",
	Short: "produce generates messages to Azure Service Bus Emulator",
	Long:  `produce generates messages to Azure Service Bus Emulator. For example:`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("produce called")

		petType, err := cmd.Flags().GetString("pet")
		if err != nil {
			fmt.Printf("Error getting pet flag: %v\n", err)
			return err
		}

		if !IsSupportedPet(petType) {
			fmt.Printf("Unsupported pet: %s\n", petType)
			return fmt.Errorf("unsupported pet: %s", petType)
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Printf("Error getting name flag: %v\n", err)
			return err
		}

		pet := petv1.Pet{
			//PetType: petv1.PetType_PET_TYPE_CAT,
			Name: name,
		}

		switch petType {
		case CAT:
			pet.PetType = petv1.PetType_PET_TYPE_CAT
		case DOG:
			pet.PetType = petv1.PetType_PET_TYPE_DOG
		case SNAKE:
			pet.PetType = petv1.PetType_PET_TYPE_SNAKE
		case HAMSTER:
			pet.PetType = petv1.PetType_PET_TYPE_HAMSTER
		default:
			pet.PetType = petv1.PetType_PET_TYPE_UNSPECIFIED
		}

		connectionString := "Endpoint=sb://localhost;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=SAS_KEY_VALUE;UseDevelopmentEmulator=true;"
		client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
		if err != nil {
			fmt.Printf("Error creating service bus client: %v\n", err)
			return err
		}

		queueName := "queue.1"

		sender, err := client.NewSender(queueName, nil)
		if err != nil {
			fmt.Printf("Error creating sender: %v\n", err)
			return err
		}
		defer sender.Close(context.TODO())

		petBytes, err := proto.Marshal(&pet)
		if err != nil {
			fmt.Printf("Error marshalling pet: %v\n", err)
			return err
		}
		sbMessage := &azservicebus.Message{
			Body: []byte(petBytes),
		}
		err = sender.SendMessage(context.TODO(), sbMessage, nil)
		if err != nil {
			fmt.Printf("Error sending message: %v\n", err)
			return err
		}

		fmt.Println("Message sent")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(produceCmd)

	// Pet flag
	produceCmd.Flags().StringP("pet", "p", "",
		fmt.Sprintf("Pet (one of %s)", strings.Join(SupportedPets(), ",")))
	// Required flag
	if err := produceCmd.MarkFlagRequired("pet"); err != nil {
		produceCmd.Printf("Error marking flag as required: %v\n", err)
		os.Exit(-1)
	}

	// Name flag
	produceCmd.Flags().StringP("name", "n", "", "Name of the pet")
	// Required flag
	if err := produceCmd.MarkFlagRequired("name"); err != nil {
		produceCmd.Printf("Error marking flag as required: %v\n", err)
		os.Exit(-1)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// produceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// produceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
