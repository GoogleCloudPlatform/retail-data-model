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

package com.google.retail;

import com.google.retail.common.pb.Country;
import io.grpc.ManagedChannel;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.BeforeEach;
import io.grpc.inprocess.InProcessChannelBuilder;
import io.grpc.inprocess.InProcessServerBuilder;
import io.grpc.testing.GrpcCleanupRule;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import com.google.retail.CountriesClient;
import com.google.retail.CountriesService;

import static org.junit.jupiter.api.Assertions.*;

/**
 * ServiceTest implements a test for a client and a server.
 */
public class ServiceTests {

  public final GrpcCleanupRule grpcCleanup = new GrpcCleanupRule();

  private CountriesClient client = null;

  /**
   * Creates a server as a test fixture
   * @throws Exception
   */
  @BeforeEach
  public void setUp() throws Exception {
    // Generate a unique in-process server name.
    String serverName = InProcessServerBuilder.generateName();

    // Create a server, add service, start, and register for automatic graceful shutdown.
    grpcCleanup.register(InProcessServerBuilder
        .forName(serverName).directExecutor().addService(new CountriesService()).build().start());

    // Create a client channel and register for automatic graceful shutdown.
    ManagedChannel channel = grpcCleanup.register(
        InProcessChannelBuilder.forName(serverName).directExecutor().build());

    // Create a HelloWorldClient using the in-process channel;
    client = new CountriesClient(channel);
  }

  /**
   * Tests the client against the server.
   */
  @Test
  @DisplayName("Test Client")
  void t0() {
    assertNotNull(client);

    var country = client.create(Country.newBuilder().setId("USA")
        .setName("United States of America")
        .setAlpha2("US")
        .setAlpha3("USA")
        .build());

    assertNotNull(country);
    assertEquals("US", country.getAlpha2());

    var list = client.list();
    assertNotNull(list);
    assertEquals(1, list.size());

    var updated = client.update(Country.newBuilder(country).setIntermediateRegion("AMER-USA-US").build());
    assertNotNull(updated);
    assertEquals("AMER-USA-US", updated.getIntermediateRegion());

    var deleted = client.delete(updated);
    Assertions.assertTrue(deleted);
  }
}
