This is a Go CLI program that validates credit card numbers.

```shell
git clone https://github.com/mariosdaskalas/goval && cd goval && go run .
```

```shell
curl "localhost:8080/validate?number=374245455400126"
The number 374245455400126 is valid according to the Luhn algorithm
```

```shell
go version go1.24.1 linux/amd64
```

Test Card Numbers

https://support.bluesnap.com/docs/test-credit-card-numbers