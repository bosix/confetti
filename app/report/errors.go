package report

import (
	"github.com/lanvard/errors"
	"github.com/lanvard/syslog/log_level"
	net "net/http"
)

/**
 * Below you will find a list of all user errors. By default these errors
 * are not logged by config `config.Errors.NoLogging`. These errors are
 * usually errors with http status 499 and lower.
 */
var UserError = errors.New("").Status(net.StatusBadRequest).Level(log_level.INFO)
var ValidationError = UserError.Status(net.StatusNotFound)
var NotFoundError = UserError.Status(net.StatusNotFound)
var PageNotFoundError = UserError.Wrap("page not found")

/**
 * This list contains errors that indicate that the system is not working
 * properly. The message is not displayed to the user on a production environment,
 * but will be logged (if MinLevel of the logger allows it).
 */
var SystemError = errors.New("").Status(net.StatusInternalServerError).Level(log_level.EMERGENCY)
