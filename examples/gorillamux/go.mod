module example-gorillamux

go 1.18

replace github.com/paul-nelson-baker/golib-regexstruct => ../..

require (
	github.com/gorilla/mux v1.8.0
	github.com/paul-nelson-baker/golib-regexstruct v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.1
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)
