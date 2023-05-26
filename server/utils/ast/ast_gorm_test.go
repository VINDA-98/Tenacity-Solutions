package ast

import (
	"github.com/VINDA-98/Tenacity-Solutions/server/global"
	"github.com/VINDA-98/Tenacity-Solutions/server/model/example"
	"testing"
)

const A = 123

func TestAddRegisterTablesAst(t *testing.T) {
	AddRegisterTablesAst("D:\\gin-vue-admin\\server\\utils\\ast_test.go", "Register", "test", "testDB", "testModel")
}

func Register() {
	test := global.GetGlobalDBByDBName("test")
	test.AutoMigrate(example.ExaFile{})
}
