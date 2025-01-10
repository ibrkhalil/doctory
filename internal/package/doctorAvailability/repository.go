package doctorAvailability

import (
	"errors"
)

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: make(map[string]DoctorAvailability),
	}
}

func (db *InMemoryDB) Create(availability DoctorAvailability) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.data[availability.ID]; exists {
		return errors.New("availability already exists")
	}

	db.data[availability.ID] = availability
	return nil
}

func (db *InMemoryDB) Read(id string) (DoctorAvailability, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	availability, exists := db.data[id]
	if !exists {
		return DoctorAvailability{}, errors.New("availability not found")
	}

	return availability, nil
}

func (db *InMemoryDB) Update(id string, availability DoctorAvailability) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.data[id]; !exists {
		return errors.New("availability not found")
	}

	db.data[id] = availability
	return nil
}

func (db *InMemoryDB) Delete(id string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.data[id]; !exists {
		return errors.New("availability not found")
	}

	delete(db.data, id)
	return nil
}
