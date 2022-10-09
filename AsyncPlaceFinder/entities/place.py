class Place:
    def __init__(self,
                 point,
                 name,
                 country,
                 osm_key,
                 osm_value,
                 osm_id):
        self.point = point
        self.name = name
        self.country = country
        self.osm_key = osm_key
        self.osm_value = osm_value
        self.osm_id = osm_id

    def __str__(self):
        return f"name: {self.name}\n" \
               f"country: {self.country}\n" \
               f"osm key: {self.osm_key}\n" \
               f"osm value: {self.osm_value}\n" \
               f"place id: {self.osm_id}\n"


def clear_places_from_extra_data(places):
    cleared_places = []
    for place in places['hits']:
        cleared_places.append(Place(
            point=place['point'],
            name=place['name'],
            country=place['country'],
            osm_key=place['osm_key'],
            osm_value=place['osm_value'],
            osm_id=place['osm_id']
        ))
    return cleared_places


def print_places(places):
    for place in places:
        print(place)


def check_id(osm_id, places):
    for place in places:
        if place.osm_id == osm_id:
            return True
    return False


def get_place_by_id(osm_id, places):
    for place in places:
        if place.osm_id == osm_id:
            return place
    return None
