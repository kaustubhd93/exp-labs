terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.0.2"
    }
  }
}

# Configure the Microsoft Azure Provider
provider "azurerm" {
  features {}
}

locals {
  subnet_cidr = cidrsubnets(var.vpc_cidr, 1, 1)
}

# Create a resource group
resource "azurerm_resource_group" "k8s-sandbox" {
  name     = "k8s-sandbox-rg"
  location = var.location
  tags = {
    environment = "poc"
  }
}

# Create a virtual network within the resource group
resource "azurerm_virtual_network" "k8s-sandbox" {
  name                = "k8s-sandbox-vnet"
  resource_group_name = azurerm_resource_group.k8s-sandbox.name
  location            = azurerm_resource_group.k8s-sandbox.location
  address_space       = [var.vpc_cidr]
}

resource "azurerm_subnet" "k8s-sandbox" {
  name                 = "k8s-sandbox-subnet"
  resource_group_name  = azurerm_resource_group.k8s-sandbox.name
  virtual_network_name = azurerm_virtual_network.k8s-sandbox.name
  address_prefixes     = local.subnet_cidr
}

resource "azurerm_network_interface" "k8s-sandbox" {
  name                = "k8s-sandbox-nic"
  location            = azurerm_resource_group.k8s-sandbox.location
  resource_group_name = azurerm_resource_group.k8s-sandbox.name

  ip_configuration {
    name                          = "k8s-sandbox-ip-config"
    subnet_id                     = azurerm_subnet.k8s-sandbox.id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.k8s-sandbox.id
  }
}

resource "azurerm_public_ip" "k8s-sandbox" {
  name                = "k8s-sandbox-public-ip"
  resource_group_name = azurerm_resource_group.k8s-sandbox.name
  location            = azurerm_resource_group.k8s-sandbox.location
  allocation_method   = "Dynamic"
}

resource "azurerm_linux_virtual_machine" "k8s-sandbox" {
  name                = "k8s-sandbox-vm"
  resource_group_name = azurerm_resource_group.k8s-sandbox.name
  location            = azurerm_resource_group.k8s-sandbox.location
  #size                = "Standard_B2s"
  size                = var.vm_size
  admin_username      = "ubuntu"
  network_interface_ids = [
    azurerm_network_interface.k8s-sandbox.id,
  ]

  custom_data = filebase64("custom_data.tpl")

  admin_ssh_key {
    username   = "ubuntu"
    public_key = file(var.ssh_key_file_path)
  }

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-jammy"
    sku       = "22_04-lts"
    version   = "latest"
  }
}

resource "azurerm_network_security_group" "k8s-sandbox" {
  name                = "k8s-sandbox-nsg"
  location            = azurerm_resource_group.k8s-sandbox.location
  resource_group_name = azurerm_resource_group.k8s-sandbox.name
}

resource "azurerm_network_security_rule" "k8s-sandbox" {
  name                        = "k8s-sandbox-ns-rule-inbound"
  priority                    = 100
  direction                   = "Inbound"
  access                      = "Allow"
  protocol                    = "Tcp"
  source_port_range           = "*"
  destination_port_range      = "*"
  source_address_prefix       = var.source_ip_cidr
  destination_address_prefix  = "*"
  resource_group_name         = azurerm_resource_group.k8s-sandbox.name
  network_security_group_name = azurerm_network_security_group.k8s-sandbox.name
}

resource "azurerm_subnet_network_security_group_association" "k8s-sandbox" {
  subnet_id                 = azurerm_subnet.k8s-sandbox.id
  network_security_group_id = azurerm_network_security_group.k8s-sandbox.id
}
