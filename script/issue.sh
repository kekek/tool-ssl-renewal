#!/usr/bin/env bash

## 前提 已安装acme.sh

## 颁发证书命令

echo "========================================== let's encrypt 申请和安装证书 =========================================="

echo ">> acme.sh 颁发ecc通配符证书"
echo "acme.sh --issue  -d ilaiyang.com.cn  -d '*.ilaiyang.com.cn'  --dns dns_dp  --keylength ec-256"
echo

echo ">> 安装ecc证书"
echo "acme.sh --installcert --ecc -d ilaiyang.com.cn \
--key-file /etc/nginx/crt/letsencrypt/ecc/privkey.pem  \
--fullchain-file /etc/nginx/crt/letsencrypt/ecc/fullchain.pem \
--reloadcmd '/usr/sbin/nginx -s reload'"
echo

echo ">> acme.sh 颁发rsa通配符证书"
echo "acme.sh --issue -d ilaiyang.com.cn -d '*.ilaiyang.com.cn' --dns dns_dp"
echo

echo ">> acme.sh 颁发rsa通配符证书 with 更新uloud hook"
echo "acme.sh --issue -d ilaiyang.com.cn -d '*.ilaiyang.com.cn' --dns dns_dp --renew-hook '/root/.acme.sh/self-hook/renew-hook.sh &>> /root/.acme.sh/self-hook/ilaiyang.com.cn.log'"
echo

echo ">> 安装rsa 证书证书"
echo "acme.sh --installcert -d ilaiyang.com.cn \
--key-file /etc/nginx/crt/letsencrypt/rsa/privkey.pem  \
--fullchain-file /etc/nginx/crt/letsencrypt/rsa/fullchain.pem \
--reloadcmd '/usr/sbin/nginx -s reload'"
echo


echo "=============== 移除不需要的证书 ==============="
echo "acme.sh --remove -d '*.ilaiyang.com.cn'"
echo "acme.sh --remove -d '*.ilaiyang.com.cn' --ecc"
echo