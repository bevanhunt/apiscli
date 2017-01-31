# APISCLI - Command-Line API discovery

## Local

### Setup
- ` brew update `
- ` brew install go `
- ` brew install glide `
-  setup GOPATH for your env file (.bashrc or .zshrc):
```bash
export GOPATH=$HOME/go
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```
- ` git clone ` this repo into ` ~/go/src `
- ` glide install ` in the local folder
- ` go build ` in the local folder

### Search Example for Trade Keyword

` ./apiscli search trade `

## Recommended Go Editor
- [Visual Studio Code](https://code.visualstudio.com/) with [Go Extension](https://github.com/Microsoft/vscode-go)

## License
Copyright (c) 2017 CA. All rights reserved.
This software may be modified and distributed under the terms of the MIT license. To learn more, see the [License](LICENSE.md).