package main

func findLinks(url string) ([]string,error)  {

	resp,err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s",url,resp.Status)
	}

	doc,err != html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Parsing %s as HTML : %v",url,err)
	}

}


// func main() {

// }
