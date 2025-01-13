package db

import (
	"sync"

	"github.com/ibrkhalil/doctory/internal/schema"
)

type SingletonDB struct {
	doctorAvailabilitySlot map[string]schema.DoctorAvailabilitySlot
	appointmentSlots       map[string]schema.AppointmentSlot
	mutex                  sync.Mutex
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
	db.mutex.Lock()
	defer db.mutex.Unlock()
	value, alreadyExists := db.doctorAvailabilitySlot[key]
	return value, alreadyExists
}

func (db *SingletonDB) GetAppointmentSlotByKey(key string) (schema.AppointmentSlot, bool) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	value, alreadyExists := db.appointmentSlots[key]
	return value, alreadyExists
}

func (db *SingletonDB) GetAllDoctorAvailabilitySlots() []schema.DoctorAvailabilitySlot {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	var response []schema.DoctorAvailabilitySlot
	for _, availabilitySlot := range db.doctorAvailabilitySlot {
		response = append(response, availabilitySlot)
	}
	return response
}

func (db *SingletonDB) GetAllAppointmentSlots() []schema.AppointmentSlot {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	var response []schema.AppointmentSlot
	for _, appointmentSlot := range db.appointmentSlots {
		response = append(response, appointmentSlot)
	}
	return response
}

func (db *SingletonDB) CancelAppointmentById(key string) bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	alreadyExistingAppointment := schema.AppointmentSlot{
		ID:           db.appointmentSlots[key].ID,
		SlotId:       db.appointmentSlots[key].SlotId,
		PatientID:    db.appointmentSlots[key].PatientID,
		PatientName:  db.appointmentSlots[key].PatientName,
		ReservedAt:   db.appointmentSlots[key].ReservedAt,
		StartingTime: db.appointmentSlots[key].StartingTime,
		State:        schema.CANCELLED_APPOINTMENT_STATE,
	}
	db.appointmentSlots[key] = alreadyExistingAppointment
	return true
}

func (db *SingletonDB) ConfirmAppointmentById(key string) bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	alreadyExistingAppointment := schema.AppointmentSlot{
		ID:           db.appointmentSlots[key].ID,
		SlotId:       db.appointmentSlots[key].SlotId,
		PatientID:    db.appointmentSlots[key].PatientID,
		PatientName:  db.appointmentSlots[key].PatientName,
		ReservedAt:   db.appointmentSlots[key].ReservedAt,
		StartingTime: db.appointmentSlots[key].StartingTime,
		State:        schema.CONFIRMED_APPOINTMENT_STATE,
	}
	db.appointmentSlots[key] = alreadyExistingAppointment
	return true
}

type slotId struct {
	sync.Mutex
	slotId int
}

func (slotId *slotId) SlotID() (id int) {
	slotId.Lock()
	defer slotId.Unlock()

	id = slotId.slotId
	slotId.slotId++
	return
}

var autoIncrementedId slotId

func NewAppointmentWithAutoIncrementedSlotID() *schema.AppointmentSlot {
	return &schema.AppointmentSlot{
		SlotId: autoIncrementedId.SlotID(),
	}
}

func Clear() {
	db := GetInstance()
	db.mutex.Lock()
	defer db.mutex.Unlock()
	for k := range db.doctorAvailabilitySlot {
		delete(db.doctorAvailabilitySlot, k)
	}
	for k := range db.appointmentSlots {
		delete(db.appointmentSlots, k)
	}
}
