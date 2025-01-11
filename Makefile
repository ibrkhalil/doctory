run:
	go run .

watch:
	 nodemon -V -e .go -w . -x go run . --count=1 --race -V --signal SIGTERM
