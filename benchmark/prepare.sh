sudo sysctl -w net.inet.ip.portrange.first=32768

curl -X POST localhost:8080/expression \
    -H 'Content-Type: application/json ' \
    -d @complex-event-1.json
