package source

type FDNS struct {
	FqDNS  string `bson:"fqdns" json:"fqdns,omitempty"`
	Name   string `bson:"name" json:"-"`
	Domain string `bson:"domain" json:"-"`
	Type   string `bson:"type" json:"type,omitempty"`
	Value  string `bson:"value" json:"value,omitempty"`
}
