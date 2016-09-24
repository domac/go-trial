package main

import (
	proto "github.com/golang/protobuf/proto"
	"github.com/phillihq/go-trial/pbtest/pb"
)

//protoc --go_out=. *.proto

func main() {

	msg := &pb.Helloworld{
		Id:  proto.Int32(101),
		Str: proto.String("Hello Linux"),
	}

	buffer, err := proto.Marshal(msg)

	if err != nil {
		println("marshal error ")
	}

	new_msg := &pb.Helloworld{}

	err = proto.Unmarshal(buffer, new_msg)

	println(new_msg.String())
}
