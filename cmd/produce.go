package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	petv1 "github.com/nickdala/azure-service-bus-emulator-go/gen/pet/v1"
	"github.com/nickdala/azure-service-bus-emulator-go/internal/cli"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

const (
	CAT  = "cat"
	DOG  = "dog"
	BIRD = "bird"
	FISH = "fish"
)

func SupportedPets() []string {
	return []string{CAT, DOG, BIRD, FISH}
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
	Long: `produce generates messages to Azure Service Bus Emulator. For example: 
	azure-service-bus-emulator-go produce -p cat -n whiskers`,
	RunE: cli.AzureServiceBusSenderClientWrapRunE(produceCommand),
}

func produceCommand(cmd *cobra.Command, args []string, sender *azservicebus.Sender) error {
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
	case BIRD:
		pet.PetType = petv1.PetType_PET_TYPE_BIRD
	case FISH:
		pet.PetType = petv1.PetType_PET_TYPE_FISH
	default:
		pet.PetType = petv1.PetType_PET_TYPE_UNSPECIFIED
	}

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
