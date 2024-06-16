test_request_product_insert:
	ab \
	-n 10 \
	-c 10 \
	-T application/json \
	-p test/product.json \
	-m POST \
	-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJhYmEwMzk4MjYyMTI0QGdtYWlsLmNvbSIsImV4cCI6MTcxMzk4MzE0NSwicHJvZmlsZV9pZCI6Nywicm9sZSI6InVzZXIiLCJzdWIiOiIxMTc5OTY0MjMxODM5MDc1NzAyNTIiLCJ1dWlkIjoiMTFiYzBmYjMtODg3Yi00MGI1LTk1MTItOGI5MTBhNWZkODAyIn0.IJDjr0TLXLhdLGW_yL_zXo-MVH16WXjFJ1TVhwM1uWE" \
	http://localhost:18883/product/api/v1/protected/product
test_get_data:
	ab \
	-n 100000 \
	-c 10000 \
	-T application/json \
	-m GET \
	http://localhost:18883/product/api/v1/ping
run_server:
	go run .
gen_code_grpc:
	protoc \
	--go_out=grpc \
	--go_opt=paths=source_relative \
	--go-grpc_out=grpc \
	--go-grpc_opt=paths=source_relative \
	proto/*.proto
export_go:
	export PATH="$PATH:$(go env GOPATH)/bin"
gen_key:
	openssl \
		req -x509 \
		-nodes \
		-days 365 \
		-newkey rsa:2048 \
		-keyout keys/server-product/private.pem \
		-out keys/server-product/public.pem \
		-config keys/server-product/san.cfg