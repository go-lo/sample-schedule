# sample-schema

This repo provides a dirt simple load test schedule to show the simplest, happiest-pathiest go-lo loadtest.

## Pre-requisites

You will need:
 1. A go-lo stack built and running
 1. [`golo-cli`](github.com/go-lo/golo-cli) installed and running
 1. The golang compiler/ ecosystem from your package manager or [from here](https://golang.org/dl/)
 1. This repo
 1. Some service you want to loadtest

## Compiling the schedule

The simplest compilation is:

```bash
$ CGO_ENABLED=0 go build -ldflags "-X 'main.endpoint=https://example.com'" -o schedule
```

This will statically compile any dependant libraries into the binary- this will help ensure your schedule can run on remote hosts without having to install extra libraries.

If, however, you're on a non-linux machine **and you're running go-lo with docker or remotely on linux** you will need to do:

```bash
$ CGO_ENABLED=0 GOOS=linux go build -ldflags "-X 'main.endpoint=https://example.com'" -o schedule
```

Note the addition of `GOOS=linux`.

## Starting the schedule

If running on localhost:

```bash
$ golo-cli -f sched.toml
```

If running on a remote host:

```bash
$ GOLO_HOSTS=10.0.0.12 golo-cli -f sched.toml --provider env
```

For further options, see [github.com/go-lo/golo-cli](github.com/go-lo/golo-cli)


## Viewing the results

Visit your chronograf node, open the `Data Explorer` tab and run:

```
SELECT count("size") AS "count_size" FROM "simple_job_runner"."autogen"."request" WHERE time > now() - 5m GROUP BY time(1s), "url"
```

This should show you a graph of requests per second by url, and is the simplest chart possible.
