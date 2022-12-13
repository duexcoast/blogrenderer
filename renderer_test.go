package blogrenderer_test

import (
	"bytes"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/duexcoast/blogrenderer"
)

func TestRenderTestRender(t *testing.T) {
	var aPost = blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

}
