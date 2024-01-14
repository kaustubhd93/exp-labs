# profile = "default"
instance_type = "t3a.micro"
environment = "dev"
ami_id = "ami-0f5ee92e2d63afc18"

# vpc config
vpc_name = "terraform-poc-vpc"
vpc_cidr = "10.10.0.0/24"
private_subnets = ["10.10.0.0/26", "10.10.0.64/26"]
public_subnets = ["10.10.0.128/26", "10.10.0.192/26"]