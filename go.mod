module github.com/Komdosh/go-bookstore-oauth-client

go 1.15

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0

require (
	github.com/Komdosh/go-bookstore-utils v0.0.0-20210116090756-a7d3cfa03af1
	github.com/go-resty/resty v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
	gopkg.in/resty.v1 v1.12.0 // indirect
)
