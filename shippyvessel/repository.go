package main

import (
	pb "github.com/vlasove/shippyvessel/proto/vessel"
	"gopkg.in/mgo.v2"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessels"
)

//Repository ...
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Close()
}

//VesselRepository ...
type VesselRepository struct {
	session *mgo.Session
}

//Close ...
func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}

//FindAvailable ...
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	// ...
	var vessel *pb.Vessel
	err := repo.collection().Find(?????).One(&vessel)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}
