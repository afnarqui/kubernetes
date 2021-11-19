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
kubectl get pods
docker run --net host -ti python:3.6-alpine sh
# python -m http.server 8081
kubectl apply -f pods/pod.yml
kubectl get pods
kubectl describe pod4
kubectl logs pod4 -c cont4
kubectl logs pod4 -c cont5
kubectl exec -ti pod4 -- sh
apk add -U curl
python -m http.server 8082
kubectl delete -f pods/pod.yml
kubectl apply -f pods/pod.yml
kubectl exec -ti pod -c cont4 -- sh
apk add -U curl
curl localhost:8082
curl localhost:8083
````

````bash
kubectl delete -f pods/pod.yml
kubectl apply -f pods/labels.yml
kubectl get pods -l app=backend
kubectl get pods -l app=front
kubectl get pods -l env=dev
````

````bash
kubectl delete -f pods/labels.yml
kubectl api-resources
kubectl apply -f replicaset/rs.yml
kubectl get pods -l app=pod-label
kubectl get replicaset
kubectl api-resources
kubectl get rs
````