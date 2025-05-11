#!/bin/bash
{
echo 'running worker-1 node bootstrap script...'
echo '**********************************'
apt-get update
apt-get upgrade -y

echo '**********************************'
echo '[✓] worker bootstraping done...'
} >> /var/log/bootstrap.log 2>&1
