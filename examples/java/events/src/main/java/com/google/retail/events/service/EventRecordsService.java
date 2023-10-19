package com.google.retail.events.service;

import com.google.retail.common.pb.IDRequest;
import com.google.retail.common.pb.TimeBoundRequest;
import com.google.retail.events.grpc.EventRecordsGrpc;
import com.google.retail.events.pb.EventRecord;
import io.grpc.stub.StreamObserver;

public class EventRecordsService extends EventRecordsGrpc.EventRecordsImplBase {
  @Override
  public void list(TimeBoundRequest request, StreamObserver<EventRecord> responseObserver) {
    super.list(request, responseObserver);
  }

  @Override
  public void findTransactionById(IDRequest request, StreamObserver<EventRecord> responseObserver) {
    super.findTransactionById(request, responseObserver);
  }
}
