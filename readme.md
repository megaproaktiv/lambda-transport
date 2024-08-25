# Lambda Transport

Transfer lambda CODE between accounts and functions

## Configuration

The source and target account and lambda are defined in the file ``.transport/config.yml`.

This is the example you find in `config-example.yml`:

```yaml
config:
  dev:
    source:
      profile: dev-profile
      region: eu-central-1
      lambda: demo
    target:
      profile: test-profile
      region: eu-central-1
      lambda: demo

  test:
    source:
      profile: dev-profile
      region: eu-central-1
      lambda: demo
    target:
      profile: prod-profile
      region: eu-central-1
      lambda: demo
```

With the paramater `stage` you select an stage entry from the config file.

So `transport --stage dev` will transfer the code from the lambda `demo` with the AWS profile `dev-profile` to the lambda `demo` with the AWS profile `test-profile`.

So first copy the example file to `.transport/config.yml` and then adjust the values.

1) `cp config-example.yml .transport/config.yml`
2) `vi .transport/config.yml`
