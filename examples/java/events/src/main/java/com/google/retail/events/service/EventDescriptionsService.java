package com.google.retail.events.service;

import com.google.protobuf.Empty;
import com.google.retail.common.pb.IDRequest;
import com.google.retail.common.pb.StatusResponse;
import io.grpc.stub.StreamObserver;
import com.google.retail.events.pb.EventDescription;
import com.google.retail.events.grpc.EventDescriptionsGrpc;

public class EventDescriptionsService extends EventDescriptionsGrpc.EventDescriptionsImplBase {

  @Override
  public void list(Empty request, StreamObserver<EventDescription> responseObserver) {
    super.list(request, responseObserver);
  }

  @Override
  public void get(IDRequest request, StreamObserver<EventDescription> responseObserver) {
    super.get(request, responseObserver);
  }

  @Override
  public StreamObserver<EventDescription> create(StreamObserver<StatusResponse> responseObserver) {
    return super.create(responseObserver);
  }

  @Override
  public StreamObserver<EventDescription> update(StreamObserver<StatusResponse> responseObserver) {
    return super.update(responseObserver);
  }

  @Override
  public StreamObserver<EventDescription> delete(StreamObserver<StatusResponse> responseObserver) {
    return super.delete(responseObserver);
  }

}