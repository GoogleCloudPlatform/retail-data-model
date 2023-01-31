package server

import (
	"log"
	"os"
	"testing"

	"github.com/rrmcguinness/retail-data-model/events/grpc"
	server2 "github.com/rrmcguinness/retail-data-model/events/pkg/server"
	grpcCore "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	ConfigFile = "-configFile=examples/configs/events-test-config.json"
	SampleRate = "-sampleRate=1"
)

var server *server2.EventServer
var conn *grpcCore.ClientConn
var client grpc.EventsClient

func flagSetup() {
	os.Args = append(os.Args, ConfigFile)
	os.Args = append(os.Args, SampleRate)
}

func serverSetup() {
	var err error
	server, err = server2.NewEventServer()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
func clientSetup() {
	var opts = make([]grpcCore.DialOption, 0)
	var err error

	opts = append(opts, grpcCore.WithTransportCredentials(insecure.NewCredentials()))

	conn, err = grpcCore.Dial("localhost:50051", opts...)
	if err != nil {
		tearDown()
		panic(err)
	}
	client = grpc.NewEventsClient(conn)
}

func tearDown() {
	// Close the clients
	if conn != nil {
		err := conn.Close()
		if err != nil {
			log.Default().Printf("Failed to close GRPC Connection to Server: %v", err)
		}
	}
	// Close the server
	if server != nil {
		server.Stop()
	}
}

func TestMain(m *testing.M) {
	flagSetup()
	serverSetup()
	server.Start()
	clientSetup()
	sig := m.Run()
	tearDown()
	os.Exit(sig)
}
