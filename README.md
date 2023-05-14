
# Mutants-meli

A simple API made with Go version go1.20.4 for the backend developer exam at MELI


## Run Locally

Clone the project

```bash
  git clone https://github.com/Garcia96/mutants-meli
```

Go to the project directory

```bash
  cd mutants-meli
```

Install dependencies

```bash
  go mod download
```

To run this project, you will need to add the following environment variables to your .env file

`port` `db_host` `db_user` `db_pass` `db_name` `db_port`


Start the server

```bash
  go run main.go
```

## API Reference

#### Mutant

```http
  POST /mutant
```
Body

```json
{
  "dna": [
        "ATGCGA",
        "CAGTGC",
        "TTATGT",
        "AGAAGG",
        "CCCCTA",
        "TCACTG"
    ]
}
```

#### Get Stats

```http
  GET /stats
```
Response

```json
{
    "count_mutant_dna": 1,
    "count_human_dna": 1,
    "ratio": 1
}
```
## Running Tests

To run tests, run the following command

```bash
  go test -v
```


## Build

To build the project, run the following command

```bash
  go build
```

## API URL

You can test the API with postman or similar tool

- http://ec2-3-139-80-95.us-east-2.compute.amazonaws.com/

