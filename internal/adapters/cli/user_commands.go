package cli

import (
	"fmt"
	"time"

	"github.com/colfarl/budgy/internal/core"
)

type UserAddCmd struct {
	Name string `arg:"" help:"Username."`
}

type UserBalancesCmd struct {
	Name string `arg:"" help:"Username."`
	StartDate string `help:"sum transactions only after mm/dd/yyyy" short:"s"`
	EndDate string `help:"sum transactions only before mm/dd/yyyy." short:"e"`
}

type UserExpenseCmd struct {
	Name string `arg:"" help:"Username."`
	StartDate string `help:"sum transactions only after mm/dd/yyyy" short:"s"`
	EndDate string `help:"sum transactions only before mm/dd/yyyy." short:"e"`
}

type UserIncomeCmd struct {
	Name string `arg:"" help:"Username."`
	StartDate string `help:"sum transactions only after mm/dd/yyyy" short:"s"`
	EndDate string `help:"sum transactions only before mm/dd/yyyy." short:"e"`
}

func getStartEndDate(start, end string) (int64, int64, error) {
	var start_time int64
	var end_time int64

	if start == "" {
		start_time = 0		
	} else {
		t, err := time.Parse("01/02/2006", start)
		if err != nil {
			return 0, 0, err
		}
		start_time = t.Unix()
	}

	if end == "" {
		end_time = time.Now().Unix()
	} else {
		t, err := time.Parse("01/02/2006", end)
		if err != nil {
			return 0, 0, err
		}
		end_time = t.Unix()
	}

	return start_time, end_time, nil
}

func (c *UserBalancesCmd) Run(binds *Context) error {	
	start, end, err := getStartEndDate(c.StartDate, c.EndDate)
	if err != nil {
		return err
	}

	binds.Store.Commands <- core.GetUserBalances{
		Username: c.Name,
		StartDate: start,
		EndDate: end,
	}	
	for {
		select {
		case <-binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v\n", binds.Store.State.Error)

		case event := <-binds.Store.Events:
			if v, ok := event.(core.UserSummed); ok {
				total := 0.0
				for _, u := range v.Accounts {
					total += u.Balance
					fmt.Printf("%s: %.2f\n", u.Name, u.Balance)
				}
				
				fmt.Printf("\nNet: %.2f\n", total)
				return nil
			}
			if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			}
			return fmt.Errorf("Unknown error while creating user\n")		
		}
	}
}

func (c *UserAddCmd) Run(binds *Context) error {
	binds.Store.Commands <- core.CreateUser{Username: c.Name}	
	for {
		select {
		case <-binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v\n", binds.Store.State.Error)

		case event := <-binds.Store.Events:
			if _, ok := event.(core.UserCreated); ok {
				fmt.Println("USER ADDED")
				return nil
			}
			if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			}
			return fmt.Errorf("Unknown error while creating user\n")		
		}
	}
}

type UserListCmd struct{}

func (c *UserListCmd) Run(binds *Context) error { 
	binds.Store.Commands <- core.LoadAllUsers{}	
	for {
		select {
		case <-binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v", binds.Store.State.Error)

		case event := <-binds.Store.Events:
			if v, ok := event.(core.UsersLoaded); ok {
				if len(v.Usernames) == 0{
					fmt.Println("No users registered")
				}
				for i, u := range v.Usernames {
					fmt.Printf("%d. %s\n", i + 1, u)
				}
				return nil
			}
			if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			}
			return fmt.Errorf("Unknown error while listing users")		
		}
	}
}

type UserDeleteCmd struct{ Name string `arg:""` }
func (c *UserDeleteCmd) Run(binds *Context) error {
	binds.Store.Commands <- core.DeleteUser{Username: c.Name}
	for {
		select {
		case <- binds.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v", binds.Store.State.Error)
		case event := <- binds.Store.Events:
			if _, ok := event.(core.UserDeleted); ok {
				fmt.Println("USER DELETED")
				return nil
			}
			if binds.Store.State.Error != nil {
				return binds.Store.State.Error
			}
			return fmt.Errorf("Unknown error while creating user")		
		}
	}
}

func (c *Context) EventListen() error {
	for {
		select {
		case <- c.Ctx.Done():
			return fmt.Errorf("TIMEOUT OCCURRED: %v", c.Store.State.Error)

		case event := <- c.Store.Events:
			switch v := event.(type) {
			case core.UserTxnsGrouped:
				total := 0.0
				for _, u := range v.Groups {
					fmt.Printf("%v: %.2f\n", u.Name, u.Amount)
					total += u.Amount
				}
				fmt.Printf("\nNet: %.2f\n", total)
				return nil

			default:
				if c.Store.State.Error != nil {
					return c.Store.State.Error
				}
				return fmt.Errorf("Unknown error while creating user")		
			}
		}
	}
}

func (c *UserExpenseCmd) Run(binds *Context) error {
	start, end, err := getStartEndDate(c.StartDate, c.EndDate)
	if err != nil {
		return err
	}

	binds.Store.Commands <- core.SumTxnsByCategory{
		Username: c.Name,
		StartDate: start,
		EndDate: end,
		Income: false,
	}			
	return binds.EventListen()
}

func (c *UserIncomeCmd) Run(binds *Context) error {
	start, end, err := getStartEndDate(c.StartDate, c.EndDate)
	if err != nil {
		return err
	}

	binds.Store.Commands <- core.SumTxnsByCategory{
		Username: c.Name,
		StartDate: start,
		EndDate: end,
		Income: true,
	}		
	return binds.EventListen()
}
