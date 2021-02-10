# Here we can install docker via ansible

First run the adduser.yaml playbook to add ansible user to the servers

ansible-playbook -i openstac_inventory.py adduser.yaml


I am going to use dynamic inventory in ansible in conjunction with openstack 

ansible-playbook -i openstack_invenotry.py docker.yaml


For proof of concept vault password is : 6056071
