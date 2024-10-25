# clean-gin
This repo represents the golang gin service implementation with repository pattern. Pull it and try it!

# Getting started

1.Spin up database with docker compose
```sh
docker compose up
```

2.Apply `db.sql` to postgres database

3.Run the cleangin application
```sh
go mod tidy
go run main.go
```

You should be able to see the following output:

```sh
$ go run main.go
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /api/v1/auth/register     --> cleangin/internal/api/handlers.UserHandler.Register-fm (5 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8123
```

Notes: 
- App will be available on port 8123 (change this in `clean-gin/internal/constants`)
- App can be tested with `rest.test.http` file (add [REST Client](https://www.linkedin.com/posts/milos-veljkovic-434271172_for-some-time-ive-been-searching-for-a-activity-7232393991516753920-Lr03?utm_source=share&utm_medium=member_desktop) extension to VSCode so you can use this)
