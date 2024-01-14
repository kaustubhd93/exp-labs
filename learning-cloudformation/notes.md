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

## Building blocks

- AWSTemplateFormatVersion: (Identifies capabilities of a template)
- Description: (Comments about template)
- Transform: (macros for template processing)
- Metadata:
- Resources: (Mandatory)
- Outputs: (view resources that have been created, similar to Terraform output)
- Conditionals:
- Rules: (validate a parameter(s) during stack creation/update)

### Template's Helpers

- References
- Functions

## CF parameters

- Parameters are like global variables for your templates which makes the template reusuable.
- a way to provide input for CF templates
- validation for parameters can be added. 

### Referring blocks in CF

- `Fn::Ref`. Since this is a little difficult to write all the time, we can use it's short form `!Ref`. Any `Fn::` can be written as `!`. If you have a shell programming background, ensure this `!` doesnt mean `not` or `false`

### SSM parameters

- There can be values like image-id that can change according to region, in this case we can use an SSM parameter (a feature of Systems Manager) where in you can use AWS public SSM parameters to compute the image id automatically based on the region. This keeps the template same across the region. 
- This can also be used to make your own custom parameter which you can refer in the template. You can make custom SSM parameters manually or by Cloudformation Resources with `Type: AWS::SSM::Parameter`.
- If you make a change in the SSM parameter directly, it can trigger an update in the linked CF stack where the SSM parameter is used. For eg: an SSM parameter like `/dev/web-server/ec2/instance-type` if changed from `t2.micro` to `t3a.micro`, it will trigger an update-stack to the CF stack where this parameter is being used. 

## DependsOn

- you can use the `DependsOn` option to specify any kind of dependencies your individual resource components have. For example:
```
Resources:
  MyS3Bucket:
    Type: AWS::S3::Bucket

  # the EC2 instance will be created after the S3 bucket
  MyInstance:
    Type: AWS::EC2::Instance
    Properties:
      ImageId: !Ref ImageId
      InstanceType: t2.micro
    DependsOn: MyS3Bucket
```

## UpdateReplacePolicy

- lets say there is property you wish to update, the update behaviour of this property in cloudformation is `Replacement`. `UpdateReplacePolicy` can control this behavior. 
- default behaviour of UpdateReplacePolicy = Delete. You can set this to Retain or Snapshot according to your requirement. 
- When you set to Retain, the resource gets removed from Cloudformation's setup.
- Similarly when you set it to Snapshot, the snapshot taken gets removed from the CF's scope.
