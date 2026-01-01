package models

import (
	"time"
)

var Todos = []Todo{
	{
		ID:        1,
		Title:     "todo1 title",
		Content:	 "todo1 content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        2,
		Title:     "todo2 title",
		Content:	 "todo2 content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
