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
    }
}
