output "instance_public_ip" {
    value = azurerm_public_ip.k8s-sandbox.ip_address
}

output "instance_private_ip" {
    value = azurerm_network_interface.k8s-sandbox.ip_configuration[0].private_ip_address
}

