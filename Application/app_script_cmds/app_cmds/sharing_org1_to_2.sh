
echo -n "org1 agree to sell"

curl --location --request POST 'localhost:3000/devices/agreetosell' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o1dev3",
    "tradeId": "tradeo1dev3",

    "tradePrice":100
}'


echo -n "org2 agree to buy"

curl --location --request POST 'localhost:4000/market/devices/interesttokens/submit' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o1dev3",
    "tradeId": "tradeo1dev3",
    "tradePrice": 100
}'

echo -n "org1 confirm sell"

curl --location --request POST 'localhost:3000/devices/confirmsell' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o1dev3",
    "tradeId": "tradeo1dev3",
    "bidderId":"Org2MSP"
}'