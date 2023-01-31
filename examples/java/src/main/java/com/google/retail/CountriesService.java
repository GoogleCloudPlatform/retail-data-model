package com.google.retail;

import com.google.protobuf.Empty;
import com.google.protobuf.Timestamp;
import com.google.retail.common.grpc.CountriesGrpc;
import com.google.retail.common.pb.Country;
import com.google.retail.common.pb.IDRequest;
import com.google.retail.common.pb.StatusResponse;
import io.grpc.stub.StreamObserver;

import java.util.Collections;
import java.util.HashMap;
import java.util.Map;

/**
 * CountriesService implements the GRPC service.
 */
public class CountriesService extends CountriesGrpc.CountriesImplBase {
  private final Map<String, Country> state = Collections.synchronizedMap(new HashMap<>());

  /**
   * Implements the list function.
   * @param request the Empty request object.
   * @param responseObserver A stream of Country objects
   */
  @Override
  public void list(Empty request, StreamObserver<Country> responseObserver) {
    state.forEach((k, v) -> {
      responseObserver.onNext(v);
    });
    responseObserver.onCompleted();
  }

  /**
   * Implements the findById
   * @param request the ID Request for the Country
   * @param responseObserver - A stream of one referencing the country.
   */
  @Override
  public void findById(IDRequest request, StreamObserver<Country> responseObserver) {
    if (state.containsKey(request.getId())) {
      responseObserver.onNext(state.get(request.getId()));
    }
    responseObserver.onCompleted();
  }

  /**
   * Implements Create
   * @param request the country to create
   * @param responseObserver the response country type, useful if decorated
   */
  @Override
  public void create(Country request, StreamObserver<Country> responseObserver) {
    state.put(request.getId(), request);
    responseObserver.onNext(state.get(request.getId()));
    responseObserver.onCompleted();
  }

  /**
   * Implements Update
   * @param request the country to update
   * @param responseObserver the
   */
  @Override
  public void update(Country request, StreamObserver<Country> responseObserver) {
    create(request, responseObserver);
  }

  /**
   * Implements Delete
   * @param request the country to delete
   * @param responseObserver the status response for if the country was deleted
   */
  @Override
  public void delete(Country request, StreamObserver<StatusResponse> responseObserver) {
    if (state.containsKey(request.getId())) {
      state.remove(request.getId());
      responseObserver.onNext(StatusResponse.newBuilder()
          .setType(StatusResponse.Type.DELETED)
          .setTs(Timestamp.newBuilder().build()).build());
      responseObserver.onCompleted();
    }
  }
}
