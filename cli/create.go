package cli

import (
	"deni1688/exercise_tracker/domain"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func getCreateCmd(service domain.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "creates a new exercise from args",
		RunE: func(cmd *cobra.Command, args []string) error {
			ex := new(domain.Exercise)

			category := args[0]
			name := args[1]
			weight, _ := strconv.ParseFloat(args[2], 64)
			duration, _ := strconv.Atoi(args[2])
			distance, _ := strconv.Atoi(args[2])
			reps, _ := strconv.Atoi(args[2])
			sets, _ := strconv.Atoi(args[2])

			ex.Category = category
			ex.Name = name
			ex.Weight = weight
			ex.Duration = duration
			ex.Distance = distance
			ex.Reps = reps
			ex.Sets = sets

			id, err := service.SaveExercise(ex)
			if err != nil {
				return err
			}

			fmt.Println("created exercise:", id)

			return nil
		},
	}
}
