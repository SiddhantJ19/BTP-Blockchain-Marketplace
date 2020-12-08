echo -n "org 1 get list of shared devices"
curl --location --request POST 'http://localhost:3000/devices/shared/list'

echo -n "Get shared devioce all data"
curl --location --request POST 'localhost:3000/devices/shared/data/all' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ownerId":"Org2MSP",
    "deviceId": "o2dev1"
}'


echo -n "Get shared devioce latest data"
curl --location --request POST 'localhost:3000/devices/shared/data/latest' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ownerId":"Org2MSP",
    "deviceId": "o2dev1"
}'