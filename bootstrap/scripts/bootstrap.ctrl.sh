#!/bin/bash
{
echo 'running control plane bootstrap script...'
echo '**********************************'
apt-get update
apt-get upgrade -y

echo '**********************************'
echo '[✓] control plane bootstrapping  done...'
} >> /var/log/bootstrap.log 2>&1
