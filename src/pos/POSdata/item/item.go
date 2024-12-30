package item

// Item
// ====
//
// structure representing a restaurant item
type Item struct {
	id        int
	name      string
	cost      string
	modifiers []Item
}
