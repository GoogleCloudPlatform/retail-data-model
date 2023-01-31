package com.google.retail;

import com.google.retail.common.pb.Country;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

/**
 * A test case for demonstrating model objects
 */
public class ModelTests {

  /**
   * A test demonstrating the creation pattern for a model object.
   */
  @Test
  @DisplayName("Test Model Object")
  void t0() {

    var usa =  Country.newBuilder().setId("USA")
        .setName("United States of America")
        .setAlpha2("US")
        .setAlpha3("USA").build();

    assertNotNull(usa);
    assertEquals("US", usa.getAlpha2());
    assertEquals("USA", usa.getAlpha3());
    assertEquals("United States of America", usa.getName());
  }
}
