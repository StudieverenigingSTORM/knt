# The new KnT backend
The new KnT backend, powered by Go.

It exposes a web API, accessible on port 5000.
See [the Swagger documentation](docs/knt-backend-doc.yml) in the [Swagger editor](https://editor.swagger.io) to learn more about what is exposed how for other applications (frontend, admin panel).

## Makefile
You can use a Make:
- `make`: build the program
- `make dev`: run the program w/o making a binary
- `make clean`: clean (remove the binary)
- `make docker`: build docker image

## Docker
To facilitate usage, Docker has been added.

Docker usage can be simplified to 3 steps:
1. `make docker`
2. `docker compose up -d`
3. `docker compose down`

## KnT Demo API keys
### user
de7d235b14f6dec69f9795e0b6c9d5b8e775919ee6f338d26d623d6a77a94da8
8d5544aa6de35e3cc221e4352a8608fd0be0665a5799b57af14c92a66631d3fe

### admin
e9391609b1c58d48242c9eae5d09f8500a8cf0a490c3fc75192df8b650fe50e9
b5c1eb7b3ff826f2876396004178cdd32e0054bcade806787e0bf07a72764f45
