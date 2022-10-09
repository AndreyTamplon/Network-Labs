class Address:
    def __init__(self,
                 xid,
                 city,
                 road,
                 state,
                 suburb,
                 country,
                 postcode,
                 country_code,
                 house_number,
                 city_district,
                 neighbourhood,
                 ):
        self.xid = xid
        self.city = city
        self.road = road
        self.state = state
        self.suburb = suburb
        self.country = country
        self.postcode = postcode
        self.country_code = country_code
        self.house_number = house_number
        self.city_district = city_district
        self.neighbourhood = neighbourhood

    def __str__(self):
        return f"city: {self.city}\n" \
               f"road: {self.road}\n" \
               f"state: {self.state}\n" \
               f"suburb: {self.suburb}\n" \
               f"country: {self.country}\n" \
               f"postcode: {self.postcode}\n" \
               f"country code: {self.country_code}\n" \
               f"house number: {self.house_number}\n" \
               f"city district: {self.city_district}\n" \
               f"neighbourhood: {self.neighbourhood}\n"



def clear_address_from_extra_data(address):
    return Address(
        xid=address['xid'],
        city=get_address_json_property(address, 'city'),
        road=get_address_json_property(address, 'road'),
        state=get_address_json_property(address, 'state'),
        suburb=get_address_json_property(address, 'suburb'),
        country=get_address_json_property(address, 'country'),
        postcode=get_address_json_property(address, 'postcode'),
        country_code=get_address_json_property(address, 'country_code'),
        house_number=get_address_json_property(address, 'house_number'),
        city_district=get_address_json_property(address, 'city_district'),
        neighbourhood=get_address_json_property(address, 'neighbourhood')
    )


def get_address_json_property(address, property):
    print(address)
    if property in address['address']:
        return address['address'][property]
    else:
        return ""



def get_by_xid(xid, addresses):
    for address in addresses:
        if address.xid == xid:
            return address
    return None
