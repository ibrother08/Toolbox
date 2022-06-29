package form

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type OtherBox struct {
	Pnl *vcl.TPanel
}

func (f *TBMainForm) CreateOtherBox() {
	f.OtherBox = &OtherBox{}
	owner := f
	parent := f.Sheet2
	box := f.OtherBox

	box.Pnl = vcl.NewPanel(owner)
	box.Pnl.SetParent(parent)
	box.Pnl.SetAlign(types.AlClient)
	box.Pnl.SetTextBuf("敬请期待")
}
