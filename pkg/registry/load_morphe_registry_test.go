package registry_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/kaloseia/morphe-go/internal/testutils"
	"github.com/kaloseia/morphe-go/pkg/registry"
	"github.com/kaloseia/morphe-go/pkg/registry/cfg"
	"github.com/kaloseia/morphe-go/pkg/yaml"
)

type LoadMorpheRegistryTestSuite struct {
	suite.Suite

	TestDirPath string

	ModelsDirPath   string
	EntitiesDirPath string
}

func TestLoadMorpheRegistryTestSuite(t *testing.T) {
	suite.Run(t, new(LoadMorpheRegistryTestSuite))
}

func (suite *LoadMorpheRegistryTestSuite) SetupTest() {
	suite.TestDirPath = testutils.GetTestDirPath()

	suite.ModelsDirPath = filepath.Join(suite.TestDirPath, "registry", "minimal", "models")
	suite.EntitiesDirPath = filepath.Join(suite.TestDirPath, "registry", "minimal", "entities")
}

func (suite *LoadMorpheRegistryTestSuite) TearDownTest() {
	suite.TestDirPath = ""
}

func (suite *LoadMorpheRegistryTestSuite) TestLoadMorpheRegistry() {
	loadHooks := registry.LoadMorpheRegistryHooks{}
	config := cfg.MorpheLoadRegistryConfig{
		RegistryModelsDirPath:   suite.ModelsDirPath,
		RegistryEntitiesDirPath: suite.EntitiesDirPath,
	}

	r, registryErr := registry.LoadMorpheRegistry(loadHooks, config)

	suite.NoError(registryErr)
	suite.NotNil(r)

	allModels := r.GetAllModels()
	suite.Len(allModels, 2)

	model0, modelExists0 := allModels["Person"]
	suite.True(modelExists0)
	suite.Equal(model0.Name, "Person")

	suite.Len(model0.Fields, 3)

	modelField00, fieldExists00 := model0.Fields["ID"]
	suite.True(fieldExists00)
	suite.Equal(modelField00.Type, yaml.ModelFieldTypeAutoIncrement)
	suite.Len(modelField00.Attributes, 1)
	suite.Equal(modelField00.Attributes[0], "mandatory")

	modelField01, fieldExists01 := model0.Fields["FirstName"]
	suite.True(fieldExists01)
	suite.Equal(modelField01.Type, yaml.ModelFieldTypeString)
	suite.Len(modelField01.Attributes, 0)

	modelField02, fieldExists02 := model0.Fields["LastName"]
	suite.True(fieldExists02)
	suite.Equal(modelField02.Type, yaml.ModelFieldTypeString)
	suite.Len(modelField02.Attributes, 0)

	suite.Len(model0.Identifiers, 2)
	modelIDs00, idsExist00 := model0.Identifiers["primary"]
	suite.True(idsExist00)
	suite.ElementsMatch(modelIDs00.Fields, []string{"ID"})

	modelIDs01, idsExist01 := model0.Identifiers["name"]
	suite.True(idsExist01)
	suite.ElementsMatch(modelIDs01.Fields, []string{"FirstName", "LastName"})

	suite.Len(model0.Related, 1)

	modelRelated00, relatedExists00 := model0.Related["ContactInfo"]
	suite.True(relatedExists00)
	suite.Equal(modelRelated00.Type, "HasOne")

	model1, modelExists1 := allModels["ContactInfo"]
	suite.True(modelExists1)
	suite.Equal(model1.Name, "ContactInfo")

	suite.Len(model1.Fields, 2)

	modelField10, fieldExists10 := model1.Fields["ID"]
	suite.True(fieldExists10)
	suite.Equal(modelField10.Type, yaml.ModelFieldTypeAutoIncrement)
	suite.Len(modelField10.Attributes, 1)
	suite.Equal(modelField10.Attributes[0], "mandatory")

	modelField11, fieldExists11 := model1.Fields["Email"]
	suite.True(fieldExists11)
	suite.Equal(modelField11.Type, yaml.ModelFieldTypeString)
	suite.Len(modelField11.Attributes, 0)

	suite.Len(model1.Identifiers, 2)
	modelID10, idExists10 := model1.Identifiers["primary"]
	suite.True(idExists10)
	suite.ElementsMatch(modelID10.Fields, []string{"ID"})

	modelIDs11, idsExist11 := model1.Identifiers["email"]
	suite.True(idsExist11)
	suite.ElementsMatch(modelIDs11.Fields, []string{"Email"})

	suite.Len(model1.Related, 1)

	modelRelated10, relatedExists10 := model1.Related["Person"]
	suite.True(relatedExists10)
	suite.Equal(modelRelated10.Type, "ForOne")
}

func (suite *LoadMorpheRegistryTestSuite) TestLoadMorpheRegistry_Failure_InvalidPaths() {
	loadHooks := registry.LoadMorpheRegistryHooks{}
	config := cfg.MorpheLoadRegistryConfig{
		RegistryModelsDirPath:   "invalid path",
		RegistryEntitiesDirPath: "invalid path",
	}

	r, registryErr := registry.LoadMorpheRegistry(loadHooks, config)

	suite.ErrorContains(registryErr, "error reading directory 'invalid path")
	suite.Nil(r)
}

func (suite *LoadMorpheRegistryTestSuite) TestLoadMorpheRegistry_StartHook_Successful() {
	loadHooks := registry.LoadMorpheRegistryHooks{
		OnRegistryLoadStart: func(config cfg.MorpheLoadRegistryConfig) (cfg.MorpheLoadRegistryConfig, error) {
			config.RegistryModelsDirPath = suite.ModelsDirPath
			config.RegistryEntitiesDirPath = suite.EntitiesDirPath
			return config, nil
		},
	}
	config := cfg.MorpheLoadRegistryConfig{
		RegistryModelsDirPath:   "invalid/path",
		RegistryEntitiesDirPath: "invalid/path",
	}

	r, registryErr := registry.LoadMorpheRegistry(loadHooks, config)

	suite.NoError(registryErr)
	suite.NotNil(r)

	allModels := r.GetAllModels()
	suite.Len(allModels, 2)

	model0, modelExists0 := allModels["Person"]
	suite.True(modelExists0)
	suite.Equal(model0.Name, "Person")

	suite.Len(model0.Fields, 3)

	modelField00, fieldExists00 := model0.Fields["ID"]
	suite.True(fieldExists00)
	suite.Equal(modelField00.Type, yaml.ModelFieldTypeAutoIncrement)
	suite.Len(modelField00.Attributes, 1)
	suite.Equal(modelField00.Attributes[0], "mandatory")

	modelField01, fieldExists01 := model0.Fields["FirstName"]
	suite.True(fieldExists01)
	suite.Equal(modelField01.Type, yaml.ModelFieldTypeString)
	suite.Len(modelField01.Attributes, 0)

	modelField02, fieldExists02 := model0.Fields["LastName"]
	suite.True(fieldExists02)
	suite.Equal(modelField02.Type, yaml.ModelFieldTypeString)
	suite.Len(modelField02.Attributes, 0)

	suite.Len(model0.Identifiers, 2)
	modelIDs00, idsExist00 := model0.Identifiers["primary"]
	suite.True(idsExist00)
	suite.ElementsMatch(modelIDs00.Fields, []string{"ID"})

	modelIDs01, idsExist01 := model0.Identifiers["name"]
	suite.True(idsExist01)
	suite.ElementsMatch(modelIDs01.Fields, []string{"FirstName", "LastName"})

	suite.Len(model0.Related, 1)

	modelRelated00, relatedExists00 := model0.Related["ContactInfo"]
	suite.True(relatedExists00)
	suite.Equal(modelRelated00.Type, "HasOne")

	model1, modelExists1 := allModels["ContactInfo"]
	suite.True(modelExists1)
	suite.Equal(model1.Name, "ContactInfo")

	suite.Len(model1.Fields, 2)

	modelField10, fieldExists10 := model1.Fields["ID"]
	suite.True(fieldExists10)
	suite.Equal(modelField10.Type, yaml.ModelFieldTypeAutoIncrement)
	suite.Len(modelField10.Attributes, 1)
	suite.Equal(modelField10.Attributes[0], "mandatory")

	modelField11, fieldExists11 := model1.Fields["Email"]
	suite.True(fieldExists11)
	suite.Equal(modelField11.Type, yaml.ModelFieldTypeString)
	suite.Len(modelField11.Attributes, 0)

	suite.Len(model1.Identifiers, 2)
	modelID10, idExists10 := model1.Identifiers["primary"]
	suite.True(idExists10)
	suite.ElementsMatch(modelID10.Fields, []string{"ID"})

	modelIDs11, idsExist11 := model1.Identifiers["email"]
	suite.True(idsExist11)
	suite.ElementsMatch(modelIDs11.Fields, []string{"Email"})

	suite.Len(model1.Related, 1)

	modelRelated10, relatedExists10 := model1.Related["Person"]
	suite.True(relatedExists10)
	suite.Equal(modelRelated10.Type, "ForOne")
}

func (suite *LoadMorpheRegistryTestSuite) TestLoadMorpheRegistry_StartHook_Failure() {
	loadHooks := registry.LoadMorpheRegistryHooks{
		OnRegistryLoadStart: func(config cfg.MorpheLoadRegistryConfig) (cfg.MorpheLoadRegistryConfig, error) {
			return config, fmt.Errorf("compile model start hook error")
		},
	}
	config := cfg.MorpheLoadRegistryConfig{
		RegistryModelsDirPath:   suite.ModelsDirPath,
		RegistryEntitiesDirPath: suite.EntitiesDirPath,
	}

	r, registryErr := registry.LoadMorpheRegistry(loadHooks, config)

	suite.ErrorContains(registryErr, "compile model start hook error")
	suite.Nil(r)
}

func (suite *LoadMorpheRegistryTestSuite) TestLoadMorpheRegistry_SuccessHook_Successful() {
	loadHooks := registry.LoadMorpheRegistryHooks{
		OnRegistryLoadSuccess: func(r registry.Registry) (registry.Registry, error) {
			personModel, personErr := r.GetModel("Person")
			if personErr != nil {
				return r, personErr
			}
			personModel.Fields["NewField"] = yaml.ModelField{
				Type: yaml.ModelFieldTypeBoolean,
			}
			r.SetModel("Person", personModel)
			return r, nil
		},
	}
	config := cfg.MorpheLoadRegistryConfig{
		RegistryModelsDirPath:   suite.ModelsDirPath,
		RegistryEntitiesDirPath: suite.EntitiesDirPath,
	}

	r, registryErr := registry.LoadMorpheRegistry(loadHooks, config)

	suite.NoError(registryErr)
	suite.NotNil(r)

	allModels := r.GetAllModels()
	suite.Len(allModels, 2)

	model0, modelExists0 := allModels["Person"]
	suite.True(modelExists0)
	suite.Equal(model0.Name, "Person")

	suite.Len(model0.Fields, 4)

	modelField00, fieldExists00 := model0.Fields["ID"]
	suite.True(fieldExists00)
	suite.Equal(modelField00.Type, yaml.ModelFieldTypeAutoIncrement)
	suite.Len(modelField00.Attributes, 1)
	suite.Equal(modelField00.Attributes[0], "mandatory")

	modelField01, fieldExists01 := model0.Fields["FirstName"]
	suite.True(fieldExists01)
	suite.Equal(modelField01.Type, yaml.ModelFieldTypeString)
	suite.Len(modelField01.Attributes, 0)

	modelField02, fieldExists02 := model0.Fields["LastName"]
	suite.True(fieldExists02)
	suite.Equal(modelField02.Type, yaml.ModelFieldTypeString)
	suite.Len(modelField02.Attributes, 0)

	modelField03, fieldExists03 := model0.Fields["NewField"]
	suite.True(fieldExists03)
	suite.Equal(modelField03.Type, yaml.ModelFieldTypeBoolean)
	suite.Len(modelField03.Attributes, 0)

	suite.Len(model0.Identifiers, 2)
	modelIDs00, idsExist00 := model0.Identifiers["primary"]
	suite.True(idsExist00)
	suite.ElementsMatch(modelIDs00.Fields, []string{"ID"})

	modelIDs01, idsExist01 := model0.Identifiers["name"]
	suite.True(idsExist01)
	suite.ElementsMatch(modelIDs01.Fields, []string{"FirstName", "LastName"})

	suite.Len(model0.Related, 1)

	modelRelated00, relatedExists00 := model0.Related["ContactInfo"]
	suite.True(relatedExists00)
	suite.Equal(modelRelated00.Type, "HasOne")

	model1, modelExists1 := allModels["ContactInfo"]
	suite.True(modelExists1)
	suite.Equal(model1.Name, "ContactInfo")

	suite.Len(model1.Fields, 2)

	modelField10, fieldExists10 := model1.Fields["ID"]
	suite.True(fieldExists10)
	suite.Equal(modelField10.Type, yaml.ModelFieldTypeAutoIncrement)
	suite.Len(modelField10.Attributes, 1)
	suite.Equal(modelField10.Attributes[0], "mandatory")

	modelField11, fieldExists11 := model1.Fields["Email"]
	suite.True(fieldExists11)
	suite.Equal(modelField11.Type, yaml.ModelFieldTypeString)
	suite.Len(modelField11.Attributes, 0)

	suite.Len(model1.Identifiers, 2)
	modelID10, idExists10 := model1.Identifiers["primary"]
	suite.True(idExists10)
	suite.ElementsMatch(modelID10.Fields, []string{"ID"})

	modelIDs11, idsExist11 := model1.Identifiers["email"]
	suite.True(idsExist11)
	suite.ElementsMatch(modelIDs11.Fields, []string{"Email"})

	suite.Len(model1.Related, 1)

	modelRelated10, relatedExists10 := model1.Related["Person"]
	suite.True(relatedExists10)
	suite.Equal(modelRelated10.Type, "ForOne")
}

func (suite *LoadMorpheRegistryTestSuite) TestLoadMorpheRegistry_SuccessHook_Failure() {
	loadHooks := registry.LoadMorpheRegistryHooks{
		OnRegistryLoadSuccess: func(registry registry.Registry) (registry.Registry, error) {
			return registry, fmt.Errorf("compile model success hook error")
		},
	}
	config := cfg.MorpheLoadRegistryConfig{
		RegistryModelsDirPath:   suite.ModelsDirPath,
		RegistryEntitiesDirPath: suite.EntitiesDirPath,
	}

	r, registryErr := registry.LoadMorpheRegistry(loadHooks, config)

	suite.ErrorContains(registryErr, "compile model success hook error")
	suite.Nil(r)
}
