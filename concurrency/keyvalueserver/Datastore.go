package main

import "errors"

type DataStore struct {
	storage map[string]int64
}

func (d *DataStore) New() *DataStore {
	return &DataStore{}
}

func (d *DataStore) AddItem(key string, value int64) {
	d.storage[key] = value
}

func (d *DataStore) IncItem(key string) (bool, error) {
	value, ok := d.storage[key]
	if !ok {
		return false, errors.New("Cannot increment. THE KEY DOES NOT EXIST!!!!")
	}
	d.storage[key] = value + 1
	return true, nil
}

func (d *DataStore) MultiplyItem(key string, multiplier int64) (bool, error) {
	value, ok := d.storage[key]
	if !ok {
		return false, errors.New("Cannot mulitply. THE KEY DOES NOT EXIST!!!!")
	}
	d.storage[key] = value * multiplier
	return true, nil
}

func (d *DataStore) GetItem(key string) (bool, int64) {
	value, ok := d.storage[key]
	return ok, value
}
