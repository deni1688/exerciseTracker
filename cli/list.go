package cli

import (
	"encoding/json"
	"fmt"
	"github.com/deni1688/exercise_tracker/domain"
	"github.com/spf13/cobra"
	"log"
)

func getListCmd(service domain.ExerciseService) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "lists the stored exercises",
		Run: func(cmd *cobra.Command, args []string) {
			results, err := service.ListExercises()
			if err != nil {
				log.Fatal(err)
				return
			}

			exercisesJSON, err := json.MarshalIndent(results, "", "  ")
			if err != nil {
				log.Fatalf(err.Error())
			}

			fmt.Println(string(exercisesJSON))
		},
	}
}
