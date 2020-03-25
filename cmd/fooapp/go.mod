module github.com/foobarfran/gommlayout/cmd/fooapp

go 1.14

require (
	github.com/foobarfran/gommlayout/pkg/fooapi v0.0.0
	github.com/gin-gonic/gin v1.6.0
)

replace github.com/foobarfran/gommlayout/pkg/fooapi => ../../pkg/fooapi
