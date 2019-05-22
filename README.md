### Getting Started

The following makes the assumption that Go is already installed on your machine. To verify, run `go version`.

The server will run by default on localhost:4000.

#### To start the development server:

First, download Realize:

`go get github.com/oxequa/realize`

Then run `realize start` to start the development server.


#### To build:

As we are using go modules, simply run `go build` inside the cloned directory. This will download required dependencies and build.


#### To run in Docker

Make sure you have `docker`, and `docker-compose` installed. Then, just run `docker-compose up --build`, which will create a container with our Postgres db, and one with our API. The API will expose port 4000 on your host machine.

### Queries and Mutations Examples

#### Querying an existing record:
```
{
  posts(id:1) {
    title
    content
  }
}
```

#### Creating a new record:
```
mutation {
  create(title: "How to create new posts", content: "Using GraphQL mutations") {
    title
  }
}
```