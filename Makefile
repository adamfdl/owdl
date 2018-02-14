build:
	docker-compose up --build -d

shell:
	@docker run --env-file=./.env --security-opt=seccomp:unconfined --name ow_bot -p 0.0.0.0:3009:3009 -v $(PWD):/go/src/github.com/adamfdl/ow_discord_leaderboard -it discord/ow_bot:latest bash || docker start -i ow_bot

attach:
	@docker attach mpg_v2_shell ||:

clean:
	-docker rm ow_bot
	@docker rmi -f discord/ow_bot