package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Parse CREATE TYPE ... AS ENUM blocks (case-insensitive, multi-line)
var enumBlock = regexp.MustCompile(`(?is)create\s+type\s+([a-zA-Z0-9_]+)\s+as\s+enum\s*\(([^)]*)\)`)

func main() {
	enumsPath := flag.String("enums", "./api/schema/db/v1/migrations/_generated_enums.sql", "path to generated enums sql")
	descriptorPath := flag.String("descriptor", "./tmp/proto.descriptor.bin", "path to file descriptor set")
	flag.Parse()

	if *descriptorPath == "" || *enumsPath == "" {
		fail("descriptor and enums args required")
	}

	// 1. Get Enums from Proto Descriptor
	protoEnums := getProtoEnums(*descriptorPath)

	// 2. Get Enums from SQL File
	dbEnums, err := readEnumFile(*enumsPath)
	if err != nil {
		fail(fmt.Sprintf("read enums: %v", err))
	}

	// 3. Compare
	for name, protoVals := range protoEnums {
		dbVals, ok := dbEnums[name] // dbEnums are lowercase keys from readEnumFile
		if !ok {
			fail(fmt.Sprintf("db enum %s component missing in sql", name))
		}
		if !sameValues(protoVals, dbVals) {
			fail(fmt.Sprintf("enum drift for %s:\n  Proto: %v\n  DB:    %v", name, protoVals, dbVals))
		}
	}

	// Optional: Check if DB has extras?
	for name := range dbEnums {
		if _, ok := protoEnums[name]; !ok {
			fmt.Printf("Warning: DB has extra enum %s not found in protos (or ignored by heuristic)\n", name)
		}
	}
}

func getProtoEnums(path string) map[string][]string {
	content, err := os.ReadFile(path)
	if err != nil {
		fail(err.Error())
	}
	var fds descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(content, &fds); err != nil {
		fail(err.Error())
	}

	res := make(map[string][]string)
	for _, fd := range fds.File {
		for _, ed := range fd.EnumType {
			targetName, ok := deduceTargetName(ed.GetName())
			if !ok {
				continue
			}
			res[targetName] = enumValues(ed)
		}
	}
	return res
}

func deduceTargetName(name string) (string, bool) {
	if strings.HasSuffix(name, "T") {
		base := strings.TrimSuffix(name, "T")
		return strings.ToLower(base) + "_type", true
	}
	return "", false
}

func enumValues(ed *descriptorpb.EnumDescriptorProto) []string {
	out := make([]string, 0, len(ed.Value))
	for _, v := range ed.Value {
		out = append(out, v.GetName())
	}
	sort.Strings(out)
	return out
}

func readEnumFile(path string) (map[string][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res := make(map[string][]string)
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 1024*1024), 1024*1024)
	var sb strings.Builder
	for sc.Scan() {
		sb.WriteString(sc.Text())
		sb.WriteByte('\n')
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	content := sb.String()
	matches := enumBlock.FindAllStringSubmatch(content, -1)
	for _, m := range matches {
		if len(m) != 3 {
			continue
		}
		name := strings.ToLower(m[1])
		raw := m[2]
		vals := splitEnumValues(raw)
		sort.Strings(vals)
		res[name] = vals
	}
	return res, nil
}

func splitEnumValues(raw string) []string {
	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		p = strings.Trim(p, "'\"")
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func sameValues(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func fail(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
