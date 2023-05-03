#!/usr/bin/bash


go build .

#In order to run as root the program file must belong to root user
sudo chown root:root $1

# To be able to make setuid syscall the program file must s in the permission byte
sudo chmod u+s $1

# run the cat /etc/sudoers as previleged user
$1 --command="cat /etc/sudoers" --user="root"