#!/bin/sh
read -p "input commit measure:"  val
echo $val
git status
git add .
git status
git commit -m"$val"
git status
git pull
#read -p "are you sure to commit ?[Y/N]"val2
set timeout 30
expect <<!
spawn git push origin master
expect "Username"
send "your name\r"
expect "Password"
send "your password\r"
expect eof
!




