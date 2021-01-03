package main

import "testing"

func Test_replaceWords(t *testing.T) {
	tests := []struct {
		name       string
		dictionary []string
		sentence   string
		want       string
	}{
		{
			sentence:   "the cattle was rattled by the battery",
			dictionary: []string{"cat", "bat", "rat"},
			want:       "the cat was rat by the bat",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceWords(tt.dictionary, tt.sentence); got != tt.want {
				t.Errorf("replaceWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
