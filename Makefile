start: 
	@docker compose up
	
stop:
	@docker compose rm -v --force --stop
	@docker rmi teacher-booking-system