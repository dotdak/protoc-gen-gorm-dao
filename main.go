package main

import (
	"flag"

	"github.com/dotdak/protoc-gen-gorm-dao/gorm"
	"github.com/dotdak/protoc-gen-gorm-dao/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func main() {
	var (
		flags        flag.FlagSet
		importPrefix = flags.String("import_prefix", "", "prefix to prepend to import paths")
	)
	importRewriteFunc := func(importPath protogen.GoImportPath) protogen.GoImportPath {
		switch importPath {
		case "context", "fmt", "math":
			return importPath
		}
		if *importPrefix != "" {
			return protogen.GoImportPath(*importPrefix) + importPath
		}
		return importPath
	}

	protogen.Options{
		ParamFunc:         flags.Set,
		ImportRewriteFunc: importRewriteFunc,
	}.Run(func(p *protogen.Plugin) error {
		for _, f := range p.Files {
			if !f.Generate || f.GoPackageName == "gorm" {
				continue
			}

			shouldGenerate := false
			for _, m := range f.Messages {
				if opt, ok := proto.GetExtension(m.Desc.Options(), gorm.E_Opts).(*gorm.GormOptions); ok && opt != nil && opt.Orm {
					shouldGenerate = true
				}
			}

			if shouldGenerate {
				_ = internal.GenerateFile(p, f)
			}
		}
		return nil
	})
}
