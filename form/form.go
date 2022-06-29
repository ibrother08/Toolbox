package form

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TBMainForm struct {
	*vcl.TForm
	Pgc    *vcl.TPageControl
	Sheet1 *vcl.TTabSheet
	Sheet2 *vcl.TTabSheet

	*ConvertBox
	*OtherBox
}

var MainForm *TBMainForm

func (f *TBMainForm) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("Toolbox by AltairWang")
	f.SetPosition(types.PoScreenCenter)
	f.SetWidth(800)
	f.SetHeight(500)

	cst := f.Constraints()
	cst.SetMinWidth(types.TConstraintSize(800))
	cst.SetMinHeight(types.TConstraintSize(500))

	f.Pgc = vcl.NewPageControl(f)
	f.Pgc.SetParent(f)
	f.Pgc.SetAlign(types.AlClient)

	f.Sheet1 = vcl.NewTabSheet(f)
	f.Sheet1.SetPageControl(f.Pgc)
	f.Sheet1.SetCaption("struct转换")

	f.Sheet2 = vcl.NewTabSheet(f)
	f.Sheet2.SetPageControl(f.Pgc)
	f.Sheet2.SetCaption("其他")

	f.CreateConvertBox()
	f.CreateOtherBox()
}
