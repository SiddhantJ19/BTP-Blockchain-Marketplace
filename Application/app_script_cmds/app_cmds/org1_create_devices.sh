# org1 create devices

curl --location --request POST 'localhost:3000/devices/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o1dev16",
    "description":"New Trial Device 001",
    "dataDescription":"Same random 01 data",
    "deviceSecret":"try001--secret"
}'

sleep 3

curl --location --request POST 'localhost:3000/devices/update' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o1dev16",
    "description":"new details of device 002",
    "on_sale":false
}'

# ###########
# sleep 3

# curl --location --request POST 'localhost:3000/devices/register' \
# --header 'Content-Type: application/json' \
# --data-raw '{
#     "deviceId": "o1dev2",
#     "description":"New org 1 Trial Device 002",
#     "dataDescription":"Same random 01 data",
#     "deviceSecret":"try002--secret"
# }'

# sleep 3

# curl --location --request POST 'localhost:3000/devices/update' \
# --header 'Content-Type: application/json' \
# --data-raw '{
#     "deviceId": "o1dev2",
#     "description":"new details of device 002",
#     "on_sale":true
# }'

# #############
# sleep 3

# curl --location --request POST 'localhost:3000/devices/register' \
# --header 'Content-Type: application/json' \
# --data-raw '{
#     "deviceId": "o1dev3",
#     "description":"New org 1 Trial Device 003",
#     "dataDescription":"Same random 03 data",
#     "deviceSecret":"try003--secret"
# }'

# sleep 3

# curl --location --request POST 'localhost:3000/devices/update' \
# --header 'Content-Type: application/json' \
# --data-raw '{
#     "deviceId": "o1dev3",
#     "description":"new details of device 002",
#     "on_sale":true
# }'
