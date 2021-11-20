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

````bash
kubectl api-resources
kubectl delete -f deployments/dep.yml
kubectl get pods
kubectl apply -f service/svc.yml
kubectl get svc
kubectl get svc -l app=front
kubectl describe svc my-service
kubectl get endpoints
kubectl get pods -l app=front -o wide
kubectl run --generator=run-pod/v1 podtest2 --image=nginx:alpine
kubectl label pods podtest2 app=pod-label
kubectl get svc -l app=front
kubectl run --rm -ti --generator=run-pod/v1 podtest4 --image=nginx:alpine -- sh
apk add -U curl
curl 10.108.25.152:8080
curl my-service:8080
ip a
````

````bash
kubectl apply -f service/nodeport.yml
kubectl get pods -l app=backend
kubectl get svc -l app=backend

````


````bash
kubectl delete -f deployments/dep.yml

````

### aws eks

````bash
docker run --rm -ti ubuntu
apt-get update
apt-get install python3-pip
pip3
pip3 install -U awscli
aws configure 
````

````bash
aws eks help
https://docs.aws.amazon.com/eks/latest/userguide/getting-started-eksctl.html
curl --silent --location "https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp
sudo mv /tmp/eksctl /usr/local/bin
sudo chmod +x /usr/local/bin/eksctl
eksctl version

eksctl create cluster --name ikatech-cluster --without-nodegroup --region us-east-1 --zones us-east-1b,us-east-1a
cat .aws/credentials
````

````bash
CloudFormation
eks > clusters
/home/afnarqui/.kube/config
rm -fr /home/afnarqui/.kube/config > no eliminar si quieren conectarse con eksctl
aws eks --region us-east-1 update-kubeconfig --name ikatech-cluster
kubectl get svc
kubectl cluster-info
kubectl run --generator=run-pod/v1 pod1 --image=nginx:alpine
kubectl get pods
kubectl describe pod id
https://docs.aws.amazon.com/eks/latest/userguide/what-is-eks.html
https://aws.amazon.com/es/ec2/instance-types/t3/
eksctl create nodegroup -h
eksctl create nodegroup \
  --cluster ikatech-cluster \
  --region us-east-1  \
  --name ikatech-workers \
  --node-type t2.medium \
  --nodes 1 \
  --nodes-min 1 \
  --nodes-max 1 \
  --asg-access

  kubectl get pods

````

## Ingress controller

````bash
https://docs.aws.amazon.com/eks/latest/userguide/alb-ingress.html
eksctl utils associate-iam-oidc-provider \
    --region us-east-1 \
    --cluster ikatech-cluster \
    --approve

curl -o iam-policy.json https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.4/docs/examples/iam-policy.json

aws iam create-policy \
--policy-name ALBIngressControllerIAMPolicy \
--policy-document "https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.4/docs/examples/iam-policy.json"


kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.4/docs/examples/rbac-role.yaml

eksctl create iamserviceaccount \
  --region us-east-1 \
  --name alb-ingress-controller \
  --namespace=kube-system \
  --cluster=ikatech-cluster \
  --attach-policy-arn=arn:aws:iam::975812830743:policy/ALBIngressControllerIAMPolicy \
  --override-existing-serviceaccounts \
  --approve

  kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.4/docs/examples/alb-ingress-controller.yaml

  kubectl edit deployment.apps/alb-ingress-controller -n kube-system
  spec:
    containers:
    - args:
     - --cluster-name=ikatech-cluster

    kubectl get pods -n kube-system     
````

````bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/v2.3.0/docs/examples/2048/2048_full.yaml
kubectl apply -f ingress/ing.yml
kubectl get all -n game-2048
kubectl port-forward pod/deployment-2048-79785cfdff-746bw 7000:80 -n game-2048
kubectl apply -f ingress/ingress.yml
kubectl get ingress -n game-2048
kubectl get -n kube-system pods
kubectl logs -f -n kube-system alb-ingress-controller-5d4f7dc9dd-lbx4b
kubectl get svc --all-namespaces
kubectl delete svc service-2048 -n game-2048
eksctl delete cluster ikatech-cluster
````

## cluster eks

````bash
eksctl create cluster --name afn-cluster

eksctl utils associate-iam-oidc-provider --cluster=afn-cluster --approve

kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-alb-ingress-controller/v1.1.4/docs/examples/rbac-role.yaml

curl -o iam_policy.json https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/v2.2.0/docs/install/iam_policy.json
	aws iam create-policy \
		--policy-name AWSLoadBalancerControllerIAMPolicy \
		--policy-document file://iam_policy.json

	eksctl create iamserviceaccount \
      --cluster=afn-cluster \
      --namespace=kube-system \
      --name=aws-load-balancer-controller \
      --attach-policy-arn=arn:aws:iam::975812830743:policy/AWSLoadBalancerControllerIAMPolicy \
      --override-existing-serviceaccounts \
      --approve


kubectl apply \
		--validate=false \
		-f https://github.com/jetstack/cert-manager/releases/download/v1.1.1/cert-manager.yaml



https://docs.aws.amazon.com/eks/latest/userguide/aws-load-balancer-controller.html


curl -Lo v2_3_0_full.yaml https://github.com/kubernetes-sigs/aws-load-balancer-controller/releases/download/v2.3.0/v2_3_0_full.yaml
````

````bash
kubectl get -n kube-system pods
kubectl logs -f -n kube-system aws-load-balancer-controller-777c5bc6b5-fgdbc
nslookup k8s-default-simplein-c51b4c8137-176125708.us-east-1.elb.amazonaws.com
sudo vim etc/host
18.211.1.204 react.afnarqui.com
kubectl delete ingress.yml
kubectl delete service.yml
kubectl delete deployment.yml
kubectl delete namespace.yml
````

````bash
https://docs.aws.amazon.com/eks/latest/userguide/horizontal-pod-autoscaler.html
https://docs.aws.amazon.com/eks/latest/userguide/metrics-server.html
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
kubectl get deployment metrics-server -n kube-system
kubectl get deployment
kubectl autoscale deployment node-back-deployment --cpu-percent=50 --min=2 --max=10
kubectl autoscale deployment react-front-deployment --cpu-percent=50 --min=2 --max=10
kubectl get hpa
kubectl get -o yaml hpa node-back-deployment
kubectl describe hpa node-back-deployment
kubectl edit hpa node-back-deployment
kubectl get deploy -o yaml node-deployment
https://aws.amazon.com/es/ec2/instance-types/t3/
kubectl get pods
watch kubectl get pods
watch kubectl get hpa
kubectl run --generator=run-pod/v1 apache-bench -i --tty --rm --image=httpd -- bash
ab -n 500000 -c 1000 http://back.afnarqui.com/
````



