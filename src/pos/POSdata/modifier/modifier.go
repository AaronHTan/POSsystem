package modifier

// Modifier
// ========
//
// modifiers that modify items
type Modifier struct {
	id           int    // id of modifier
	name         string // name of modifier, for display in front end
	cost         int    // cost of modifier
	global       bool   // global represents whether it can modify all items or not
	availability bool   // availability decides whether the item is available or not
	items        []int  // items is a list of item ids which it modifies, if it is not global
}
