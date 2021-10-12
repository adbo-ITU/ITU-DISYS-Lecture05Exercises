package main

import (
	pb "adbo/timeservice/service"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var wg sync.WaitGroup

func main() {
	logPath := fmt.Sprintf("client_%d.log", 0) // time.Now().Unix())
	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer logFile.Close()
	logOutputs := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(logOutputs)

	addresses := []string{"10.26.18.24:8080", "127.0.0.1:8080", "10.26.31.80:8080"}
	clients := make([]pb.TimeClient, len(addresses))

	for i, address := range addresses {
		conn, err := connectToServer(address)
		if err != nil {
			return
		}
		defer conn.Close()
		clients[i] = pb.NewTimeClient(conn)
	}

	for i, client := range clients {
		wg.Add(1)
		go fetchTimeWithLogger(client, addresses[i])
	}
	wg.Wait()
}

func connectToServer(address string) (*grpc.ClientConn, error) {
	log.Printf("Connecting to server %s...\n", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		return nil, err
	}
	return conn, nil
}

func fetchTimeWithLogger(client pb.TimeClient, logTag string) (*pb.Time, error) {
	defer wg.Done()
	start := time.Now()
	log.Printf("(%s) Requesting Now...\n", logTag)
	currentTime, err := fetchTime(client)

	end := time.Now()
	elapsed := end.Sub(start)

	if err != nil {
		log.Fatalf("(%s) Failed when requesting Now: %v\n", logTag, err)
		return nil, err
	}

	parsedTime, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", currentTime.Time)
	if err != nil {
		log.Fatalf("(%s) Failed to parse time string: %v\n", logTag, err)
		return nil, err
	}
	delta := parsedTime.Sub(start)

	cristian := parsedTime.Add(end.Sub(start) / 2)

	log.Printf(`(%s) Received Now response:
  - Address: %s
  - Start time: %s
  - Round trip: %s
  - Returned time: %s
  - Time delta: %s
  - Synchronized time (Cristian): %s
`,
		logTag, logTag, start, elapsed, parsedTime, delta, cristian)
	return currentTime, nil
}

// DO NOT CALL DIRECTLY
func fetchTime(client pb.TimeClient) (*pb.Time, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return client.Now(ctx, &emptypb.Empty{})
}
