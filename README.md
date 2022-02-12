# Ent GraphQL Connections

Minimal example on how to use [GraphQL Cursor Connections][1] using [Ent][2].

### Usage

Inside the repository exec `go run ./cmd/users`:

```shell
{
  "edges": [
    {
      "node": {
        "id": 1,
        "name": "a8m"
      },
      "cursor": {
        "ID": 1,
        "Value": null
      }
    }
  ],
  "pageInfo": {
    "hasNextPage": true,
    "hasPreviousPage": false,
    "startCursor": {
      "ID": 1,
      "Value": null
    },
    "endCursor": {
      "ID": 1,
      "Value": null
    }
  },
  "totalCount": 3
}
{
  "edges": [
    {
      "node": {
        "id": 2,
        "name": "bodo"
      },
      "cursor": {
        "ID": 2,
        "Value": null
      }
    }
  ],
  "pageInfo": {
    "hasNextPage": true,
    "hasPreviousPage": false,
    "startCursor": {
      "ID": 2,
      "Value": null
    },
    "endCursor": {
      "ID": 2,
      "Value": null
    }
  },
  "totalCount": 3
}
{
  "edges": [
    {
      "node": {
        "id": 3,
        "name": "yoni"
      },
      "cursor": {
        "ID": 3,
        "Value": null
      }
    }
  ],
  "pageInfo": {
    "hasNextPage": false,
    "hasPreviousPage": false,
    "startCursor": {
      "ID": 3,
      "Value": null
    },
    "endCursor": {
      "ID": 3,
      "Value": null
    }
  },
  "totalCount": 3
}
```

[1]: https://relay.dev/graphql/connections.htm
[2]: https://entgo.io/docs/tutorial-todo-gql-paginate
