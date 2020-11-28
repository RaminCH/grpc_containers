package main

import (
	pb "github.com/RaminCH/lec5/server/proto/consigment"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/reflection"

	_ "context"
	_ "log"
	_ "net"
	_ "sync"
)

const (
	port = ":50001"
)

//obyekt kotoriy budet xranit info o vnutrennix konfiguratsiyax
type repository interface {
	Create(*pb.Command) (*pb.Command, error)
}

func main() {

}
