# hayesgopong
Classic pong game written in golang compile to javavscript.


![Screenshot](hayesgopong.png)

# Compiling

Just install gopherjs and build with it.

```console
$ go get github.com/gopherjs/gopherjs
$ go get github.com/gen2brain/beeep
$ go get github.com/gopherjs/gopherwasm/js
$ gopherjs build github.com/KeithHayes23/hayesgopong
$ gopherjs serve github.com/KeithHayes23/hayesgopong
