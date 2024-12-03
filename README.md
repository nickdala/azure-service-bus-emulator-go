# Azure Service Bus Emulator Example in GO

This repository contains an example of how to use the Azure Service Bus Emulator in Go. The emulator allows you to develop and test your applications locally without needing to connect to the actual Azure Service Bus service.

## Prerequisites

- Docker installed on your machine.
- Go installed on your machine.

**Note:** There is a [Dev Container](https://code.visualstudio.com/docs/remote/containers) that can be used to build the project in a container. This is useful if you don't have Go installed on your machine. You can open the project in Visual Studio Code and select the "Reopen in Container" option. Or you can use [CodeSpaces](https://github.com/features/codespaces) to build the project in the cloud.

![Open in CodeSpaces](./images/codespaces.png)

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/nickdala/azure-service-bus-emulator-go

cd azure-service-bus-emulator-go
```

2. Run the Docker container for the Azure Service Bus Emulator:

```bash
docker compose -f ./docker/docker-compose.yaml up -d
```

**Note:** To change the SQL password, edit the `./docker/.env` file and set *MSSQL_SA_PASSWORD* to a secure password per [Microsoft's documentation](https://learn.microsoft.com/sql/relational-databases/security/strong-passwords?view=sql-server-linux-ver16).

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments

- [Azure SDK for Go](https://github.com/Azure/azure-sdk-for-go) for interacting with Azure services.
- [Azure Service Bus Emulator](https://learn.microsoft.com/en-us/azure/service-bus-messaging/overview-emulator) for local development and testing.

Feel free to fork this repository and make it your own!