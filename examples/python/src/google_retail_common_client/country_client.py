
import logging
import time

import grpc
from api.common.model_pb2 import Country
from google.protobuf.empty_pb2 import Empty
from api.common import service_pb2_grpc
import typing
from typing import AsyncIterable

def getCountriesStub(channel) -> service_pb2_grpc.CountriesStub:
  return service_pb2_grpc.CountriesStub(channel)


class CountryClient():
  def __init__(self, address: str):
    self.channel = grpc.insecure_channel('localhost:50051')
    self.stub = service_pb2_grpc.CountriesStub(self.channel)

  def Create(self, country: Country) -> Country:
    logging.info("------------- Saving Country ----------------------")
    return self.stub.Create(country)

  def Update(self, country: Country) -> Country:
    logging.info("------------ Updating Countries --------------------")
    return self.stub.Update(country)

  def Delete(self, country: Country) -> bool:
    logging.info("------------ Deleting Countries --------------------")
    return self.stub.Delete(country)

  def List(self) -> AsyncIterable[Country]:
    logging.info("------------ Listing Countries --------------------")
    return self.stub.List(Empty())