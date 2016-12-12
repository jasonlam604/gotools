# GoTools

* [Overview](#overview)
* [User Guide](#user-guide)
      * [Local IP](#myip)
      * [GOPATH & Bin config](#gopath-and-bin)
* [Build and Install](#build-and-install)
* [License](#license)
      
## Overview

Simple command line tools for GoLang, not really intended to be an active project, just a project
I'm adding to when I need quick and dirty solutions.

## User Guide


#### myip

```bash
~/bin/gotools myip
```

will display your ip

#### GoPath and Bin

```bash
~/bin/gotools gopath
```

When this command is called your current working directory is set at the $GOPATH and $GOPATH/bin
is add to your PATH.  Any existing $GOPATH/bin defined in $PATH will be removed first.

## Build and Install

Download dependency [Cobra](https://github.com/spf13/cobra)

```bash
go get -v github.com/spf13/cobra/cobra
```

Intall

```bash
go install github.com/jasonlam604/gotools
```

Under ~/bin you should now see the binary *gotools*

## License

Since there is a dependency on [Cobra](https://github.com/spf13/cobra) which is released under Apache 2.0 License to 
keep things simple this project is released under the same license.