# 树莓派的k8s配置
## **一、树莓派基本配置**
### 1.更新树莓派源
```
cat << EOF > /etc/apt/sources.list
deb http://mirrors.tuna.tsinghua.edu.cn/raspbian/raspbian/ buster main non-free contrib rpi
deb-src http://mirrors.tuna.tsinghua.edu.cn/raspbian/raspbian/ buster main non-free contrib rpi
EOF
cat << EOF > /etc/apt/sources.list.d/raspi.list
deb http://mirrors.tuna.tsinghua.edu.cn/raspberrypi/ buster main ui
EOF
```
### 2.禁用swap（此处永久禁用swap）
```
dphys-swapfile swapoff

dphys-swapfile uninstall

update-rc.d dphys-swapfile remove

rm -f /etc/init.d/dphys-swapfile

service dphys-swapfile stop

systemctl disable dphys-swapfile.service
```
### 3.启用桥接
```
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF

sudo sysctl --system
```
修改 /etc/systctl.conf
```
net.ipv4.ip_forward = 1
```

## **二、安装**
### 1.安装docker
```
curl -fsSL https://get.docker.com -o get-docker.sh

sudo sh get-docker.sh
```
### 2.修改docker配置，替换镜像源
```
cat > /etc/docker/daemon.json <<EOF
{
  "registry-mirrors": ["https://kyzgcmt0.mirror.aliyuncs.com"],
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2",
}
EOF

sed -i '$ s/$/ cgroup_enable=cpuset cgroup_enable=memory cgroup_memory=1 swapaccount=1/' /boot/cmdline.txt
```
### 3.安装k8s
```
curl -s https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add -

echo "deb https://mirrors.aliyun.com/kubernetes/apt kubernetes-xenial main" | tee -a /etc/apt/sources.list.d/kubernetes.list

apt-get update

apt-get install -y kubelet kubeadm kubectl

apt-mark hold kubelet kubeadm kubectl
```
### 4.下载所需镜像
查询所需镜像
```
kubeadm config images list
```
根据查询结果修改以下脚本内容（镜像的版本和coredns的源可能会与脚本中不同）
```
# pull kubernetes images from hub.docker.com
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy:$KUBE_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager:$KUBE_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver:$KUBE_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler:$KUBE_VERSION
# pull aliyuncs mirror docker images
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/pause:$PAUSE_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:$CORE_DNS_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/etcd:$ETCD_VERSION

# retag to k8s.gcr.io prefix
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy:$KUBE_VERSION  k8s.gcr.io/kube-proxy:$KUBE_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager:$KUBE_VERSION k8s.gcr.io/kube-controller-manager:$KUBE_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver:$KUBE_VERSION k8s.gcr.io/kube-apiserver:$KUBE_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler:$KUBE_VERSION k8s.gcr.io/kube-scheduler:$KUBE_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/pause:$PAUSE_VERSION k8s.gcr.io/pause:$PAUSE_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/coredns/coredns:$CORE_DNS_VERSION k8s.gcr.io/coredns:$CORE_DNS_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/etcd:$ETCD_VERSION k8s.gcr.io/etcd:$ETCD_VERSION

# untag origin tag, the images won't be delete.
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy:$KUBE_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager:$KUBE_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver:$KUBE_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler:$KUBE_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/pause:$PAUSE_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/coredns/coredns:$CORE_DNS_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/etcd:$ETCD_VERSION
```
## **三、初始化**
### 1.生成令牌
```
TOKEN=$(sudo kubeadm token generate)
```
### 2.初始化控制平面
kubernetes-version可以根据需要修改
```
kubeadm init --token=${TOKEN} --kubernetes-version=v1.22.1 --pod-network-cidr=10.244.0.0/16
```
运行完成后执行
```
mkdir -p $HOME/.kube

sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config

sudo chown $(id -u):$(id -g) $HOME/.kube/config
```
### 3.配置flannel
如果flannel镜像拉取失败可以根据yml文件手动拉取
```
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
```
## **四、安装dashboard**
对应的dashboard版本可以通过 https://github.com/kubernetes/dashboard/releases 查看
### 1.下载并配置dashboard（在master节点执行）
```
wget https://raw.githubusercontent.com/kubernetes/dashboard/v2.3.1/aio/deploy/recommended.yaml

kubectl apply -f recommended.yaml
```
### 2.查看当前服务
```
kubectl get pods --all-namespaces

kubectl get svc --all-namespaces
```
### 3.删除服务并重新配置
```
kubectl delete service kubernetes-dashboard --namespace=kubernetes-dashboard
```
新建配置文件
```
vim dashboard-svc.yaml

#内容
kind: Service
apiVersion: v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
spec:
  type: NodePort
  ports:
    - port: 443
      targetPort: 8443
  selector:
    k8s-app: kubernetes-dashboard
```
更新配置文件
```
kubectl apply -f dashboard-svc.yaml
```
再次查看服务，获得运行端口号
```
kubectl get svc --all-namespaces
```
### 4.创建管理员角色
```
vim dashboard-svc-account.yaml
 
# 内容
apiVersion: v1
kind: ServiceAccount
metadata:
  name: dashboard-admin
  namespace: kube-system
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: dashboard-admin
subjects:
  - kind: ServiceAccount
    name: dashboard-admin
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io
```
更新
```
kubectl apply -f dashboard-svc-account.yaml
```
获得token
```
kubectl get secret -n kube-system |grep admin|awk '{print $1}'

kubectl describe secret dashboard-admin-token-bwgjv -n kube-system|grep '^token'|awk '{print $2}'
```
任意nodeIP:port都可以访问
## **五、安装metrics**
### 1.下载yaml配置文件
```
wget https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```
### 2.修改配置文件
在deployment.spec.template.spec.containers.args下添加
```
- --kubelet-insecure-tls
```
修改deployment.spec.template.spec.containers.image为
```
registry.cn-hangzhou.aliyuncs.com/google_containers/metrics-server:[version]
```
### 3.启动metrics-server
```
kubectl apply -f components.yaml
```