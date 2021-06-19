# ext2 : Create a vpc and it's subnets.
#Declare providers and region
terraform {
    required_providers {
      aws = {
          source = "hashicorp/aws"
          version = "~> 3.0"
      }
    }
}

provider "aws" {
    region = "ap-south-1"
}

# Creating a vpc
resource "aws_vpc" "tfvpc" {
    cidr_block = "192.168.1.0/24"
    tags = {
        Name = "tflabs"
        mode_of_creation = "Terraform"
    }
}

# Creating a subnet
resource "aws_subnet" "tfsubnet-1" {
    vpc_id = aws_vpc.tfvpc.id
    cidr_block = "192.168.1.0/26"
    tags = {
        Name = format("%s-public-1a", aws_vpc.tfvpc.tags.Name)
        mode_of_creation = "Terraform"
    }
    availability_zone = "ap-south-1a"
    
}

resource "aws_subnet" "tfsubnet-2" {
    vpc_id = aws_vpc.tfvpc.id
    cidr_block = "192.168.1.64/26"
    tags = {
        Name = format("%s-public-1b", aws_vpc.tfvpc.tags.Name)
        mode_of_creation = "Terraform"
    }
    availability_zone = "ap-south-1b"
}

resource "aws_subnet" "tfsubnet-3" {
    vpc_id = aws_vpc.tfvpc.id
    cidr_block = "192.168.1.128/26"
    tags = {
        Name = format("%s-private-1a", aws_vpc.tfvpc.tags.Name)
        mode_of_creation = "Terraform"
    }
    availability_zone = "ap-south-1a"
}

resource "aws_subnet" "tfsubnet-4" {
    vpc_id = aws_vpc.tfvpc.id
    cidr_block = "192.168.1.192/26"
    tags = {
        Name = format("%s-private-1b", aws_vpc.tfvpc.tags.Name)
        mode_of_creation = "Terraform"
    }
    availability_zone = "ap-south-1b"
}

# Timestamp : 1:13:16