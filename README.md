# url-shortener

This is a simple URL shortener written in Go, with MongoDB as the database.
It's implementation is inspired by the URL shortener system design case study done in Alex Xu's popular book [here](https://www.amazon.com/System-Design-Interview-insiders-Second/dp/B08CMF2CQF)

Here are some useful steps to set it up:

```
docker-compose up -d // setup
docker-compose down // teardown

go test ./tests // testing
```

The mongo shell tutorial [here](https://www.mongodb.com/docs/v4.4/mongo/) will also be useful for debugging
