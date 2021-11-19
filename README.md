## KUBERNETES

```bash
kubectl run --generator=run-pod/v1 podtest2 --image=nginx:alpine
kubectl get pods
kubectl describe pod podtest
kubectl delete pod podtest
kubectl api-resources
kubectl get pod podtest2 -o yaml
docker ps -a
kubect exec -ti podtest2
curl 172.17.0.3
kubectl exec -ti podtest2 -- sh
pdw
whoami
cd /usr/share/nginx/html
cat cd /usr/share/nginx/html/index.html
kubectl logs podtest2 -f
````


````bash
kubectl apply -f pods/pod.yml
````