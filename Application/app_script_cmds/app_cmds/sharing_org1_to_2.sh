#!/bin/bash
echo -n "org1 agree to sell"
revokeTime=$(date -d "+3 minutes" +%s)

curl --location --request POST 'localhost:3000/devices/agreetosell' \
-H "Accept: application/json" \
--header 'Content-Type: application/json' \
--data @<(cat <<EOF
{
    "deviceId": "o1dev29",
    "tradeId": "tradeo1dev15",
    "revoke_time": $revokeTime,
    "tradePrice":100
}
EOF
)


echo -n "org2 agree to buy"

curl --location --request POST 'localhost:4000/market/devices/interesttokens/submit' \
--header 'Content-Type: application/json' \
--data @<(cat <<EOF
{
    "deviceId": "o1dev29",
    "tradeId": "tradeo1dev15",
    "revoke_time": $revokeTime,
    "tradePrice":100
}
EOF
)

echo -n "org1 confirm sell"

curl --location --request POST 'localhost:3000/devices/confirmsell' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o1dev29",
    "tradeId": "tradeo1dev15",   
    "bidderId":"Org2MSP"
}'