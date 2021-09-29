package cli

import (
	"deni1688/exercise_tracker/domain"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func getCreateCmd(service domain.ExerciseService) *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "creates a new exercise from args",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			data, err := cmd.Flags().GetString("data")
			if err != nil {
				log.Fatal(err)
				return
			}

			var ex *domain.ExerciseRequest
			err = json.Unmarshal([]byte(data), &ex)
			if err != nil {
				log.Fatal(err)
				return
			}

			id, err := service.SaveExercise(ex)
			if err != nil {
				log.Fatal(err)
				return
			}

			fmt.Println("created exercise:", id)
		},
	}

	createCmd.Flags().StringP(
		"data",
		"d",
		"'{}'",
		"--data='{...exercise props}'",
	)

	return createCmd
}
