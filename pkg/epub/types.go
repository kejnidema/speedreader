package epub

type PackageDocument struct {
	metadata              MetaData
	manifest              string
	spine                 SpinePlane
	collections           []string
	manifestFallbackChain string
}

type MetaData struct {
	DCIdentifier              DCIdentifier
	DCTitle                   DCTitle
	DCLanguage                DCLanguage
	DCContributor             DCElement
	DCCoverage                DCElement
	DCCreator                 DCElement
	DCDate                    DCElement
	DCDescription             DCElement
	DCFormat                  DCElement
	DCPublisher               DCElement
	DCRelation                DCElement
	DCRights                  DCElement
	DCSource                  DCElement
	DCSubject                 DCElement
	DCType                    DCElement
	DublinCoreOptionalElement string
	meta                      Meta
	OPF2meta                  string //legacy
	link                      Link
}

type Manifest struct {
	id   string
	item Item
}

type Item struct {
	fallback     string
	href         string
	id           string
	mediaOverlay string
	mediaType    string
	properties   []string
}

type Spine struct {
	id                       string
	pageProgressionDirection string
	toc                      string // legacy
	itemref                  ItemRef
}

type ItemRef struct {
	id         string
	idref      string
	linear     string
	properties []string
}

type Collection struct {
	dir      string
	id       string
	role     string
	XMLLang  string
	metadata MetaData
	// collection Collection
	link Link
}

type nav struct {
	h1 string
	h2 string
	h3 string
	h4 string
	h5 string
	h6 string
	ol []string
	li []struct {
		span string
		a    string
		li   string
	}
	span string
	a    string
}

type smil struct {
	version    string
	id         string
	epubPrefix string
	head       Head
	body       Body
}

type Head struct {
	metadata MetaData // this is false for this I think
}

type Body struct {
	epubType    string
	id          string
	epubTextRef string
	seq         Seq
	par         Par
}

type Seq struct {
	epubType    string
	id          string
	epubTextRef string
	seq         *Seq
	par         Par
}

type Par struct {
	epubType string
	id       string
	text     Text
	audio    Audio
}

type Text struct {
	src string
	id  string
}

type Audio struct {
	id        string
	src       string
	clipBegin string
	clipEnd   string
}

type MustSetProperties struct {
	mathml          string
	remoteResources string
	scripted        string
	svg             string
	propSwitch      string
}

type Link struct {
	href       string
	hrefLang   string
	id         string
	mediaType  string
	properties []string
	refines    string
	rel        string
}

type Meta struct {
	dir      string
	id       string
	property string
	refines  string
	scheme   string
	XMLLang  string
}

type DCElement struct {
	dir     string
	id      string
	XMLLang string
}

type DCIdentifier struct {
	id string
}

type DCTitle struct {
	dir     string
	id      string
	XMLLang string
}

type DCLanguage struct {
	id string
}

type mimetype struct {
}

type MetaInf struct {
	container  string
	signatures string
	encryption string
	metadata   string
	rights     string
	manifest   string
}

/*
*

	Holds all resources of EPUB (publication resources and linked resources)
*/
type ManifestPlane struct {
}

/*
*

	Holds all resources necessary to render the spine of the EPUB (EPUB content documents and foreign content documents)
*/
type SpinePlane struct {
}

/*
*

	Holds all resources necessary to render the EPUB and foreign content documents (core media type resources, foreign resources, and exempt resources)
*/
type ContentPlane struct {
}
