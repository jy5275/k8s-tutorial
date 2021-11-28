# Deploy my own application from source code to monitored deployment
Official document is here https://kubernetes.io/zh/docs/tasks/configure-pod-container/pull-image-private-registry/

Steps:

1. Develop source code locally

2. Build image: `docker build -f Dockerfile -t frozenlight/go-web .`

3. Login dockerhub: `docker login`

4. Push image to dockerhub `docker push frozenlight/go-web`. Might take several minutes. Check at https://hub.docker.com

5. Initiate minikube cluster `minikube start --driver=docker`

6. Install helm charts
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add stable https://kubernetes-charts.storage.googleapis.com/
helm repo update
helm install prometheus prometheus-community/kube-prometheus-stack
```

7. Expose ports
```
kubectl port-forward deployment.apps/prometheus-grafana 3000
kubectl port-forward statefulset.apps/prometheus-prometheus-kube-prometheus-prometheus 9090
```

8. Create credential for minikube to pull image from dockerhub
```
kubectl create secret docker-registry regcred \
  --docker-server=https://index.docker.io/v2/ \
  --docker-username=frozenlight \
  --docker-password=Jyeecs_2017 \
  --docker-email=527507046@qq.com
```

9. Create k8s deployment and service `kubectl apply -f go-web.yaml`

10. Create k8s serviceMonitor `kubectl apply -f svcm.yaml`
Meaning of each entry in serviceMonitor config file is here https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/user-guides/getting-started.md

11. Add panel, promQL=`covid_num{job="go-web-svc"}`