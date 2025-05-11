#!/bin/bash
{
echo 'running devbox bootstrap script...'
echo '**********************************'
apt-get update
apt-get upgrade -y

echo '**********************************'
echo '[✓] devbox bootstrap done...'
} >> /var/log/bootstrap.log 2>&1
