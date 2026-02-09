package cli

import (
	"fmt"

	"github.com/colfarl/budgy/internal/core"
)

type UserAddCmd struct {
	Name string `arg:"" help:"Username."`
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
