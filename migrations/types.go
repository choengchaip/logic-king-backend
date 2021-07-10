package main

import (
	"errors"
	"fmt"
	"log"
)

type IMigration interface {
	AddBatch(batch IMigrationBatch)
	Up() error
	Down() error
}

type Migration struct {
	batches []IMigrationBatch
}

type IMigrationBatch interface {
	Name() string
	Up() error
	Down() error
}

func NewMigration() IMigration {
	return &Migration{}
}

func (m *Migration) AddBatch(batch IMigrationBatch) {
	m.batches = append(m.batches, batch)
}

func (m *Migration) Up() error {
	if m.batches == nil {
		return errors.New("BatchIsEmpty")
	}

	for _, batch := range m.batches {
		log.Println(fmt.Sprintf("%s is migrating up", batch.Name()))
		err := batch.Up()
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Migration) Down() error {
	if m.batches == nil {
		return errors.New("BatchIsEmpty")
	}

	for _, batch := range m.batches {
		log.Println(fmt.Sprintf("%s is migrating down", batch.Name()))
		err := batch.Down()
		if err != nil {
			return err
		}
	}

	return nil
}
