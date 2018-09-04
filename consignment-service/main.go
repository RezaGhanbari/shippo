// consignment-service/main.go
package main
import (
	pb "shippo/consignment-service/proto/consignment"
)
const (
	port = ":50051"
)


type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}
