# docker compose up consul sm-db -d


# curl -T data.json http://192.168.49.2:32086/v1/kv/grpc-svc
# curl -T auth.json http://192.168.49.2:32086/v1/kv/auth

# curl -T data.json http://127.0.0.1:8500/v1/kv/grpc-svc
curl -T auth.json http://127.0.0.1:8500/v1/kv/auth
curl -T app.json http://127.0.0.1:8500/v1/kv/app

