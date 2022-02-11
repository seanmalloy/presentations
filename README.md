## Presentations
Here is where I keep all of my presentations. I'm using
the [Go present tool](https://godoc.org/golang.org/x/tools/present). PDFs
of the presentations can be found [here](pdf).

## Customize
Files in the `base` directory can be updated to change to over all look
and feel of all presentations in this repo.

## Install
```
$ go get golang.org/x/tools/cmd/present
$ go install golang.org/x/tools/cmd/present@latest
```
## Run
```
$ git clone git@github.com:seanmalloy/presentations.git
$ cd presentations
$ present -base ./base
```
