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
	resp, _, err := client.Get(path, &schedules)
	if err != nil {
		println(err.Error())
	}
	return &schedules, resp, err
}

func (client *Client) GetSchedule(bucketKey string, testID string, scheduleID string) (*Schedule, *http.Response, error) {
	var schedule = Schedule{}
	path := fmt.Sprintf("buckets/%s/tests/%s/schedules/%s", bucketKey, testID, scheduleID)
	resp, _, err := client.Get(path, &schedule)
	if err != nil {
		println(err.Error())
	}
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
	resp, _, err := client.Post(path, &schedule, &newSchedule)
	if err != nil {
		println(err.Error())
	}
	return &newSchedule, resp, err
}

func (client *Client) UpdateSchedule(bucketKey string, testID string, schedule *Schedule) (*Schedule, *http.Response, error) {
	var newSchedule = Schedule{}
	if schedule.ID == "" {
		err := errors.New("ID must not be empty when updating a schedule")
		return &newSchedule, &http.Response{}, err
	}
	if schedule.EnvironmentID == "" {
		err := errors.New("EnvironmentID must not be empty when updating a schedule")
		return &newSchedule, &http.Response{}, err
	}
	if schedule.Interval == "" {
		err := errors.New("Interval must not be empty when updating schedule")
		return &newSchedule, &http.Response{}, err
	}
	path := fmt.Sprintf("buckets/%s/tests/%s/schedules/%s", bucketKey, testID, schedule.ID)
	resp, _, err := client.Put(path, &schedule, &newSchedule)
	if err != nil {
		println(err.Error())
	}
	return &newSchedule, resp, err
}

func (client *Client) DeleteSchedule(bucketKey string, testID string, scheduleID string) (*http.Response, error) {
	path := fmt.Sprintf("buckets/%s/tests/%s/schedules/%s", bucketKey, testID, scheduleID)
	resp, err := client.Delete(path)
	if err != nil {
		println(err.Error())
	}
	return resp, err
}
