test_request_product_insert:
	ab \
	-n 10000 \
	-c 100 \
	-T application/json \
	-p test/product.json \
	-m POST \
	-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJhYmEwMzk4MjYyMTI0QGdtYWlsLmNvbSIsImV4cCI6MTcxMzkwNDY3MSwicHJvZmlsZV9pZCI6Nywicm9sZSI6InVzZXIiLCJzdWIiOiIxMTc5OTY0MjMxODM5MDc1NzAyNTIiLCJ1dWlkIjoiNjUzMGJiN2MtMTViYi00NGZlLWI2OTQtNzcwZmZhMDg1MmViIn0.OB2vJQpNN6eT6VoDIPk4M8Q2KP7_LtUXuez0_kVAo2k" \
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