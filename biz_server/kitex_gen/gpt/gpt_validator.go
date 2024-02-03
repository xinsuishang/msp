// Code generated by Validator v0.2.3. DO NOT EDIT.

package gpt

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *ApplicationReq) IsValid() error {
	if p.ModelId <= int32(0) {
		return fmt.Errorf("field ModelId gt rule failed, current value: %v", p.ModelId)
	}
	if len(p.SessionId) < int(5) {
		return fmt.Errorf("field SessionId min_len rule failed, current value: %d", len(p.SessionId))
	}
	if len(p.SessionId) > int(100) {
		return fmt.Errorf("field SessionId max_len rule failed, current value: %d", len(p.SessionId))
	}
	if len(p.Prompt) < int(1) {
		return fmt.Errorf("field Prompt min_len rule failed, current value: %d", len(p.Prompt))
	}
	if len(p.Prompt) > int(2000) {
		return fmt.Errorf("field Prompt max_len rule failed, current value: %d", len(p.Prompt))
	}
	return nil
}
func (p *ApplicationResp) IsValid() error {
	return nil
}
func (p *ApplicationListReq) IsValid() error {
	if p.Page != nil {
		if err := p.Page.IsValid(); err != nil {
			return fmt.Errorf("field Page not valid, %w", err)
		}
	}
	return nil
}
func (p *ApplicationInfo) IsValid() error {
	return nil
}
func (p *ApplicationListResp) IsValid() error {
	if p.Page != nil {
		if err := p.Page.IsValid(); err != nil {
			return fmt.Errorf("field Page not valid, %w", err)
		}
	}
	return nil
}
