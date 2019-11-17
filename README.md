### CMS

App built with Go and using some [Buffalo](https://gobuffalo.io) tools.

## Clone and run the server

```bash
$ git clone https://github.com/gustavocd/cms.git && cd cms
```

```bash
$ go run cmd/cms.go
```

## Available endpoints

| Endpoint             | Description       | HTTP method |
| -------------------- | ----------------- | ----------- |
| `/api/v1/pages`      | Get all the pages | `GET`       |
| `/api/v1/pages/{id}` | Get a single page | `GET`       |
| `/api/v1/pages`      | Create a page     | `POST`      |
| `/api/v1/pages/{id}` | Edit a page       | `PUT`       |
| `/api/v1/pages/{id}` | Delete a page     | `DELETE`    |

Note: the main reason to build this app is for putting in practice some tricks explained in this article [How I write HTTP web services after eight years](https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831).
