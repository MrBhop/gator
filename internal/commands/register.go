package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/MrBhop/BlogAggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Usage: %v <username>\n", cmd.name)
	}
	userName := cmd.args[0]

	return register(s, userName)
}

func register(s *state, userName string) error {
	user, err := s.DataBase.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: userName,
	})
	if err != nil {
		return fmt.Errorf("Could not create user '%v'\n", userName)
	}

	fmt.Printf("User '%v' was created\n", userName)
	fmt.Println(userToString(user))

	return login(s, userName)
}

func userToString(item database.User) string {
	output := fmt.Sprintln("User{")
	output += fmt.Sprintf("ID: %v\n", item.ID)
	output += fmt.Sprintf("CreatedAt: %v\n", item.CreatedAt)
	output += fmt.Sprintf("UpdatedAt: %v\n", item.UpdatedAt)
	output += fmt.Sprintf("Name: %v\n", item.Name)
	output += fmt.Sprintln("}")
	return output
}
