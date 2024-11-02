# Cli Weather API
This is a command-line weather application written in Go using https://www.weatherapi.com/

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/cli-weather-api.git
   cd cli-weather-api
    ```
2. Add an .env file in the top directory, specifying your API key like so:
    ```.env
    WEATHER_API_KEY=your_api_key
    ```
3. Install the necessary modules:
    ```bash
    go mod tidy
    ```

## Usage
- To use the app, simply run:
    ```bash
    go run main.go [location] [options]
    ```
- location (required): the city for which you want to check the weather
- options (optional)
  - `-c`: for the current weather (default if left blank)
  - `-f`: for the forecast weather

## Examples

Command:
```bash
go run main.go Warsaw
```
Response:
```
-------------------------------
City: Warsaw
Country: Poland
Current time: 02-11-2024 13:12
-------------------------------
Temperature: 8 ℃
Condition: Sunny

```


Command:
```bash
go run main.go Warsaw -f
```
Response:
```
-------------------------------
City: Warsaw
Country: Poland
Current time: 02-11-2024 13:12
-------------------------------
14:00 - 9 ℃ | Sunny
15:00 - 9 ℃ | Sunny
16:00 - 8 ℃ | Sunny
17:00 - 8 ℃ | Clear
18:00 - 7 ℃ | Overcast
19:00 - 7 ℃ | Overcast
20:00 - 6 ℃ | Clear
21:00 - 6 ℃ | Partly Cloudy
22:00 - 6 ℃ | Partly Cloudy
23:00 - 6 ℃ | Partly Cloudy
```