# Leo's Music Shop API in Go (Gin)

This is another very cool API for Leo's Music Shop.

For now this API only sells keyboards and pianos.

## Motivation

This can help in Golang integrations with our clients. 

## Architecture

- [Gin](https://github.com/gin-gonic/gin)
- [Godotenv](https://github.com/joho/godotenv)
- [Gorm](https://gorm.io/)

## Running

First create a `.env` file with the following:

```ini
SEGMENT_WRITE_KEY=your-source-write-key
```

Install all required packages:

```sh
go get .
```

Then run it:

```sh
go run .
```