# Learnyougo

This is a go port of the original [learnyounode](https://github.com/workshopper/learnyounode) workshop, that helped thousands of developers all around the world to learn nodejs on their own or during hosted workshop sessions.

Our initiative is to bring the same workshopper experience to the go community. Our goal is to make learning go more approachable for everyone.

## Building

We're bundling assets into the binary file using the [go-bindata package](https://github.com/a-urth/go-bindata) package. To run the project first generate the `bindata.go` file using `go-bindata ./exercises`
