package modularize

import "testing"

type TestStruct struct {
	Teste		string
}

func TestResources_Inject(t *testing.T) {
	var resources Resources
	testResource := TestStruct{Teste: "AAA"}
	resources.SetResource("teste", testResource)
	var testeInject TestStruct
	resources.Inject(&testeInject)
	if testeInject.Teste != "AAA" {
		t.Error("Error in inject value in struct")
	}
}

type TestStruct2 struct {
	Teste	string
}

func TestResources_InjectInvalid(t *testing.T) {
	var resources Resources
	testResource := TestStruct2{Teste: "BBB"}
	resources.SetResource("teste", testResource)
	var testeInject TestStruct
	resources.Inject(&testeInject)
	if testeInject.Teste != "" {
		t.Error("Error in inject value in struct")
	}
}

type ComplexType struct {
	ServerName		string		`inject:"server"`
	Test			*TestStruct	`inject:"test"`
}

func TestResources_Complex(t *testing.T) {
	var resources Resources
	resources.SetResource("server", "server_name")
	var complexTest ComplexType
	resources.Inject(&complexTest)
	if complexTest.ServerName != "server_name" {
		t.Error("Error in inject value in complex struct")
	}
}

func TestResources_ComplexPtr(t *testing.T) {
	var resources Resources
	resources.SetResource("test", &TestStruct{Teste: "AAA"})
	var complexTest ComplexType
	resources.Inject(&complexTest)
	if complexTest.Test.Teste != "AAA" {
		t.Error("Error in inject value in complex struct")
	}
}