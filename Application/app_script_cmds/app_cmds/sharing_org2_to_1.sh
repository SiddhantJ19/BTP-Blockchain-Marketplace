
echo -n "org2 agree to sell"

curl --location --request POST 'localhost:4000/devices/agreetosell' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o2dev1",
    "tradeId": "tradeo2dev1",

    "tradePrice":250
}'


echo -n "org1 agree to buy"

curl --location --request POST 'localhost:3000/market/devices/interesttokens/submit' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o2dev1",
    "tradeId": "tradeo2dev1",
    "tradePrice": 250
}'

echo -n "org2 confirm sell"

curl --location --request POST 'localhost:4000/devices/confirmsell' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o2dev1",
    "tradeId": "tradeo2dev1",
    "bidderId":"Org1MSP"
}'