package main

import (
	"database/sql"
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"log"
)

var dsn = fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
	InstanceConnectionName,
	DatabaseName,
	DatabaseUser,
	Password)

// getAlertEventTime returns the alert with the given id.
func getAlertEventTime(id int) string {
	db, err := sql.Open("cloudsqlpostgres", dsn)
	if err != nil {
		log.Fatalf("failed to open DB: %v\n", err)
	}
	defer db.Close()
	row, err := db.Query("SELECT event_time FROM alert where id = $1", id)
	if err != nil {
		log.Fatalf("failed to query alert: %v\n", err)
	}
	var et string
	for row.Next() {
		if err = row.Scan(&et); err != nil {
			log.Fatalf("failed to recover event time from alert: %v\n", err)
		}
	}
	_ = row.Close()
	return et
}

// fetchImage returns the image related to the alert id.
func fetchImage(alertId int, fPath string) {
	db, err := sql.Open("cloudsqlpostgres", dsn)
	if err != nil {
		log.Fatalf("failed to open DB: %v\n", err)
	}
	defer db.Close()
	row, err := db.Query("SELECT data FROM image LEFT JOIN alert a on image.id = a.image_id WHERE a.id =  $1", alertId)
	if err != nil {
		log.Fatalf("failed to query image: %v\n", err)
	}
	var img []byte
	for row.Next() {
		if err = row.Scan(&img); err != nil {
			log.Fatalf("failed to recover image from alert: %v\n", err)
		}
	}
	_ = row.Close()
	bytesToJpeg(img, fPath)
}

// checkAlert marks the alert with `id` as sent in the DB.
func checkAlert(id int) {
	db, err := sql.Open("cloudsqlpostgres", dsn)
	if err != nil {
		log.Fatalf("failed to open DB: %v\n", err)
	}
	defer db.Close()
	_, err = db.Exec("UPDATE alert SET notification_sent = true WHERE id = $1", id)
	if err != nil {
		log.Fatalf("failed to update notification_set to true: %v\n", err)
	}
}
