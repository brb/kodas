List iptables triggered rules.

```
$ iptables-save -c > /tmp/before
# send some packets
$ iptables-save -c > /tmp/after
$ ./iptables-diff /tmp/before /tmp/after

+1 -A PREROUTING -m comment --comment "kubernetes service portals" -j KUBE-SERVICES
+9 -A OUTPUT -m comment --comment "kubernetes service portals" -j KUBE-SERVICES
<..>
```
