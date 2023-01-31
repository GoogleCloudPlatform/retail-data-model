import logging
from typing import AsyncIterable

from api.common.model_pb2 import Country
from api.common.model_pb2 import IDRequest
from api.common.model_pb2 import StatusResponse
from google.protobuf.empty_pb2 import Empty

from api.common import service_pb2_grpc


class MockCountryServer(service_pb2_grpc.CountriesServicer):
  def __init__(self):
    self.state : dict[str, Country] = {}

  def FindById(self, request: IDRequest, unused_context) -> Country:
    return self.state(request.id)

  def Create(self, request: Country, unused_context) -> Country:
    logging.debug("Saving Country")
    self.state[request.id] = request
    logging.debug("Saved Country")
    return request

  def Update(self, request: Country, unused_context) -> Country:
    self.Delete(request, unused_context)
    return self.Create(request, unused_context)

  def Delete(self, request: Country, unused_context) -> IDRequest:
    del self.state[request.id]
    status = StatusResponse()
    status.code = 204
    return status

  def List(self, request: Empty, context) -> AsyncIterable[Country]:
    logging.debug("Listing countries")
    for key in self.state:
      yield self.state[key]
