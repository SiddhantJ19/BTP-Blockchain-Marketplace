curl --location --request POST 'localhost:3000/users/connect' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userName":"admin"
}'