package pkg

import (
	"fmt"
	"io/ioutil"
	"jiuhuche120/jselenium"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type Snake struct {
	suite.Suite
}

func Run(t *testing.T, s suite.TestingSuite) {
	if jselenium.Config == nil {
		t.Log("warning: configuration not initialized")
		t.Fail()
		return
	}
	suite.Run(t, s)
}

func (suite Snake) Equal(expected interface{}, actual interface{}, args ...interface{}) {
	var page *Page
	var msgAndArgs []interface{}
	if args != nil {
		page = args[0].(*Page)
		msgAndArgs = args[1:]
	}
	if !suite.Assert().Equal(expected, actual, msgAndArgs) && page != nil {
		err := suite.doSome(page)
		if err != nil {
			suite.Assert().Error(err)
		}
	}
}

func (suite Snake) NotEqual(expected interface{}, actual interface{}, args ...interface{}) {
	var page *Page
	var msgAndArgs []interface{}
	if args != nil {
		page = args[0].(*Page)
		msgAndArgs = args[1:]
	}
	if !suite.Assert().NotEqual(expected, actual, msgAndArgs) && page != nil {
		err := suite.doSome(page)
		if err != nil {
			suite.Assert().Error(err)
		}
	}
}

func (suite Snake) Nil(object interface{}, args ...interface{}) {
	var page *Page
	var msgAndArgs []interface{}
	if args != nil {
		page = args[0].(*Page)
		msgAndArgs = args[1:]
	}
	if !suite.Assert().Nil(object, msgAndArgs) && page != nil {
		err := suite.doSome(page)
		if err != nil {
			suite.Assert().Error(err)
		}
	}
}

func (suite Snake) NotNil(object interface{}, args ...interface{}) {
	var page *Page
	var msgAndArgs []interface{}
	if args != nil {
		page = args[0].(*Page)
		msgAndArgs = args[1:]
	}
	if !suite.Assert().NotNil(object, msgAndArgs) && page != nil {
		err := suite.doSome(page)
		if err != nil {
			suite.Assert().Error(err)
		}
	}
}

func (suite Snake) GreaterOrEqual(e1 interface{}, e2 interface{}, args ...interface{}) {
	var page *Page
	var msgAndArgs []interface{}
	if args != nil {
		page = args[0].(*Page)
		msgAndArgs = args[1:]
	}
	if !suite.Assert().GreaterOrEqual(e1, e2, msgAndArgs) && page != nil {
		err := suite.doSome(page)
		if err != nil {
			suite.Assert().Error(err)
		}
	}
}

func (suite Snake) LessOrEqual(e1 interface{}, e2 interface{}, args ...interface{}) {
	var page *Page
	var msgAndArgs []interface{}
	if args != nil {
		page = args[0].(*Page)
		msgAndArgs = args[1:]
	}
	if !suite.Assert().LessOrEqual(e1, e2, msgAndArgs) && page != nil {
		err := suite.doSome(page)
		if err != nil {
			suite.Assert().Error(err)
		}
	}
}

func (suite Snake) doSome(page *Page) error {
	err := suite.Screenshot(page)
	if err != nil {
		return err
	}
	return nil
}

func (suite Snake) Screenshot(page *Page) error {
	path := filepath.Join(jselenium.Config.ReportPath, fmt.Sprintf("%v", time.Now().Format("2006-01-02")))
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0777)
		if err != nil {
			return err
		}
	}
	filename := strings.Split(suite.T().Name(), "/")[1] + ".png"
	data, err := page.Screenshot()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(path, filename), data, 0777)
	if err != nil {
		return err
	}
	return nil
}
