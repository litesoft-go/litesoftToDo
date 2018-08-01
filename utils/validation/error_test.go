package validation

import (
	"testing"
	"github.com/pkg/errors"
)

func TestNew_Wrapping_IsValidation(t *testing.T) {
	zError := NewErrorf("T%s", "est")
	if !IsValidation(zError) || (zError.Error() != "Test") {
		t.Errorf("NewError(...) either did not report as Validation OR output Issue: '%s'", zError.Error())
	}

	err := errors.New("DC")
	if IsValidation(err) || (err.Error() != "DC") {
		t.Errorf("errors.New(...) either reported as Validation OR output Issue: '%s'", err.Error())
	}

	err = zError.WithCause(err)
	if !IsValidation(err) || (err.Error() != "Test: DC") {
		t.Errorf("NewError(...).WithCause(...) either did not report as Validation OR output Issue: '%s'", err.Error())
	}

	err = errors.Wrap(err, "Wrapped")
	if !IsValidation(err) || (err.Error() != "Wrapped: Test: DC") {
		t.Errorf("Wrapped( NewError(...).WithCause(...) ) either did not report as Validation OR output Issue: '%s'", err.Error())
	}
}
