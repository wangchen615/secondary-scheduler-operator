Simple Steps to Start
============
1. Log in to your OpenShift cluster and clone the operator's code
```
~ cd ~/code/Go/src/github.com/openshift/
~ git clone https://github.com/openshift/secondary-scheduler-operator.git
~ cd secondary-scheduler-operator
```

2. Build your own secondary-scheduler-operator image.
```
~ export QUAY_REGISTRY=${your_quay_user_id}
~ docker login quay.io
~ docker build -t quay.io/${QUAY_REGISTRY}/secondary-scheduler-operator:4.9 .
~ docker push quay.io/${QUAY_REGISTRY}/secondary-scheduler-operator:4.9
```

3. Copy the `deploy` folder from the secondary-scheduler-operator repo to a temparory  folder.
```
~ mkdir _tmp
~ cp -r ./deploy/*.yaml ./_tmp/
```


4. Replace the operator image and prometheus info in yamls.
```
~ ./hack/trimaran-example.sh
```

5. Create all resources needed for the secondary-scheduler-operator
```
~ oc create -f _tmp
```

