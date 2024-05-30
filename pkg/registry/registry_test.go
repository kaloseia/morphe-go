package registry_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/kaloseia/morphe-go/internal/testutils"

	"github.com/kaloseia/morphe-go/pkg/registry"
	"github.com/kaloseia/morphe-go/pkg/yaml"
)

type RegistryTestSuite struct {
	suite.Suite

	TestDirPath string

	ModelsDirPath   string
	EntitiesDirPath string
}

func TestRegistryTestSuite(t *testing.T) {
	suite.Run(t, new(RegistryTestSuite))
}

func (suite *RegistryTestSuite) SetupTest() {
	suite.TestDirPath = testutils.GetTestDirPath()

	suite.ModelsDirPath = filepath.Join(suite.TestDirPath, "registry", "models")
	suite.EntitiesDirPath = filepath.Join(suite.TestDirPath, "registry", "entities")
}

func (suite *RegistryTestSuite) TearDownTest() {
	suite.TestDirPath = ""
}

func (suite *RegistryTestSuite) TestLoadModelsFromDirectory() {
	registry := registry.NewRegistry()

	modelsErr := registry.LoadModelsFromDirectory(suite.ModelsDirPath)

	suite.Nil(modelsErr)
	suite.Len(registry.Models, 4)
	suite.Len(registry.Entities, 0)

	model0, modelExists0 := registry.Models["Company"]
	suite.True(modelExists0)
	suite.Equal(model0.Name, "Company")

	suite.Len(model0.Fields, 4)

	modelField00, fieldExists00 := model0.Fields["UUID"]
	suite.True(fieldExists00)
	suite.Equal(modelField00.Type, yaml.ModelFieldTypeUUID)
	suite.Len(modelField00.Attributes, 2)
	suite.Equal(modelField00.Attributes[0], "immutable")
	suite.Equal(modelField00.Attributes[1], "mandatory")

	modelField01, fieldExists01 := model0.Fields["ID"]
	suite.True(fieldExists01)
	suite.Equal(modelField01.Type, yaml.ModelFieldTypeAutoIncrement)
	suite.Len(modelField01.Attributes, 1)
	suite.Equal(modelField01.Attributes[0], "mandatory")

	modelField02, fieldExists02 := model0.Fields["FoundedAt"]
	suite.True(fieldExists02)
	suite.Equal(modelField02.Type, yaml.ModelFieldTypeTime)
	suite.Len(modelField02.Attributes, 0)

	modelField03, fieldExists03 := model0.Fields["Name"]
	suite.True(fieldExists03)
	suite.Equal(modelField03.Type, yaml.ModelFieldTypeString)
	suite.Len(modelField03.Attributes, 0)

	suite.Len(model0.Identifiers, 2)

	modelIDs00, idsExist00 := model0.Identifiers["primary"]
	suite.True(idsExist00)
	suite.ElementsMatch(modelIDs00.Fields, []string{"ID"})

	modelIDs01, idsExist01 := model0.Identifiers["entity"]
	suite.True(idsExist01)
	suite.ElementsMatch(modelIDs01.Fields, []string{"UUID"})

	suite.Len(model0.Related, 2)

	modelRelated00, relatedExists00 := model0.Related["Address"]
	suite.True(relatedExists00)
	suite.Equal(modelRelated00.Type, "HasOne")

	modelRelated01, relatedExists01 := model0.Related["Person"]
	suite.True(relatedExists01)
	suite.Equal(modelRelated01.Type, "HasMany")

	model1, modelExists1 := registry.Models["Address"]
	suite.True(modelExists1)
	suite.Equal(model1.Name, "Address")

	suite.Len(model1.Fields, 3)

	modelField10, fieldExists10 := model1.Fields["ID"]
	suite.True(fieldExists10)
	suite.Equal(modelField10.Type, yaml.ModelFieldTypeAutoIncrement)
	suite.Len(modelField10.Attributes, 1)
	suite.Equal(modelField10.Attributes[0], "mandatory")

	modelField11, fieldExists11 := model1.Fields["Street"]
	suite.True(fieldExists11)
	suite.Equal(modelField11.Type, yaml.ModelFieldTypeString)
	suite.Len(modelField11.Attributes, 0)

	modelField12, fieldExists12 := model1.Fields["HouseNumber"]
	suite.True(fieldExists12)
	suite.Equal(modelField12.Type, yaml.ModelFieldTypeString)
	suite.Len(modelField12.Attributes, 0)

	suite.Len(model1.Identifiers, 2)
	modelIDs10, idsExist10 := model1.Identifiers["primary"]
	suite.True(idsExist10)
	suite.ElementsMatch(modelIDs10.Fields, []string{"ID"})

	modelIDs11, idsExist11 := model1.Identifiers["street"]
	suite.True(idsExist11)
	suite.ElementsMatch(modelIDs11.Fields, []string{"Street", "HouseNumber"})

	suite.Len(model1.Related, 1)

	modelRelated10, relatedExists10 := model1.Related["Company"]
	suite.True(relatedExists10)
	suite.Equal(modelRelated10.Type, "ForOne")

	model2, modelExists2 := registry.Models["ContactInfo"]
	suite.True(modelExists2)
	suite.Equal(model2.Name, "ContactInfo")

	suite.Len(model2.Fields, 3)

	modelField20, fieldExists20 := model2.Fields["ID"]
	suite.True(fieldExists20)
	suite.Equal(modelField20.Type, yaml.ModelFieldTypeAutoIncrement)
	suite.Len(modelField20.Attributes, 1)
	suite.Equal(modelField20.Attributes[0], "mandatory")

	modelField21, fieldExists21 := model2.Fields["Email"]
	suite.True(fieldExists21)
	suite.Equal(modelField21.Type, yaml.ModelFieldTypeString)
	suite.Len(modelField21.Attributes, 1)
	suite.Equal(modelField21.Attributes[0], "mandatory")

	modelField22, fieldExists22 := model2.Fields["PhoneNumber"]
	suite.True(fieldExists22)
	suite.Equal(modelField22.Type, yaml.ModelFieldTypeString)
	suite.Len(modelField22.Attributes, 0)

	suite.Len(model2.Identifiers, 1)
	modelID20, idExists20 := model2.Identifiers["primary"]
	suite.True(idExists20)
	suite.ElementsMatch(modelID20.Fields, []string{"ID"})

	suite.Len(model2.Related, 1)

	modelRelated20, relatedExists20 := model2.Related["Person"]
	suite.True(relatedExists20)
	suite.Equal(modelRelated20.Type, "ForOne")

	model3, modelExists3 := registry.Models["Person"]
	suite.True(modelExists3)
	suite.Equal(model3.Name, "Person")

	suite.Len(model3.Fields, 4)

	modelField30, fieldExists30 := model3.Fields["UUID"]
	suite.True(fieldExists30)
	suite.Equal(modelField30.Type, yaml.ModelFieldTypeUUID)
	suite.Len(modelField30.Attributes, 2)
	suite.Equal(modelField30.Attributes[0], "immutable")
	suite.Equal(modelField30.Attributes[1], "mandatory")

	modelField31, fieldExists31 := model3.Fields["ID"]
	suite.True(fieldExists31)
	suite.Equal(modelField31.Type, yaml.ModelFieldTypeAutoIncrement)
	suite.Len(modelField31.Attributes, 1)
	suite.Equal(modelField31.Attributes[0], "mandatory")

	modelField32, fieldExists32 := model3.Fields["FirstName"]
	suite.True(fieldExists32)
	suite.Equal(modelField32.Type, yaml.ModelFieldTypeString)
	suite.Len(modelField32.Attributes, 0)

	modelField33, fieldExists33 := model3.Fields["LastName"]
	suite.True(fieldExists33)
	suite.Equal(modelField33.Type, yaml.ModelFieldTypeString)
	suite.Len(modelField33.Attributes, 0)

	suite.Len(model3.Identifiers, 3)
	modelIDs30, idsExist10 := model3.Identifiers["primary"]
	suite.True(idsExist10)
	suite.ElementsMatch(modelIDs30.Fields, []string{"ID"})

	modelIDs31, idsExist11 := model3.Identifiers["entity"]
	suite.True(idsExist11)
	suite.ElementsMatch(modelIDs31.Fields, []string{"UUID"})

	modelIDs32, idsExist12 := model3.Identifiers["name"]
	suite.True(idsExist12)
	suite.ElementsMatch(modelIDs32.Fields, []string{"FirstName", "LastName"})

	suite.Len(model3.Related, 2)

	modelRelated30, relatedExists30 := model3.Related["Company"]
	suite.True(relatedExists30)
	suite.Equal(modelRelated30.Type, "ForOne")

	modelRelated31, relatedExists31 := model3.Related["ContactInfo"]
	suite.True(relatedExists31)
	suite.Equal(modelRelated31.Type, "HasOne")
}

func (suite *RegistryTestSuite) TestLoadModelsFromDirectory_InvalidDirPath() {
	registry := registry.NewRegistry()

	modelsErr := registry.LoadModelsFromDirectory("####INVALID/DIR/PATH####")

	suite.NotNil(modelsErr)
	modelsErrMsg := modelsErr.Error()
	suite.Contains(modelsErrMsg, "error reading directory")
	suite.Contains(modelsErrMsg, "####INVALID/DIR/PATH####")
	suite.Len(registry.Models, 0)
	suite.Len(registry.Entities, 0)
}

func (suite *RegistryTestSuite) TestLoadModelsFromDirectory_ConflictingName() {
	registry := registry.NewRegistry()

	registry.Models["Company"] = yaml.Model{Name: "Company"}

	modelsErr := registry.LoadModelsFromDirectory(suite.ModelsDirPath)

	suite.NotNil(modelsErr)
	modelsErrMsg := modelsErr.Error()
	suite.Contains(modelsErrMsg, "model name 'Company' already exists in registry")

	conflictPath := filepath.Join(suite.ModelsDirPath, "company.mod")
	suite.Contains(modelsErrMsg, conflictPath)
}

func (suite *RegistryTestSuite) TestLoadEntitiesFromDirectory() {
	registry := registry.NewRegistry()

	entitiesErr := registry.LoadEntitiesFromDirectory(suite.EntitiesDirPath)

	suite.Nil(entitiesErr)
	suite.Len(registry.Models, 0)
	suite.Len(registry.Entities, 2)

	entity0, entityExists0 := registry.Entities["Company"]
	suite.True(entityExists0)
	suite.Equal(entity0.Name, "Company")

	suite.Len(entity0.Fields, 6)

	entityField00, fieldExists00 := entity0.Fields["UUID"]
	suite.True(fieldExists00)
	suite.Equal(entityField00.Type, yaml.ModelFieldPath("Company.UUID"))
	suite.Len(entityField00.Attributes, 2)
	suite.Equal(entityField00.Attributes[0], "immutable")
	suite.Equal(entityField00.Attributes[1], "mandatory")

	entityField01, fieldExists01 := entity0.Fields["ID"]
	suite.True(fieldExists01)
	suite.Equal(entityField01.Type, yaml.ModelFieldPath("Company.ID"))
	suite.Len(entityField01.Attributes, 0)

	entityField02, fieldExists02 := entity0.Fields["FoundedAt"]
	suite.True(fieldExists02)
	suite.Equal(entityField02.Type, yaml.ModelFieldPath("Company.FoundedAt"))
	suite.Len(entityField02.Attributes, 0)

	entityField03, fieldExists03 := entity0.Fields["Name"]
	suite.True(fieldExists03)
	suite.Equal(entityField03.Type, yaml.ModelFieldPath("Company.Name"))
	suite.Len(entityField03.Attributes, 0)

	entityField04, fieldExists04 := entity0.Fields["ZipCode"]
	suite.True(fieldExists04)
	suite.Equal(entityField04.Type, yaml.ModelFieldPath("Company.Address.ZipCode"))
	suite.Len(entityField04.Attributes, 0)

	entityField05, fieldExists05 := entity0.Fields["City"]
	suite.True(fieldExists05)
	suite.Equal(entityField05.Type, yaml.ModelFieldPath("Company.Address.City"))
	suite.Len(entityField05.Attributes, 0)

	suite.Len(entity0.Related, 1)

	entityRelated00, relatedExists00 := entity0.Related["Person"]
	suite.True(relatedExists00)
	suite.Equal(entityRelated00.Type, "HasMany")

	entity1, entityExists1 := registry.Entities["Person"]
	suite.True(entityExists1)
	suite.Equal(entity1.Name, "Person")

	suite.Len(entity1.Fields, 5)

	entityField10, fieldExists10 := entity1.Fields["UUID"]
	suite.True(fieldExists10)
	suite.Equal(entityField10.Type, yaml.ModelFieldPath("Person.UUID"))
	suite.Len(entityField10.Attributes, 2)
	suite.Equal(entityField10.Attributes[0], "immutable")
	suite.Equal(entityField10.Attributes[1], "mandatory")

	entityField11, fieldExists11 := entity1.Fields["ID"]
	suite.True(fieldExists11)
	suite.Equal(entityField11.Type, yaml.ModelFieldPath("Person.ID"))
	suite.Len(entityField11.Attributes, 0)

	entityField12, fieldExists12 := entity1.Fields["FirstName"]
	suite.True(fieldExists12)
	suite.Equal(entityField12.Type, yaml.ModelFieldPath("Person.FirstName"))
	suite.Len(entityField12.Attributes, 0)

	entityField13, fieldExists13 := entity1.Fields["LastName"]
	suite.True(fieldExists13)
	suite.Equal(entityField13.Type, yaml.ModelFieldPath("Person.LastName"))
	suite.Len(entityField13.Attributes, 0)

	entityField14, fieldExists14 := entity1.Fields["Email"]
	suite.True(fieldExists14)
	suite.Equal(entityField14.Type, yaml.ModelFieldPath("Person.ContactInfo.Email"))
	suite.Len(entityField14.Attributes, 0)

	suite.Len(entity1.Related, 1)

	entityRelated10, relatedExists10 := entity1.Related["Company"]
	suite.True(relatedExists10)
	suite.Equal(entityRelated10.Type, "ForOne")
}

func (suite *RegistryTestSuite) TestLoadEntitiesFromDirectory_InvalidDirPath() {
	registry := registry.NewRegistry()

	entitiesErr := registry.LoadEntitiesFromDirectory("####INVALID/DIR/PATH####")

	suite.NotNil(entitiesErr)
	entitiesErrMsg := entitiesErr.Error()
	suite.Contains(entitiesErrMsg, "error reading directory")
	suite.Contains(entitiesErrMsg, "####INVALID/DIR/PATH####")
	suite.Len(registry.Models, 0)
	suite.Len(registry.Entities, 0)
}

func (suite *RegistryTestSuite) TestLoadEntitiesFromDirectory_ConflictingName() {
	registry := registry.NewRegistry()

	registry.Entities["Company"] = yaml.Entity{Name: "Company"}

	entitiesErr := registry.LoadEntitiesFromDirectory(suite.EntitiesDirPath)

	suite.NotNil(entitiesErr)
	entitiesErrMsg := entitiesErr.Error()
	suite.Contains(entitiesErrMsg, "entity name 'Company' already exists in registry")

	conflictPath := filepath.Join(suite.EntitiesDirPath, "company.ent")
	suite.Contains(entitiesErrMsg, conflictPath)
}

func (suite *RegistryTestSuite) TestDeepClone() {
	registry := registry.NewRegistry()

	modelsErr := registry.LoadModelsFromDirectory(suite.ModelsDirPath)
	suite.Nil(modelsErr)

	entitiesErr := registry.LoadEntitiesFromDirectory(suite.EntitiesDirPath)
	suite.Nil(entitiesErr)

	registryClone := registry.DeepClone()

	suite.NotSame(registry, registryClone)
	suite.Equal(registry, registryClone)
}

func (suite *RegistryTestSuite) TestDeepClone_Empty() {
	registry := registry.NewRegistry()

	registryClone := registry.DeepClone()

	suite.NotSame(registry, registryClone)
	suite.Equal(registry, registryClone)
}
