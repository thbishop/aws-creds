## aws-creds

This utility reads AWS credentials from AWS credential file profiles and exports
them as environment variables.

## Install

Download the latest binary or
`brew tap thbishop/aws-creds && brew install aws-creds` if you're on OSX

## Usage

```sh
aws-creds --profile foo-bar
```

Where `--profile` is a profile in your `~/.aws/credentials` file.

If you want to use aws-creds with another tool, you would do:
```sh
$(aws-creds --profile foo) && ./other-tool
```

## Contribute
* Fork the project
* Make your feature addition or bug fix (with tests and docs) in a topic branch
* Make sure tests pass
* Send a pull request and I'll get it integrated

## LICENSE
See [LICENSE](LICENSE)
