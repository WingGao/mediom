package controllers

import "github.com/revel/revel"

func init() {
	revel.InterceptMethod((*App).Before, revel.BEFORE)
	revel.InterceptMethod((*App).After, revel.AFTER)
	revel.InterceptMethod((*Accounts).Before, revel.BEFORE)
	// revel.InterceptMethod((*Accounts).After, revel.AFTER)
}
