# validate-alloy-syntax

Performs basic syntax validation of [Grafana Alloy](https://github.com/grafana/alloy) configurations (matching parentheses, only valid characters, etc...). Does not validate things at the service or component level. 

## Usage

### Example Workflow
```yaml
name: My Workflow
on: [push, pull_request]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4
    - name: Validate Alloy Syntax
      uses: spartan0x117/validate-alloy-syntax@master
      with:
        directory: ./alloy_configs
```

### Inputs

| Input                                                | Description                                   |
|------------------------------------------------------|-----------------------------------------------|
| `directory`  | The directory containing Alloy configurations to validate.                            |

### Outputs

This action does not produce any outputs.
