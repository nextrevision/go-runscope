package runscope

import (
	"errors"
	"fmt"
	"net/http"
)

type Schedule struct {
	ID            string `json:"id"`
	Note          string `json:"note"`
	Interval      string `json:"interval"`
	EnvironmentID string `json:"environment_id"`
}

func (client *Client) ListSchedules(bucketKey string, testID string) (*[]Schedule, *http.Response, error) {
	var schedules = []Schedule{}
	path := fmt.Sprintf("buckets/%s/tests/%s/schedules", bucketKey, testID)
	resp, err := client.Get(path, &schedules)
	return &schedules, resp, err
}

func (client *Client) GetSchedule(bucketKey string, testID string, scheduleID string) (*Schedule, *http.Response, error) {
	var schedule = Schedule{}
	path := fmt.Sprintf("buckets/%s/tests/%s/schedules/%s", bucketKey, testID, scheduleID)
	resp, err := client.Get(path, &schedule)
	return &schedule, resp, err
}

func (client *Client) NewSchedule(bucketKey string, testID string, schedule *Schedule) (*Schedule, *http.Response, error) {
	var newSchedule = Schedule{}
	if schedule.EnvironmentID == "" {
		err := errors.New("EnvironmentID must not be empty when creating new schedules")
		return &newSchedule, &http.Response{}, err
	}
	if schedule.Interval == "" {
		err := errors.New("Interval must not be empty when creating new schedules")
		return &newSchedule, &http.Response{}, err
	}
	path := fmt.Sprintf("buckets/%s/tests/%s/schedules", bucketKey, testID)
	resp, err := client.Post(path, &schedule, &newSchedule)
	return &newSchedule, resp, err
}

func (client *Client) UpdateSchedule(bucketKey string, testID string, scheduleID string, schedule *Schedule) (*Schedule, *http.Response, error) {
	var newSchedule = Schedule{}
	if schedule.EnvironmentID == "" {
		err := errors.New("EnvironmentID must not be empty when updating a schedule")
		return &newSchedule, &http.Response{}, err
	}
	if schedule.Interval == "" {
		err := errors.New("Interval must not be empty when updating schedule")
		return &newSchedule, &http.Response{}, err
	}
	path := fmt.Sprintf("buckets/%s/tests/%s/schedules/%s", bucketKey, testID, scheduleID)
	resp, err := client.Put(path, &schedule, &newSchedule)
	return &newSchedule, resp, err
}

func (client *Client) DeleteSchedule(bucketKey string, testID string, scheduleID string) (*http.Response, error) {
	path := fmt.Sprintf("buckets/%s/tests/%s/schedules/%s", bucketKey, testID, scheduleID)
	resp, err := client.Delete(path)
	return resp, err
}
