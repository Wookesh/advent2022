package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

const test = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

func main() {
	data, err := os.ReadFile("./advent2022/19/input.txt")
	if err != nil {
		log.Fatalf("os.ReadFile() failed: %v", err)
	}

	//data = []byte(test)

	t1 := time.Now()
	resultOne := partOne(string(data))
	log.Printf("time: %v", time.Now().Sub(t1))
	log.Printf("ans 1: %v", resultOne)

	t2 := time.Now()
	resultTwo := partTwo(string(data))
	log.Printf("time: %v", time.Now().Sub(t2))
	log.Printf("ans 2: %v", resultTwo)
}

type Blueprint struct {
	ID                int
	OreRobotCost      Resources
	ClayRobotCost     Resources
	ObsidianRobotCost Resources
	GeodesRobotCost   Resources
}

type Resources struct {
	Ore      int
	Clay     int
	Obsidian int
	Geode    int
}

func (r *Resources) Add(o *Resources) Resources {
	return Resources{
		Ore:      r.Ore + o.Ore,
		Clay:     r.Clay + o.Clay,
		Obsidian: r.Obsidian + o.Obsidian,
		Geode:    r.Geode + o.Geode,
	}
}

func partOne(s string) int {
	var blueprints []*Blueprint
	for _, l := range strings.Split(s, "\n") {
		var (
			id     int
			oore   int
			core   int
			obore  int
			obclay int
			gore   int
			gobs   int
		)
		_, err := fmt.Sscanf(l, "Blueprint %v: Each ore robot costs %v ore. Each clay robot costs %v ore. Each obsidian robot costs %v ore and %v clay. Each geode robot costs %v ore and %v obsidian.", &id, &oore, &core, &obore, &obclay, &gore, &gobs)
		if err != nil {
			log.Fatal(err)
		}

		blueprints = append(blueprints, &Blueprint{
			ID:                id,
			OreRobotCost:      Resources{Ore: oore},
			ClayRobotCost:     Resources{Ore: core},
			ObsidianRobotCost: Resources{Ore: obore, Clay: obclay},
			GeodesRobotCost:   Resources{Ore: gore, Obsidian: gobs},
		})
	}

	totalQualityLevel := 0
	for _, blueprint := range blueprints {
		//log.Printf("Blueprint: %#v", blueprint)
		geodes := simulate(blueprint, Simulation{collects: Resources{Ore: 1}, iteration: 1, maxIterations: 24}, make(map[Simulation]int), maxUse(blueprint))
		qualityLevel := blueprint.ID * geodes
		log.Printf("Blueprint: %v, geodes: %v, qualityLevel: %v", blueprint.ID, geodes, qualityLevel)
		totalQualityLevel += qualityLevel
	}

	return totalQualityLevel
}

func maxUse(b *Blueprint) Resources {
	return Resources{
		Ore:      max(b.OreRobotCost.Ore, b.ClayRobotCost.Ore, b.ObsidianRobotCost.Ore, b.GeodesRobotCost.Ore),
		Clay:     max(b.OreRobotCost.Clay, b.ClayRobotCost.Clay, b.ObsidianRobotCost.Clay, b.GeodesRobotCost.Clay),
		Obsidian: max(b.OreRobotCost.Obsidian, b.ClayRobotCost.Obsidian, b.ObsidianRobotCost.Obsidian, b.GeodesRobotCost.Obsidian),
		Geode:    math.MaxInt,
	}
}

func max(k ...int) int {
	if len(k) == 0 {
		return 0
	}
	z := k[0]
	for _, v := range k {
		if z < v {
			z = v
		}
	}
	return z
}

type Simulation struct {
	r             Resources
	collects      Resources
	iteration     int
	maxIterations int
}

func simulate(blueprint *Blueprint, simulation Simulation, seen map[Simulation]int, maxUse Resources) int {
	if simulation.iteration > simulation.maxIterations {
		return simulation.r.Geode
	}
	if geodes, ok := seen[simulation]; ok {
		return geodes
	}

	var nextSimulations []Simulation

	if canBuild(&blueprint.GeodesRobotCost, &simulation.r) {
		nextSimulations = append(nextSimulations, Simulation{
			r:             build(&blueprint.GeodesRobotCost, &simulation.r),
			collects:      simulation.collects.Add(&Resources{Geode: 1}),
			iteration:     simulation.iteration + 1,
			maxIterations: simulation.maxIterations,
		})
	} else if canBuild(&blueprint.ObsidianRobotCost, &simulation.r) && maxUse.Obsidian > simulation.collects.Obsidian {
		nextSimulations = append(nextSimulations, Simulation{
			r:             build(&blueprint.ObsidianRobotCost, &simulation.r),
			collects:      simulation.collects.Add(&Resources{Obsidian: 1}),
			iteration:     simulation.iteration + 1,
			maxIterations: simulation.maxIterations,
		})
	} else if canBuild(&blueprint.ClayRobotCost, &simulation.r) && maxUse.Clay > simulation.collects.Clay {
		nextSimulations = append(nextSimulations, Simulation{
			r:             build(&blueprint.ClayRobotCost, &simulation.r),
			collects:      simulation.collects.Add(&Resources{Clay: 1}),
			iteration:     simulation.iteration + 1,
			maxIterations: simulation.maxIterations,
		})
	}
	if canBuild(&blueprint.OreRobotCost, &simulation.r) && maxUse.Ore > simulation.collects.Ore {
		nextSimulations = append(nextSimulations, Simulation{
			r:             build(&blueprint.OreRobotCost, &simulation.r),
			collects:      simulation.collects.Add(&Resources{Ore: 1}),
			iteration:     simulation.iteration + 1,
			maxIterations: simulation.maxIterations,
		})
	}
	nextSimulations = append(nextSimulations, Simulation{
		r:             simulation.r,
		collects:      simulation.collects,
		iteration:     simulation.iteration + 1,
		maxIterations: simulation.maxIterations,
	})

	maxGeode := 0
	for _, sim := range nextSimulations {
		sim.r = sim.r.Add(&simulation.collects)
		geodes := simulate(blueprint, sim, seen, maxUse)
		if geodes > maxGeode {
			maxGeode = geodes
		}
	}

	seen[simulation] = maxGeode

	return maxGeode
}

func canBuild(cost, r *Resources) bool {
	return cost.Ore <= r.Ore && cost.Clay <= r.Clay && cost.Obsidian <= r.Obsidian
}

func build(cost, r *Resources) Resources {
	return Resources{
		Ore:      r.Ore - cost.Ore,
		Clay:     r.Clay - cost.Clay,
		Obsidian: r.Obsidian - cost.Obsidian,
		Geode:    r.Geode,
	}
}

func partTwo(s string) int {
	var blueprints []*Blueprint
	for _, l := range strings.Split(s, "\n") {
		var (
			id     int
			oore   int
			core   int
			obore  int
			obclay int
			gore   int
			gobs   int
		)
		_, err := fmt.Sscanf(l, "Blueprint %v: Each ore robot costs %v ore. Each clay robot costs %v ore. Each obsidian robot costs %v ore and %v clay. Each geode robot costs %v ore and %v obsidian.", &id, &oore, &core, &obore, &obclay, &gore, &gobs)
		if err != nil {
			log.Fatal(err)
		}

		blueprints = append(blueprints, &Blueprint{
			ID:                id,
			OreRobotCost:      Resources{Ore: oore},
			ClayRobotCost:     Resources{Ore: core},
			ObsidianRobotCost: Resources{Ore: obore, Clay: obclay},
			GeodesRobotCost:   Resources{Ore: gore, Obsidian: gobs},
		})
	}

	totalQualityLevel := 0
	for _, blueprint := range blueprints[:3] {
		geodes := simulate(blueprint, Simulation{collects: Resources{Ore: 1}, iteration: 1, maxIterations: 32}, make(map[Simulation]int), maxUse(blueprint))
		log.Printf("Blueprint: %v, geodes: %v", blueprint.ID, geodes)
		totalQualityLevel *= geodes
	}

	return totalQualityLevel
}
