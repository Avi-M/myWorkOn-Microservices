package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	if len(feed.Items) != 1 {
		t.Error("item was not added")
	}
}
func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	res := feed.GetAll()
	if len(res) != 1 {
		t.Error("item wasn't added")
	}

}
