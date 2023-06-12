# sysctl -w net.inet.ip.portrange.first=32768

curl -X POST localhost:8080/expression \
    -H 'Content-Type: application/json ' \
    -d @queries/query-1.json

curl -X POST localhost:8080/expression \
    -H 'Content-Type: application/json ' \
    -d @queries/query-2.json

curl -X POST localhost:8080/expression \
    -H 'Content-Type: application/json ' \
    -d @queries/query-3.json

curl -X POST localhost:8080/expression \
    -H 'Content-Type: application/json ' \
    -d @queries/query-4.json
