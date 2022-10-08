class Attraction:
    def __init__(self,
                 name,
                 dist,
                 rate,
                 kinds,
                 xid
                 ):
        self.xid = xid
        self.name = name
        self.dist = dist
        self.rate = rate
        self.kinds = kinds

    def __str__(self):
        return f"name: {self.name}\n" \
               f"dist: {self.dist}\n" \
               f"rate: {self.rate}\n" \
               f"kinds: {self.kinds}\n" \
               f"xid: {self.xid}\n"


def clear_attractions_from_extra_data(attractions):
    cleared_attractions = []
    for attraction in attractions['features']:
        cleared_attractions.append(Attraction(
            xid=get_attraction_json_property(attraction, 'xid'),
            name=get_attraction_json_property(attraction, 'name'),
            dist=get_attraction_json_property(attraction, 'dist'),
            rate=get_attraction_json_property(attraction, 'rate'),
            kinds=get_attraction_json_property(attraction, 'kinds')
        ))
        if attraction['properties']['name'] == "":
            cleared_attractions[-1].name = "no name"
    return cleared_attractions


def get_attraction_json_property(attraction, property):
    if property in attraction['properties']:
        return attraction['properties'][property]
    else:
        return ""
