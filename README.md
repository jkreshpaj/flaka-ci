<img src="https://i.imgur.com/e07eZQv.png" align="right" />

## Minimalistic - Zero configuration CI/CD
![build](https://camo.githubusercontent.com/30ae0cf6825132db112b4208a5776454bf0cc330/68747470733a2f2f73656d6170686f72656170702e636f6d2f6170692f76312f70726f6a656374732f64346363613530362d393962652d343464322d623139652d3137366633366563386366312f3132383530352f736869656c64735f62616467652e737667)    ![version](https://camo.githubusercontent.com/872e8e7b7893bb2335c27be1f7cac90227dfd255/68747470733a2f2f62616467652e667572792e696f2f67682f626f656e6e656d616e6e2532466261646765732e737667) ![love](https://camo.githubusercontent.com/d9ce827af4ec2b7b3c52ce4595bbb354d8b21405/68747470733a2f2f6261646765732e66726170736f66742e636f6d2f6f732f76312f6f70656e2d736f757263652e7376673f763d313032)

### Install
As a go package
```sh
go get github.com/jkreshpaj/flaka-ci
```
Or just download the binary
```sh
curl -LJO https://github.com/jkreshpaj/flaka-ci/raw/master/build/flaka-ci-<PLATFORM> --output flaka-ci
chmod +x flaka-ci
./flaka-ci
```

### Options
```
flaka-ci --help
Run flaka-ci [arg] to start server

Usage:
  flaka-ci [flags]

Flags:
  -c, --config string   Configuration file (default "flaka-ci.yml")
  -d, --detach          Detached mode runs FlakaCI in background
  -h, --help            help for flaka-ci
  -n, --notify string   Webhook url to send automatic log messages
  -p, --port string     FlakaCI server port (default "7000")
```

### Example

#### Folder structure

```
 my-app
    service-1
    service-2
    service-3
    flaka-ci.yml
```

The only configuration you need to have with FlakaCI is a '.yml' file to specify the services that are in your project and commands to execute after the service is updated.
By default FlakaCI is configured to watch for changes only on master branch.

#### flaka-ci.yml

```
services:
    servicename1:
        path: /service-1
        build: true
    servicename2:
        path: /service-2
        command:
            - echo Command 1 Service 2
    test:
        path: /service-3
        build: true
        command:
            - echo Testing errors
            - ng build //This will throw an error
            - echo Other command
```

In flaka-ci.yml you specify all your services. Every services must be initialized with git and have its own repo. Service name can be whatever you want except when your are using it with docker-compose it must be the same.
Under each service you have the options:
- path: Directory of the service
- build: If you are using this service with docker-compose and you want to rebuild it on update
- command: Commands list to be executed after service is updated and container is rebuild (in case you have it). Commands are executed one by one in the service directory.
#### Starting the server
To start the server you just need to run ```flaka-ci``` and optional flags.
By default FlakaCI runs on port 7000 and in the current process.
If you want to run it in background you can use ```--detach``` or ```-d``` flag.
```
$ flaka-ci -d
FlakaCI started in background.
```

##### Logs
If you are runnig  in background you can see what services are beeing updated and logs of each command bu running ```flaka-ci logs```
Additionally you can tell FlakaCI to send a noitfication using slack. All you need to do is create a simple slack app and when you run ```flaka-ci``` pass your webhook url in ```--notify``` or ```-n``` flag.

```
$ flaka-ci -d -n https://hooks.slack.com/services/AOIE302573/AFSV31095732t/BASG32095302ugv2bo43vbr
```

<img src="https://i.imgur.com/TJ1luLi.png" align="center" />

testing5
