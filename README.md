# aws-test
A little tool to fastly test if AWS IAM keys are valid

The idea for this tool came after a conversation with [Grisolfi](https://github.com/Grisolfi) and [rodrigoramosrs](https://github.com/rodrigoramosrs) about testing AWS related credentials.

This tool will perform a "sts get-caller-identity" without needing to have awscli installed on the machine or having to change your environment variables.

Testing credentials:

```
echo AKIAEXAMPLE000000 askcaExampleSecretGSkdsmcfklams/asdSDasmgkasd123 | aws-test
```

## Installation

First, you'll need to [install go](https://golang.org/doc/install).

Then run this command to download + compile aws-test:
```
go install github.com/edivangalindo/aws-test@latest
```

You can now run `~/go/bin/aws-test`. If you'd like to just run `aws-test` without the full path, you'll need to `export PATH="/go/bin/:$PATH"`. You can also add this line to your `~/.bashrc` file if you'd like this to persist.
