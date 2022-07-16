package vmmanagerapi

import (
	"encoding/json"
	"fmt"

	"github.com/gadost/go-vmmanager-api/types"
)

func (a *Api) Recipes() (*types.Recipes, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/recipe", requestTypeGet, DefaultService)

	var r *types.Recipes
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) Recipe(rid int) (*types.RecipeResponse, error) {
	uri := fmt.Sprintf("/recipe/%d", rid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var r *types.RecipeResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) RecipeNew(params *types.RecipeNew) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, "/recipe", requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) RecipeUpdate(rid int, params *types.RecipeNew) (*types.Task, error) {
	uri := fmt.Sprintf("/recipe/%d", rid)
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) RecipeDelete(rid int) (*types.DeletedResponse, error) {
	uri := fmt.Sprintf("/recipe/%d", rid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var r *types.DeletedResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) RecipeParams(rid int) ([]byte, error) {
	uri := fmt.Sprintf("/recipe/%d/params", rid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	return bodyResp, err
}

func (a *Api) RecipeSave(params *types.RecipesSaveRequest) (*types.Task, error) {
	uri := "/recipe"
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ScriptVariables() (*types.ScriptVariablesResponse, error) {
	bodyResp, err := a.NewRequest(NilPayload, "/script_variable", requestTypeGet, DefaultService)

	var r *types.ScriptVariablesResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ScriptVariableNew(params *types.ScriptVariable) (*types.Task, error) {
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, "/script_variable", requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ScriptVariableUpdate(scid int, params *types.ScriptVariable) (*types.Task, error) {
	uri := fmt.Sprintf("/script_variable/%d", scid)
	payload, _ := json.Marshal(params)
	bodyResp, err := a.NewRequest(payload, uri, requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ScriptVariable(scid int) (*types.ScriptVariable, error) {
	uri := fmt.Sprintf("/script_variable/%d", scid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeGet, DefaultService)

	var r *types.ScriptVariable
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ScriptVariableDelete(scid int) (*types.DeletedResponse, error) {
	uri := fmt.Sprintf("/script_variable/%d", scid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypeDelete, DefaultService)

	var r *types.DeletedResponse
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ScriptVariableEnable(scid int) (*types.Task, error) {
	uri := fmt.Sprintf("/script_variable/%d/enable", scid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}

func (a *Api) ScriptVariableDisable(scid int) (*types.Task, error) {
	uri := fmt.Sprintf("/script_variable/%d/disable", scid)
	bodyResp, err := a.NewRequest(NilPayload, uri, requestTypePost, DefaultService)

	var r *types.Task
	json.Unmarshal(bodyResp, &r)
	return r, err
}
