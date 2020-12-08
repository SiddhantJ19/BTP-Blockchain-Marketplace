org1 create devices
sleep 3

curl --location --request POST 'localhost:4000/devices/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o2dev1",
    "description":"New Org2 Trial Device 001",
    "dataDescription":"Same random 01 data",
    "deviceSecret":"try001--secret"
}'

sleep 3

curl --location --request POST 'localhost:4000/devices/update' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o2dev1",
    "description":"new details of device 001",
    "on_sale":true
}'

sleep 3
##########3

curl --location --request POST 'localhost:4000/devices/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o2dev2",
    "description":"New org 2 Trial Device 002",
    "dataDescription":"Same random 01 data",
    "deviceSecret":"try002--secret"
}'

sleep 3

curl --location --request POST 'localhost:4000/devices/update' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o2dev2",
    "description":"new details of device 002",
    "on_sale":false
}'

sleep 3
########

curl --location --request POST 'localhost:4000/devices/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o2dev3",
    "description":"New org 2 Trial Device 003",
    "dataDescription":"Same random 03 data",
    "deviceSecret":"try003--secret"
}'

sleep 3

curl --location --request POST 'localhost:4000/devices/update' \
--header 'Content-Type: application/json' \
--data-raw '{
    "deviceId": "o2dev3",
    "description":"new details of device 003",
    "on_sale":true
}'

#######

# sleep 3
# curl --location --request POST 'localhost:4000/devices/register' \
# --header 'Content-Type: application/json' \
# --data-raw '{
#     "deviceId": "o2dev4",
#     "description":"New org 2 Trial Device 003",
#     "dataDescription":"Same random 03 data",
#     "deviceSecret":"try004--secret"
# }'

# sleep 3

# curl --location --request POST 'localhost:4000/devices/update' \
# --header 'Content-Type: application/json' \
# --data-raw '{
#     "deviceId": "o2dev4",
#     "description":"new details of device 003",
#     "on_sale":true
# }'

# #######

# sleep 3
# curl --location --request POST 'localhost:4000/devices/register' \
# --header 'Content-Type: application/json' \
# --data-raw '{
#     "deviceId": "o2dev5",
#     "description":"New org 2 Trial Device 003",
#     "dataDescription":"Same random 03 data",
#     "deviceSecret":"try004--secret"
# }'

# sleep 3

# curl --location --request POST 'localhost:4000/devices/update' \
# --header 'Content-Type: application/json' \
# --data-raw '{
#     "deviceId": "o2dev5",
#     "description":"new details of device 003",
#     "on_sale":true
# }'

# #######

# sleep 3
# curl --location --request POST 'localhost:4000/devices/register' \
# --header 'Content-Type: application/json' \
# --data-raw '{
#     "deviceId": "o2dev6",
#     "description":"New org 2 Trial Device 003",
#     "dataDescription":"Same random 03 data",
#     "deviceSecret":"try004--secret"
# }'

# sleep 3

# curl --location --request POST 'localhost:4000/devices/update' \
# --header 'Content-Type: application/json' \
# --data-raw '{
#     "deviceId": "o2dev6",
#     "description":"new details of device 003",
#     "on_sale":true
# }'