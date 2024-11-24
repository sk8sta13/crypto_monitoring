# Crypto Monitoring

Receive notifications when a crypto reaches the value determined in the settings.

The program reads the settings and searches the Binance API for the value of the cryptocurrency, the quote value is tested and if it reaches the value you defined in the configuration, an alert is triggered in the operating system's notification area.

When making changes to the settings, it is not necessary to stop the program and run it again, the program checks if there have been any changes, if so, it already takes the new settings and starts working with them.

## Settings

See an example configuration:

```bash
{
    "interval": 5,
    "requests": [
        {
            "url": "https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT",
            "rules": [
                {"operator": "=", "value": 97000.0},
                {"operator": "<=", "value": 57000}
            ]
        },
        {
            "url": "https://api.binance.com/api/v3/ticker/price?symbol=DOGEUSDT",
            "rules": [
                {"operator": ">=", "value": 2.5},
                {"operator": "=", "value": 2}
            ]
        }
    ]
}
```

**Interval** is the time in seconds that the program will wait to make each request.
**Operator** is a rule applied to compare values, which came from the API with what you set in the configuration. Os operadores (<, <=, =, >=, >, <>).

I will not explain the other settings as I believe they are self-explanatory.

## Test

https://github.com/user-attachments/assets/5fce164e-8633-475c-9c64-4dcf60d2ea4f

## Updates

Firstly, I intend to create an installer that adds this program as a Linux service.

I also want to make a modification in order to have a specific query interval for each cryptocurrency.
