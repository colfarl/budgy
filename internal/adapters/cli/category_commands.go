package cli

import (
	"fmt"

	"github.com/colfarl/budgy/internal/core"
)

type CategoryAddCmd struct {
	Name 		string 	`arg:"" help:"Name of Category"`
	IsIncome	bool	`help:"Does category describe income or expenses"`
}

type CategoryDeleteCmd struct {
	CategoryID	int64 	`arg:"" help:"ID of category to delete"`
}

type CategoryListCmd struct {}


func (c *CategoryAddCmd) Run(binds *Context) error {
	binds.Store.Commands <- core.CreateCategory{
		Name: c.Name,
		IsIncome: c.IsIncome,
	}

	for {
		select {
		case <- binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v", binds.Store.State.Error)
		case event := <- binds.Store.Events:
			if v, ok := event.(core.CategoryCreated); ok {
				fmt.Printf("Category: ID: %d, Name: %s, Income: %v\n", v.Category.ID, v.Category.Name, v.Category.IsIncome)
				return nil
			} else if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			} else {
				return fmt.Errorf("Unknown error occured while creating txn")
			}

		}
	}
}

func (c *CategoryDeleteCmd) Run(binds *Context) error {
	binds.Store.Commands <- core.DeleteCategory{
		ID: c.CategoryID,
	}
	for {
		select {
		case <- binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v", binds.Store.State.Error)
		case event := <- binds.Store.Events:
			if v, ok := event.(core.CategoryDeleted); ok {
				fmt.Printf("DELETED CATEGORY: ID - %d\n", v.ID)
				return nil
			} else if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			} else {
				return fmt.Errorf("Unknown error occured while creating txn")
			}

		}
	}
}

func (c *CategoryListCmd) Run(binds *Context) error {
	binds.Store.Commands <- core.LoadAllCategories{}

	for {
		select {
		case <- binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v", binds.Store.State.Error)
		case event := <- binds.Store.Events:
			if v, ok := event.(core.CategoriesLoaded); ok {
				if len(v.Categories) == 0 {
					fmt.Println("No Categories :(")
					return nil
				}

				fmt.Println("Expense Categories")
				count := 0
				for i := range v.Categories {
					if v.Categories[i].IsIncome {continue}
					fmt.Printf("	%d. ID: %d; Name: %s\n", i + 1, v.Categories[i].ID, v.Categories[i].Name)
					count += 1
				}

				if count == 0 {
					fmt.Println("	No Expenses (I wish)")
				}
				

				fmt.Println("Income Categories")
				count = 0
				for i := range v.Categories {
					if !v.Categories[i].IsIncome {continue}
					fmt.Printf("	%d. ID: %d; Name: %s\n", i + 1, v.Categories[i].ID, v.Categories[i].Name)
					count += 1
				}

				if count == 0 {
					fmt.Println("	No Income (uh oh...)")
				}
				

				return nil
			} else if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			} else {
				return fmt.Errorf("Unknown error occured while creating txn")
			}

		}
	}
}



