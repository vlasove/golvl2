package hospital

// переделать на объект пациента
type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func NewPatient(name string) *Patient {
	return &Patient{
		name: name,
	}
}
