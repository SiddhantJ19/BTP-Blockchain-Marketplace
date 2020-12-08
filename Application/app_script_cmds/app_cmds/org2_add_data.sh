
# add data to dev 3
curl --location --request POST 'localhost:4000/devices/data/add' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o2dev1",
    "data":"o2d1 data before sharing "
}'