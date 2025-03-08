package models

// RegisterModels holds all models to be migrated automatically
var RegisterModels = []interface{}{
	&User{},
	&Post{},
	&Role{},
	&Permission{},
	&RoleHasPermissions{},
}
