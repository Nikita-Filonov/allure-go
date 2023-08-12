package teardown

import (
	"errors"
	"github.com/Nikita-Filonov/allure-go"
	"testing"
)

func TestDeferredAfterFailNow(t *testing.T) {
	allure.Test(t,
		allure.Description("Break now"),
		allure.Action(func() {
			//panic("panic at the before statement! (disco)")
			allure.FailNow(errors.New("break now statement"))

			allure.Step(allure.Description("Step after break"),
				allure.Action(func() {}))
		}))
	defer func() {
		allure.AfterTest(t, allure.Action(func() {
			allure.Step(allure.Description("Running after a panic"))
		}))
	}()
}
