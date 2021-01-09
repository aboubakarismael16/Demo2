package main

import "fmt"

type AuthorDetail interface {
	detail()
}

type AuthorArticle interface {
	article()
	picked()
}

type FinalDetail interface {

	AuthorArticle

	//even you call the method
	//the output is same when you call the interface
	detail()
	//article()
	cdetail()
}

type author struct {
	name      string
	branch    string
	college   string
	year      int
	salary    int
	particles int
	tarticles int
	pick int
	post string
	cid int
}

func (a author) detail()  {
	fmt.Printf("\nauthor name: %s",a.name)
	fmt.Printf("\nBranch : %s and passing year : %d",a.branch,a.year)
	fmt.Printf("\nCollege name: %s",a.college)
	fmt.Printf("\nSalary : %d",a.salary)
	fmt.Printf("\nPublished articles: %d",a.particles)

}

func (a author) article()  {
	pendingArtilces := a.tarticles - a.particles
	fmt.Printf("\nPending articles: %d", pendingArtilces)

}

func (a author) picked()  {
	fmt.Printf("\nTotal number of picked articles: %d", a.pick)

}

func (a author) cdetail()  {
	fmt.Printf("\nAuthor Id: %d",a.cid)
	fmt.Printf("\nAuthor Post: %s",a.post)

}

func main() {

	value := author{
		name:      "ismael",
		branch:    "engineering",
		college:   "NEPU",
		year:      2020,
		salary:    2000,
		particles: 12,
		tarticles: 20,
		cid: 180006150105,
		post: "PDG",
		pick: 3000,

	}

	var v FinalDetail = value
	v.detail()
	v.article()
	v.cdetail()
	v.picked()

}
