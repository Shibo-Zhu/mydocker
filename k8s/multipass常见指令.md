以下是 **Multipass** 的常见指令和用途：  

### **1. 基础管理命令**  
- **启动 Multipass 服务**  
  ```bash
  sudo snap start multipass
  ```  
- **停止 Multipass 服务**  
  ```bash
  sudo snap stop multipass
  ```  
- **查看 Multipass 服务状态**  
  ```bash
  sudo snap services multipass
  ```

### **2. 实例管理**  
- **启动一个新实例**  
  ```bash
  multipass launch --name my-instance
  ```  
- **指定 Ubuntu 版本启动实例**  
  ```bash
  multipass launch --name my-instance 20.04
  ```  
- **列出所有实例**  
  ```bash
  multipass list
  ```  
- **查看实例详细信息**  
  ```bash
  multipass info my-instance
  ```  
- **启动实例**  
  ```bash
  multipass start my-instance
  ```  
- **停止实例**  
  ```bash
  multipass stop my-instance
  ```  
- **删除实例**  
  ```bash
  multipass delete my-instance
  ```  
- **永久删除实例及数据**  
  ```bash
  multipass purge
  ```

### **3. 网络和 Shell 相关**  
- **进入实例 Shell**  
  ```bash
  multipass shell my-instance
  ```  
- **执行远程命令**  
  ```bash
  multipass exec my-instance -- ls -l
  ```  
- **查看实例 IP 地址**  
  ```bash
  multipass list
  ```  

### **4. 文件共享和挂载**  
- **挂载主机目录到实例**  
  ```bash
  multipass mount /host/path my-instance:/target/path
  ```  
- **卸载挂载目录**  
  ```bash
  multipass umount my-instance:/target/path
  ```  
- **将文件拷贝到实例中**  
  ```bash
  multipass transfer localfile.txt my-instance:/home/ubuntu/
  ```  
- **从实例中拷贝文件到本地**  
  ```bash
  multipass transfer my-instance:/home/ubuntu/file.txt ./
  ```

### **5. 镜像管理和快照**  
- **查看可用镜像**  
  ```bash
  multipass find
  ```  
- **创建快照**  
  ```bash
  multipass snapshot my-instance
  ```  
- **恢复快照**  
  ```bash
  multipass restore my-instance
  ```  

### **6. 资源管理**  
- **创建实例时分配资源**  
  ```bash
  multipass launch --name my-instance --cpus 2 --mem 4G --disk 20G
  ```  
- **动态调整资源**（需要重启实例）  
  ```bash
  multipass set local.driver=qemu
  ```  

### **7. 日志查看与调试**  
- **查看日志**  
  ```bash
  multipass get log
  ```  

这些指令涵盖了 Multipass 的主要功能。如果需要更详细的帮助，可以运行：  
```bash
multipass help
```


可以使用 `multipass transfer` 命令将宿主机上的 YAML 文件复制到 Multipass 虚拟机中。以下是详细步骤：  

### **1. 确保虚拟机正在运行**  
```bash
multipass list
```
检查虚拟机状态。如果虚拟机未运行，可以启动它：  
```bash
multipass start my-instance
```

### **2. 将 YAML 文件复制到虚拟机**  
假设宿主机上的 YAML 文件路径是 `/home/user/config.yaml`，要将其复制到虚拟机 `my-instance` 的 `/home/ubuntu/` 目录：  
```bash
multipass transfer /home/user/config.yaml my-instance:/home/ubuntu/
```

### **3. 验证文件已复制**  
登录到虚拟机：  
```bash
multipass shell my-instance
```
然后检查文件：  
```bash
ls -l /home/ubuntu/config.yaml
```

### **4. 从虚拟机复制文件到宿主机**  
如果需要从虚拟机将 YAML 文件复制回宿主机：  
```bash
multipass transfer my-instance:/home/ubuntu/config.yaml /home/user/
```

### **注意事项**  
1. 确保虚拟机名称和路径正确。  
2. 如果文件较大，可以考虑打包压缩再传输，提高效率。  
3. 如果出现权限问题，可以尝试将文件放到虚拟机的 `/tmp` 或 `/home/ubuntu` 等目录下，避免写入权限限制。  

以上步骤即可完成 YAML 文件的传输。