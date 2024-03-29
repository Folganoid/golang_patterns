package pkg

type VendingMachine struct {
	hasItem			State
	itemRequested	State
	hasMoney		State
	noItem			State
	currentState	State
	itemCount		int
	itemPrice		int
}

func (v *VendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

func (v *VendingMachine) AddItem(count int) error {
	return v.currentState.AddItem(count)
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

func (v *VendingMachine) setState (s State) {
	v.currentState = s
}

func (v *VendingMachine) incrementItemCount(count int) {
	v.itemCount += count
}

func NewVendingMachine(count int, price int) *VendingMachine {

	v := &VendingMachine{
		itemCount: count,
		itemPrice: price,
	}

	hasItemState := &HasItemState{
		vendingMachine: v,
	}

	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}

	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}

	noItemState := &NoItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}