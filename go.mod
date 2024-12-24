module github.com/lightningsdk/example

go 1.19

require (
	github.com/lightningsdk/blog v1.0.1 //
	github.com/lightningsdk/core v1.0.1 //
	github.com/lightningsdk/db v1.0.1 //
	github.com/lightningsdk/ui v1.0.1 //
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	dario.cat/mergo v1.0.0 // indirect
	github.com/creasty/defaults v1.7.0 // indirect
	github.com/jmoiron/sqlx v1.4.0 // indirect
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
	golang.org/x/net v0.20.0 // indirect
)

replace github.com/lightningsdk/ui => ../ui

replace github.com/lightningsdk/core => ../core

replace github.com/lightningsdk/blog => ../blog

replace github.com/lightningsdk/db => ../db
