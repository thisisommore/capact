# populator register ocf-manifests

Populates the OCF manifests into Neo4j database. It reads manifest from remote or local path, converts them into JSON and uploads to database.

## Prerequisites

- [Go](https://golang.org)
- Running Kubernetes cluster with Capact installed

## Build

To build the binary install required [Prerequisites](https://capact.io/community/development/development-guide/#prerequisites) and run:

```shell
make build-tool-populator
```

It creates a binary for your platform in the `bin` directory. For example, for Linux systems, it is `bin/populator-linux-amd64`.

## Usage

> **CAUTION:**  In order to run DB populator manually, make sure the populator inside development cluster is disabled.
> To disable it, run `ENABLE_POPULATOR=false make dev-cluster-update`

It requires one argument, which is a path to directory with Hub manifests. Internally it uses [go-getter](https://github.com/hashicorp/go-getter) so it can download manifests from different locations and in different formats.

To be able to use it locally when Capact is running in a Kubernetes cluster, two ports need to
be forwarded:

```shell
kubectl -n capact-system port-forward svc/neo4j-neo4j 7687:7687
kubectl -n capact-system port-forward svc/neo4j-neo4j 7474:7474
```

To run it and use manifests, for example from the [`hub-manifests`](https://github.com/capactio/hub-manifests) repo, run:

```shell
./bin/populator-linux-amd64 register ocf-manifests {PATH_TO_THE_MAIN_DIRECTORY_OF_THE_REPO}
```

To use manifests from private git repo, private key, encoded in base64 format, is needed.
For example command to download manifests from Capact repo would look like this:
```shell
export SSHKEY=`base64 -w0 ~/.ssh/id_rsa`
./populator register ocf-manifests git@github.com:capactio/capact.git?sshkey=$SSHKEY
```

For better performance populator starts HTTP server to serve manifests converted to JSON files.
Neo4j needs access to this JSON files. `APP_JSON_PUBLISH_ADDR` environment variable should be set
so populator can send a correct link to a Neo4j:

```shell
APP_JSON_PUBLISH_ADDR=http://{HOST_IP} ./populator register ocf-manifests .
```
Replace `HOST_IP` with your computer IP

## Configuration

You can set the following environment variables to configure the Hub database populator:

| Name                                | Required | Default   | Description                                                                                                                                                           |
| ----------------------------------- | -------- | --------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| APP_NEO4J_ADDR                       | no       | `neo4j://localhost:7687` | Neo4j address                                                                                                                                         |
| APP_NEO4J_USER                       | no       | `neo4j`                  | Neo4j admin user                                                                                                                                      |
| APP_NEO4J_PASSWORD                   | yes      |                          | Neo4h admin password                                                                                                                                  |
| APP_JSON_PUBLISH_ADDR                | yes      |                          | Address on which populator will serve JSON files                                                                                                      |
| APP_JSON_PUBLISH_PORT                | no       | `8080`                   | Port number on which populator will be listening                                                                                                      |
| APP_MANIFESTS_PATH                   | no       | `manifests`            | Path to a directory in a repository where manifests are stored                                                                                        |
| APP_UPDATE_ON_GIT_COMMIT        | no       | `false`                  | Flag to make populator populate data only when there are new changes in a repository                                                                  |
| APP_LOGGER_DEV_MODE                  | no       | `false`                  | Enable development mode logging                                                                                                                       |
