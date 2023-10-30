# profile = "default"

environment = "poc"

vpc_config = {
    tgw-poc-vpc-01 = {
        vpc_name = "tgw-poc-vpc-01"
        vpc_cidr = "10.10.0.0/24"
        private_subnets = ["10.10.0.0/26", "10.10.0.64/26"]
        public_subnets = ["10.10.0.128/26", "10.10.0.192/26"]
    },
    tgw-poc-vpc-02 = {
        vpc_name = "tgw-poc-vpc-02"
        vpc_cidr = "10.11.0.0/24"
        private_subnets = ["10.11.0.0/26", "10.11.0.64/26"]
        public_subnets = ["10.11.0.128/26", "10.11.0.192/26"]
    },
    tgw-poc-vpc-03 = {
        vpc_name = "tgw-poc-vpc-03"
        vpc_cidr = "10.12.0.0/24"
        private_subnets = ["10.12.0.0/26", "10.12.0.64/26"]
        public_subnets = ["10.12.0.128/26", "10.12.0.192/26"]
    }
}