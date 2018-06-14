# goroyale
A Golang wrapper for the Clash Royale API at https://royaleapi.com/.

## Installing
If you have Go installed you can run this command.
```sh
go get github.com/meliorihi/goroyale
```

## Ratelimits
If you hit the RoyaleAPI ratelimit, the lib will just refuse to run your request and return an error message in this format:
```
ratelimit, retry in: <milliseconds>
```
There are a few minor issues with this.

- Some endpoints work differently with ratelimits (such as /version)
- It's not super easy to work with

In a future version this will be improved.