class Weather:
    def __init__(self,
                 main,
                 description,
                 temp,
                 feels_like,
                 temp_min,
                 temp_max,
                 pressure,
                 humidity,
                 wind_speed,
                 wind_deg,
                 clouds,
                 sunrise,
                 sunset,
                 visibility):
        self.main = main
        self.description = description
        self.temp = temp
        self.feels_like = feels_like
        self.temp_min = temp_min
        self.temp_max = temp_max
        self.pressure = pressure
        self.humidity = humidity
        self.wind_speed = wind_speed
        self.wind_deg = wind_deg
        self.clouds = clouds
        self.sunrise = sunrise
        self.sunset = sunset
        self.visibility = visibility

    def __str__(self):
        return f"main: {self.main}\n" \
               f"description: {self.description}\n" \
               f"temp: {self.temp}\n" \
               f"feels like: {self.feels_like}\n" \
               f"temp min: {self.temp_min}\n" \
               f"temp max: {self.temp_max}\n" \
               f"pressure: {self.pressure}\n" \
               f"humidity: {self.humidity}\n" \
               f"wind speed: {self.wind_speed}\n" \
               f"wind deg: {self.wind_deg}\n" \
               f"clouds: {self.clouds}\n" \
               f"sunrise: {self.sunrise}\n" \
               f"sunset: {self.sunset}\n" \
               f"visibility: {self.visibility}\n"


def clear_weather_from_extra_data(weather):
    return Weather(
        main=weather['weather'][0]['main'],
        description=weather['weather'][0]['description'],
        temp=weather['main']['temp'],
        feels_like=weather['main']['feels_like'],
        temp_min=weather['main']['temp_min'],
        temp_max=weather['main']['temp_max'],
        pressure=weather['main']['pressure'],
        humidity=weather['main']['humidity'],
        wind_speed=weather['wind']['speed'],
        wind_deg=weather['wind']['deg'],
        clouds=weather['clouds']['all'],
        sunrise=weather['sys']['sunrise'],
        sunset=weather['sys']['sunset'],
        visibility=weather['visibility'],
    )


