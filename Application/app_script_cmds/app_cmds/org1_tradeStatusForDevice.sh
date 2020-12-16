echo -n "Get Trade status"
curl --location --request POST 'localhost:3000/devices/tradeStatus' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o1dev37"
}'