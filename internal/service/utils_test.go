package service

import (
	"bytes"
	"testing"

	"github.com/LaQuannT/go-inventory/internal/model"
)

func TestPromptForData(t *testing.T) {
	testCases := []struct {
		desc       string
		prompt     string
		byteString string
		want       string
	}{
		{
			desc:       "returns expected input from prompt",
			prompt:     "breakfast food",
			byteString: "pancakes\n",
			want:       "pancakes",
		},
		{
			desc:       "returns lower case input from uppercase input",
			prompt:     "coffee",
			byteString: "FLAT WHITE\n",
			want:       "flat white",
		},
		{
			desc:       "returns expected input trimming extra space",
			prompt:     "color",
			byteString: "red  \n",
			want:       "red",
		},
		{
			desc:       "returns empty string when no input is given",
			prompt:     "name",
			byteString: "\n",
			want:       "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := bytes.NewReader([]byte(tc.byteString))
			data, err := promptForData(reader, tc.prompt)
			if err != nil {
				t.Fatalf("unexpected input error: %s", err)
			}

			if data != tc.want {
				t.Errorf("expected %q, got %q", tc.want, data)
			}
		})
	}
}

func TestValidateInput(t *testing.T) {
	// Name prompt case covers Brand and sku
	testCases := []struct {
		desc       string
		prompt     string
		byteString string
		want       string
		expectErr  bool
	}{
		{
			desc:       "returns a name",
			prompt:     "Name",
			byteString: "john\n",
			want:       "john",
		},
		{
			desc:       "returns a number amount as a string",
			prompt:     "Amount",
			byteString: "21\n",
			want:       "21",
		},
		{
			desc:       "returns error if name input is empty",
			prompt:     "Name",
			byteString: "\n",
			expectErr:  true,
		},
		{
			desc:       "returns error if amount input is empty",
			prompt:     "Amount",
			byteString: "\n",
			expectErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			reader := bytes.NewReader([]byte(tc.byteString))
			input, err := validateInput(reader, tc.prompt)
			if err != nil && !tc.expectErr {
				t.Fatalf("unexpected error validating input: %s", err)
			}

			if tc.expectErr {
				if err == nil {
					t.Error("expected error to be thrown, no error thrown")
				}
			} else {
				if input != tc.want {
					t.Errorf("expected %q, got %q", tc.want, input)
				}
			}
		})
	}
}

func TestDisplayData(t *testing.T) {
	testCases := []struct {
		desc        string
		item        *model.Item
		expectedStr string
	}{
		{
			desc: "return expected display string",
			item: &model.Item{
				Name:     "iphone 12 pro",
				Brand:    "apple",
				Sku:      "aap12p21",
				Category: "mobile",
				Location: "warehouse",
				Amount:   2,
			},
			expectedStr: "[AAP12P21] Name: Iphone 12 Pro | Brand: Apple | Category: Mobile | location: Warehouse | Stock: 2\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			output := new(bytes.Buffer)
			displayData(tc.item, output)

			if output.String() != tc.expectedStr {
				t.Errorf("expected output %q, got %q", tc.expectedStr, output.String())
			}
		})
	}
}
