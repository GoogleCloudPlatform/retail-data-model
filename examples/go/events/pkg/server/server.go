package server

import (
	"flag"
	"fmt"
	"log"
	"net"

	eventsGrpc "github.com/rrmcguinness/retail-data-model/events/grpc"
	"github.com/rrmcguinness/retail-data-model/events/pkg/module"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	host     = flag.String("host", "localhost", "The host interface to listen on")
	port     = flag.Int("port", 50051, "The server port")
)

type EventServer struct {
	opts        []grpc.ServerOption
	started     bool
	useTls      bool
	certFile    string
	keyFile     string
	Host        string
	Port        int32
	eventModule *module.EventModule
	server      *grpc.Server
	quit        chan bool
}

func NewEventServer() (out *EventServer, err error) {
	flag.Parse()

	out = &EventServer{
		opts:     make([]grpc.ServerOption, 0),
		useTls:   *tls,
		certFile: *certFile,
		keyFile:  *keyFile,
		Host:     *host,
		Port:     int32(*port),
		quit:     make(chan bool),
	}
	out.PrepareGRPCOpts()

	eventModule, err := module.Load()

	if ok := out.fatalErrorCheck("failed to create module", err); ok {
		out.eventModule = eventModule

		out.server = grpc.NewServer(out.opts...)

		eventsGrpc.RegisterEventsServer(out.server, out.eventModule.EventServer)
		eventsGrpc.RegisterEventRecordsServer(out.server, out.eventModule.EventRecordServer)
		eventsGrpc.RegisterEventDescriptionsServer(out.server, out.eventModule.EventDescriptionService)
	}
	return out, err
}

func (server *EventServer) fatalErrorCheck(msg string, err error) bool {
	if err != nil {
		if server.eventModule != nil && server.eventModule.GetDefaultLog() != nil {
			server.eventModule.GetDefaultLog().Fatalf(msg, err)
		} else {
			log.Default().Fatalf(msg, err)
		}
		return false
	}
	return true
}

func (server *EventServer) PrepareGRPCOpts() {
	if *tls {
		if *certFile != "" && *keyFile != "" {
			creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
			if ok := server.fatalErrorCheck("failed to generate credentials", err); ok {
				server.opts = append(server.opts, grpc.Creds(creds))
			}
		} else {
			// Anonymous credentials
		}
	}
}

// Start starts the server listening on the registered port.
func (server *EventServer) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Host, server.Port))
	if ok := server.fatalErrorCheck("failed to listen on port", err); ok {
		go func() {
			for {
				select {
				case <-server.quit:
					log.Printf("Stopping Server: %s:%d", server.Host, server.Port)
					server.server.GracefulStop()
					server.started = false
					return
				default:
					err = server.server.Serve(lis)
					if ok := server.fatalErrorCheck("failed to listen", err); ok {
						server.started = true
					}
				}
			}
		}()
	}
}

// Stop gracefully shuts down the server
func (server *EventServer) Stop() {
	server.quit <- true
}
