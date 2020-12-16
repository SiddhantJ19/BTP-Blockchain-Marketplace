
# add data to dev 3
curl --location --request POST 'localhost:3000/devices/data/add' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o1dev37",
    "data":"o1dev37 data after sharing"
}'