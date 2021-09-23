package cli

import (
	"deni1688/exercise_tracker/domain"
	"fmt"
	"github.com/spf13/cobra"
)

func getCreateCmd(service domain.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "creates a new exercise from args",
		RunE: func(cmd *cobra.Command, args []string) error {
			ex := new(domain.Exercise)

			id, err := service.SaveExercise(ex)
			if err != nil {
				return err
			}

			fmt.Println("created exercise:", id)

			return nil
		},
	}
}
