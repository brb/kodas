#!/bin/sh

# Tears down k8s cluster provisioned with kubeadm.
# Usage: sh ./teardown.sh [ssh-user-host]

set -eu

teardown() {
    ssh $1 'sudo -s -- <<\EOF
systemctl stop kubelet;
docker rm -f -v $(docker ps -q);
find /var/lib/kubelet | xargs -n 1 findmnt -n -t tmpfs -o TARGET -T | uniq | xargs -r umount -v;
rm -r -f /etc/kubernetes /var/lib/kubelet /var/lib/etcd;
systemctl reboot
EOF'
}

for i in "$@"; do
    echo "Tearing down $i..."
    teardown $i
done
