#!/bin/bash

#git checkout -b main
git add .
git commit -m "debug"
git push -u origin main
git push -u github main

#内测版本
#git checkout -b myAdmin_limit_v2.12.10
#git push -u origin myAdmin_limit_v2.12.10

#公测版本
#git checkout -b myAdmin_release_v2.3.7
#git push -u origin myAdmin_release_v2.3.7

#删除 本地 release 分支
#git branch -d release
#创建 本地 release 分支 并切换
#git checkout -b release
#从远程拉取 对应分支到本地 release分支
#git pull origin myAdmin_release_v2.3.7:release