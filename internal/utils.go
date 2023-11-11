package internal

import "google.golang.org/protobuf/compiler/protogen"

func GetMessageByName(g *protogen.GeneratedFile, f *protogen.File, name string) *protogen.Message {
	for _, m := range f.Messages {
		if string(m.Desc.Name()) == name {
			return m
		}
	}
	return nil
}

func GetFieldByName(g *protogen.GeneratedFile, m *protogen.Message, name string) *protogen.Field {
	for _, field := range m.Fields {
		if string(field.Desc.Name()) == name {
			return field
		}
	}
	return nil
}

func Coalesce[K comparable](vals ...K) K {
	var zero K
	for _, v := range vals {
		if v != zero {
			return v
		}
	}

	return zero
}
