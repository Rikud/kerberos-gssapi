package main

import (
	"github.com/Rikud/kerberos-gssapi/gssapi"

	"github.com/jcmturner/gokrb5/v8/keytab"
)

func loadKeyTab() {
	ktFile := "/etc/krb5.keytab" // TODO: make it configurable
	keytab.Load(ktFile)
}

// keytab.Load -> verifyAPReq -> continueAcceptor (gss_accept_sec_context?) -> Continue

func main() {
	// TODO: load keytab: gsskrb5_register_acceptor_identity
	loadKeyTab()

	// TODO: load cert: gss_acquire_cred
	/*server := */
	gssapi.NewMech("kerberos_v5")

	// TODO: start listen

	// TODO: gss_accept_sec_context

	// TODO: gss_wrap

	// TODO: gss_unwrap
}
