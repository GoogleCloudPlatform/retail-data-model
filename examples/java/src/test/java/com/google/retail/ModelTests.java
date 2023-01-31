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
