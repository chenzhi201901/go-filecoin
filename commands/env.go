package commands

import (
	"context"

	cmds "gx/ipfs/Qma6uuSyjkecGhMFFLfzyJDPyoDtNJSHJNweDccZhaWkgU/go-ipfs-cmds"

	"github.com/filecoin-project/go-filecoin/api"
	"github.com/filecoin-project/go-filecoin/plumbing"
)

// Env is the environment passed to commands. Implements cmds.Environment.
type Env struct {
	ctx         context.Context
	api         api.API
	plumbingAPI *plumbing.API
}

var _ cmds.Environment = (*Env)(nil)

// Context returns the context of the environment.
func (ce *Env) Context() context.Context {
	return ce.ctx
}

// API returns the associated FilecoinAPI object.
func (ce *Env) API() api.API {
	return ce.api
}

// GetAPI returns the Filecoin API object of the environment.
func GetAPI(env cmds.Environment) api.API {
	ce := env.(*Env)
	return ce.API()
}

// GetPlumbingAPI returns the api2.Filecoin interface from the environment.
func GetPlumbingAPI(env cmds.Environment) *plumbing.API {
	ce := env.(*Env)
	return ce.plumbingAPI
}
