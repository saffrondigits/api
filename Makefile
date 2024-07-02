up:
	migrate -path sql -database "postgresql://root:Secret@localhost:3333/golang_class?sslmode=disable" -verbose up
down:
	migrate -path sql -database "postgresql://root:Secret@localhost:3333/golang_class?sslmode=disable" -verbose down 1
mig-gen:
	migrate create -ext sql -dir sql -seq alter_password