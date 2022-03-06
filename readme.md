Sample CURL request for add family member:

curl --location --request POST 'http://localhost:8080/add' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "first_name": "Sushil",
    "last_name":"Shelake"
}'


Sample CURL for search member:

curl --location --request GET 'http://localhost:8080/search?name=Rahul' \
--data-raw ''