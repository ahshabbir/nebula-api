package main

import (
	"fmt"
	"log"
	"syscall"
	"time"

	"github.com/tebeka/selenium"
	"golang.org/x/term"
)

func getDriver(port int, driver string) selenium.WebDriver {
	if driver != "firefox" && driver != "chrome" {
		log.Fatalf("invalid driver name: '%s'", driver)
	}

	wd, err := selenium.NewRemote(
		selenium.Capabilities{"browserName": driver},
		fmt.Sprintf("http://localhost:%d", port),
	)
	if err != nil {
		log.Fatalf("could not open webdriver at port '%d': %v", port, err)
	}

	return wd
}

func getUTDCreds() (netid, password string, err error) {
	fmt.Print("Enter UTD NetID: ")
	_, err = fmt.Scanln(&netid)
	if err != nil {
		return "", "", fmt.Errorf("failed to get netid: %v", err)
	}

	fmt.Print("Enter UTD Password: ")
	bytepw, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", fmt.Errorf("failed to get password: %v", err)
	}
	password = string(bytepw)

	return netid, password, nil

}

func loadLoginPage(wd selenium.WebDriver) error {
	err := wd.Get("https://coursebook.utdallas.edu/")
	if err != nil {
		return fmt.Errorf("failed to load coursebook login: %v", err)
	}

	time.Sleep(3 * time.Second)

	loginLink, err := wd.FindElement("id", "pauth_link")
	if err != nil {
		return fmt.Errorf("failed to find pauth_link object: %v", err)
	}

	if err := loginLink.Click(); err != nil {
		return fmt.Errorf("failed to click on login link: %v", err)
	}

	time.Sleep(3 * time.Second)

	return nil
}

func loadCoursebook(wd selenium.WebDriver) error {
	netid, password, err := getUTDCreds()
	if err != nil {
		return fmt.Errorf("failed to get UTD creds from user: %v", err)
	}

	if err = loadLoginPage(wd); err != nil {
		return fmt.Errorf("failed to load login page: %v", err)
	}

	nidInput, err := wd.FindElement("id", "netid")
	if err != nil {
		return fmt.Errorf("failed to get netid object: %v", err)
	}

	err = nidInput.SendKeys(netid)
	if err != nil {
		return fmt.Errorf("failed to send netid to netid input: %v", err)
	}

	passwordInput, err := wd.FindElement("id", "password")
	if err != nil {
		return fmt.Errorf("failed to get netid object: %v", err)
	}

	err = passwordInput.SendKeys(password)
	if err != nil {
		return fmt.Errorf("failed to send password to password input: %v", err)
	}

	submitBtn, err := wd.FindElement("id", "login-button")
	if err != nil {
		return fmt.Errorf("failed to load login button: %v", err)
	}

	err = submitBtn.Click()
	if err != nil {
		return fmt.Errorf("failed to submit login: %v", err)
	}

	return nil
}

func main() {
	wd := getDriver(4444, "firefox")
	defer wd.Close()

	if err := loadCoursebook(wd); err != nil {
		log.Fatalf("failed to load coursebook: %v", err)
	}

	time.Sleep(10 * time.Second)
}
