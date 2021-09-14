
compose:
	@docker-compose up

client:
	@PGPASSWORD=postgres psql -h 192.168.59.103 -U postgres -p 5432

.PHONY: client compose