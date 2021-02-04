# currency_exchange_rates_collector
currency_exchange_collector
Requirements not were obvious, but i tried to do all test task items
1) Routes

    1.1) `/price` - refreshes whole informations all fsyms and tsyms in config.yml and returns requested fsyms and tsyms currency exchange rates
        Query parameters: fsyms - list of cryptocurrencies joined by comma
                          tsyms - list of noncrypto currencies joined by comma

    1.2) `/refresh` - refreshes whole informations all fsyms and tsyms and returns it

    1.3) `/ws` registers in websocket hub and immediately triggers refresh by all fsyms and tsyms in config.yml
    Notes: `price` and `refresh` triggers sends refreshed currency exchange rates to all websocket clients

2) Maybe i did not correcly undertand of websocket usage what you wanted

3) In production i would use GRPC instead HTTP 1.0
