package errormessage

import (
	"errors"
	"fmt"
	regexpsyntax "regexp/syntax"

	"github.com//config"
	"github.com//config/configStructs"
	"github.com/danielpickens/tesseract/misc"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
)

// formatError wraps error with a detailed message that is meant for the user.
// While the errors are meant to be displayed, they are not meant to be exported as classes outsite of CLI.
func FormatError(err error) error {
	var errorNew error
	if k8serrors.IsForbidden(err) {
		errorNew = fmt.Errorf("insufficient permissions: %w. "+
			"supply the required permission or control %s's access to namespaces by setting %s "+
			"in the config file or setting the targeted namespace with --%s %s=<NAMEPSACE>",
			err,
			misc.Software,
			configStructs.ReleaseNamespaceLabel,
			config.SetCommandName,
			configStructs.ReleaseNamespaceLabel)
	} else if syntaxError, isSyntaxError := asRegexSyntaxError(err); isSyntaxError {
		errorNew = fmt.Errorf("regex %s is invalid: %w", syntaxError.Expr, err)
	} else {
		errorNew = err
	}

	return errorNew
}

func asRegexSyntaxError(err error) (*regexpsyntax.Error, bool) {
	var syntaxError *regexpsyntax.Error
	return syntaxError, errors.As(err, &syntaxError)
}

func asK8sError(err error) (*k8serrors.StatusError, bool) {
	var statusError *k8serrors.StatusError
	return statusError, errors.As(err, &statusError)
}