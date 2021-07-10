package main

import (
	"errors"
	"fmt"
	"log"
)

type ISeed interface {
	AddBatch(batch ISeedBatch)
	Up() error
	Down() error
}

type Seed struct {
	batches []ISeedBatch
}

type ISeedBatch interface {
	Name() string
	Up() error
	Down() error
}

func NewSeed() ISeed {
	return &Seed{}
}

func (s *Seed) AddBatch(batch ISeedBatch) {
	s.batches = append(s.batches, batch)
}

func (s *Seed) Up() error {
	if s.batches == nil {
		return errors.New("BatchIsEmpty")
	}

	for _, batch := range s.batches {
		log.Println(fmt.Sprintf("%s is migrating up", batch.Name()))
		err := batch.Up()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Seed) Down() error {
	if s.batches == nil {
		return errors.New("BatchIsEmpty")
	}

	for _, batch := range s.batches {
		log.Println(fmt.Sprintf("%s is migrating down", batch.Name()))
		err := batch.Down()
		if err != nil {
			return err
		}
	}

	return nil
}
