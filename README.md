<img src="https://i.imgur.com/e07eZQv.png" align="right" />

# Simple CI/CD tool in GO
![build](https://camo.githubusercontent.com/30ae0cf6825132db112b4208a5776454bf0cc330/68747470733a2f2f73656d6170686f72656170702e636f6d2f6170692f76312f70726f6a656374732f64346363613530362d393962652d343464322d623139652d3137366633366563386366312f3132383530352f736869656c64735f62616467652e737667)    ![version](https://camo.githubusercontent.com/872e8e7b7893bb2335c27be1f7cac90227dfd255/68747470733a2f2f62616467652e667572792e696f2f67682f626f656e6e656d616e6e2532466261646765732e737667) ![love](https://camo.githubusercontent.com/d9ce827af4ec2b7b3c52ce4595bbb354d8b21405/68747470733a2f2f6261646765732e66726170736f66742e636f6d2f6f732f76312f6f70656e2d736f757263652e7376673f763d313032)

## Install
As a go package
```sh
$ go get github.com/jkreshpaj/flaka-ci
```
Or just download the binary
```sh
$ curl -LJO https://github.com/jkreshpaj/flaka-ci/raw/master/flaka-ci
$ chmod 777 flaka-ci
$ ./flaka-ci
```

## Options
```
$ flaka-ci --help
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

testing commands2
