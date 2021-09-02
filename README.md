# Num To Words

Implements a routine that takes an integer number in the range 1 - 9999 and convert that
number to words.

[![Build Status](https://github.com/dikaeinstein/numtowords/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/dikaeinstein/numtowords/actions)
[![Coverage Status](https://coveralls.io/repos/github/dikaeinstein/numtowords/badge.svg?branch=main)](https://coveralls.io/github/dikaeinstein/numtowords?branch=main)

## Prerequisites

- make
- GO 1.17+

### Question 2

The terraform configuration files solution for the question is stored in the
`terraform` directory. To create the infrastructure for the web site:

```sh
    cd terraform
    terraform init
    terraform plan
    terraform apply
```

Retrieve the website endpoint from the terraform output to test the website.
After testing, delete the s3 bucket files and destroy the infrastructure:
`terraform destroy`
