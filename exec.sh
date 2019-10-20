docker build -t go-boilerplate .

docker run -d --rm --network="host" --name gogram -e AUTH_KEY=$AUTH_KEY go-boilerplate