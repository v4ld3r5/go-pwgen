# go-pwgen
Demo app inspired by `pwgen` that provides a REST API. Do not use, this is only for playground :)

# Usage
As a pre-requisite you need to have golang available in your `$PATH`. I like to use `goenv` to manage local versions. You will find a `.go-version` file to lock in this version:

```bash
$ go version
go version go1.19.3 darwin/amd64
```

Simply clone this repository:

```bash
git clone https://github.com/v4ld3r5/go-pwgen.git
cd go-pwgen
```

Get the dependencies:

```bash
go mod download
```

And you can just use `go run` to launch it. It will launch the HTTP server `gin`and listen locally at port `8080`:

```bash
go run .

...
[GIN-debug] Listening and serving HTTP on :8080
```

From a different terminal, you can use `curl` to send a `POST` request. The data to be provided in the request in JSON format as follows:
- `minLen`: minimum password length 
- `schar`: number of special characters in the password
- `nchar`: number of numbers in the password
- `num`: number of passwords that must be created

Example:

```json
curl -s http://localhost:8080/pwgen \
-H "Content-Type: application/json" -X POST \
--data '{"minLen":8, "schar":1, "nchar":2, "num":4}' | jq
```

Will generate an output similar to the following. I've used `jq` at the end to obtain an easier to read output:

```json
[
  "fdzP2Y=kaG4yhp",
  "V~sRX9yLUz2ETO",
  "9YjVZ(HI1SsPkU"
  "kDhTcC+60dlavn"
]
```

# Deployment

You can build this util into a container and run it with `docker` and `kubernetes`. Use the provided `Dockerfile`:

```bash
docker build -t go-pwgen:latest .
```

To run it on a `k8s` you can try (not tested):

```bash
kubectl apply -f ./k8s
```
