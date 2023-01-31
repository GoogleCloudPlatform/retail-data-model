package com.google.retail;

import com.google.retail.enums.Country;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.DisplayName;

import static org.junit.jupiter.api.Assertions.*;

/**
 * An enumeration test case
 */
public class EnumTests {

  /**
   * Demonstrates how to use an enumeration.
   */
  @Test
  @DisplayName("Enum Test")
  void t0() {
    var usa = Country.USA;
    assertNotNull(usa);
  }
}
