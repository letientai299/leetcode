// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "401f29470b49690fd28eb7184d0e8ad2"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"106843b192ca942032d276d812961979": "1f8b08000000000000ff2a484cce4e4c4f55c84dcccce3e2aaaed64bcecfcd4dcd2ba9ade5aaaed655d04bce4f49adad05040000ffff2bbf2f1126000000",
		"382a9fa8a851a254b7602b9b8cf7b7a7": "1f8b08000000000000ff52cecc4bce294d4955502a2dc9482cced0cb50e2e2aaaed64bcecfcd4dcd2ba9ade5aaaed655d04bce4f49adad05040000ffff10a8a9862d000000",
		"6730ba81dcbfe9a4a9b3049204163b12": "1f8b08000000000000ff4ccdb18a84301006e0dea7981baef0e0b807b03e16b6d9c6d6266b460964478d93050979f7256a5448913ff3e59f5000e030925362067ea817610538799a53fe57a2f03791b772463d2dcd5841001423966aebfba443f83b728c0871fd317972cb5ee616b85696df87afa01667b8fffa81d0301cacbc8853efe8643737b010ebbbdede57b95ddb75225bd0a6eb4cebad2c79a8a966338e2473ee04b08afbb426e7a41a8ef960118b4f000000ffff7c3f831830010000",
		"71cf1877a3831b42abd359681c335a40": "1f8b08000000000000ff2a484cce4e4c4f55c8494d2d49ce4f49d5abaed6838ad5d65a73715557eb25e7e7e6a6e695d4d6725557eb2ae88154d5d67201020000ffff21bd407c39000000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("default box", "./tpl")
		b.SetResolver("c_file.tpl", packr.Pointer{ForwardBox: gk, ForwardPath: "382a9fa8a851a254b7602b9b8cf7b7a7"})
		b.SetResolver("golang_file.tpl", packr.Pointer{ForwardBox: gk, ForwardPath: "106843b192ca942032d276d812961979"})
		b.SetResolver("java_file.tpl", packr.Pointer{ForwardBox: gk, ForwardPath: "71cf1877a3831b42abd359681c335a40"})
		b.SetResolver("question_data.tpl", packr.Pointer{ForwardBox: gk, ForwardPath: "6730ba81dcbfe9a4a9b3049204163b12"})
	}()

	return nil
}()