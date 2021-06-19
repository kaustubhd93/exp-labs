# Provider specific declaration block.
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


# Create a VPC
resource "aws_vpc" "tfvpc" {
    cidr_block = "192.168.1.0/24"
    enable_dns_hostnames = true
    tags = {
        Name = "tflabs"
        mode_of_creation = "Terraform"
    }
}

# Create subnets
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

# Create an internet gateway
resource "aws_internet_gateway" "tfigw" {
    vpc_id = aws_vpc.tfvpc.id

    tags = {
        Name = format("%s-igw", aws_vpc.tfvpc.tags.Name)
        mode_of_creation = "Terraform"
    }
}

# Create a Elastic IP for NAT gateway 
resource "aws_eip" "tfeip" {
    vpc = true
    depends_on = [
      aws_internet_gateway.tfigw
    ]
}

# Create a NAT gateway for private subnet
resource "aws_nat_gateway" "tfnatgw" {
    allocation_id = aws_eip.tfeip.id
    subnet_id = aws_subnet.tfsubnet-1.id

    tags = {
        Name = format("%s-natgw", aws_vpc.tfvpc.tags.Name)
        mode_of_creation = "Terraform"
    }

    depends_on = [
      aws_internet_gateway.tfigw
    ]
}

# Create custom route table 
resource "aws_route_table" "tf-rt-public" {
    vpc_id = aws_vpc.tfvpc.id
    # The local default route is automatically created. 
    # There is no option to mention it explicitly either.
    route {
        # Destination
        cidr_block = "0.0.0.0/0"
        # Target
        gateway_id = aws_internet_gateway.tfigw.id
    }
    tags = {
        Name = format("rt-%s-public", aws_vpc.tfvpc.tags.Name)
        mode_of_creation = "Terraform"
    }
  
}

resource "aws_route_table" "tf-rt-private" {
    vpc_id = aws_vpc.tfvpc.id
    route {
        # Destination
        cidr_block = "0.0.0.0/0"
        # Target
        nat_gateway_id = aws_nat_gateway.tfnatgw.id
    }
    tags = {
        Name = format("rt-%s-private", aws_vpc.tfvpc.tags.Name)
        mode_of_creation = "Terraform"
    }
  
}

# Associate subnet with route table. 
resource "aws_route_table_association" "tf-rt-public-assc-1" {
    subnet_id = aws_subnet.tfsubnet-1.id
    route_table_id = aws_route_table.tf-rt-public.id
}

resource "aws_route_table_association" "tf-rt-public-assc-2" {
    subnet_id = aws_subnet.tfsubnet-2.id
    route_table_id = aws_route_table.tf-rt-public.id
}

resource "aws_route_table_association" "tf-rt-public-assc-3" {
    subnet_id = aws_subnet.tfsubnet-3.id
    route_table_id = aws_route_table.tf-rt-private.id
}

resource "aws_route_table_association" "tf-rt-public-assc-4" {
    subnet_id = aws_subnet.tfsubnet-4.id
    route_table_id = aws_route_table.tf-rt-private.id
}

# Create security group to allow port 22, 80, 443
resource "aws_security_group" "tf-sg-web-ssh" {
    name = format("%s-sg-web-ssh", aws_vpc.tfvpc.tags.Name)
    description = "Allow Web and SSH traffic"
    vpc_id = aws_vpc.tfvpc.id

    ingress {
        description = "TLS"
        from_port = 443
        to_port = 443
        protocol = "tcp"
        cidr_blocks = ["103.140.27.185/32"]
    }

    ingress {
        description = "HTTP"
        from_port = 80
        to_port = 80
        protocol = "tcp"
        cidr_blocks = ["103.140.27.185/32"]
    }

    ingress {
        description = "SSH"
        from_port = 22
        to_port = 22
        protocol = "tcp"
        cidr_blocks = ["103.140.27.185/32"]
    }

    egress {
        description = "Allow all outbound"
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }

    tags = {
        Name = format("%s-sg-web-ssh", aws_vpc.tfvpc.tags.Name)
        mode_of_creation = "Terraform"
    }
}

# Launch an ubuntu server and install apache on it


