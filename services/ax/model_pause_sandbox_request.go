package ax

type PauseSandboxRequest struct {
	SandboxID     *string `json:"-"`
	HibernateMode *string `json:"-"`
}
