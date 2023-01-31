import unittest

from api.enums.country_pb2 import Country


class EnumTestCase(unittest.TestCase):

  def test_country_enum(self):
    country = Country.USA
    country_name = Country.Name(Country.USA)
    self.assertEqual(236, country)
    self.assertEqual("USA", country_name)

if __name__ == '__main__':
  unittest.main()
