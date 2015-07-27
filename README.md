# Wrengo

NOTE: This README is a little out of date as per the current build.
I'll be updating soon.

Wrengo is a simple hobby project of mine that has been built on top of
[The Wren Programming Language](https://github.com/munificent/wren). Wren is
an embeddable scripting language that's been written in C. Wrengo, as you may
have already guessed by the name, creates an API written in both Go and Wren that
extends the functionality of Wren to do various other things.

#### Installing the binary (currently darwin-amd64 only)

```bash
$ wget https://github.com/hazbo/wrengo/releases/download/latest/wrengo-latest-dev-darwin-amd64.tar.gz
$ tar xzvf wrengo-latest-dev-darwin-amd64.tar.gz
$ cd wrengo-latest-dev-darwin-amd64
$ ./wrengo --version
```

Or just download [the latest dev release](https://github.com/hazbo/wrengo/releases/download/latest/wrengo-latest-dev-darwin-amd64.tar.gz).

Running: `./wrengo run file.wren`

#### Building from source

You'll need any build requirements required from [Wren](https://github.com/munificent/wren)
itself.

```bash
$ git clone https://github.com/hazbo/wrengo
$ cd wrengo
$ export GOPATH=$PWD/vendor
$ make
```

Wrengo will be located at `build/wrengo`. To install: `sudo make install`

#### Usage

Unlike Wren, Wrengo programs have a `main` entry point. See below for a
simple example of us reading data from a text file:

```javascript
// Application entry point
var main = new Fn {
	var f = new File
	var c = f.Read("myfile.txt")

	IO.print(c)
}
```

This project is brand new and is still very much in it's initial development
stage. Pull requests and improvements are very welcome. I'm still writing
the API docs. For information on implementing your own Go functionality into
this, please [refer to the wiki](https://github.com/hazbo/wrengo/wiki/Hacking).

### License

MIT
