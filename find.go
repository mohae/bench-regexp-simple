package benchregexp

import (
	"errors"
	"regexp"
	"strings"
)

// precompiled version
var reg *regexp.Regexp

func init() {
	// don't do things in init that could err
	var err error
	reg, err = regexp.Compile(`VMStateChangeTime\D(.*)`)
	if err != nil {
		panic(err)
	}
}

func findWRegExp() (string, error) {
	r, err := regexp.Compile(`VMStateChangeTime\D(.*)`)
	if err != nil {
		return "", err
	}
	str := r.FindStringSubmatch(vminfo)

	return strings.Join(str[1:], ""), nil
}

func findWRegExpPrecompiled() (string, error) {
	str := reg.FindStringSubmatch(vminfo)
	return strings.Join(str[1:], ""), nil
}

func findWIndex() (string, error) {
	i := strings.Index(vminfo, "VMStateChangeTime")
	if i < 0 {
		return "", errors.New("VMStateChangeTime: not found")
	}
	// skip the =
	i += 18
	// make
	b := make([]byte, 0, 31)
	for {
		if vminfo[i] == '\n' {
			break
		}
		b = append(b, vminfo[i])
		i++
	}
	return string(b), nil
}

func findWIndexBuf(buf []byte) error {
	i := strings.Index(vminfo, "VMStateChangeTime")
	if i < 0 {
		return errors.New("VMStateChangeTime: not found")
	}
	// skip the =
	i += 18
	var ndx int
	for {
		if vminfo[i] == '\n' {
			break
		}
		buf[ndx] = vminfo[i]
		i++
		ndx++
	}
	return nil

}
