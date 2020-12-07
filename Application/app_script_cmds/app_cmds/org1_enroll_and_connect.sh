# org 1 enroll and connect admin
curl --location --request POST 'localhost:3000/users/admin/enroll'
curl --location --request POST 'localhost:3000/users/connect' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userName":"admin"
}'
