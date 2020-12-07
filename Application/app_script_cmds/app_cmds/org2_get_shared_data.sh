echo -n "org 2 get list of shared devices"
curl --location --request POST 'http://localhost:4000/devices/shared/list'

echo -n "oprg 2 Get shared devioce all data"
curl --location --request POST 'localhost:4000/devices/shared/data/all' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ownerId":"Org1MSP",
    "deviceId": "o1dev3"
}'


echo -n "org 2 Get shared devioce latest data"
curl --location --request POST 'localhost:4000/devices/shared/data/latest' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ownerId":"Org1MSP",
    "deviceId": "o1dev3"
}'