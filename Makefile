DB_URL=mysql://root:ligmaballs@tcp(localhost:3306)/go_simple_restful

dbup:
	migrate -database "$(DB_URL)" -path db/migrations up $(n)

dbdown:
	migrate -database "$(DB_URL)" -path db/migrations down $(n)

dbcreate:
	migrate create -ext sql -dir db/migrations -seq $(name)

dbclean:
	migrate -database "$(DB_URL)" -path db/migrations drop -f

dbv:
	migrate -database "$(DB_URL)" -path db/migrations version