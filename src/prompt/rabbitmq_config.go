package prompts

import (
	"errors"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/asaskevich/govalidator"
)

type NewCtxPromptContent struct {
	Name      string
	Host      string
	User      string
	Password  string
	Port      string
	AdminPort string
}

// the questions to ask
var simpleQs = []*survey.Question{
	{
		Name: "Name",
		Prompt: &survey.Input{
			Message: "Name:",
		},
		Validate: func(input interface{}) error {
			if !govalidator.IsAlpha(input.(string)) {
				return errors.New("Context Name can only have alpha characters")
			}
			return nil
		},
	},
	{
		Name: "Host",
		Prompt: &survey.Input{
			Message: "Host:",
		},
		Validate: func(input interface{}) error {
			if !govalidator.IsDNSName(input.(string)) {
				return errors.New("Please enter a valid DNS address")
			}
			return nil
		},
	},
	{
		Name: "User",
		Prompt: &survey.Input{
			Message: "User:",
		},
		Validate: func(input interface{}) error {
			if !govalidator.IsAlphanumeric(input.(string)) {
				return errors.New("Rabbitmq Username can only have alphanumeric characters")
			}
			return nil
		},
	},
	{
		Name: "Password",
		Prompt: &survey.Password{
			Message: "Password:",
		},
		Validate: survey.Required,
	},
	{
		Name: "Port",
		Prompt: &survey.Input{
			Message: "Port:",
		},
		Validate: func(input interface{}) error {
			if !govalidator.IsNumeric(input.(string)) {
				return errors.New("Port must be numeric")
			}
			return nil
		},
	},
	{
		Name: "AdminPort",
		Prompt: &survey.Input{
			Message: "AdminPort:",
		},
		Validate: func(input interface{}) error {
			if !govalidator.IsNumeric(input.(string)) {
				return errors.New("AdminPort must be numeric")
			}
			return nil
		},
	},
}

func PromptContext() *NewCtxPromptContent {
	answers := &NewCtxPromptContent{}
	err := survey.Ask(simpleQs, answers)

	if err != nil {
		log.Panic(err.Error())
		return nil
	}

	return answers
}
