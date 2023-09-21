package com.google.retail.events.service;

import com.google.retail.common.pb.StatusResponse;
import com.google.retail.events.grpc.EventsGrpc;
import com.google.retail.events.pb.Event;
import io.grpc.stub.StreamObserver;

public class EventsService extends EventsGrpc.EventsImplBase {
  @Override
  public StreamObserver<Event> emit(StreamObserver<StatusResponse> responseObserver) {
    return super.emit(responseObserver);
  }
}
