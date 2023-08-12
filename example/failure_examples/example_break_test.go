package failure_examples

import (
	"github.com/Nikita-Filonov/allure-go"
	"github.com/pkg/errors"
	"testing"
)

func TestBreakAllure(t *testing.T) {
	allure.Test(t, allure.Description("Test with Allure error in it"), allure.Action(func() {
		allure.Step(allure.Description("Step 1"), allure.Action(func() {}))
		allure.Step(allure.Description("Step 2"), allure.Action(func() {
			allure.Break(errors.New("Error message"))
		}))
	}))
}
