# tfmap

Command line tool to convert environment variables into a Terraform map variable. 

Example: `MAP_foo=bar MAP_abc=123` = `{foo:bar, abc:123}`
## Installation

#### Locally
```sh
go get -u github.com/shanesavoie/tfmap
```

#### Dockerfile
```sh
FROM golang as tfmap
RUN go get -u github.com/shanesavoie/tfmap

FROM <your-image>
COPY --from=tfmap /go/bin/tfmap /usr/local/bin/tfmap
```

## Usage
* `--help` - Will list the available options.
* `--whitelist=MAP_` - Will only take in environment variables with the specefied prefix.
* `--export=TF_VARS_my_map` - Will export to the specified environment variable.

## Example
```sh
$ MAP_foo=bar MAP_abc=123 tfmap --whitelist=MAP_ --export=TF_VARS_ecs_environment
export TF_VARS_ecs_environment={"abc":"123","foo":"bar"}
```

## Apply to Terraform
```sh
eval "$(tfmap)"
terraform ... 
```
