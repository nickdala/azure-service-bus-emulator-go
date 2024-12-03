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

3. Verify that the Docker container is running:

```bash
docker ps
```

You should see something like the following:

```
vscode ➜ /workspaces/azure-service-bus-emulator-go $ docker ps
CONTAINER ID   IMAGE                                                          COMMAND                  CREATED          STATUS         PORTS                                                 NAMES
de1d403a5e9c   mcr.microsoft.com/azure-messaging/servicebus-emulator:latest   "/ServiceBus_Emulato…"   13 minutes ago   Up 5 seconds   0.0.0.0:5672->5672/tcp, :::5672->5672/tcp, 8080/tcp   servicebus-emulator
da806c13de70   mcr.microsoft.com/azure-sql-edge:latest                        "/opt/mssql/bin/perm…"   13 minutes ago   Up 6 seconds   1401/tcp, 1433/tcp                                    sqledge
```

## Teardown

To stop and remove the Docker container, run:

```bash
docker compose -f ./docker/docker-compose.yaml down
```

To remove any volumes created by the Docker container, run:

```bash
docker compose -f ./docker/docker-compose.yaml down -v
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments

- [Azure SDK for Go](https://github.com/Azure/azure-sdk-for-go) for interacting with Azure services.
- [Azure Service Bus Emulator](https://learn.microsoft.com/en-us/azure/service-bus-messaging/overview-emulator) for local development and testing.

Feel free to fork this repository and make it your own!