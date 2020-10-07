[![Go Report Card](https://goreportcard.com/badge/github.com/fpaupier/notifyMask)](https://goreportcard.com/report/github.com/fpaupier/notifyMask)

# Notification service

Subscribe to a Kafka topic of notifications to be sent (_alert for someone not wearing their mask_), fetch the alert from a Cloud SQL instance,
send an email to a system administrator and record that the alert has been sent in the database.

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

## Related Projects

This repository hosts the code for the last part of the project; notifying an admin that someone is not wearing their mask.
The previous projects are:

- [pi-mask-detection](https://github.com/fpaupier/pi-mask-detection) focuses on the detection of whether someone is wearing their mask or not, as seen per the Raspberry Pi.

- [alertDispatcher](https://github.com/fpaupier/alertDispatcher) is a Go module designed to run at the edge, especially a Raspberry Pi 4 B with 4Go of RAM.
The [alertDispatcher](https://github.com/fpaupier/alertDispatcher) polls the local SQLite event store and publishes them to a Kafka topic. 
 
- [alertIngress](https://github.com/fpaupier/alertIngress) is a Go module designed to run on a server, consuming from 
a Kafka topic where edge devices pushes their events. For each event consumed, the alert Ingress archives it in PostgresSQL and publishes a message
to be consumed by the notification service.
 