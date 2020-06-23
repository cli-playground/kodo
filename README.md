## Kodo - Command line interface for OpenShift
Kodo is a command line interface to build and deploy applications on Openshift.

## Setup
- `git clone https://github.com/cli-playground/kodo.git` 
- `cd kodo`
- To build run,  `go build -o bin/kodo main.go`

## Contributing
 1. Fork from https://github.com/cli-playground/kodo.git
 2. Create Your feature branch (`git checkout -b feature/fooBar`)
 3. Commit your changes (`git commit -am 'Add some fooBar'`)
 4. Push to the branch (`git push origin feature/fooBar`)
 5. Before creating pull request, Make sure to pull all upstream latest changes
 5. Create a new pull request

## Dependencies
 - You can refer go.mod for list of dependencies.

## Commands List

Before running commands, Set Path to $HOME/kodo/bin/.

1. count : Command to count resources like pods running on cluster\
    `kodo count <resources>` \
    For example, `kodo count pods --server=<server url> --token=<token>` will output count of running pods of a given cluster.

2. deploy : Command to deploy an image \
    `kodo deploy --image=<image> --replicas=<no of replicas> --port=<port number> --token=<token> --server=<cluster url>  --namespace=<namespace>` 

3. build: Command to create new BuildConfig and ImageStream from Dockerfile in github repo \
    `kodo build --source=<github url> --namespace=<namespace> 
 --token=<Token>--server=<cluster url>` 

4. help: Command to help user to list all available commands and flags\
    `kodo help`

5. version : Command to check current version \
    `kodo version`
