mock:
	mockery \
		--name TxInterface \
		--name DBInterface \
		--name StmtInterface \
		--name RowInterface \
		--name RowsInterface \
		--dir sql \
		--output sql/mocks

protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/pbTimestamp/timestamp.proto
