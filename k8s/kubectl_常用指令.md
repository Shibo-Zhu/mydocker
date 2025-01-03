# 将yaml文件发布到集群中
```bash
kubectl apply -f config.yaml
```
# 查看deployment状态
```bash 
kubectl get deployments
```

# 查看pod状态
```bash
kubectl get pods
```
# 查看pod日志
```bash
kubectl logs pod-name
```
# 检查pod的详细信息
```bash
kubectl describe pod pod-name
```
# 检查pod的详细信息流式输出
```bash
kubectl describe pod pod-name
```

# 删除deployment
```bash
kubectl delete deployment deployment-name
```
# 删除pod
```bash
kubectl delete pod pod-name
```