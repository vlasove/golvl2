package hospital

import "log"

type Visitor interface {
	visitForMedical(*Patient)
	visitForDoctor(*Patient)
	visitForReception(*Patient)
	visitForCashier(*Patient)
}

type GeneralVisitor struct{}

var generalVisitor = &GeneralVisitor{}

func (gv *GeneralVisitor) visitForMedical(p *Patient) {
	log.Println("		medical visitor working")
	if p.medicineDone {
		log.Println("medicine already given to patient")
		return
	}
	log.Println("medical giving medicine to patient")
	p.medicineDone = true
}

func (gv *GeneralVisitor) visitForDoctor(p *Patient) {
	log.Println("		doctor visitor working")
	if p.doctorCheckUpDone {
		log.Println("doctor checkup already done")
		return
	}
	log.Println("doctor checking patient")
	p.doctorCheckUpDone = true
}

func (gv *GeneralVisitor) visitForReception(p *Patient) {
	log.Println("		reception visitor working")
	if p.registrationDone {
		log.Println("patient registration already done")
		return
	}
	log.Println("reception registering patient")
	p.registrationDone = true
}

func (gv *GeneralVisitor) visitForCashier(p *Patient) {
	log.Println("		cashier visitor working")
	if p.paymentDone {
		log.Println("payment done")
		return
	}
	log.Println("cashier getting money from patient patient")
}
