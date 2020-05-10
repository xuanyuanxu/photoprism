package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLabelBySlug(t *testing.T) {
	t.Run("files found", func(t *testing.T) {
		label, err := LabelBySlug("flower")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "Flower", label.LabelName)
	})

	t.Run("no files found", func(t *testing.T) {
		label, err := LabelBySlug("111")

		assert.Error(t, err, "record not found")
		assert.Empty(t, label.ID)
	})
}

func TestLabelByUUID(t *testing.T) {
	t.Run("files found", func(t *testing.T) {
		label, err := LabelByUUID("lt9k3pw1wowuy3c5")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "COW", label.LabelName)
	})

	t.Run("no files found", func(t *testing.T) {
		label, err := LabelByUUID("111")

		assert.Error(t, err, "record not found")
		assert.Empty(t, label.ID)
	})
}

func TestLabelThumbBySlug(t *testing.T) {
	t.Run("files found", func(t *testing.T) {
		file, err := LabelThumbBySlug("flower")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "exampleFileName.jpg", file.FileName)
	})

	t.Run("no files found", func(t *testing.T) {
		file, err := LabelThumbBySlug("cow")

		assert.Error(t, err, "record not found")
		t.Log(file)
	})
}

func TestLabelThumbByUUID(t *testing.T) {
	t.Run("files found", func(t *testing.T) {
		file, err := LabelThumbByUUID("lt9k3pw1wowuy3c4")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "exampleFileName.jpg", file.FileName)
	})

	t.Run("no files found", func(t *testing.T) {
		file, err := LabelThumbByUUID("14")

		assert.Error(t, err, "record not found")
		t.Log(file)
	})
}
