TABLE_NAME ?= "users"

create-migration:
	goose -dir db/migrations create create_${TABLE_NAME}_table sql
	#usage : make create-migration TABLE_NAME=products


create-migration_user:
	goose -dir db/migrations create create_users_table sql
	
goose-up:
	goose -dir db/migrations postgres "postgres://postgres:root@localhost:5432/attempt2?sslmode=disable" up
	# cd sql/schema && goose postgres postgres://postgres:root@localhost:5432/attempt2 up

goose-down:
	goose -dir db/migrations postgres "postgres://postgres:root@localhost:5432/attempt2?sslmode=disable" down

goose-status:
	goose -dir db/migrations postgres "postgres://postgres:root@localhost:5432/attempt2?sslmode=disable"  status
