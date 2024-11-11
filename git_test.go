/*
 * Copyright © 2024 weidongkl <weidongkx@gmail.com>
 */

package git

import (
	"testing"
)

func TestClone(t *testing.T) {
	cloneObj := NewClone("https://gitee.com/os_build/lorax.git")
	//cloneObj.SetAuth("username", "password").SetBranch("sos")
	//cloneObj.SetBranch("master")
	err := cloneObj.Clone()
	if err != nil {
		t.Error(err)
	}
}

func TestPull(t *testing.T) {
	pullObj := NewPull("/path/to/local/repo")
	pullObj.SetAuth("username", "password")
	err := pullObj.Pull()
	if err != nil {
		t.Error(err)
	}
}
