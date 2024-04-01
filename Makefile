test_request_product_insert:
	ab \
	-n 100 \
	-c 10 \
	-T application/json \
	-p test/product.json \
	-m POST \
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