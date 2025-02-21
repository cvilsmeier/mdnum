package main

import (
	"os"
	"testing"
)

func TestBuilder(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		input := ""
		want := ""
		have := Convert(input)
		if have != want {
			t.Fatalf("\nwant %q\nhave %q", want, have)
		}
	})
	t.Run("oneEmptyLine", func(t *testing.T) {
		input := "\n"
		want := "\n"
		have := Convert(input)
		if have != want {
			t.Fatalf("\nwant %q\nhave %q", want, have)
		}
	})
	t.Run("oneEmptyLineWithNewline", func(t *testing.T) {
		input := "\n\n"
		want := "\n\n"
		have := Convert(input)
		if have != want {
			t.Fatalf("\nwant %q\nhave %q", want, have)
		}
	})
	t.Run("emptyLines", func(t *testing.T) {
		input := "\n\n\n"
		want := "\n\n\n"
		have := Convert(input)
		if have != want {
			t.Fatalf("\nwant %q\nhave %q", want, have)
		}
	})
	t.Run("level1", func(t *testing.T) {
		input := "# 1. one\n" +
			"# 6. two\n" +
			"#       5.        three"
		want := "# 1. one\n" +
			"# 2. two\n" +
			"# 3. three"
		have := Convert(input)
		if have != want {
			t.Fatalf("\nwant %q\nhave %q", want, have)
		}
	})
	t.Run("level2", func(t *testing.T) {
		input := "\n" +
			"# 1. a\n" +
			"# 6. b\n" +
			"## 6.6. b1\n" +
			"## 6.7. b2\n" +
			"# 5. c\n"
		want := "\n" +
			"# 1. a\n" +
			"# 2. b\n" +
			"## 2.1. b1\n" +
			"## 2.2. b2\n" +
			"# 3. c\n"
		have := Convert(input)
		if have != want {
			t.Fatalf("\nwant %q\nhave %q", want, have)
		}
	})
	t.Run("level4", func(t *testing.T) {
		input := "" +
			"# 1. 1\n" +
			"### 2.2.2. 1_1_1\n" +
			"### 3.3.3. 1_1_2"
		want := "" +
			"# 1. 1\n" +
			"### 1.1.1. 1_1_1\n" +
			"### 1.1.2. 1_1_2"
		have := Convert(input)
		if have != want {
			t.Fatalf("\nwant %q\nhave %q", want, have)
		}
	})
	t.Run("level6", func(t *testing.T) {
		input := "" +
			"# 1. 1\n" +
			"# 1. 2\n" +
			"## 1.1. 2_1\n" +
			"## 1.1. 2_2\n" +
			"# 1. 3\n" +
			"## 1.1. 3_1\n" +
			"## 1.1. 3_2\n" +
			"#### 1.1.1.1. 3_2_1_1\n" +
			"#### 1.1.1.1. 3_2_1_2\n" +
			"# 1. 4\n" +
			"\n" +
			"\n"
		want := "" +
			"# 1. 1\n" +
			"# 2. 2\n" +
			"## 2.1. 2_1\n" +
			"## 2.2. 2_2\n" +
			"# 3. 3\n" +
			"## 3.1. 3_1\n" +
			"## 3.2. 3_2\n" +
			"#### 3.2.1.1. 3_2_1_1\n" +
			"#### 3.2.1.2. 3_2_1_2\n" +
			"# 4. 4\n" +
			"\n" +
			"\n"
		have := Convert(input)
		if have != want {
			t.Fatalf("\nwant %q\nhave %q", want, have)
		}
	})
	t.Run("small.md", func(t *testing.T) {
		input := mustReadFile("testdata/small.md")
		have := Convert(input)
		want := mustReadFile("testdata/small.golden.md")
		if have != want {
			t.Fatalf("\nwant %q\nhave %q", want, have)
		}
	})
}

func mustReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}
