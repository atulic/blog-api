### Getting Started

The following makes the assumption that Go is already installed on your machine. To verify, run `go version`. Alternatively, we can run in a Docker container.

The server will run by default on localhost:4000.

#### To build:

As we are using go modules, simply run `go build` inside the cloned directory. This will download required dependencies and generate an executable binary file.

#### To run in Docker

Make sure you have `docker`, and `docker-compose` installed. Then, in the `server` directory run `docker-compose up --build`, which will create a container with our Postgres db, and one with our API. The API will expose port 4000 on your host machine.

#### To start the Frontend

The React Frontend is not currently being run in a Docker container. To start it, `cd` into the `client` directory, then run `npm install` followed by `npm start`.

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

#### Updating an existing record

```
mutation {
  update(id: 1, title: "I've updated the title", content: "And the content too!") {
    title
  }
}
```

#### Deleting an existing record:

```
mutation {
  delete(id:1) {
    title
  }
}
```
