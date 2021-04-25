# Declare required providers
terraform {
    required_providers {
      aws = {
          source = "hashicorp/aws"
          version = "~> 3.0"
      }
    }
}

# Configure provider 
provider "aws" {
    region = "ap-south-1"
}

# Create a resource
# ami_id_ubuntu_18 = ami-0db0b3ab7df22e366
# poc-terraform is not the name of the instance, it's scope is upto 
# the instance only. 
# If this resource block is commented and then you run terraform apply then
# the corresponding ec2 instance will get terminated as terraform will identify
# the desired configuration and apply it accordingly.
resource "aws_instance" "poc" {
    ami = "ami-0db0b3ab7df22e366"
    instance_type = "t2.micro"
    tags = {
      "Name" = "poc-terraform"
      "mode_of_creation" = "terraform"
      "exercise" = "launching_instance_in_existing_vpc"
    }
    subnet_id = "subnet-0ed68e775e623b876"
    vpc_security_group_ids = [ "sg-0f5df33a8ceda0e30" ]
    # Override the subnet setting of auto assigning a public ipv4 address.
    associate_public_ip_address = true
    key_name = aws_key_pair.instance_ssh_key.key_name
    root_block_device {
      tags = {
      "Name" = "poc-terraform"
      "mode_of_creation" = "terraform"
      "exercise" = "launching_instance_in_existing_vpc"
      } 
    }
}


resource "aws_key_pair" "instance_ssh_key" {
  key_name = "kd-personal-mum"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDKudlNV0oiXDrv2X09gZaM6RKZI9a5TNw8VKq5uunITOe0enF1V4imKfMv+HFtu2dh2Yq3IeISmugBbSiruw2hQwjNkF8epawOBzYt4Cas0a0t7tnllXLNhkuXRRJ526fB8FDZbngcZGfDgCy/+J+dUUitLRx71JUS0mjEAApEMk03NN6x25LAL7Db2kj0EWGrYvBKqlYWI1oTZBn5IKJkqhVyiXS7LIl+OZrikAm8it4vjAFCoJESYiGMwcobbH3bKUDOUniLoPyFEf3Iyqr7MpQOpO1VCsyEDJBmJuo4wtVKBrgcMhMPuGJFbDKUmxujpJtnw6Y3rS2xl8/3kmXSDMxfKd/xOOPAEplWxYDbWVwehLouOM9O3K7X6fFw5mrHFrUBgf+hJjCJjnhoI9LfUU5vWULzG08oHfLCbScUQq2HEygoHvejb/k/MHSTzQLEej6HJ/L5J4tuRLQYQU4YBqZi2Xht8l8J8uM3+TfrmV+QgozIInFPUUi9lKQjsbc= kaustubh@techjunkie"
}

# Timestamp : 54:46

# Create a vpc 
# resource "aws_vpc" "poc" {
#   cidr_block = "10.31.0.0/16"
#   tags = {
#     "Name" = "poc"
#   }
# }

