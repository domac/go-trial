package main

import snap "github.com/mreiferson/go-snappystream"
import "os"
import "io"
import "flag"

type Options struct {
	Encode bool
	Decode bool
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var en = flag.Bool("encode", false, "encode stdin (the default)")
var de = flag.Bool("decode", false, "decode stdin")

func main() {
	flag.Parse()
	opts := &Options{
		Decode: *en,
		Encode: *de,
	}

	if opts.Decode {
		decode()
	} else {
		encode()
	}
}

func encode() {
	_, err := io.Copy(snap.NewWriter(os.Stdout), os.Stdin)
	check(err)
}

func decode() {
	_, err := io.Copy(os.Stdout, snap.NewReader(os.Stdin, snap.VerifyChecksum))
	check(err)
}
