#!/bin/bash
set -e

echo "[sentinel] configuring iptables"

iptables -C INPUT -j NFQUEUE --queue-num 0 2>/dev/null || \
iptables -I INPUT -j NFQUEUE --queue-num 0

iptables -C OUTPUT -j NFQUEUE --queue-num 0 2>/dev/null || \
iptables -I OUTPUT -j NFQUEUE --queue-num 0

echo "[sentinel] done"
