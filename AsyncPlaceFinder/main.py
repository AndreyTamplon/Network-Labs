import json
import asyncio
import aiohttp
from entities.place import clear_places_from_extra_data, get_place_by_id, print_places
from entities.weather import clear_weather_from_extra_data
from entities.attraction import clear_attractions_from_extra_data
from entities.address import clear_address_from_extra_data, get_by_xid


def input_place_name():
    print("Enter the place: ")
    place_name = input()
    return place_name


async def request_places(place_name):
    async with aiohttp.ClientSession() as session:
        with open("config.json") as cfg:
            config = json.load(cfg)
        url = (config["places_api"]["url"]).format(place_name,
                                                   config["places_api"]["locale"],
                                                   config["places_api"]["key"])
        response = await session.get(url)
        return await response.text(encoding="UTF-8")


def get_places(place_name):
    return clear_places_from_extra_data(json.loads(asyncio.run(request_places(place_name))))


def get_place_id(places):
    print("To get the weather, provide the place id: ")
    place_id = int(input())
    while True:
        place_user_interested_in = get_place_by_id(place_id, places)
        if place_user_interested_in is None:
            print("Wrong id, try again: ")
            place_id = int(input())
        else:
            break
    return place_user_interested_in


async def request_weather(place):
    async with aiohttp.ClientSession() as session:
        with open("config.json") as cfg:
            config = json.load(cfg)
        url = (config["weather_api"]["url"]).format(place.point["lat"],
                                                    place.point["lng"],
                                                    config["weather_api"]["key"])
        response = await session.get(url)
        return await response.text(encoding="UTF-8")


def get_weather(place_user_interested_in):
    weather = clear_weather_from_extra_data(json.loads(asyncio.run(request_weather(place_user_interested_in))))
    return weather


def get_radius():
    print("To get the attractions, provide the radius: ")
    while True:
        radius = int(input())
        if radius < 0:
            print("Wrong radius, it should be positive")
        else:
            break
    return radius


async def request_attractions(place, radius):
    async with aiohttp.ClientSession() as session:
        with open("config.json") as cfg:
            config = json.load(cfg)
        url = (config["attractions_api"]["url"]).format(radius,
                                                        place.point["lng"],
                                                        place.point["lat"],
                                                        config["attractions_api"]["key"])
        response = await session.get(url)
        return await response.text(encoding="UTF-8")


def get_attractions(place_user_interested_in, radius):
    return clear_attractions_from_extra_data(
        json.loads(asyncio.run(request_attractions(place_user_interested_in, radius))))


def create_tasks(session, attractions):
    with open("config.json") as cfg:
        config = json.load(cfg)
    tasks = []
    for attraction in attractions:
        address_url = (config["attractions_api"]["address_url"]).format(attraction.xid,
                                                                        config["attractions_api"]["key"])
        tasks.append(session.get(address_url))
    return tasks


async def request_addresses(attractions):
    addresses = []
    async with aiohttp.ClientSession() as session:
        tasks = create_tasks(session, attractions)
        responses = await asyncio.gather(*tasks)
        addresses.append([await response.json() for response in responses])
        return addresses


def get_addresses(attractions):
    raw_addresses = asyncio.run(request_addresses(attractions))
    addresses = [clear_address_from_extra_data(raw_address) for raw_address in raw_addresses[0]]
    return addresses


def print_attractions_with_address(attractions, addresses):
    for attraction in attractions:
        address = get_by_xid(attraction.xid, addresses)
        print(attraction)
        print(address)


def main():
    place_name = input_place_name()
    asyncio.set_event_loop_policy(asyncio.WindowsSelectorEventLoopPolicy())  # it's necessary for Windows
    places = get_places(place_name)
    print_places(places)

    place_user_interested_in = get_place_id(places)
    weather = get_weather(place_user_interested_in)
    print(weather)

    radius = get_radius()
    attractions = get_attractions(place_user_interested_in, radius)

    addresses = get_addresses(attractions)
    print_attractions_with_address(attractions, addresses)


if __name__ == '__main__':
    main()
