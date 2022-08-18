local_database_setup:
	@mysql.server start

clean:
	@mysql.server stop