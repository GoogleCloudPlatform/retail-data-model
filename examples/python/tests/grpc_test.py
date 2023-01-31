import logging
import unittest
from concurrent import futures
from threading import Thread

import grpc

from api.common.model_pb2 import Country
from country_client import CountryClient

import country_service
from api.common import service_pb2_grpc

def client_test():
  client = CountryClient('localhost:50051')

  country = Country(
    id="US",
    name="United States of America",
    alpha2="US",
    alpha3="USA")

  client.Create(country)

  for c in client.List():
    print("\t"+c.name)

class ServerThread(Thread):
  def __init__(self):
    Thread.__init__(self, name = "server_thread", daemon=True)
    self.server = grpc.server(futures.ThreadPoolExecutor(max_workers=2))
    service_pb2_grpc.add_CountriesServicer_to_server(country_service.MockCountryServer(),
                                                     self.server)
    self.server.add_insecure_port('localhost:50051')

  def run(self):
    logging.debug("\n################# Starting Server ######################\n")
    self.server.start()
    # Start, but only run for ten seconds
    self.server.wait_for_termination(timeout=5)
    logging.debug("\n################# Stopping Server ######################\n")

  def stop(self):
    self.server.stop()


class ClientServerTestCase(unittest.TestCase):
  def test_server_and_client(self):
    logging.basicConfig(level=logging.INFO)

    # Start the server
    server = ServerThread()
    server.start()

    # Execute the client
    worker = Thread(target=client_test, name="t2")
    worker.start()

    # Wait for them to complete
    worker.join()
    server.join()

if __name__ == '__main__':
  unittest.main()
