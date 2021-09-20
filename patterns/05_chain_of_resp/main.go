package main

import "github.com/vlasove/golvl2/patterns/05_chain_of_resp/hospital"

// Визитор добавлен!
func main() {
	cashierHandler := &hospital.Cashier{}
	medicalHandler := hospital.NewMedical(cashierHandler)
	doctorHandler := hospital.NewDoctor(medicalHandler)
	receptionHandler := hospital.NewReception(doctorHandler)

	patient := hospital.NewPatient("Vasya")
	receptionHandler.Execute(patient)
}
