package components

import (
	"bytes"
	"fmt"
	"github.com/adrg/frontmatter"
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/components"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"golang.org/x/exp/slices"
	"log"
	"os"
	"path/filepath"
	"sort"

	"strings"
)

const (
	omitPage = "omit"
)

const (
	public = iota
	private
	internal
	user
	maintainer
	supervisor
	admin
)

var roleRanks = map[string]int{
	"public":     public,
	"private":    private,
	"internal":   internal,
	"user":       user,
	"maintainer": maintainer,
	"supervisor": supervisor,
	"admin":      admin,
}

// roleToRank returns the rank of a given role string
func roleToRank(role string) (int, bool) {
	rank, ok := roleRanks[role]
	return rank, ok
}

// compareRoles compares two role strings and returns:
// -1 if role1 < role2
//
//	0 if role1 == role2
//	1 if role1 > role2
func CompareRoles(role1 string, comparison []string) (int, error) {
	if len(comparison) == 0 {
		return 1, nil //they are allowed in
	}
	rank1, ok1 := roleToRank(role1)
	if !ok1 {
		return 0, fmt.Errorf("invalid role: %s", role1)
	}

	//if the length of the comparison is 1, then the user can view if they are greater role
	if len(comparison) == 1 {
		rank2, ok2 := roleToRank(comparison[0])
		if !ok2 {
			return 0, fmt.Errorf("invalid role: %s", comparison[0])
		}
		if rank1 < rank2 {
			return -1, nil
		} else if rank1 == rank2 {
			return 0, nil
		} else {
			return 1, nil
		}
	}
	//if its greater than 1, they need to match exactly
	if slices.Contains(comparison, role1) {
		return 0, nil //there are one of the allowed roles
	}
	return -1, nil //they are not allowed in
}

type matter struct {
	Group string   `yaml:"group"` //if no group it goes top level
	Name  string   `yaml:"name"`  //the file name currently
	Roles []string `yaml:"roles"` //the permissions
}

type link struct {
	Name  string
	Group string // You can decide how to populate this as it wasn't clear from the Description
	URL   string
}

// GenerateLinks traverses the given directory of markdown files,
// extracts the roles from the frontmatter, and returns a list of links
// for the files where the roles contain "user".
func GenerateLinks(dir string) ([]link, error) {
	var links []link

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".md") {
			content, readErr := os.ReadFile(path)
			if readErr != nil {
				return readErr
			}
			role := "user"
			var m matter
			_, err := frontmatter.Parse(bytes.NewReader(content), &m)
			if err != nil {
				log.Fatal(err)
			}
			result, err := CompareRoles(role, m.Roles)
			if err != nil {
				fmt.Println(err)
				return err
			}
			if len(m.Roles) == 0 || result > -1 {
				links = append(links, link{
					Name:  info.Name()[:len(info.Name())-len(filepath.Ext(info.Name()))],
					URL:   "/documents/" + info.Name(),
					Group: m.Group, // Assign this based on your criteria
				})
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return links, nil
}
func createSideBar(links []link) app.HTMLDiv {
	sideBarGroups := make(map[string][]string)
	for _, l := range links {
		if l.Group != omitPage { //don't show in side bar as an option - does not make it unlinkable
			sideBarGroups[l.Group] = append(sideBarGroups[l.Group], l.Name)
		}
	}
	keys := make([]string, 0, len(sideBarGroups))

	for k := range sideBarGroups {
		keys = append(keys, k)
	}
	sort.Strings(keys) //now we have the sorted keys, we can create the side bar in alphabetical order
	var accordions []app.UI
	for _, k := range keys {
		v := sideBarGroups[k]
		pathId := strings.ReplaceAll(strings.ToLower(k), " ", "-")
		accordions = append(accordions, components.AccordionMenu(pathId, k, v))
	}
	return components.Accordion(accordions...)
}
