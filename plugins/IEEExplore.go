package plugins

import (
	"encoding/json"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/koffeinsource/go-URLextract/webpage"
	"github.com/koffeinsource/go-klogger"
)

type ieeeInfo struct {
	UserInfo struct {
		Institute           bool `json:"institute"`
		Member              bool `json:"member"`
		Individual          bool `json:"individual"`
		Guest               bool `json:"guest"`
		SubscribedContent   bool `json:"subscribedContent"`
		FileCabinetContent  bool `json:"fileCabinetContent"`
		FileCabinetUser     bool `json:"fileCabinetUser"`
		ShowPatentCitations bool `json:"showPatentCitations"`
		ShowGet802Link      bool `json:"showGet802Link"`
		ShowOpenURLLink     bool `json:"showOpenUrlLink"`
	} `json:"userInfo"`
	Authors []struct {
		Name        string `json:"name"`
		Affiliation string `json:"affiliation"`
	} `json:"authors"`
	Isbn []struct {
		Format string `json:"format"`
		Value  string `json:"value"`
	} `json:"isbn"`
	ArticleNumber string `json:"articleNumber"`
	DbTime        string `json:"dbTime"`
	Metrics       struct {
		CitationCountPaper  int `json:"citationCountPaper"`
		CitationCountPatent int `json:"citationCountPatent"`
		TotalDownloads      int `json:"totalDownloads"`
	} `json:"metrics"`
	PdfURL          string `json:"pdfUrl"`
	PurchaseOptions struct {
		ShowOtherFormatPricingTab bool `json:"showOtherFormatPricingTab"`
		ShowPdfFormatPricingTab   bool `json:"showPdfFormatPricingTab"`
		PdfPricingInfoAvailable   bool `json:"pdfPricingInfoAvailable"`
		OtherPricingInfoAvailable bool `json:"otherPricingInfoAvailable"`
		MandatoryBundle           bool `json:"mandatoryBundle"`
		OptionalBundle            bool `json:"optionalBundle"`
		PdfPricingInfo            []struct {
			MemberPrice    string `json:"memberPrice"`
			NonMemberPrice string `json:"nonMemberPrice"`
			PartNumber     string `json:"partNumber"`
			Type           string `json:"type"`
		} `json:"pdfPricingInfo"`
	} `json:"purchaseOptions"`
	FormulaStrippedArticleTitle string `json:"formulaStrippedArticleTitle"`
	Title                       string `json:"title"`
	Abstract                    string `json:"abstract"`
	PublicationTitle            string `json:"publicationTitle"`
	EndPage                     string `json:"endPage"`
	StartPage                   string `json:"startPage"`
	Doi                         string `json:"doi"`
	RightsLink                  string `json:"rightsLink"`
	DisplayPublicationTitle     string `json:"displayPublicationTitle"`
	PdfPath                     string `json:"pdfPath"`
	Keywords                    []struct {
		Type string   `json:"type"`
		Kwd  []string `json:"kwd"`
	} `json:"keywords"`
	AllowComments                   bool   `json:"allowComments"`
	PubLink                         string `json:"pubLink"`
	IssueLink                       string `json:"issueLink"`
	StandardTitle                   string `json:"standardTitle"`
	IsJournal                       bool   `json:"isJournal"`
	IsConference                    bool   `json:"isConference"`
	DateOfInsertion                 string `json:"dateOfInsertion"`
	IsStandard                      bool   `json:"isStandard"`
	Publisher                       string `json:"publisher"`
	PublicationDate                 string `json:"publicationDate"`
	ConferenceDate                  string `json:"conferenceDate"`
	IsACM                           bool   `json:"isACM"`
	IsOpenAccess                    bool   `json:"isOpenAccess"`
	IsEphemera                      bool   `json:"isEphemera"`
	AccessionNumber                 string `json:"accessionNumber"`
	HTMLLink                        string `json:"htmlLink"`
	IsEarlyAccess                   bool   `json:"isEarlyAccess"`
	IsBook                          bool   `json:"isBook"`
	IsDynamicHTML                   bool   `json:"isDynamicHtml"`
	IsFreeDocument                  bool   `json:"isFreeDocument"`
	IsSMPTE                         bool   `json:"isSMPTE"`
	IsCustomDenial                  bool   `json:"isCustomDenial"`
	PersistentLink                  string `json:"persistentLink"`
	IsStaticHTML                    bool   `json:"isStaticHtml"`
	IsNotDynamicOrStatic            bool   `json:"isNotDynamicOrStatic"`
	HTMLAbstractLink                string `json:"htmlAbstractLink"`
	IsPromo                         bool   `json:"isPromo"`
	JournalDisplayDateOfPublication string `json:"journalDisplayDateOfPublication"`
	ChronOrPublicationDate          string `json:"chronOrPublicationDate"`
	CopyrightYear                   string `json:"copyrightYear"`
	CitationCount                   string `json:"citationCount"`
	CopyrightOwner                  string `json:"copyrightOwner"`
	XplorePubID                     string `json:"xplore-pub-id"`
	Issue                           string `json:"issue"`
	XploreIssue                     string `json:"xplore-issue"`
	ContentType0                    string `json:"contentType"`
	ArticleID                       string `json:"articleId"`
	IsNumber                        string `json:"isNumber"`
	Lastupdate                      string `json:"lastupdate"`
	Sections                        struct {
		Abstract   string `json:"abstract"`
		Authors    string `json:"authors"`
		Figures    string `json:"figures"`
		Multimedia string `json:"multimedia"`
		References string `json:"references"`
		Citedby    string `json:"citedby"`
		Keywords   string `json:"keywords"`
		Footnotes  string `json:"footnotes"`
		Disclaimer string `json:"disclaimer"`
	} `json:"sections"`
	MlHTMLFlag        string `json:"ml_html_flag"`
	PublicationNumber string `json:"publicationNumber"`
	EphemeraFlag      string `json:"ephemeraFlag"`
	ChronDate         string `json:"chronDate"`
	MlTime            string `json:"mlTime"`
	ContentType1      string `json:"content_type"`
	MediaPath         string `json:"mediaPath"`
	OnlineDate        string `json:"onlineDate"`
	OpenAccessFlag    string `json:"openAccessFlag"`
	HTMLFlag          string `json:"html_flag"`
	SubType           string `json:"subType"`
	PublicationYear   string `json:"publicationYear"`
}

// IEEExplore parse an IEEE xplore webpage
func IEEExplore(i *webpage.Info, sourceURL string, doc *goquery.Document, log klogger.KLogger) {
	if !strings.Contains(sourceURL, "ieeexplore.ieee.org/document/") {
		return
	}
	log.Infof("Running IEEExplore plugin.")
	html, err := doc.Html()
	if err != nil {
		log.Infof("IEEExplore plugin: could not convert doc to html. " + sourceURL)
		return
	}

	const jsonStartStr = "global.document.metadata="

	pos := strings.Index(html, jsonStartStr)
	if pos == -1 {
		log.Infof("IEEExplore plugin: could not find json start. " + sourceURL)
		return
	}

	pos += len(jsonStartStr)

	jsonText := html[pos:]
	pos = strings.Index(jsonText, "\n")
	if pos == -1 {
		log.Infof("IEEExplore plugin: could not find json newline. " + sourceURL)
		return
	}

	jsonText = jsonText[:pos-1] // -1 because of a ';' at the end of the line

	var info ieeeInfo
	err = json.Unmarshal([]byte(jsonText), &info)
	if err != nil {
		log.Infof("IEEExplore plugin: error parsing json. " + sourceURL)
		return
	}

	i.Caption = info.Title

	i.Description = info.Abstract
	if info.Doi != "" {
		i.URL = "http://dx.doi.org/" + info.Doi
	}
}
