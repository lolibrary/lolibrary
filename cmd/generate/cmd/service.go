package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

type template struct {
	// Name is the service name, with any "service" prefix stripped.
	Name string

	// ServiceName is the "service.foo" version of a name
	ServiceName string

	// ServiceShortName is the "s-foo" version of a name.
	ServiceShortName string

	// ProtoPackageName is the package name for the protobuf file (and the filename), minus the "proto" suffix.
	ProtoPackageName string

	// ProtoServiceName is the name used in the generated protobuf service.
	ProtoServiceName string
}

var serviceCmd = &cobra.Command{
	Use:     "service SERVICENAME",
	Aliases: []string{"s", "svc"},
	Args:    cobra.ExactArgs(1),
	Short:   "Generates a service using template.service as a template.",
	Run: func(cmd *cobra.Command, args []string) {
		t := getTemplate(args[0])
		baseTemplate, targetDirectory := getDirectories(t.ServiceName)

		walk(t, baseTemplate, targetDirectory)
	},
}

func walk(t template, baseTemplate, targetDirectory string) {
	err := filepath.Walk(baseTemplate, func(filepath string, info os.FileInfo, err error) error {
		target := path.Join(targetDirectory, strings.TrimPrefix(filepath, baseTemplate))
		if info.IsDir() {
			if err := os.Mkdir(target, 0755); err != nil {
				log.Fatalf("😬 couldn't create service directory: %v\n", err)
			}

			return nil
		}

		contents, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Fatalf("😬 Error reading template file: %v\n", err)
		}

		switch {
		case strings.HasSuffix(filepath, ".go"):
			contents = replaceGoContents(contents, t)
		case strings.HasSuffix(filepath, ".yml"):
			contents = replaceYAMLContents(contents, t)
		case strings.HasSuffix(filepath, ".proto"):
			target = strings.Replace(target, "template.proto", t.ProtoPackageName+".proto", 1)
			contents = replaceProtoContents(contents, t)
		}

		if err := ioutil.WriteFile(target, contents, 0644); err != nil {
			log.Fatalf("Failed to write file to %v: %v", target, err)
		}

		return nil
	})
	if err != nil {
		log.Fatalf("😬 Error while generating service: %v\n", err)
	}
}

func replaceGoContents(content []byte, t template) []byte {
	str := string(content)
	str = strings.ReplaceAll(str, "template.service", t.ServiceName)

	return []byte(str)
}

func replaceYAMLContents(content []byte, t template) []byte {
	str := string(content)

	str = strings.ReplaceAll(str, "template.service", t.ServiceName)
	str = strings.ReplaceAll(str, "s-template", t.ServiceShortName)

	return []byte(str)
}

func replaceProtoContents(content []byte, t template) []byte {
	str := string(content)

	str = strings.ReplaceAll(str, "templatepackagename", t.ProtoPackageName)
	str = strings.ReplaceAll(str, "template_service_name", t.ProtoServiceName)

	return []byte(str)
}

func getDirectories(service string) (string, string) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatal("😬 GOPATH not set, please set it first.")
	}

	root := path.Join(gopath, "/src/github.com/lolibrary/lolibrary")

	return path.Join(root, "template.service"), path.Join(root, service)
}

func validateServiceName(svc string) string {
	switch {
	case !strings.HasPrefix(svc, "service."):
		log.Fatalf("😬 services should be in the format %v.%v\n", aurora.Blue("service"), aurora.Green("foo"))
	}

	return strings.TrimPrefix(svc, "service.")
}

func getServiceNames(svc string) (string, string) {
	return "service." + svc, "s-" + svc
}

func getProtoNames(svc string) (string, string) {
	protoPackage := strings.ReplaceAll(svc, "-", "")
	protoService := strings.ReplaceAll(svc, "-", "")

	return protoPackage, protoService
}

func getTemplate(input string) template {
	name := validateServiceName(input)
	long, short := getServiceNames(name)
	protoPackage, protoService := getProtoNames(name)

	return template{
		Name:             name,
		ServiceName:      long,
		ServiceShortName: short,
		ProtoPackageName: protoPackage,
		ProtoServiceName: protoService,
	}
}
