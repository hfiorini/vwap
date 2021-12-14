# VWAP calculation engine
The goal of this project is to create a real-time WWAP (volume-weighted average price) calculation engine. 
It uses the coinbase websocket feed to stream in trade executions and update the WWAP for each trading pair as updates become available. 

## Design
This application is composed by a service layer, to orchestrate the whole execution of the process.
A calculation layer (calculation-engine) to store the items retrieved by the web socket and perform the VWAP calculation 
and a third layer, the web socket client, using the coinbase exchange, but it could be easily replaced by any other.

## Calculation

Based on this formula:
<img width="673" alt="screenshot" src="https://i.ibb.co/616Z068/vwap.jpg">

Everytime an item is received from the web socket there are three maps that are being updated:
* A map with the sum of  (volume x Price), by product id
* A map with the sum of volume by product id
* A map with the result of  SUM(volume x Price)/ SUM(volume)

This approach is used to avoid going through the whole items list on each update

## Configuration
A configuration file is provided with the following values:
* CoinbaseUrl: The url for the web socket client
* TradingPairs: A string array with these values: BTC-USD,ETH-USD,ETH-BTC
* MaxSize: the max number of items for the VWAP calculation.

If needed, these values could be taken from any other data-sources, like a Secret Manager, a DB, or a text file

## To run
```
go run main.go
```

## Considerations
There are some things that could be a "nice to have" for future improvements
* Add mocks for some tests
* Add a library to handle decimal values
* Add a log library

