#  Copyright 2022 Google LLC
#
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

# TODO - Fix with the GRPC issue

# import logging
# from typing import AsyncIterable
#
# from api.common.model_pb2 import Country
# from api.common.model_pb2 import IDRequest
# from api.common.model_pb2 import StatusResponse
# from google.protobuf.empty_pb2 import Empty
#
# from api.common import service_pb2_grpc
#
#
# class MockCountryServer(service_pb2_grpc.CountriesServicer):
#   def __init__(self):
#     self.state : dict[str, Country] = {}
#
#   def FindById(self, request: IDRequest, unused_context) -> Country:
#     return self.state(request.id)
#
#   def Create(self, request: Country, unused_context) -> Country:
#     logging.debug("Saving Country")
#     self.state[request.id] = request
#     logging.debug("Saved Country")
#     return request
#
#   def Update(self, request: Country, unused_context) -> Country:
#     self.Delete(request, unused_context)
#     return self.Create(request, unused_context)
#
#   def Delete(self, request: Country, unused_context) -> IDRequest:
#     del self.state[request.id]
#     status = StatusResponse()
#     status.code = 204
#     return status
#
#   def List(self, request: Empty, context) -> AsyncIterable[Country]:
#     logging.debug("Listing countries")
#     for key in self.state:
#       yield self.state[key]
