package src

import (
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

const PATH_TO_VARIABLES = "./"

var (
	Application_id, Token string
)

func Initialize() string {
	// Reading the environment variables
	mod, diag := tfconfig.LoadModule(PATH_TO_VARIABLES)
	variables := mod.Variables

	// Checking the error
	if diag.HasErrors() {
		panic(diag.Err())
	}

	token, ok := variables["token"].Default.(string)
	application_id, ok := variables["application_id"].Default.(string)

	if !ok {
		panic("Application_id or token are not recognized, non-existent or wrongly formatted")
	}

	// Setting the globals
	Application_id = application_id
	Token = token

	// Returning the token

	return token
}
