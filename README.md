## OADP Tests in Ginkgo

### Getting Started ###

This is a preview only repo for easing the process to get started with Go lang and writing OADP test cases using Go and Ginkgo

#### Pre-requisites

1. OCP cluster should be up and running

2. kubeconfig file generated for your OCP cluster

3. If you don't have the kubeconfig file already, follow below steps -
   
   - Open terminal and run `export KUBECONFIG=kubeconfig`
   
   - Login to your OCP using `oc login` command
   
   - In the current directory, you should be able to see the kubeconfig file generated for your OCP cluster

4. OADP operator must be installed

#### Instructions

1. Clone the repository

    `git clone https://github.com/nitishSr/oadp-go-tests.git`

2. Go to data folder

    `cd ~/oadp/data/`
    
3. Write your aws credentials value against the key in file aws_credentials

4. Go to tests/velero folder

   `cd ~/oadp/tests/velero/`
   
5. In the terminal set the KUBECONFIG env variable to point to kubeconfig file
   
   `export KUBECONFIG=/path/to/kubeconfig`
   
6. Run velero CR test which creates a Velero CRD in OADP operator's namespace

   `go test`
