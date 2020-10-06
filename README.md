[![Go Report Card](https://goreportcard.com/badge/github.com/fpaupier/notifyMask)](https://goreportcard.com/report/github.com/fpaupier/notifyMask)

# Notification service

Subscribe to a Kafka topic of notifications to be sent (_alert for someone not wearing a mask_), fetch the alert from a Cloud SQL instance,
send an email to a system administrator and check that the alert has been sent in the database.

# Running on local 

1. Install locally
````shell script
go install .
````

2. Run the service:
```shell script
$GOPATH/bin/notifyMask
```

# Deploying to Google Cloud Platform (_GCP_)

I assume you already have an existing GCP project.

1. First build and publish the image to Google cloud (make sure the storage bucket write access on your project) 
````shell script
gcloud builds submit --tag gcr.io/YOUR_PROJECT_NAME/notify-mask
````

2. Then, create a Cloud Compute Engine instance using the image you just published. Create a permanent IP address instead of using an ephemeral one.

3. Copy the instance external IP address and go to your Cloud SQL instance dashboard. 

4. Go to the `connection` tab, click `+ Add network` and paste your instance IP address. 

