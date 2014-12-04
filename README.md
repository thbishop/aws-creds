## aws-creds
[![Build Status](https://travis-ci.org/thbishop/aws-creds.svg?branch=master)](https://travis-ci.org/thbishop/aws-creds)

This utility reads credentials from the AWS credential file profiles and prints
them for use with other tools.

## Install

Download the latest binary or
`brew tap thbishop/aws-creds && brew install aws-creds` if you're on OSX

## Usage

```sh
aws-creds --profile foo-bar
```

Where `--profile` is a profile in your `~/.aws/credentials` file. If you do not
provide a profile, then it will use `default`.

There are two modes of printing the credentials.

The first is if you want one line style env vars with `env`.
```sh
env $(aws-creds --profile foo) ./other-tool
```
This is the equivalent of doing:
```sh
env AWS_ACCESS_KEY_ID=foo AWS_SECRET_ACCESS_KEY=bar ./other-tool
```

The second mode will prepend the `export` command to each item. This is useful
should you just want to export the env vars in your shell.
```sh
$(aws-creds --profile foo --export) && ./other-tool
```
This is the equivalent of doing:
```sh
export AWS_ACCESS_KEY_ID=foo
export AWS_SECRET_ACCESS_KEY=bar && ./other-tool
```


## Contribute
* Fork the project
* Make your feature addition or bug fix (with tests and docs) in a topic branch
* Make sure tests pass
* Send a pull request and I'll get it integrated

## LICENSE
See [LICENSE](LICENSE)
