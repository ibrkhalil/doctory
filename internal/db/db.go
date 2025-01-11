package db

import (
	"sync"

	"github.com/ibrkhalil/doctory/internal/schema"
)

type SingletonDB struct {
	doctorAvailabilitySlot map[string]schema.DoctorAvailabilitySlot
	appointmentSlots       map[string]schema.AppointmentSlot
	mutex                  sync.RWMutex
}

var singletonInstance *SingletonDB
var once sync.Once

func GetInstance() *SingletonDB {
	once.Do(func() {
		singletonInstance = &SingletonDB{
			doctorAvailabilitySlot: make(map[string]schema.DoctorAvailabilitySlot),
			appointmentSlots:       make(map[string]schema.AppointmentSlot),
		}
	})
	return singletonInstance
}

func (db *SingletonDB) SetDoctorAvailabilitySlot(key string, value schema.DoctorAvailabilitySlot) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.doctorAvailabilitySlot[key] = value
}

func (db *SingletonDB) SetAppointmentSlots(key string, value schema.AppointmentSlot) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.appointmentSlots[key] = value
}

func (db *SingletonDB) GetDoctorAvailabilitySlotByKey(key string) (schema.DoctorAvailabilitySlot, bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	value, alreadyExists := db.doctorAvailabilitySlot[key]
	return value, alreadyExists
}

func (db *SingletonDB) GetAppointmentSlotByKey(key string) (schema.AppointmentSlot, bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	value, alreadyExists := db.appointmentSlots[key]
	return value, alreadyExists
}

func (db *SingletonDB) GetAllDoctorAvailabilitySlots() []schema.DoctorAvailabilitySlot {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	var response []schema.DoctorAvailabilitySlot
	for _, v := range db.doctorAvailabilitySlot {
		response = append(response, v)
	}
	return response
}

func (db *SingletonDB) GetAllAppointmentSlots() []schema.AppointmentSlot {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	var response []schema.AppointmentSlot
	for _, v := range db.appointmentSlots {
		response = append(response, v)
	}
	return response
}
