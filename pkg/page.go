package pkg

import (
	"fmt"
	"jiuhuche120/jselenium"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/log"
)

var _ selenium.WebDriver = (*Page)(nil)

type Page struct {
	Driver selenium.WebDriver
}

func NewPage() Page {
	_, err := selenium.NewChromeDriverService(jselenium.Config.Driver, jselenium.Config.Port, []selenium.ServiceOption{}...)
	if err != nil {
		panic(err)
	}
	var caps selenium.Capabilities
	if jselenium.Config.DriverType == "chrome" {
		caps = map[string]interface{}{"browserName": "chrome"}
	} else {
		caps = map[string]interface{}{"browserName": "firefox"}
	}
	driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", jselenium.Config.Port))
	if err != nil {
		panic(err)
	}
	return Page{
		Driver: driver,
	}
}

func (p Page) Status() (*selenium.Status, error) {
	return p.Driver.Status()
}

func (p Page) NewSession() (string, error) {
	return p.Driver.NewSession()
}

func (p Page) SessionId() string {
	return p.Driver.SessionId()
}

func (p Page) SessionID() string {
	return p.Driver.SessionID()
}

func (p Page) SwitchSession(sessionID string) error {
	return p.Driver.SwitchSession(sessionID)
}

func (p Page) Capabilities() (selenium.Capabilities, error) {
	return p.Driver.Capabilities()
}

func (p Page) SetAsyncScriptTimeout(timeout time.Duration) error {
	return p.Driver.SetAsyncScriptTimeout(timeout)
}

func (p Page) SetImplicitWaitTimeout(timeout time.Duration) error {
	return p.Driver.SetImplicitWaitTimeout(timeout)
}

func (p Page) SetPageLoadTimeout(timeout time.Duration) error {
	return p.Driver.SetPageLoadTimeout(timeout)
}

func (p Page) Quit() error {
	return p.Driver.Quit()
}

func (p Page) CurrentWindowHandle() (string, error) {
	return p.Driver.CurrentWindowHandle()
}

func (p Page) WindowHandles() ([]string, error) {
	return p.Driver.WindowHandles()
}

func (p Page) CurrentURL() (string, error) {
	return p.Driver.CurrentURL()
}

func (p Page) Title() (string, error) {
	return p.Driver.Title()
}

func (p Page) PageSource() (string, error) {
	return p.Driver.PageSource()
}

func (p Page) Close() error {
	return p.Driver.Close()
}

func (p Page) SwitchFrame(frame interface{}) error {
	return p.Driver.SwitchFrame(frame)
}

func (p Page) SwitchWindow(name string) error {
	return p.Driver.SwitchWindow(name)
}

func (p Page) CloseWindow(name string) error {
	return p.Driver.CloseWindow(name)
}

func (p Page) MaximizeWindow(name string) error {
	return p.Driver.MaximizeWindow(name)
}

func (p Page) ResizeWindow(name string, width, height int) error {
	return p.Driver.ResizeWindow(name, width, height)
}

func (p Page) Get(url string) error {
	return p.Driver.Get(url)
}

func (p Page) Forward() error {
	return p.Driver.Forward()
}

func (p Page) Back() error {
	return p.Driver.Back()
}

func (p Page) Refresh() error {
	return p.Driver.Back()
}

func (p Page) FindElement(by, value string) (selenium.WebElement, error) {
	return p.Driver.FindElement(by, value)
}

func (p Page) FindElements(by, value string) ([]selenium.WebElement, error) {
	return p.Driver.FindElements(by, value)
}

func (p Page) ActiveElement() (selenium.WebElement, error) {
	return p.Driver.ActiveElement()
}

func (p Page) DecodeElement(bytes []byte) (selenium.WebElement, error) {
	return p.Driver.DecodeElement(bytes)
}

func (p Page) DecodeElements(bytes []byte) ([]selenium.WebElement, error) {
	return p.Driver.DecodeElements(bytes)
}

func (p Page) GetCookies() ([]selenium.Cookie, error) {
	return p.Driver.GetCookies()
}

func (p Page) GetCookie(name string) (selenium.Cookie, error) {
	return p.Driver.GetCookie(name)
}

func (p Page) AddCookie(cookie *selenium.Cookie) error {
	return p.Driver.AddCookie(cookie)
}

func (p Page) DeleteAllCookies() error {
	return p.Driver.DeleteAllCookies()
}

func (p Page) DeleteCookie(name string) error {
	return p.Driver.DeleteCookie(name)
}

func (p Page) Click(button int) error {
	return p.Driver.Click(button)
}

func (p Page) DoubleClick() error {
	return p.Driver.DoubleClick()
}

func (p Page) ButtonDown() error {
	return p.Driver.ButtonDown()
}

func (p Page) ButtonUp() error {
	return p.Driver.ButtonUp()
}

func (p Page) SendModifier(modifier string, isDown bool) error {
	return p.Driver.SendModifier(modifier, isDown)
}

func (p Page) KeyDown(keys string) error {
	return p.Driver.KeyDown(keys)
}

func (p Page) KeyUp(keys string) error {
	return p.Driver.KeyUp(keys)
}

func (p Page) Screenshot() ([]byte, error) {
	return p.Driver.Screenshot()
}

func (p Page) Log(typ log.Type) ([]log.Message, error) {
	return p.Driver.Log(typ)
}

func (p Page) DismissAlert() error {
	return p.Driver.DismissAlert()
}

func (p Page) AcceptAlert() error {
	return p.Driver.AcceptAlert()
}

func (p Page) AlertText() (string, error) {
	return p.Driver.AlertText()
}

func (p Page) SetAlertText(text string) error {
	return p.Driver.SetAlertText(text)
}

func (p Page) ExecuteScript(script string, args []interface{}) (interface{}, error) {
	return p.Driver.ExecuteScript(script, args)
}

func (p Page) ExecuteScriptAsync(script string, args []interface{}) (interface{}, error) {
	return p.Driver.ExecuteScriptAsync(script, args)
}

func (p Page) ExecuteScriptRaw(script string, args []interface{}) ([]byte, error) {
	return p.Driver.ExecuteScriptRaw(script, args)
}

func (p Page) ExecuteScriptAsyncRaw(script string, args []interface{}) ([]byte, error) {
	return p.Driver.ExecuteScriptAsyncRaw(script, args)
}

func (p Page) WaitWithTimeoutAndInterval(condition selenium.Condition, timeout, interval time.Duration) error {
	return p.Driver.WaitWithTimeoutAndInterval(condition, timeout, interval)
}

func (p Page) WaitWithTimeout(condition selenium.Condition, timeout time.Duration) error {
	return p.Driver.WaitWithTimeout(condition, timeout)
}

func (p Page) Wait(condition selenium.Condition) error {
	return p.Driver.Wait(condition)
}
