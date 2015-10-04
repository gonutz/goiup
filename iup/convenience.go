package iup

import "errors"

func StartProgram(guiCode, startWindowHandleID string) error {
	if err := Open(); err != nil {
		return err
	}
	if err := LoadBuffer(guiCode); err != nil {
		return err
	}
	return RunMainDialog(startWindowHandleID)
}

func RunMainDialog(handleName string) error {
	dialog := GetHandle(handleName)
	if dialog == nil {
		return errors.New("handle " + handleName + " not found")
	}
	dialog.SetCallback("CLOSE_CB", func() int { return CLOSE })
	if ret := dialog.Show(); ret == ERROR {
		return errors.New("Show() on " + handleName + " returned an error")
	}
	if ret := MainLoop(); ret == ERROR {
		return errors.New("MainLoop() after showing " + handleName + " returned an error")
	}
	return nil
}
