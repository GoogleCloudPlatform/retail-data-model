/*
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package server

import (
	"errors"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Server is a test server, not to be used as a production instance.
type Server struct {
	Listener net.Listener
	Server   *grpc.Server
	Started  bool
	quit     chan bool
}

// NewServer builds a server object
func NewServer(options []grpc.ServerOption) *Server {
	// Get a random port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer(options...)
	return &Server{Listener: l, Server: server, Started: false, quit: make(chan bool)}
}

// Start starts the server listener
func (t *Server) Start() {
	go func() {
		for {
			select {
			case <-t.quit:
				log.Printf("Stopping Server: %v", t.Listener.Addr())
				t.Server.GracefulStop()
				return
			default:
				log.Printf("Starting Server %v", t.Listener.Addr())
				if err := t.Server.Serve(t.Listener); err != nil {
					log.Fatalf("server exited with error %v", err)
				}
			}
		}
	}()
	t.Started = true
}

func (t *Server) Stop() {
	t.quit <- true
}

// Address returns the address used by the server since it's a random port
func (t *Server) Address() (string, error) {
	if t.Listener != nil && t.Server != nil && t.Started {
		return t.Listener.Addr().String(), nil
	}
	return "", errors.New("server is currently not listening on a port")
}
