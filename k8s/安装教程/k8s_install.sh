#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status.

# Step 1: Disable swap and load kernel modules
sudo swapoff -a
sudo sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab

# Load necessary kernel modules
sudo modprobe overlay
sudo modprobe br_netfilter

# Ensure the modules are loaded on boot
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
overlay
br_netfilter
EOF

# Load kernel parameters
sudo sysctl --system

echo "swap disabled and kernel modules loaded."
echo "installing containerd..."

# Step 2: Install Containerd
sudo apt update && sudo apt install -y curl gnupg2 software-properties-common apt-transport-https ca-certificates

# Add containerd repository based on architecture
ARCH=$(dpkg --print-architecture)
if [ "$ARCH" == "amd64" ]; then
  sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmour -o /etc/apt/trusted.gpg.d/containerd.gpg
  sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
elif [ "$ARCH" == "arm64" ]; then
  sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmour -o /etc/apt/trusted.gpg.d/containerd.gpg
  sudo add-apt-repository "deb [arch=arm64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
else
  echo "Unsupported architecture: $ARCH"
  exit 1
fi

# Install containerd
sudo apt install -y containerd.io

# Configure containerd to use systemd as cgroup driver
sudo containerd config default | sudo tee /etc/containerd/config.toml >/dev/null 2>&1
sudo sed -i 's/SystemdCgroup = false/SystemdCgroup = true/g' /etc/containerd/config.toml

# Update pause image for compatibility in China
sudo sed -i 's|sandbox_image = .*|sandbox_image = "registry.aliyuncs.com/google_containers/pause:3.10"|' /etc/containerd/config.toml

# Restart containerd to apply changes
sudo systemctl restart containerd

# Add non-root user to containerd group
sudo groupadd -f containerd
sudo usermod -aG containerd $(whoami)
newgrp containerd

echo "Containerd installed successfully."
echo "installing Kubernetes components..."

# Step 3: Install Kubernetes components
# Use Alibaba Cloud repository for better access in China
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.31/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg

echo "deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.31/deb/ /" | sudo tee /etc/apt/sources.list.d/kubernetes.list

# Update package list and install Kubernetes components
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl

# Hold versions to prevent unintended upgrades
sudo apt-mark hold kubelet kubeadm kubectl

echo "Node setup completed successfully."