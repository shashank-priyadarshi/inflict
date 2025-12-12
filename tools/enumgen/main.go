package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type enumSpec struct {
	name   string
	values []string
}

func main() {
	descriptorPath := flag.String("descriptor", "", "path to file descriptor set")
	outPath := flag.String("out", "", "path to write enum ddl")
	flag.Parse()

	if *descriptorPath == "" || *outPath == "" {
		fmt.Fprintln(os.Stderr, "descriptor and out are required")
		os.Exit(1)
	}

	setBytes, err := os.ReadFile(*descriptorPath)
	if err != nil {
		exitErr(err)
	}

	var fds descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(setBytes, &fds); err != nil {
		exitErr(err)
	}

	var enums []enumSpec
	seen := make(map[string]bool)

	for _, fd := range fds.File {
		for _, ed := range fd.EnumType {
			targetName, ok := deduceTargetName(ed.GetName())
			if !ok {
				continue
			}
			if seen[targetName] {
				continue
			}
			enums = append(enums, enumSpec{
				name:   targetName,
				values: enumValues(ed),
			})
			seen[targetName] = true
		}
	}

	if len(enums) == 0 {
		exitErr(fmt.Errorf("no target enums found"))
	}

	// Sort enums by name to ensure deterministic output
	sort.Slice(enums, func(i, j int) bool {
		return enums[i].name < enums[j].name
	})

	var buf bytes.Buffer
	for i, e := range enums {
		if i > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString("CREATE TYPE ")
		buf.WriteString(e.name)
		buf.WriteString(" AS ENUM (\n    '")
		buf.WriteString(strings.Join(e.values, "',\n    '"))
		buf.WriteString("'\n);\n")
	}

	if err := os.WriteFile(*outPath, buf.Bytes(), 0o644); err != nil {
		exitErr(err)
	}
}

func enumValues(ed *descriptorpb.EnumDescriptorProto) []string {
	out := make([]string, 0, len(ed.Value))
	for _, v := range ed.Value {
		// Skip the zero value if it looks like UNSPECIFIED or similar,
		// but typically we just dump all values.
		// However, standard style often has 0 as unknown/unspecified.
		// For DB Enums, we usually only want the valid payloads.
		// Looking at existing files, 'Earner' = 0. So we keep all.
		out = append(out, v.GetName())
	}
	return out
}

func deduceTargetName(name string) (string, bool) {
	if strings.HasSuffix(name, "T") {
		// e.g. AmountT -> amount_type
		base := strings.TrimSuffix(name, "T")
		return strings.ToLower(base) + "_type", true
	}
	return "", false
}

func exitErr(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
