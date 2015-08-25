# Anybar

Anybar is a command line tool and Go library, which provides the ability to send commands to an Anybar.app instance. 

**Note:** You must have Anybar.app installed, to install go [here](https://github.com/tonsky/AnyBar).

**Note:** This project is only for OS X.

# Installation

> go get github.com/umahmood/anybar

> cd $GOPATH/src/github.com/umahmood/anybar

> go install ./...

# Usage

To change the dot color to blue:

> anybar -cmd=blue

If the Anybar.app is not running on the default port, you can specify it:

> anybar -port=1740 -cmd=green

# Documenation

If you would like to use the library with in your application go here:

> http://godoc.org/github.com/umahmood/anybar

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
