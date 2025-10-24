`git clone git@github.com:aler93/AddressOnGo.git`

`cp .env.example .env`
- Set a password for Redis [REDIS_PASSWORD] (if using docker-compose)

`cp conf.json.example conf.json`
- Set a location for log file and the password (go will use to connect)

`docker compose up -d`
- To create the Redis connection, if necessary

`go build -o bin/AddressOnGo`
`./bin/AddressOnGo {zipcode}`
