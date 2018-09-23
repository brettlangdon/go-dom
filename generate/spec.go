package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type Spec struct {
	Arguments       []Argument      `json:"arguments"`
	ExtraAttributes ExtraAttributes `json:"extAttrs"`
	IDLType         IDLType         `json:"idlType"`
	Includes        []Includes      `json:"includes"`
	Inheritance     Inheritance     `json:"inheritance"`
	RawMembers      []Member        `json:"members"`
	Name            string          `json:"name"`
	Partials        []Spec          `json:"partials"`
	Trivia          interface{}     `json:"trivia"`
	Type            string          `json:"type"`
	Values          []Value         `json:"values"`
}
type SpecMap map[string]Spec

func LoadSpecs(f string) (specs SpecMap, err error) {
	fp, err := os.Open(f)
	if err != nil {
		return
	}
	defer fp.Close()

	bytes, err := ioutil.ReadAll(fp)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &specs)
	return
}

func (s Spec) Shortname() string {
	return string(strings.ToLower(s.Name)[0])
}

func (s Spec) Receiver() string {
	return fmt.Sprintf("(%s %s)", s.Shortname(), s.Name)
}

func (s Spec) Filename() string {
	return strings.ToLower(s.Name) + ".go"
}

func (s Spec) ResolveInheritance(specs SpecMap) (inheritance []Spec, err error) {
	i := s.Inheritance
	for i.Name != "" {
		// This wasn't defined in any of our sources :(
		if i.Name == "MouseEvent" {
			return
		}
		spec, ok := specs[i.Name]
		if !ok {
			err = fmt.Errorf("Could not find spec for inheritance %q", i.Name)
			return
		}
		inheritance = append(inheritance, spec)
		i = spec.Inheritance
	}
	return
}

func (s Spec) Members() (mems []Member) {
	memMap := make(map[string][]Member)
	for _, m := range s.RawMembers {
		if m.Title() == "" {
			continue
		}
		memMap[m.Title()] = append(memMap[m.Title()], m)
	}

	for _, members := range memMap {
		if len(members) == 1 {
			// Do nothing
		} else if len(members) == 2 {
			if len(members[1].Body.Arguments) > len(members[0].Body.Arguments) {
				members[1].Body.Name.Value = fmt.Sprintf("%sWithArgs", members[1].Body.Name.Value)
				members[1].Body.Name.Escaped = fmt.Sprintf("%sWithArgs", members[1].Body.Name.Escaped)
			} else if len(members[0].Body.Arguments) > len(members[1].Body.Arguments) {
				members[0].Body.Name.Value = fmt.Sprintf("%sWithArgs", members[0].Body.Name.Value)
				members[0].Body.Name.Escaped = fmt.Sprintf("%sWithArgs", members[0].Body.Name.Escaped)
			}
		} else if s.Name == "WebSocket" && members[0].Title() == "Send" {
			for i, m := range members {
				members[i].Body.Name.Value = fmt.Sprintf("%sWith%s", m.Body.Name.Value, m.Body.Arguments[0].IDLType.IDLType)
				members[i].Body.Name.Escaped = fmt.Sprintf("%sWith%s", m.Body.Name.Escaped, m.Body.Arguments[0].IDLType.IDLType)
			}
		}
		mems = append(mems, members...)
	}

	return
}

func (s Spec) ResolveMembers(specs SpecMap, includeParents bool) (mems []Member, err error) {
	parents, err := s.ResolveInheritance(specs)
	if err != nil {
		return
	}

	memMap := make(map[string]Member)

	if includeParents {
		for _, p := range parents {
			for _, m := range p.Members() {
				memMap[m.Title()] = m
			}
		}
	}

	for _, i := range s.Includes {
		p, ok := specs[i.Includes]
		if !ok {
			err = fmt.Errorf("Cannot include unknown type %q", i.Includes)
			return
		}

		for _, m := range p.Members() {
			memMap[m.Title()] = m
		}
	}

	for _, m := range s.Members() {
		memMap[m.Title()] = m
	}

	for _, m := range memMap {
		mems = append(mems, m)
	}

	sort.Slice(mems, func(i, j int) bool { return mems[i].Title() < mems[j].Title() })
	return
}
