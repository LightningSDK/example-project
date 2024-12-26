module github.com/lightningsdk/example

go 1.22.0

toolchain go1.23.4

require (
	github.com/lightningsdk/core v1.0.1 //
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require golang.org/x/exp v0.0.0-20241217172543-b2144cdd0a67 // indirect

replace github.com/lightningsdk/ui => ../ui

replace github.com/lightningsdk/core => ../core

replace github.com/lightningsdk/blog => ../blog

replace github.com/lightningsdk/db => ../db
