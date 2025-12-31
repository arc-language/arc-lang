package irgen

// DeferAction is a closure that generates IR when executed.
type DeferAction func(g *Generator)

// DeferStack manages the LIFO order of defer statements for the current function.
type DeferStack struct {
	actions []DeferAction
}

func NewDeferStack() *DeferStack {
	return &DeferStack{
		actions: make([]DeferAction, 0),
	}
}

// Add pushes a new defer action onto the stack
func (ds *DeferStack) Add(action DeferAction) {
	ds.actions = append(ds.actions, action)
}

// Emit executes the actions in reverse order (LIFO)
// This is called immediately before a 'ret' instruction is generated.
func (ds *DeferStack) Emit(g *Generator) {
	for i := len(ds.actions) - 1; i >= 0; i-- {
		ds.actions[i](g)
	}
}