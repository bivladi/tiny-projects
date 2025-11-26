package main

import "testing"

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

func Example_main() {
	main()
	// Output:
	// Here are the books in common:
	// - The Handmaid's Tale by Margaret Atwood
}

func TestLoadBookworms_Success(t *testing.T) {
	type testCase struct {
		bookwormsFile string
		want          []Bookworm
		wantErr       bool
	}
	tests := map[string]testCase{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			bookwormsFile: "testdata/no_such_file.json",
			want:          nil,
			wantErr:       true,
		},
		"invalid JSON": {
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantErr:       true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(tc.bookwormsFile)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected err, got nothing")
				}
				return
			}
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if !equalBookworms(t, got, tc.want) {
				t.Fatalf("different result: got %v, expected %v", got, tc.want)
			}
		})
	}
}

func TestBooksCount(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  map[Book]uint
	}{
		"nominal use case": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{
				handmaidsTale: 2,
				theBellJar:    1,
				oryxAndCrake:  1,
				janeEyre:      1,
			},
		},
		"no bookworms": {
			input: []Bookworm{},
			want:  map[Book]uint{},
		},
		"bookworm without books": {
			input: []Bookworm{
				{Name: "Serji", Books: []Book{}},
				{Name: "Vasya", Books: []Book{handmaidsTale}},
			},
			want: map[Book]uint{
				handmaidsTale: 1,
			},
		},
		"bookworm with 2 same books": {
			input: []Bookworm{
				{Name: "Jack", Books: []Book{handmaidsTale, handmaidsTale, oryxAndCrake}},
				{Name: "Sam", Books: []Book{oryxAndCrake, handmaidsTale}},
			},
			want: map[Book]uint{
				handmaidsTale: 3,
				oryxAndCrake:  2,
			},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)
			if !equalBooksCount(t, tc.want, got) {
				t.Fatalf(
					"got a different list of books: %v, expected %v",
					got,
					tc.want,
				)
			}
		})
	}
}

func TestFindCommonBooks(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  []Book
	}{
		"1 common books": {
			input: []Bookworm{
				{Name: "John", Books: []Book{handmaidsTale, janeEyre, oryxAndCrake}},
				{Name: "Sam", Books: []Book{handmaidsTale, theBellJar}},
			},
			want: []Book{handmaidsTale},
		},
		"0 common books": {
			input: []Bookworm{
				{Name: "John", Books: []Book{janeEyre, oryxAndCrake}},
				{Name: "Sam", Books: []Book{handmaidsTale, theBellJar}},
			},
			want: []Book{},
		},
		"3 bookworms have the same books on their shelves": {
			input: []Bookworm{
				{Name: "John", Books: []Book{janeEyre, oryxAndCrake}},
				{Name: "Sam", Books: []Book{janeEyre, oryxAndCrake}},
				{Name: "Ann", Books: []Book{janeEyre, oryxAndCrake}},
			},
			want: []Book{janeEyre, oryxAndCrake},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := findCommonBooks(tc.input)
			if !equalBooks(t, got, tc.want) {
				t.Fatalf("expected %v got %v", tc.want, got)
			}
		})
	}
}

func equalBookworms(t *testing.T, bookworms, target []Bookworm) bool {
	t.Helper()

	if len(bookworms) != len(target) {
		return false
	}
	for i := range bookworms {
		if bookworms[i].Name != target[i].Name {
			return false
		}
		if !equalBooks(t, bookworms[i].Books, target[i].Books) {
			return false
		}
	}
	return true
}

func equalBooks(t *testing.T, books, target []Book) bool {
	t.Helper()
	if len(books) != len(target) {
		return false
	}
	for i := range books {
		if books[i] != target[i] {
			return false
		}
	}
	return true
}

func equalBooksCount(t *testing.T, got, want map[Book]uint) bool {
	t.Helper()

	if len(got) != len(want) {
		return false
	}
	for book, targetCount := range want {
		count, ok := got[book]
		if !ok || targetCount != count {
			return false
		}
	}
	return true
}
