import unittest

from api.common.model_pb2 import Country
from api.enums.country_pb2 import Country as Ctry
from api.enums.locale_pb2 import Locale


class ModelTestCase(unittest.TestCase):
  def test_model(self):
    en_us = Locale.EN_US
    US = Ctry.USA

    country = Country()
    country.id = Ctry.Name(US)
    country.name = "United States of America"
    country.alpha2 = "US"
    country.alpha3 = "USA"
    self.assertEqual("USA", country.id)
    self.assertEqual("US", country.alpha2)

if __name__ == '__main__':
  unittest.main()
