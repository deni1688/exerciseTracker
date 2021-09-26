package cli

import (
	"deni1688/exercise_tracker/domain"
	"github.com/spf13/cobra"
)

type ExerciseCLI interface {
	Execute() error
}

type exerciseCLI struct {
	rootCmd *cobra.Command
}

func (ec *exerciseCLI) Execute() error {
	return ec.rootCmd.Execute()
}



func NewExerciseCLI(service domain.ExerciseService) ExerciseCLI {
	rootCmd := &cobra.Command{
		Use:   "exercise tracker cli",
		Short: "cli tool for managing exercises",
	}

	rootCmd.AddCommand(getCreateCmd(service))
	rootCmd.AddCommand(getListCmd(service))

	return &exerciseCLI{rootCmd}
}



