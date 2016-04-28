package runscope

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Schedule struct {
	ID            string `json:"id"`
	Note          string `json:"note"`
	Interval      string `json:"interval"`
	EnvironmentID string `json:"environment_id"`
}

func (client *Client) ListSchedules(bucketKey string, testID string) (*[]Schedule, error) {
	var schedules = []Schedule{}

	path := fmt.Sprintf("buckets/%s/tests/%s/schedules", bucketKey, testID)
	content, err := client.Get(path)
	if err != nil {
		return &schedules, err
	}

	err = unmarshal(content, &schedules)
	return &schedules, err
}

func (client *Client) GetSchedule(bucketKey string, testID string, scheduleID string) (*Schedule, error) {
	var schedule = Schedule{}

	path := fmt.Sprintf("buckets/%s/tests/%s/schedules/%s", bucketKey, testID, scheduleID)
	content, err := client.Get(path)
	if err != nil {
		return &schedule, err
	}

	err = unmarshal(content, &schedule)
	return &schedule, err
}

func (client *Client) NewSchedule(bucketKey string, testID string, schedule *Schedule) (*Schedule, error) {
	var newSchedule = Schedule{}

	if schedule.EnvironmentID == "" {
		err := errors.New("EnvironmentID must not be empty when creating new schedules")
		return &newSchedule, err
	}
	if schedule.Interval == "" {
		err := errors.New("Interval must not be empty when creating new schedules")
		return &newSchedule, err
	}

	path := fmt.Sprintf("buckets/%s/tests/%s/schedules", bucketKey, testID)
	data, err := json.Marshal(schedule)
	if err != nil {
		return &newSchedule, err
	}

	content, err := client.Post(path, data)
	if err != nil {
		return &newSchedule, err
	}

	err = unmarshal(content, &newSchedule)
	return &newSchedule, err
}

func (client *Client) UpdateSchedule(bucketKey string, testID string, scheduleID string, schedule *Schedule) (*Schedule, error) {
	var newSchedule = Schedule{}

	if schedule.EnvironmentID == "" {
		err := errors.New("EnvironmentID must not be empty when updating a schedule")
		return &newSchedule, err
	}
	if schedule.Interval == "" {
		err := errors.New("Interval must not be empty when updating schedule")
		return &newSchedule, err
	}

	path := fmt.Sprintf("buckets/%s/tests/%s/schedules/%s", bucketKey, testID, scheduleID)
	data, err := json.Marshal(schedule)
	if err != nil {
		return &newSchedule, err
	}

	content, err := client.Put(path, data)
	if err != nil {
		return &newSchedule, err
	}

	err = unmarshal(content, &newSchedule)
	return &newSchedule, err
}

func (client *Client) DeleteSchedule(bucketKey string, testID string, scheduleID string) error {
	path := fmt.Sprintf("buckets/%s/tests/%s/schedules/%s", bucketKey, testID, scheduleID)
	return client.Delete(path)
}
