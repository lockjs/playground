# Colorcli

Source: https://codesource.io/create-a-cli-in-golang-with-cobra/

## Run

```sh
docker_run \
    -v `pwd`/go:/go \ #Override the default GOPATH locaton
    golang:1.17 \
    <cmd> # e.g. bash
```

`go` commands should be run in the `/src` directory

## Notes

The initial setup instructions in the tutorial no longer work. Needed to follow the [Cobra Generator](https://github.com/spf13/cobra/blob/master/cobra/README.md) instructions instead