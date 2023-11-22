
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
# import time
#
# import grpc
# from api.common.model_pb2 import Country
# from google.protobuf.empty_pb2 import Empty
# from api.common import service_pb2_grpc
# import typing
# from typing import AsyncIterable
#
# def getCountriesStub(channel) -> service_pb2_grpc.CountriesStub:
#   return service_pb2_grpc.CountriesStub(channel)
#
#
# class CountryClient():
#   def __init__(self, address: str):
#     self.channel = grpc.insecure_channel('localhost:50051')
#     self.stub = service_pb2_grpc.CountriesStub(self.channel)
#
#   def Create(self, country: Country) -> Country:
#     logging.info("------------- Saving Country ----------------------")
#     return self.stub.Create(country)
#
#   def Update(self, country: Country) -> Country:
#     logging.info("------------ Updating Countries --------------------")
#     return self.stub.Update(country)
#
#   def Delete(self, country: Country) -> bool:
#     logging.info("------------ Deleting Countries --------------------")
#     return self.stub.Delete(country)
#
#   def List(self) -> AsyncIterable[Country]:
#     logging.info("------------ Listing Countries --------------------")
#     return self.stub.List(Empty())