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

````bash
kubectl describe rs
kubectl get pod 
kubectl get pod rs-test-98vxm -o yaml
kubectl get rs rs-test -o yaml
````

````bash
kubectl delete -f replicaset/rs.yml
kubectl run --generator=run-pod/v1 podtest1 --image=nginx:alpine
kubectl run --generator=run-pod/v1 podtest2 --image=nginx:alpine
kubectl label pods podtest1 app=pod-label
kubectl label pods podtest2 app=pod-label
kubectl describe pod podtest1
kubectl get pods -l app=pod-label
kubectl apply -f replicaset/rs.yml
kubectl get pods
````

````bash
kubectl delete -f replicaset/rs.yml
kubectl apply -f deployments/dep.yml
kubectl get deployment
kubectl get deployment --show-labels
kubectl rollout status deployment deployment-test
kubectl get rs --show-labels
kubectl get pods --show-labels
kubectl get all
kubectl get rs deployment-test-b7c99d94b -o yaml
````


````bash
kubectl apply -f deployments/dep.yml
kubectl rollout status deployment deployment-test
kubectl describe deployment deployment-test
kubectl get pods
kubectl rollout history deployment deployment-test
kubectl apply -f deployments/dep.yml
kubectl get rs
````


````bash
kubectl delete -f deployments/dep.yml
kubectl apply -f deployments/dep.yml
kubectl get deploy
kubectl rollout history deployment deployment-test
kubectl apply -f deployments/dep.yml
kubectl get pods
kubectl rollout undo deployment deployment-test --to-revision=2
kubectl rollout status deployment deployment-test
````



