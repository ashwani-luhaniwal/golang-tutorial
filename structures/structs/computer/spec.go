
package computer

type Spec struct {	// exported struct
	Maker string	// exported field
	model string	// unexported field
	Price int		// exported field
}