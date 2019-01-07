package store_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/language"

	"github.com/jansorg/tom/go-tom/store"
	"github.com/jansorg/tom/go-tom/testSetup"
)

func TestTags(t *testing.T) {
	ctx, err := testSetup.CreateTestContext(language.English)
	require.NoError(t, err)
	defer testSetup.CleanupTestContext(ctx)

	p, err := ctx.Store.AddProject(store.Project{Name: "p1"})
	require.NoError(t, err)

	frame, err := ctx.Store.AddFrame(store.NewStartedFrame(p))
	require.NoError(t, err)

	tag, err := ctx.Store.AddTag(store.Tag{Name: "Tag1"})
	require.NoError(t, err)
	assert.NotEmpty(t, tag)
	assert.NotEmpty(t, tag.ID)

	assert.False(t, frame.HasTag(tag))
	frame.AddTags(tag)
	assert.True(t, frame.HasTag(tag))

	oldID := tag.ID
	tag, err = ctx.Store.UpdateTag(*tag)
	require.NoError(t, err)
	assert.EqualValues(t, oldID, tag.ID)

	// rename tag and test HasTag
	tag.Name = "New Tag"
	tag, err = ctx.Store.UpdateTag(*tag)
	assert.EqualValues(t, "New Tag", tag.Name)
	assert.True(t, frame.HasTag(tag))

	// remove tag
	frame.RemoveTagID(tag.ID)
	assert.False(t, frame.HasTag(tag))
}
