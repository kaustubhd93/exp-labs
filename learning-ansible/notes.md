# Key features

- agentless
- requires only python to to be installed on target machine
- works on ssh so updates to multiple vms can be slow. 

## Configuration precedence

- environment variables, for eg: `$ANSIBLE_CONFIG=/opt/ansible-web.cfg`
- ansible.cfg in the current directory where ansible is being run from.
- .ansible.cfg in home directory
- default ansible config file /etc/ansible/ansible.cfg
> Note: no need to have all values in config, only once which you want to override. 

