# you can find openstack playbook here
# 

What I am going to do is initiating 3 instances on top of openstack and their associate playbooks are stored in this folder. 

You need to install openstacksdk 
   
    pip install ansible openstacksdk

Run the playbook :

    ansible-playbook openstack.yaml --ask-vault-pass

    For proof of concept vault password is: 6056071





