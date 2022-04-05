package customLog

import "github.com/sirupsen/logrus"

func Fields(pkg, fun string, extra ...string) logrus.Fields {
	fields := logrus.Fields{"package": pkg, "where": fun}
	if len(extra) > 0 {
		fields["extra"] = extra[0]
	}
	return fields
}
