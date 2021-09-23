package cli

import (
	"deni1688/exercise_tracker/domain"
	"github.com/lensesio/tableprinter"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func getListCmd(service domain.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "lists the stored exercises",
		Run: func(cmd *cobra.Command, args []string){
			results, err := service.ListExercises()
			if err != nil {
				log.Fatal(err)
				return
			}

			tableprinter.Print(os.Stdout, results)
		},
	}
}
