package main

import (
	"strconv"
	"sync"
	"time"
)

type (
	world struct {
		maxX, maxY, time, maxTime, Ubertraveled int
		ubers                                   []*Uber // ah... Pointers, Sweet XD
		clients                                 map[int][]*passenger
		waitingclients                          []*passenger
		Runtime                                 string `json:"runtime"`
		X                                       []int  `json:"x"` // For parse to a JSON :)
		Y                                       []int  `json:"y"`
		log                                     string
		end                                     bool
		filterWaitingClients                    worldFilterWaitingClients
		addClient                               worldAddClient
		clientsToWaitingList                    worldclientstowaitinglist
		uberForClient                           worlduberforclient
		getAvalaibleUbers                       worldGetAvalaibleUbers
		runwWithoutPram                         worldRunWithoutPram
		runWithPram                             worldRunWithPram
		instantSave                             worldInstantSave
	}
	worldFilterWaitingClients func(world2 *world)
	worldAddClient            func(world2 *world, client2 *passenger)
	worldclientstowaitinglist func(world2 *world)
	worlduberforclient        func(world2 *world, client2 *passenger, ubers *[]*Uber) bool
	worldGetAvalaibleUbers    func(world2 *world) []*Uber
	worldRunWithoutPram       func(world2 *world)
	worldRunWithPram          func(world2 *world)
	worldInstantSave          func(world2 *world)
)

func createWorld(maxTime int) *world {
	w := world{
		maxX:           1000,
		maxY:           1000,
		time:           0,
		maxTime:        maxTime,
		Ubertraveled:   0,
		ubers:          make([]*Uber, 0),
		clients:        make(map[int][]*passenger),
		waitingclients: make([]*passenger, 0),
		X:              make([]int, maxTime),
		Y:              make([]int, maxTime),
		filterWaitingClients: func(world2 *world) {
			oldlist := world2.waitingclients
			newlist := make([]*passenger, 0)
			for _, client := range oldlist {
				if client.waiting {
					newlist = append(newlist, client)
				}
			}
			world2.waitingclients = newlist
		},
		addClient: func(world2 *world, client2 *passenger) {
			if client2.time < world2.time {
				client2.time = world2.time
			}
			world2.clients[client2.time] = append(world2.clients[client2.time], client2)
		},
		clientsToWaitingList: func(world2 *world) {
			world2.waitingclients = append(world2.waitingclients, world2.clients[world2.time]...)
		},
		getAvalaibleUbers: func(world2 *world) []*Uber {
			oldList := world2.ubers
			newlist := make([]*Uber, 0)
			for _, uber := range oldList {
				if uber.avalaible {
					newlist = append(newlist, uber)
				}
			}
			return newlist
		},
		uberForClient: func(world2 *world, client2 *passenger, ubers *[]*Uber) bool {
			total := 0.0
			if len(*ubers) == 0 { // No ubers avalaible
				return false
			}
			for _, uber := range *ubers {
				total += DistanceBetween(client2, uber)
			}
			probs := make(map[*Uber]float64)
			for _, uber := range *ubers {
				if total == 0 {
					probs[uber] = 1
				} else {
					probs[uber] = DistanceBetween(client2, uber) / total
				}
			}
			uber := selectConditionedUber(probs)
			uber.setClient(uber, client2)
			return true
		},
		runwWithoutPram: func(world2 *world) {
			start := time.Now()
			for world2.time < world2.maxTime {
				world2.filterWaitingClients(world2)
				world2.clientsToWaitingList(world2)
				for _, client := range world2.waitingclients {
					ubers := world2.getAvalaibleUbers(world2)
					if len(ubers) == 0 {
						break
					}
					if !world2.uberForClient(world2, client, &ubers) {
						break // Because there's no more ubers avalaible
					}
				}
				for _, uber := range world2.ubers {
					if uber.avalaible {
						continue
					}
					uber.checkMove(uber)
				}
				world2.time += 1
				world2.instantSave(world2)
			}
			end := time.Now()
			world2.Runtime = end.Sub(start).String()
		},
		runWithPram: func(world2 *world) {
			start := time.Now()
			for world2.time < world2.maxTime {
				world2.filterWaitingClients(world2)
				world2.clientsToWaitingList(world2)
				var wg = new(sync.WaitGroup)
				for _, client := range world2.waitingclients {
					ubers := world2.getAvalaibleUbers(world2)
					if len(ubers) == 0 {
						break
					}
					wg.Add(1)
					go func(client *passenger, ubers []*Uber) {
						defer wg.Done()
						if !world2.uberForClient(world2, client, &ubers) {
							return // Because there's no more ubers avalaible
						}
					}(client, ubers)
				}
				wg.Wait()
				for _, uber := range world2.ubers {
					if uber.avalaible {
						continue
					}
					wg.Add(1)
					go func(uber *Uber) {
						defer wg.Done()
						uber.checkMove(uber)
					}(uber)
				}
				wg.Wait()
				world2.time += 1
				world2.instantSave(world2)
				//log.Println(world2.time, world2.Ubertraveled)
			}
			end := time.Now()
			world2.Runtime = end.Sub(start).String()
		},
		instantSave: func(world2 *world) {
			world2.log += strconv.Itoa(world2.time) + " " + strconv.Itoa(world2.Ubertraveled) + "\n"
			world2.X[world2.time-1] = world2.time
			world2.Y[world2.time-1] = world2.Ubertraveled
			if world2.time == 12000 {
				world2.end = true
			}
			//world2.X = append(world2.X, world2.time)
			//world2.Y = append(world2.Y, world2.Ubertraveled)
		},
	}
	return &w
}
