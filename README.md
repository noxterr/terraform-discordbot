### Discord bot written with HASHICORP's terraform

Copy the `variables .example.tfvars` file into a file named `variables.tf`

Change the variables inside `variables.tf` with the real data

Run those commands
```bash
terraform init &&
terraform fmt &&
terraform validate
```

The expected result is a success message with `Success! The configuration is valid.` and a bunch of new files