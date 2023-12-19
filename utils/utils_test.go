package utils

import (
	"bytes"
	"errors"
	"testing"
)

func TestDataPrompt(t *testing.T) {
	testCases := []struct {
		desc       string
		prompt     string
		byteString string
		want       string
	}{
		{
			desc:       "handles-input-data",
			prompt:     "breakfast food",
			byteString: "pancakes\n",
			want:       "pancakes",
		},
		{
			desc:       "handles-blank-input",
			prompt:     "name",
			byteString: "\n",
			want:       "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			reader := bytes.NewReader([]byte(tC.byteString))
			data, err := dataPrompt(reader, tC.prompt)
			if err != nil {
				t.Errorf("unexpected input error: %s", err)
			}

			if data != tC.want {
				t.Errorf("expected (%s), got (%s)", tC.want, data)
			}
		})
	}
}

func TestValidateInputData(t *testing.T) {
	testCases := []struct {
		desc       string
		prompt     string
		byteString string
		want       string
		attempt    int
		expectErr  bool
	}{
		{
			desc:       "valid-input-data",
			prompt:     "name",
			byteString: "john\n",
			want:       "john",
			attempt:    1,
		},
		{
			desc:       "invalid-input-data",
			prompt:     "favorite color",
			byteString: "\n",
			want:       "",
			attempt:    1,
			expectErr:  true,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			reader := bytes.NewReader([]byte(tC.byteString))
			input, err := ValidateInputData(reader, tC.prompt, tC.attempt)
			if err != nil {
				switch {
				case tC.expectErr:
					if !errors.Is(err, ErrExceedsAttempts) {
						t.Errorf("expected error (%s), got (%s)", ErrExceedsAttempts, err)
					}
				default:
					t.Errorf("unexpected error validating input: %s", err)
				}
			}

			if input != tC.want {
				t.Errorf("expected (%s), got (%s)", tC.want, input)
			}
		})
	}
}

func TestValidateInputDataChange(t *testing.T) {
	testCases := []struct {
		desc         string
		prompt       string
		originalData string
		byteString   string
		want         string
	}{
		{
			desc:         "valid-changed-data",
			prompt:       "name",
			originalData: "james",
			byteString:   "john\n",
			want:         "john",
		},
		{
			desc:         "data-unchanged",
			prompt:       "do I need coffee",
			originalData: "yes",
			byteString:   "\n",
			want:         "yes",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			reader := bytes.NewReader([]byte(tC.byteString))
			data, err := ValidateInputDataChange(reader, tC.prompt, tC.originalData)
			if err != nil {
				t.Errorf("unexpected error validating updated input: %s", err)
			}

			if data != tC.want {
				t.Errorf("expected (%s), got (%s)", tC.want, data)
			}
		})
	}
}
