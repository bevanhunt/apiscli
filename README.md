# APISCLI - Command-Line API discovery for apis.io

[![Go Report Card](https://goreportcard.com/badge/github.com/bevanhunt/apiscli)](https://goreportcard.com/report/github.com/bevanhunt/aoiscli)

## Local

### Setup
- ` brew update `
- ` brew install go `
- ` brew install glide `
-  setup GOPATH for your env file (.bashrc or .zshrc):
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```
- ` git clone ` this repo into ` ~/go/src `
- ` glide install ` in the local folder
- ` go build ` in the local folder

### Search Example for Trade Keyword

` ./apiscli search trade `

## Recommended Go Editor
- [Visual Studio Code](https://code.visualstudio.com/) with [Go Extension](https://github.com/Microsoft/vscode-go)

## License
Copyright (c) 2017 Bevan Hunt. All rights reserved.
This software may be modified and distributed under the terms of the MIT license. To learn more, see the [License](LICENSE.md).