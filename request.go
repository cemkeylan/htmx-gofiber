package htmx

import (
	"github.com/gofiber/fiber/v3"
	"github.com/samber/lo"
)

func hasHeaderTrue(header []string) bool {
	return lo.ContainsBy(header, func(item string) bool {
		return item == "true"
	})
}

func getHeaderValue(header []string) (string, bool) {
	return lo.First(header)
}

// IsHTMX returns true if the given request
// was made by HTMX.
//
// This can be used to add special logic for HTMX requests.
//
// Checks if header 'HX-Request' is 'true'.
func IsHTMX(c fiber.Ctx) bool {
	return hasHeaderTrue(c.Req().GetHeaders()[HeaderRequest])
}

// IsBoosted returns true if the given request
// was made via an element using 'hx-boost'.
//
// This can be used to add special logic for boosted requests.
//
// Checks if header 'HX-Boosted' is 'true'.
//
// For more info, see https://htmx.org/attributes/hx-boost/
func IsBoosted(c fiber.Ctx) bool {
	return hasHeaderTrue(c.Req().GetHeaders()[HeaderBoosted])
}

// IsHistoryRestoreRequest returns true if the given request
// is for history restoration after a miss in the local history cache.
//
// Checks if header 'HX-History-Restore-Request' is 'true'.
func IsHistoryRestoreRequest(c fiber.Ctx) bool {
	return hasHeaderTrue(c.Req().GetHeaders()[HeaderHistoryRestoreRequest])
}

// GetCurrentURL returns the current URL that HTMX made this request from.
//
// Returns false if header 'HX-Current-URL' does not exist.
func GetCurrentURL(c fiber.Ctx) (string, bool) {
	return lo.First(c.Req().GetHeaders()[HeaderCurrentURL])
}

// GetPrompt returns the user response to an hx-prompt from a given request.
//
// Returns false if header 'HX-Prompt' does not exist.
//
// For more info, see https://htmx.org/attributes/hx-prompt/
func GetPrompt(c fiber.Ctx) (string, bool) {
	return lo.First(c.Req().GetHeaders()[HeaderPrompt])
}

// GetTarget returns the ID of the target element if it exists from a given request.
//
// Returns false if header 'HX-Target' does not exist.
//
// For more info, see https://htmx.org/attributes/hx-target/
func GetTarget(c fiber.Ctx) (string, bool) {
	return lo.First(c.Req().GetHeaders()[HeaderTarget])
}

// GetTriggerName returns the 'name' of the triggered element if it exists from a given request.
//
// Returns false if header 'HX-Trigger-Name' does not exist.
//
// For more info, see https://htmx.org/attributes/hx-trigger/
func GetTriggerName(c fiber.Ctx) (string, bool) {
	return lo.First(c.Req().GetHeaders()[HeaderTriggerName])
}

// GetTrigger returns the ID of the triggered element if it exists from a given request.
//
// Returns false if header 'HX-Trigger' does not exist.
//
// For more info, see https://htmx.org/attributes/hx-trigger/
func GetTrigger(c fiber.Ctx) (string, bool) {
	return lo.First(c.Req().GetHeaders()[HeaderTrigger])
}
