package main

import (
	"MY_GO_CODES/Grpc_Rest_api/data"
	"MY_GO_CODES/Grpc_Rest_api/proto"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("GRPC SERVER ACTIVED")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		//Error while creating a tcp connection
		log.Fatalln(err)
	}
	//setting up the server as GRPC connection
	s := grpc.NewServer()
	//registering the GRPC connection as grpc server to implement methods from
	proto.RegisterEmpServiceServer(s, &server{})
	//reflection.Register(s)
	ers := s.Serve(lis)
	if ers != nil {
		log.Fatalf("failed to server: %v", ers)
	}
}
func (*server) SearchData(ctx context.Context, req *proto.EmpUnaryRequest) (*proto.EmpUnaryResponse, error) {
	fmt.Println("Searching data function active")
	fmt.Println(req)
	received := req.GetUnaryinput()
	slice := data.Stored("searching")
	var name string
	var mail string
	var mobile string
	for _, val := range slice {
		if received == val.Empid {
			name = val.Empname
			mail = val.Empmail
			mobile = val.Empmobile
			break
		}
	}
	res := &proto.EmpUnaryResponse{
		Unaryoutput: &proto.Employee{
			EmployeeId:     received,
			EmployeeName:   name,
			EmployeeMail:   mail,
			EmployeeMobile: mobile,
		},
	}
	fmt.Println("searching is done")
	return res, nil
}
func (*server) ShowAllData(req *proto.EmpStreamRequest, stream proto.EmpService_ShowAllDataServer) error {
	fmt.Println("starting the ShowAllData method")
	fmt.Println("the received request: ", req.GetStreaminput())
	var id string
	var name string
	var mail string
	var mobile string
	slice := data.Stored("sending all data")
	for _, val := range slice {
		id = val.Empid
		name = val.Empname
		mail = val.Empmail
		mobile = val.Empmobile
		res := &proto.EmpStreamResponse{
			Streamoutput: &proto.Employee{
				EmployeeId:     id,
				EmployeeName:   name,
				EmployeeMail:   mail,
				EmployeeMobile: mobile,
			},
		}
		//fmt.Println(res)
		stream.Send(res)
	}
	fmt.Println("Printed all data")
	return nil
}
