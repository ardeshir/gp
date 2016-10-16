package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {

	plants := []PowerPlant {
		PowerPlant{hydro, 300, active},
		PowerPlant{wind, 30, active},
		PowerPlant{wind, 25, inactive},
		PowerPlant{wind, 35, active},
		PowerPlant{solar, 45, unavailable},
		PowerPlant{solar, 40, inactive},
	} // end of structs	


  grid := PowerGrid{300, plants} 
  

  if option, err := requestOption(); err == nil {
	fmt.Println("")

	switch option {
		case "1":
			grid.generatePlantReport()
		case "2":
		// TODO 
		}
  } else {
	 fmt.Println(err.Error())
  } 
  
}

 func requestOption() (option string, err error) {
 	
  fmt.Println("1) Generate Power Plant Report")
  fmt.Println("2) Generate Power Grid Report")
 
  fmt.Print("Please choose an Option: ")
  fmt.Scanln(&option)

  if option != "1" && option != "2" {
	err = errors.New("Invalid option selected")
  }
  
   return 
 }

func generatePlantCapacityReport(plantCapacities ...float64) {
  for idx, cap := range plantCapacities {
	fmt.Printf("Plant %d capacity: %.0f\n", idx, cap)
  }
} // end of genPlantCapRep


func generatePowerGridReport(activePlant []int, plantCapacities []float64, gridLoad float64) {
   capacity := 0.
 
	for _, plantId := range activePlant {
		capacity += plantCapacities[plantId]
	}
	fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
	fmt.Printf("%-20s%.0f\n", "Load: ", gridLoad)
	fmt.Printf("%-20s.1f\n", "Utilization: ", gridLoad/capacity*100)

	
} // end of genPowerGridRep




//   ADDING NEW TYPES

type PlantType string

const (
	hydro PlantType = "Hydro"
	wind  PlantType	= "Wind"
	solar PlantType = "Solar"
)

type PlantStatus string

const (
	active PlantStatus   = "Active"
	inactive PlantStatus = "Inactive"
	unavailable			 = "Unavailable" 
)

type PowerPlant struct {
	plantType PlantType
	capacity  float64
	status    PlantStatus
	
}

type PowerGrid struct {
 load float64	
 plants []PowerPlant
}

func (pg *PowerGrid) generatePlantReport() {
	for idx, p := range pg.plants {
			label := fmt.Sprintf("%s%d", "Plant #", idx)
			fmt.Println(label)
			fmt.Println(strings.Repeat("-", len(label)))
			fmt.Printf("%-20s%s\n", "Type:", p.plantType)
			fmt.Printf("%-20s%.0f\n", "Capacity:", p.capacity)
			fmt.Printf("%-20s%s\n", "Status:", p.status)
			fmt.Println("")
	}

}