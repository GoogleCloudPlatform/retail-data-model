package com.google.retail;

import static com.google.retail.common.grpc.CountriesGrpc.newBlockingStub;

import com.google.protobuf.Empty;
import com.google.retail.common.grpc.CountriesGrpc;
import com.google.retail.common.pb.Country;
import com.google.retail.common.pb.StatusResponse;
import io.grpc.Channel;

import java.util.ArrayList;
import java.util.List;

/**
 * CountriesClient is used for communicating with a server via a channel and executing the
 * RPC calls.
 */
public class CountriesClient {

  /**
   * Identifies a blocking stub, best used for the test case, the non-blocking stub
   * should be considered for async workflows.
   */
  private final CountriesGrpc.CountriesBlockingStub blockingStub;

  /**
   * Constuctor for creating the CountryClient from a channel
   * @param channel the channel to be communicated across
   */
  public CountriesClient(Channel channel) {
    blockingStub = newBlockingStub(channel);
  }

  /**
   * Lists countries from the service implementation
   * @return a list of countries
   */
  public List<Country> list() {
    var emptyRequest = Empty.newBuilder().build();
    var response = blockingStub.list(emptyRequest);
    final List<Country> out = new ArrayList<>();
    while (response.hasNext()) {
      out.add(response.next());
    }
    return out;
  }

  /**
   * Creates a country on the server
   * @param country the country to add
   * @return the country after it's executed from the service
   */
  public Country create(Country country) {
    var response = blockingStub.create(country);
    return response;
  }

  /**
   * Updates a country on the server
   * @param country the changeset of the country
   * @return the country after it's been updated
   */
  public Country update(Country country) {
    var response = blockingStub.update(country);
    return response;
  }

  /**
   * Deletes a country from the server, the deletion strategy is up to the service implementation.
   * @param country the country to delete
   * @return true if the country was deleted
   */
  public boolean delete(Country country) {
    var response = blockingStub.delete(country);
    return response.getType() == StatusResponse.Type.SUCCESS ||
        response.getType() == StatusResponse.Type.DELETED;
  }

}
