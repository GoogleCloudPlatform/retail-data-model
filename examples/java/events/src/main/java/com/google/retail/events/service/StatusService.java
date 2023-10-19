package com.google.retail.events.service;

import com.google.protobuf.Empty;
import com.google.retail.common.pb.HealthCheckResponse;
import com.google.retail.events.grpc.StatusGrpc;
import io.grpc.stub.StreamObserver;

public class StatusService extends StatusGrpc.StatusImplBase {
  @Override
  public void get(Empty request, StreamObserver<HealthCheckResponse> responseObserver) {
    responseObserver.onNext(HealthCheckResponse.newBuilder().setStatus(HealthCheckResponse.ServingStatus.SERVING).build());
  }
}
