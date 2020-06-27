package formatter

import "testing"

func TestCsv_Parse(t *testing.T) {
	list := [][]string{
		{"11111", "Here is the first line"},
	}

	fmt := Csv{}
	result := fmt.Parse(list)
	expected := "11111,Here is the first line\n"

	if result != expected {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestCsv_Parse_withDoubleQuotes(t *testing.T) {
	list := [][]string{
		{"33333", `Testing with "double quotes"`},
	}

	fmt := Csv{}
	result := fmt.Parse(list)
	expected := "33333,Testing with \"\"double quotes\"\"\n"

	if result != expected {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestCsv_Parse_withCommas(t *testing.T) {
	list := [][]string{
		{"44444", `What if we have commas, like this?`},
	}

	fmt := Csv{}
	result := fmt.Parse(list)
	expected := "44444,\"What if we have commas, like this?\"\n"

	if result != expected {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestCsv_ToArray(t *testing.T) {
	raw := "11111,Hello first line"
	fmt := Csv{}

	expected := [][]string{
		{"11111", "Hello first line"},
	}
	result := fmt.ToArray(raw)

	if !equals(result, expected) {
		t.Errorf("\nExpected:\n%s\nGot:\n%s\n", expected, result)
	}
}

func TestCsv_ToArray_withDoubleQuotes(t *testing.T) {
	raw := `11111,Hello with ""double quotes""`
	fmt := Csv{}

	expected := [][]string{
		{"11111", "Hello with \"double quotes\""},
	}
	result := fmt.ToArray(raw)

	if !equals(result, expected) {
		t.Errorf("\nExpected:\n%s\nGot:\n%s\n", expected, result)
	}
}

func TestCsv_ToArray_withCommas(t *testing.T) {
	raw := `11111,Hello with ""double, quotes""`
	fmt := Csv{}

	expected := [][]string{
		{"11111", "Hello with \"double, quotes\""},
	}
	result := fmt.ToArray(raw)

	if !equals(result, expected) {
		t.Errorf("\nExpected:\n%s\nGot:\n%s\n", expected, result)
	}
}

func equals(a [][]string, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}
