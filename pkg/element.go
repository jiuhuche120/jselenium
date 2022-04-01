package pkg

import (
	"github.com/tebeka/selenium"
)

var _ selenium.WebElement = (*Element)(nil)

type Element struct {
	by     string
	value  string
	driver selenium.WebDriver
	index  uint64
}

func (e Element) Click() error {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return err
	}
	return elements[e.index].Click()
}

func NewElement(by, value string, driver selenium.WebDriver) Element {
	return Element{
		by:     by,
		value:  value,
		driver: driver,
		index:  0,
	}
}

func (e Element) SendKeys(keys string) error {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return err
	}
	return elements[e.index].SendKeys(keys)
}

func (e Element) Submit() error {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return err
	}
	return elements[e.index].Submit()
}

func (e Element) Clear() error {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return err
	}
	return elements[e.index].Clear()
}

func (e Element) MoveTo(xOffset, yOffset int) error {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return err
	}
	return elements[e.index].MoveTo(xOffset, yOffset)
}

func (e Element) FindElement(by, value string) (selenium.WebElement, error) {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return nil, err
	}
	return elements[e.index].FindElement(by, value)
}

func (e Element) FindElements(by, value string) ([]selenium.WebElement, error) {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return nil, err
	}
	return elements[e.index].FindElements(by, value)
}

func (e Element) TagName() (string, error) {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return "", err
	}
	return elements[e.index].TagName()
}

func (e Element) Text() (string, error) {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return "", err
	}
	return elements[e.index].Text()
}

func (e Element) IsSelected() (bool, error) {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return false, err
	}
	return elements[e.index].IsSelected()
}

func (e Element) IsEnabled() (bool, error) {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return false, err
	}
	return elements[e.index].IsEnabled()
}

func (e Element) IsDisplayed() (bool, error) {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return false, err
	}
	return elements[e.index].IsDisplayed()
}

func (e Element) GetAttribute(name string) (string, error) {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return "", err
	}
	return elements[e.index].GetAttribute(name)
}

func (e Element) Location() (*selenium.Point, error) {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return nil, err
	}
	return elements[e.index].Location()
}

func (e Element) LocationInView() (*selenium.Point, error) {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return nil, err
	}
	return elements[e.index].LocationInView()
}

func (e Element) Size() (*selenium.Size, error) {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return nil, err
	}
	return elements[e.index].Size()
}

func (e Element) CSSProperty(name string) (string, error) {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return "", err
	}
	return elements[e.index].CSSProperty(name)
}

func (e Element) Screenshot(scroll bool) ([]byte, error) {
	elements, err := e.driver.FindElements(e.by, e.value)
	if err != nil {
		return []byte{}, err
	}
	return elements[e.index].Screenshot(scroll)
}

type Elements struct {
	Element
}

func NewElements(by, value string, driver selenium.WebDriver) Elements {
	return Elements{
		Element{by: by, value: value, driver: driver},
	}
}

func (e Elements) Get(index uint64) {
	e.index = index
}
