package main

import (
	"os"
	"testing"
)

func TestSplitLines(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		lines := splitLines("")
		if len(lines) != 0 {
			t.Fatalf("wrong %v", len(lines))
		}
	})
	t.Run("oneLine", func(t *testing.T) {
		lines := splitLines("\n")
		if len(lines) != 1 {
			t.Fatalf("wrong %v", len(lines))
		}
		if lines[0] != "\n" {
			t.Fatalf("wrong %v", lines[0])
		}
	})
	t.Run("twoLines", func(t *testing.T) {
		lines := splitLines("a\nb")
		if len(lines) != 2 {
			t.Fatalf("wrong %v", len(lines))
		}
		if lines[0] != "a\n" {
			t.Fatalf("wrong %v", lines[0])
		}
		if lines[1] != "b" {
			t.Fatalf("wrong %v", lines[1])
		}
	})
	t.Run("endsWithNewline", func(t *testing.T) {
		lines := splitLines("a\nb\n")
		if len(lines) != 2 {
			t.Fatalf("wrong %v", len(lines))
		}
		if lines[0] != "a\n" {
			t.Fatalf("wrong %v", lines[0])
		}
		if lines[1] != "b\n" {
			t.Fatalf("wrong %v", lines[1])
		}
	})
}

func TestBuilder(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		b := NewBuilder()
		have := b.Finish()
		want := ""
		if have != want {
			t.Fatalf("\nwant %q\nhave %q", want, have)
		}
	})
	t.Run("oneEmptyLine", func(t *testing.T) {
		b := NewBuilder()
		b.Add("")
		have := b.Finish()
		want := ""
		if have != want {
			t.Fatalf("\nwant %q\nhave %q", want, have)
		}
	})
	t.Run("oneEmptyLineWithNewline", func(t *testing.T) {
		b := NewBuilder()
		b.Add("\n")
		have := b.Finish()
		want := "\n"
		if have != want {
			t.Fatalf("\nwant %q\nhave %q", want, have)
		}
	})
	t.Run("emptyLines", func(t *testing.T) {
		b := NewBuilder()
		b.Add("\n")
		b.Add("\n")
		b.Add("\n")
		b.Add("")
		want := "\n" +
			"\n" +
			"\n"
		have := b.Finish()
		if have != want {
			t.Fatalf("\nwant %q\nhave %q", want, have)
		}
	})
	t.Run("level1", func(t *testing.T) {
		b := NewBuilder()
		b.Add("# 1. one\n")
		b.Add("# 6. two\n")
		b.Add("#       5.        three")
		want := "# 1. one\n" +
			"# 2. two\n" +
			"# 3. three"
		have := b.Finish()
		if have != want {
			t.Fatalf("\nwant %q\nhave %q", want, have)
		}
	})
	t.Run("level2", func(t *testing.T) {
		b := NewBuilder()
		b.Add("# 1. a\n")
		b.Add("# 6. b\n")
		b.Add("## 6.6. b1\n")
		b.Add("## 6.7. b2\n")
		b.Add("# 5. c\n")
		want := "# 1. a\n" +
			"# 2. b\n" +
			"## 2.1. b1\n" +
			"## 2.2. b2\n" +
			"# 3. c\n"
		have := b.Finish()
		if have != want {
			t.Fatalf("\nwant %q\nhave %q", want, have)
		}
	})
	t.Run("level4", func(t *testing.T) {
		b := NewBuilder()
		b.Add("# 1. 1\n")
		b.Add("### 1.1.1. 1_1_1\n")
		b.Add("### 1.1.1. 1_1_2")
		want := "# 1. 1\n" +
			"### 1.1.1. 1_1_1\n" +
			"### 1.1.2. 1_1_2"
		have := b.Finish()
		if have != want {
			t.Fatalf("\nwant %q\nhave %q", want, have)
		}
	})
	t.Run("level6", func(t *testing.T) {
		b := NewBuilder()
		b.Add("# 1. 1\n")
		b.Add("# 1. 2\n")
		b.Add("## 1.1. 2_1\n")
		b.Add("## 1.1. 2_2\n")
		b.Add("# 1. 3\n")
		b.Add("## 1.1. 3_1\n")
		b.Add("## 1.1. 3_2\n")
		b.Add("#### 1.1.1.1. 3_2_1_1\n")
		b.Add("#### 1.1.1.1. 3_2_1_2\n")
		b.Add("# 1. 4\n")
		want := "# 1. 1\n" +
			"# 2. 2\n" +
			"## 2.1. 2_1\n" +
			"## 2.2. 2_2\n" +
			"# 3. 3\n" +
			"## 3.1. 3_1\n" +
			"## 3.2. 3_2\n" +
			"#### 3.2.1.1. 3_2_1_1\n" +
			"#### 3.2.1.2. 3_2_1_2\n" +
			"# 4. 4\n"
		have := b.Finish()
		if have != want {
			t.Fatalf("\nwant %q\nhave %q", want, have)
		}
	})
	t.Run("small.md", func(t *testing.T) {
		lines := splitLines(mustReadFile("testdata/small.md"))
		b := NewBuilder()
		for _, line := range lines {
			b.Add(line)
		}
		have := b.Finish()
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
