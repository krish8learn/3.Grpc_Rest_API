package main

import (
	"MY_GO_CODES/Grpc_Rest_api/proto"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Rest Api server is activated")
	//setting up connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := proto.NewEmpServiceClient(conn)
	//setting up http server
	router := gin.Default()
	router.GET("/streaminput", func(c *gin.Context) {
		streaminput := c.Param("streaminput")
		//contacting server and printing out response
		req := &proto.EmpStreamRequest{
			Streaminput: streaminput,
		}
		//calling the function of grpc server
		got, err := client.ShowAllData(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		for {
			msg, err := got.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				c.JSON(http.StatusExpectationFailed, gin.H{
					"error": err.Error(),
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint("ID: ", msg.GetStreamoutput().GetEmployeeId(), ",",
					"Name: ", msg.GetStreamoutput().GetEmployeeName(), ",",
					"Mail: ", msg.GetStreamoutput().GetEmployeeMail(), ",",
					"Mobile: ", msg.GetStreamoutput().GetEmployeeMobile()),
			})
			//fmt.Println(msg.GetStreamoutput().GetEmployeeId())
		}
	})
	router.GET("/Employee/:Employee_id", func(c *gin.Context) {
		Employee_id := c.Param("Employee_id")
		re := &proto.EmpUnaryRequest{
			Unaryinput: Employee_id,
		}
		//calling the function of grpc server
		got, err := client.SearchData(c, re)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint("-->searchid:", got.GetUnaryoutput().GetEmployeeId(), ",",
				"Name: ", got.GetUnaryoutput().GetEmployeeName(), ",",
				"Mail:", got.GetUnaryoutput().GetEmployeeMail(), ",",
				"Mobile:", got.GetUnaryoutput().GetEmployeeMobile()),
		})
	})
	if erv := router.Run("localhost:8052"); erv != nil {
		log.Fatalln(err)
	}
}
