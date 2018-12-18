package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/knakk/rdf"
	"gopkg.in/resty.v1"
)

func main() {
	fmt.Println("Jena Loader to test Lucene index on large sets")

	tpl := graphBuild()

	err := jenaLoad(tpl, "default")
	if err != nil {
		fmt.Println(err)
	}
}

func jenaLoad(tpl, graph string) error {

	resp, err := resty.R().
		SetBody(tpl).
		SetHeader("Content-Type", "application/n-triples").
		Put("http://localhost:3030/ds")

		// explore response object
	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
	fmt.Printf("\nResponse Body: %v", resp) // or resp.String() or string(resp.Body())

	return nil
}

func graphBuild() string {

	tr := []rdf.Triple{}

	newsub, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata.org/id/resource/csdco/feature/X")) // Sprintf a correct URI here
	newpred, _ := rdf.NewIRI("http://opencoredata.org/id/voc/csdco/v1/ngdc_ship")
	//	newobj, _ := rdf.NewLiteral("simple object")
	newobj, _ := rdf.NewLiteral(testText())
	newtriple := rdf.Triple{Subj: newsub, Pred: newpred, Obj: newobj}
	tr = append(tr, newtriple)

	buf := new(bytes.Buffer)

	// Write triples to a file
	var inoutFormat rdf.Format
	inoutFormat = rdf.NTriples // Turtle NQuads
	enc := rdf.NewTripleEncoder(buf, inoutFormat)
	err := enc.EncodeAll(tr)
	// err = enc.Encode(newtriple)
	enc.Close()
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func testText() string {
	text := `The evolving funding landscape at academic and research institutions has had a major impact on career
	opportunities for scientists, particularly those who are early-career. As a result of grant dollars being
	increasingly awarded to a disproportionately small number of established investigators and institutes
	(Larivière et al., 2010), intellectual discovery has become captured by a privileged few (Parker et al.
	2010), leading to greater bias in scientific research, diminished scientific productivity (Barjak &
	Robinson, 2008), and less potential for breakthrough discoveries (Fleming 2007; Nielsen 2012). Such a
	lack of social diversity and equity is a major challenge in science, technology, engineering, and
	mathematics (STEM; NCSES 2017; Sheltzer & Smith 2018). Solutions are often sought out by proposing
	adjustments to the “career pipeline”, but these same issues in STEM continue to be unresolved.

	new resources, tools, and infrastructure (courtesy of STEM advances) such as lab space, journal access,
and high performance computing, either publicly available, or available for rent, that allow science to
thrive outside of traditional institutions (the orange, next outermost circle) (Vasbinder 2017). In addition,
bottom-up changes are already being driven by early career scientists themselves in many different ways
(Arbesman & Wilkins 2012; McDowell 2014; Nicholson 2015: Hansen et al. 2018).
Many postdocs and adjunct scientists already have the majority of tools that they need to do independent
science, such as deep training and understanding of their field, a body of work that demonstrates their
scientific ability, pre-existing networks of colleagues with similar intellectual interests, and the Internet to
collaborate and share. By moving beyond the existing pipeline model of academic science, the ecosystem
vision provides the space, flexibility, and diversity that science needs to be more responsive to both local
and broader complex scales affecting science.
To demonstate how an ecosystem model would work in practice, we present a set of conceptual ​ design
patterns​ loosely inspired by commons-based approaches (Ostrom 1990; Bauwens 2005; Bollier &
Helfrich 2014), systems-thinking approaches (Meadows 2008), and the sustainable livelihoods framework
(DFID 2006). We acknowledge specific social movements and grassroots changes that are occuring
today, and demonstrate how science now has the means to be more egalitarian, inclusive, and diverse by
being less dependent on their institutional settings. We recognize, however, that major institutional
reforms are needed to realize this vision to its fullest, so we also address the changing role of institutions
within this vision. Note that we chose the term ​ ecosystem ​ deliberately to resonate with many of the
phenomena that exist in biological ecosystems: diversity, resilience, multiplicity of scales, dynamic
feedback loops etc., and we use some of these concepts when framing each of the design patterns, but we
don’t claim a one-to-one correspondence with biological ecosystems.
1. Fundamental development of the scientist. ​ Basic necessities (i.e., Maslow’s hierarchy of needs) are
fundamental to any human livelihood, and certainly for a scientist to be able to develop and flourish. To
truly allow independent scientists to develop, however, a strong set of progressive social policies such as
universal health care, basic income, and high-quality free education, are needed to strengthen the core of
the ecosystem (Lehdonvirta 2017; Standing 2017). The ecosystem concept recognizes that the journey of
a scientist through training is often an indirect path with many more career development influences than
the direct path a “pipeline” implies (​ Figure 1)​ . Instead, an individual learns foundational knowledge,
explores ideas, and gathers experience through a journey that is influenced by a broad range of interests, a
balance of personal and professional goals, and adaptation to the challenges of life overall.
Such a student might attend the traditional classes expected in their field, explore other fields of interest
(e.g., fine arts, social activism), and gather experience through travel, work, internships, and
volunteering--within their field and outside of it. Along the way, they might explore other career (or life)
choices, and perhaps return to academia completely, or explore specific scientific questions from a new
perspective in another career choice outside of traditional academic institutions. Overall, the ecosystem
model emphasizes that​ there is no right way to become a scientist ​ , for example there should be no need to
become an Assistant Professor before claiming independence. The diversity of experiences and
perspectives are key to advancing STEM development in novel and more inclusive ways.
	`

	return text
}
