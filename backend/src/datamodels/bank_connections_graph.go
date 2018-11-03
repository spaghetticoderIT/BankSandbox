package datamodels

import (
	"errors"
	"time"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/utils"
)

type BankConnectionRoute struct {
	FromBIC  string        `bson:"from" json:"from"`
	ToBIC    string        `bson:"to" json:"to"`
	FeesCost uint8         `bson:"feesCost" json:"feesCost"`
	TimeCost time.Duration `bson:"timeCost" json:"timeCost"`
}

func NewBankConnectionRoute(fromBIC string, toBIC string, feesCost uint8, timeCost time.Duration) *BankConnectionRoute {
	route := new(BankConnectionRoute)
	route.FromBIC = fromBIC
	route.ToBIC = toBIC
	route.FeesCost = feesCost
	route.TimeCost = timeCost
	return route
}

type dijkstraBy string

const (
	Time     dijkstraBy = "time"
	Fees     dijkstraBy = "fees"
	distance dijkstraBy = "distance"
)

type BankConnectionsGraph struct {
	BanksBICs []string               `bson:"banks" json:"banks"`
	Routes    []*BankConnectionRoute `bson:"routes" json:"routes"`
}

func NewBankConnectionsGraph() *BankConnectionsGraph {
	graph := new(BankConnectionsGraph)
	graph.BanksBICs = make([]string, 0)
	graph.Routes = make([]*BankConnectionRoute, 0)
	return graph
}

func (graph *BankConnectionsGraph) PushNewBank(bic string) {
	graph.BanksBICs = append(graph.BanksBICs, bic)
}

func (graph *BankConnectionsGraph) CreateNewRoute(route *BankConnectionRoute) error {
	// Check if route.FromBIC exists in graph banks
	if !utils.IsInArray(route.FromBIC, graph.BanksBICs) {
		return errors.New(`The given route could not be appended to the graph because
		 route's from Bank doesn't match any graph's bank`)
	}

	// Check if route.ToBIC exists in graph banks
	if !utils.IsInArray(route.ToBIC, graph.BanksBICs) {
		return errors.New(`The given route could not be appended to the graph because
		 route's to Bank doesn't match any graph's bank`)
	}

	graph.Routes = append(graph.Routes, route)
	return nil
}

// TODO
func (graph *BankConnectionsGraph) FindShortestRoute(from string, to string) ([]string, error) {
	return []string{}, nil
}

// TODO
func (graph *BankConnectionsGraph) FindCheapestRoute(from string, to string) ([]string, error) {
	return []string{}, nil
}

// TODO
func (graph *BankConnectionsGraph) FindFastesRoute(from string, to string) ([]string, error) {
	return []string{}, nil
}

// TODO
func dijkstra(graph *BankConnectionsGraph, from string, to string, costType dijkstraBy) []string {
	return []string{}
}