# Key things

- upload the template to s3 bucket, cloudformation refers the same and creates the stack.
- cannot edit the template in place, it has to be reuploaded.
- either upload the template into CF console while creating the stack or upload it on s3 and
  then refer the s3 url in the console.
- if you upload the template into the CF console directly, cloudformation will create a bucket
  with a unique id in your account and store all tempates over there.
- template can be in json form or yml. yml is more readable and easier to write.

# Syntax of a CF template:

```
Resources:
  # logical id
  MyAWSResource:
    Type: aws_resource_type_string
    Properties:
      key: value
```

## CF update behaviour

### update with no interruption

- no change in physical ID. 
- no interruption.

### update with some interruption

- no chage in physical ID. 
- interruption to services present. For eg: updating instance type of an ec2 instance will 

### replacement

- physical ID changes in this behaviour
- a new resource is created and the old one gets deleted. The old references are now moved to the the new one.
> This update behaviour is very critical so 

## Parameters common to CF templates.

- Tags (will be applied to all resources involved)
- Permissions (IAM role required for creating cloudformation resources)
- Notification Options (relay CF events to SNS topics)
- Timeout (by default there is no )
- Rollback on Failure (default behaviour of CF)
- Rollback configuration (rollback if any cloudwatch alarm configured goes in alarm state.)
- Stack Policy
- Termination Protection (so that stack doesn't get accidentally deleted.)
- Quick-start link (to get stacks up and running quickly from the AWS CloudFormation console)


