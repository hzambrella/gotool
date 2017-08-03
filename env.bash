#/!bin/bash

# 设定git库地址转换, 以便解决部分包的库被墙的问题
git config --global url."git@git.ot24.net:".insteadOf "https://git.ot24.net"
git config --global url."https://github.com/golang/".insteadOf "https://go.googlesource.com/"

PWDDIR=`pwd`
#export PJPATH=$PWDDIR:$PWDDIR/tool:$PWDDIR/patu:$PWDDIR/tableWriter:$PWDDIR/ini
export PJPATH=$PWDDIR
export LIBS=$PWDDIR/libs
export GOLIBS=$PWDDIR/golibs
export GOPATH=$GOLIBS:$LIBS:$PJPATH
export LOGCFG=$PWDDIR/src/silog
echo "gopath:"$GOPATH
